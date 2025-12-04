package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// List of known malicious Chrome extension IDs from the ShadyPanda campaign
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
	fmt.Println("🔍 Scanning Chrome extensions for ShadyPanda malware...")
	fmt.Println("=" + string(make([]byte, 60)) + "=")

	// Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting home directory: %v\n", err)
		os.Exit(1)
	}

	// Chrome extensions directory on Mac
	chromeExtensionsDir := filepath.Join(homeDir, "Library", "Application Support", "Google", "Chrome", "Default", "Extensions")

	// Check if Chrome extensions directory exists
	if _, err := os.Stat(chromeExtensionsDir); os.IsNotExist(err) {
		fmt.Printf("❌ Chrome extensions directory not found at:\n   %s\n", chromeExtensionsDir)
		fmt.Println("\nThis could mean:")
		fmt.Println("  • Chrome is not installed")
		fmt.Println("  • Chrome hasn't been run yet")
		fmt.Println("  • Extensions are in a different profile")
		os.Exit(0)
	}

	fmt.Printf("📂 Checking directory: %s\n\n", chromeExtensionsDir)

	// Create a map for faster lookup
	maliciousMap := make(map[string]bool)
	for _, id := range maliciousExtensions {
		maliciousMap[id] = true
	}

	// Scan the extensions directory
	entries, err := os.ReadDir(chromeExtensionsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading extensions directory: %v\n", err)
		os.Exit(1)
	}

	foundMalicious := []string{}
	installedCount := 0

	for _, entry := range entries {
		if entry.IsDir() {
			installedCount++
			extensionID := entry.Name()

			// Check if this extension ID is in our malicious list
			if maliciousMap[extensionID] {
				foundMalicious = append(foundMalicious, extensionID)
			}
		}
	}

	// Display results
	fmt.Printf("📊 Total Chrome extensions found: %d\n", installedCount)
	fmt.Printf("🛡️  Malicious extensions checked: %d\n\n", len(maliciousExtensions))

	if len(foundMalicious) > 0 {
		fmt.Println("⚠️  ALERT: MALICIOUS EXTENSIONS DETECTED!")
		fmt.Println("=" + string(make([]byte, 60)) + "=")
		fmt.Printf("\n🚨 Found %d malicious extension(s):\n\n", len(foundMalicious))

		for i, id := range foundMalicious {
			fmt.Printf("%d. %s", i+1, id)
			if id == "eagiakjmjnblliacokhcalebgnhellfi" {
				fmt.Printf(" (Clean Master)")
			}
			fmt.Println()

			// Show the full path
			extensionPath := filepath.Join(chromeExtensionsDir, id)
			fmt.Printf("   Path: %s\n\n", extensionPath)
		}

		fmt.Println("⚡ RECOMMENDED ACTIONS:")
		fmt.Println("  1. Remove these extensions immediately from Chrome")
		fmt.Println("  2. Go to chrome://extensions in your browser")
		fmt.Println("  3. Enable 'Developer mode' to see extension IDs")
		fmt.Println("  4. Remove any extensions matching the IDs above")
		fmt.Println("  5. Change your passwords across all accounts")
		fmt.Println("  6. Run a full antivirus scan")
		fmt.Println()

		os.Exit(1)
	} else {
		fmt.Println("✅ GOOD NEWS: No malicious extensions detected!")
		fmt.Println("\nYour Chrome installation appears to be clean from the")
		fmt.Println("ShadyPanda malware campaign extensions.")
	}

	fmt.Println()
}
