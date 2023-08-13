package main

import (
	"fmt"
	"os"
	"os/exec"
)

func installPackage(packageName string) {
	cmd := exec.Command("apt", "install", packageName, "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error installing package:", err)
	}
}

func removePackage(packageName string) {
	cmd := exec.Command("apt", "remove", packageName, "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error removing package:", err)
	}
}

func main() {
	packageName := "example-package"

	installPackage(packageName)
	fmt.Println("Package installed.")

	removePackage(packageName)
	fmt.Println("Package removed.")
}
