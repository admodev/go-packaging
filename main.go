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
	entryPoint := fmt.Sprintf("%s/main.go", binaryPath)

	os.MkdirAll(binaryPath, os.ModePerm)
	os.OpenFile(entryPoint, os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Binary folders created successfully!")

	// Functional domain packages
	fmt.Println("⏱️ Creating internal folders...")
	os.MkdirAll("./internal/rename_me/", os.ModePerm)
	os.MkdirAll("./internal/config/", os.ModePerm)
	os.OpenFile("./internal/rename_me/example_file.go", os.O_RDONLY|os.O_CREATE, 0666)
	os.OpenFile("./internal/rename_me/example_file_test.go", os.O_RDONLY|os.O_CREATE, 0666)
	os.OpenFile("./internal/config/config.go", os.O_RDONLY|os.O_CREATE, 0666)
	os.OpenFile("./internal/config/config_test.go", os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Internal folders created successfully!")

	// Reusable exported code
	fmt.Println("⏱️ Creating pkg folders...")
	os.MkdirAll("./pkg/example_package/", os.ModePerm)
	os.OpenFile("./pkg/example_package/example_package_file.go", os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Pkg folders created successfully!")
}
