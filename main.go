package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("❌ FATAL ERROR ‼️ Please, pass your app name as an argument to the script...")
		os.Exit(1)
	}

	appName := os.Args[1]

	fmt.Println("📦 Creating package oriented project structure...")

	// binary folders (cmd...)
	fmt.Println("⏱️ Creating binary folders...")

	binaryPath := fmt.Sprintf("./cmd/%s/", appName)

	os.MkdirAll(binaryPath, os.ModePerm)
	fmt.Println("✔️ Binary folders created successfully!")
}
