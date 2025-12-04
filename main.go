package main

import (
	"fmt"
	"os"
	"path/filepath"
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

func main() {
	printHeader()

	chromeExtensionsDirectory, err := getChromeExtensionsDirectory()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting Chrome extensions directory: %v\n", err)
		os.Exit(1)
	}

	if err := verifyDirectoryExists(chromeExtensionsDirectory); err != nil {
		printDirectoryNotFound(chromeExtensionsDirectory)
		os.Exit(0)
	}

	fmt.Printf("📂 Checking directory: %s\n\n", chromeExtensionsDirectory)

	maliciousMap := buildMaliciousExtensionsMap()

	foundMalicious, installedCount, err := scanExtensions(chromeExtensionsDirectory, maliciousMap)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning extensions: %v\n", err)
		os.Exit(1)
	}

	printScanSummary(installedCount)

	if len(foundMalicious) > 0 {
		printMaliciousExtensionsAlert(foundMalicious, chromeExtensionsDirectory)
		os.Exit(1)
	}

	printCleanResult()
}

func printHeader() {
	fmt.Println("🔍 Scanning Chrome extensions for ShadyPanda malware...")
	fmt.Println(separator)
}

func getChromeExtensionsDirectory() (string, error) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	return filepath.Join(homeDirectory, "Library", "Application Support", "Google", "Chrome", "Default", "Extensions"), nil
}

func verifyDirectoryExists(directory string) error {
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %w", err)
	}
	return err
}

func printDirectoryNotFound(directory string) {
	fmt.Printf("❌ Chrome extensions directory not found at:\n   %s\n", directory)
	fmt.Println("\nThis could mean:")
	fmt.Println("  • Chrome is not installed")
	fmt.Println("  • Chrome hasn't been run yet")
	fmt.Println("  • Extensions are in a different profile")
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

func printScanSummary(installedCount int) {
	fmt.Printf("📊 Total Chrome extensions found: %d\n", installedCount)
	fmt.Printf("🛡️  Malicious extensions checked: %d\n\n", len(maliciousExtensions))
}

func printMaliciousExtensionsAlert(foundMalicious []string, baseDirectory string) {
	fmt.Println("⚠️  ALERT: MALICIOUS EXTENSIONS DETECTED!")
	fmt.Println(separator)
	fmt.Printf("\n🚨 Found %d malicious extension(s):\n\n", len(foundMalicious))

	for i, extensionID := range foundMalicious {
		fmt.Printf("%d. %s", i+1, extensionID)
		if extensionID == "eagiakjmjnblliacokhcalebgnhellfi" {
			fmt.Printf(" (Clean Master)")
		}
		fmt.Println()

		extensionPath := filepath.Join(baseDirectory, extensionID)
		fmt.Printf("   Path: %s\n\n", extensionPath)
	}

	printRemovalInstructions()
}

func printRemovalInstructions() {
	fmt.Println("⚡ RECOMMENDED ACTIONS:")
	fmt.Println("  1. Remove these extensions immediately from Chrome")
	fmt.Println("  2. Go to chrome://extensions in your browser")
	fmt.Println("  3. Enable 'Developer mode' to see extension IDs")
	fmt.Println("  4. Remove any extensions matching the IDs above")
	fmt.Println("  5. Change your passwords across all accounts")
	fmt.Println("  6. Run a full antivirus scan")
	fmt.Println()
}

func printCleanResult() {
	fmt.Println("✅ GOOD NEWS: No malicious extensions detected!")
	fmt.Println("\nYour Chrome installation appears to be clean from the")
	fmt.Println("ShadyPanda malware campaign extensions.")
	fmt.Println()
}
