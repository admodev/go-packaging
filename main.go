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

	var version string = "v1.0.0"

	appName := os.Args[1]

	fmt.Println("✨ Packaging version", version)
	fmt.Println("📦 Creating package oriented project structure...")

	// binary folders (cmd...)
	fmt.Println("⏱️ Creating binary folders...")

	binaryPath := fmt.Sprintf("./cmd/%s/", appName)
	mainFile := fmt.Sprintf("%s/main.go", binaryPath)

	os.MkdirAll(binaryPath, os.ModePerm)
	os.OpenFile(mainFile, os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Binary folders created successfully!")
}
