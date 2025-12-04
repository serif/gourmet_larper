package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const separator = "============================================================="

var maliciousExtensions = []string{
	"bpgaffohfacaamplbbojgbiicfgedmoi",
	"cdgonefipacceedbkflolomdegncceid",
	"cihbmmokhmieaidfgamioabhhkggnehm",
	"eagiakjmjnblliacokhcalebgnhellfi", // Clean Master
	"eaokmbopbenbmgegkmoiogmpejlaikea",
	"gipnpcencdgljnaecpekokmpgnhgpela",
	"gnhgdhlkojnlgljamagoigaabdmfhfeg",
	"hlcjkaoneihodfmonjnlnnfpdcopgfjk",
	"hmhifpbclhgklaaepgbabgcpfgidkoei",
	"ibiejjpajlfljcgjndbonclhcbdcamai",
	"ijcpbhmpbaafndchbjdjchogaogelnjl",
	"imdgpklnabbkghcbhmkbjbhcomnfdige",
	"ineempkjpmbdejmdgienaphomigjjiej",
	"jbnopeoocgbmnochaadfnhiiimfpbpmf",
	"lehjnmndiohfaphecnjhopgookigekdk",
	"lhiehjmkpbhhkfapacaiheolgejcifgd",
	"llkncpcdceadgibhbedecmkencokjajg",
	"lnlononncfdnhdfmgpkdfoibmfdehfoj",
	"mljmfnkjmcdmongjnnnbbnajjdbojoci",
	"nagbiboibhbjbclhcigklajjdefaiidc",
	"nmfbniajnpceakchicdhfofoejhgjefb",
	"nnnklgkfdfbdijeeglhjfleaoagiagig",
	"ocffbdeldlbilgegmifiakciiicnoaeo",
	"ofkopmlicnffaiiabnmnaajaimmenkjn",
	"ogjneoecnllmjcegcfpaamfpbiaaiekh",
	"olaahjgjlhoehkpemnfognpgmkbedodk",
	"ondhgmkgppbdnogfiglikgpdkmkaiggk",
}

type browserInfo struct {
	name      string
	directory string
}

type profileScanResult struct {
	browserName      string
	profileName      string
	foundMalicious   []string
	installedCount   int
	extensionsPath   string
}

func main() {
	printHeader()

	browsers, err := discoverBrowsers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error discovering browsers: %v\n", err)
		os.Exit(1)
	}

	if len(browsers) == 0 {
		printNoBrowsersFound()
		os.Exit(0)
	}

	allProfiles := []profileScanResult{}

	for _, browser := range browsers {
		profiles, err := discoverBrowserProfiles(browser.directory)
		if err != nil {
			continue
		}

		maliciousMap := buildMaliciousExtensionsMap()
		results := scanAllProfiles(browser.name, profiles, maliciousMap)
		allProfiles = append(allProfiles, results...)
	}

	if len(allProfiles) == 0 {
		fmt.Println("⚠️  No browser profiles found to scan.")
		os.Exit(0)
	}

	printAllResults(allProfiles)

	if hasMaliciousExtensions(allProfiles) {
		os.Exit(1)
	}
}

func scanAllProfiles(browserName string, profiles []string, maliciousMap map[string]bool) []profileScanResult {
	results := []profileScanResult{}

	for _, profilePath := range profiles {
		profileName := filepath.Base(profilePath)
		extensionsDirectory := filepath.Join(profilePath, "Extensions")

		if err := verifyDirectoryExists(extensionsDirectory); err != nil {
			continue
		}

		foundMalicious, installedCount, err := scanExtensions(extensionsDirectory, maliciousMap)
		if err != nil {
			continue
		}

		results = append(results, profileScanResult{
			browserName:    browserName,
			profileName:    profileName,
			foundMalicious: foundMalicious,
			installedCount: installedCount,
			extensionsPath: extensionsDirectory,
		})
	}

	return results
}

func printHeader() {
	fmt.Println("🔍 Scanning browser extensions for ShadyPanda malware...")
	fmt.Printf("Platform: %s\n", runtime.GOOS)
	fmt.Println(separator)
}

func discoverBrowsers() ([]browserInfo, error) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	browsers := []browserInfo{}

	// Platform-specific paths
	var chromePath, bravePath string
	
	switch runtime.GOOS {
	case "darwin": // macOS
		chromePath = filepath.Join(homeDirectory, "Library", "Application Support", "Google", "Chrome")
		bravePath = filepath.Join(homeDirectory, "Library", "Application Support", "BraveSoftware", "Brave-Browser")
	case "windows":
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData == "" {
			localAppData = filepath.Join(homeDirectory, "AppData", "Local")
		}
		chromePath = filepath.Join(localAppData, "Google", "Chrome", "User Data")
		bravePath = filepath.Join(localAppData, "BraveSoftware", "Brave-Browser", "User Data")
	case "linux":
		chromePath = filepath.Join(homeDirectory, ".config", "google-chrome")
		bravePath = filepath.Join(homeDirectory, ".config", "BraveSoftware", "Brave-Browser")
	default:
		return nil, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	if err := verifyDirectoryExists(chromePath); err == nil {
		browsers = append(browsers, browserInfo{
			name:      "Chrome",
			directory: chromePath,
		})
	}

	if err := verifyDirectoryExists(bravePath); err == nil {
		browsers = append(browsers, browserInfo{
			name:      "Brave",
			directory: bravePath,
		})
	}

	return browsers, nil
}

func verifyDirectoryExists(directory string) error {
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %w", err)
	}
	return err
}

func printNoBrowsersFound() {
	fmt.Println("❌ No supported browsers found.")
	fmt.Println("\nSupported browsers:")
	fmt.Println("  • Google Chrome")
	fmt.Println("  • Brave Browser")
	fmt.Println("\nMake sure at least one of these browsers is installed and has been run.")
	
	// Show expected paths for debugging
	homeDir, _ := os.UserHomeDir()
	fmt.Println("\nExpected browser locations for your platform:")
	
	switch runtime.GOOS {
	case "darwin":
		fmt.Printf("  Chrome: %s\n", filepath.Join(homeDir, "Library", "Application Support", "Google", "Chrome"))
		fmt.Printf("  Brave:  %s\n", filepath.Join(homeDir, "Library", "Application Support", "BraveSoftware", "Brave-Browser"))
	case "windows":
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData == "" {
			localAppData = filepath.Join(homeDir, "AppData", "Local")
		}
		fmt.Printf("  Chrome: %s\n", filepath.Join(localAppData, "Google", "Chrome", "User Data"))
		fmt.Printf("  Brave:  %s\n", filepath.Join(localAppData, "BraveSoftware", "Brave-Browser", "User Data"))
	case "linux":
		fmt.Printf("  Chrome: %s\n", filepath.Join(homeDir, ".config", "google-chrome"))
		fmt.Printf("  Brave:  %s\n", filepath.Join(homeDir, ".config", "BraveSoftware", "Brave-Browser"))
	}
}

func discoverBrowserProfiles(browserDirectory string) ([]string, error) {
	entries, err := os.ReadDir(browserDirectory)
	if err != nil {
		return nil, fmt.Errorf("failed to read browser directory: %w", err)
	}

	profiles := []string{}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		if name == "Default" || strings.HasPrefix(name, "Profile ") {
			profilePath := filepath.Join(browserDirectory, name)
			profiles = append(profiles, profilePath)
		}
	}

	return profiles, nil
}

func buildMaliciousExtensionsMap() map[string]bool {
	maliciousMap := make(map[string]bool)
	for _, extensionID := range maliciousExtensions {
		maliciousMap[extensionID] = true
	}
	return maliciousMap
}

func scanExtensions(directory string, maliciousMap map[string]bool) ([]string, int, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read extensions directory: %w", err)
	}

	foundMalicious := []string{}
	installedCount := 0

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		installedCount++
		extensionID := entry.Name()

		if maliciousMap[extensionID] {
			foundMalicious = append(foundMalicious, extensionID)
		}
	}

	return foundMalicious, installedCount, nil
}

func printAllResults(results []profileScanResult) {
	totalExtensions := 0
	totalMalicious := 0
	browserCounts := make(map[string]int)

	for _, result := range results {
		totalExtensions += result.installedCount
		totalMalicious += len(result.foundMalicious)
		browserCounts[result.browserName]++
	}

	fmt.Println("Scan Summary:")
	for browser, count := range browserCounts {
		fmt.Printf("  • %s: %d profile(s)\n", browser, count)
	}
	fmt.Printf("\n📊 Total profiles scanned: %d\n", len(results))
	fmt.Printf("📦 Total extensions found: %d\n", totalExtensions)
	fmt.Printf("🛡️  Malicious extensions checked: %d\n\n", len(maliciousExtensions))

	if totalMalicious > 0 {
		printMaliciousResults(results)
		return
	}

	printCleanResult()
}

func printMaliciousResults(results []profileScanResult) {
	fmt.Println("⚠️  ALERT: MALICIOUS EXTENSIONS DETECTED!")
	fmt.Println(separator)

	for _, result := range results {
		if len(result.foundMalicious) == 0 {
			continue
		}

		fmt.Printf("\n🚨 Browser: %s | Profile: %s\n", result.browserName, result.profileName)
		fmt.Printf("Found %d malicious extension(s):\n\n", len(result.foundMalicious))

		for i, extensionID := range result.foundMalicious {
			fmt.Printf("  %d. %s", i+1, extensionID)
			if extensionID == "eagiakjmjnblliacokhcalebgnhellfi" {
				fmt.Printf(" (Clean Master)")
			}
			fmt.Println()

			extensionPath := filepath.Join(result.extensionsPath, extensionID)
			fmt.Printf("     Path: %s\n\n", extensionPath)
		}
	}

	printRemovalInstructions()
}

func printRemovalInstructions() {
	fmt.Println("⚡ RECOMMENDED ACTIONS:")
	fmt.Println("  1. Remove these extensions immediately from your browser")
	fmt.Println("  2. Go to chrome://extensions (Chrome) or brave://extensions (Brave)")
	fmt.Println("  3. Enable 'Developer mode' to see extension IDs")
	fmt.Println("  4. Remove any extensions matching the IDs above")
	fmt.Println("  5. Change your passwords across all accounts")
	fmt.Println("  6. Run a full antivirus scan")
	fmt.Println()
}

func printCleanResult() {
	fmt.Println("✅ GOOD NEWS: No malicious extensions detected!")
	fmt.Println("\nAll scanned browser profiles appear to be clean from the")
	fmt.Println("ShadyPanda malware campaign extensions.")
	fmt.Println()
}

func hasMaliciousExtensions(results []profileScanResult) bool {
	for _, result := range results {
		if len(result.foundMalicious) > 0 {
			return true
		}
	}
	return false
}
