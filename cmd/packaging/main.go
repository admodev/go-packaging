package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func helpMessage(softwareVersion string) {
	fmt.Printf("📦 Packaging version %s\n", softwareVersion)
	fmt.Println("To create a new project with a package oriented structure run the")
	fmt.Println("packaging command with your organization and project name")
	fmt.Println("as arguments, for example: packaging myorg newproject")
}

func main() {
	var version string = "v1.0.0"

	if len(os.Args) == 2 && os.Args[1] == "--help" {
		helpMessage(version)
		return
	}

	if len(os.Args) < 3 {
		log.Fatalf("❌ FATAL ERROR ‼️ Please, pass your organization name and app name as arguments.\n\nUsage:\n\tpackaging <org_name> <app_name>")
	}

	orgName := os.Args[1]
	appName := os.Args[2]

	fmt.Println("✨ Packaging version", version)
	fmt.Println("📦 Creating package oriented project structure...")

	// binary folders (cmd...)
	fmt.Println("⏱️ Creating binary folders...")

	binaryPath := fmt.Sprintf("./cmd/%s/", appName)
	entryPoint := fmt.Sprintf("%s/main.go", binaryPath)

	os.MkdirAll(binaryPath, os.ModePerm)
	os.OpenFile(entryPoint, os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Binary folders and files created successfully!")

	// Functional domain packages
	fmt.Println("⏱️ Creating internal folders...")
	os.MkdirAll("./internal/rename_me/", os.ModePerm)
	os.MkdirAll("./internal/config/", os.ModePerm)
	os.OpenFile("./internal/rename_me/example_file.go", os.O_RDONLY|os.O_CREATE, 0666)
	os.OpenFile("./internal/rename_me/example_file_test.go", os.O_RDONLY|os.O_CREATE, 0666)
	os.OpenFile("./internal/config/config.go", os.O_RDONLY|os.O_CREATE, 0666)
	os.OpenFile("./internal/config/config_test.go", os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Internal folders and files created successfully!")

	// Reusable exported code
	fmt.Println("⏱️ Creating pkg folders...")
	os.MkdirAll("./pkg/example_package/", os.ModePerm)
	os.OpenFile("./pkg/example_package/example_package_file.go", os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Pkg folders and files created successfully!")

	// HTTP/GRPC handlers and middlewares
	fmt.Println("⏱️ Creating api folders...")
	os.MkdirAll("./api/handler/", os.ModePerm)
	os.OpenFile("./api/handler/health_check.go", os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Api folders and files created successfully!")

	// Dev tools for building, linting and testing the project
	fmt.Println("⏱️ Creating scripts folders...")
	os.MkdirAll("./scripts/", os.ModePerm)
	os.OpenFile("./scripts/build.sh", os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Scripts folders and files created successfully!")

	// Environment scoped configurations
	fmt.Println("⏱️ Creating configs folders...")
	os.MkdirAll("./configs/", os.ModePerm)
	os.OpenFile("./configs/dev.yaml", os.O_RDONLY|os.O_CREATE, 0666)
	os.OpenFile("./configs/qa.yaml", os.O_RDONLY|os.O_CREATE, 0666)
	os.OpenFile("./configs/pre.yaml", os.O_RDONLY|os.O_CREATE, 0666)
	os.OpenFile("./configs/prod.yaml", os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Configs folders and files created successfully!")

	// Ideally for running the project in containers or kubernetes
	fmt.Println("⏱️ Creating deployments folders...")
	os.MkdirAll("./deployments/k8s/", os.ModePerm)
	os.OpenFile("./deployments/docker-compose.yaml", os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Deployments folders and files created successfully!")

	// Test suite
	fmt.Println("⏱️ Creating test folders...")
	os.MkdirAll("./test/", os.ModePerm)
	os.OpenFile("./test/integration_test.go", os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Test folders and files created successfully!")

	// Makefile for standardized cross-platform builds
	fmt.Println("⏱️ Creating Makefile...")
	os.OpenFile("./Makefile", os.O_RDONLY|os.O_CREATE, 0666)
	fmt.Println("✔️ Makefile created successfully!")

	// Create a go module file
	fmt.Println("⏱️ Creating go.mod file...")
	
	goModuleName := fmt.Sprintf("%s/%s", orgName, appName)
	goModuleInitCommand := exec.Command("go", "mod", "init", goModuleName)

	output, err := goModuleInitCommand.CombinedOutput()

	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	fmt.Printf("Go module initialized:\n%s\n", output)
	fmt.Println("✔️ Project structure generated successfully!")
}
