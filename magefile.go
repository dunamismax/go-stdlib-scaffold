//go:build mage

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	// goCmd is the go command.
	goCmd = "go"
)

// Default sets the default target for Mage.
var Default = Build.All

// --- Namespaces ---

type (
	// Build groups functions for building the application.
	Build mg.Namespace
	// Run groups functions for running the application.
	Run mg.Namespace
)

// --- Build Targets ---

// All builds all binaries and assets.
func (Build) All() {
	mg.SerialDeps(Build.CSS, Build.JS, Build.Server, Build.CLI)
}

// Server builds the main server binary.
func (Build) Server() error {
	fmt.Println("Building server binary...")
	return sh.Run(goCmd, "build", "-o", "bin/server", "./cmd/server")
}

// CLI builds the CLI binary.
func (Build) CLI() error {
	fmt.Println("Building CLI binary...")
	return sh.Run(goCmd, "build", "-o", "bin/cli", "./cmd/cli")
}

// CSS builds the Tailwind CSS assets.
func (Build) CSS() error {
	mg.Deps(installFrontendDeps) // Ensures dependencies are installed before building
	fmt.Println("Building CSS assets...")
	return sh.RunV("npm", "run", "build:css")
}

// JS builds the JavaScript assets.
func (Build) JS() error {
	mg.Deps(installFrontendDeps) // Ensures dependencies are installed before building
	fmt.Println("Building JS assets...")
	return sh.RunV("npm", "run", "build:js")
}

// --- Application Execution ---

// Dev builds and runs the server.
func Dev() error {
	mg.Deps(Build.All)
	fmt.Println("Starting server...")
	return sh.RunV("./bin/server")
}

// Server runs the main server binary.
func (Run) Server() error {
	fmt.Println("Running server...")
	return sh.RunV("./bin/server")
}

// CLI runs the CLI binary with the given arguments.
func (Run) CLI(args string) error {
	fmt.Printf("Running CLI with args: %s\n", args)
	return sh.RunV("./bin/cli", strings.Fields(args)...)
}

// --- Quality & Verification ---

// Check runs all checks: linting, testing, and vulnerability scanning.
func Check() {
	mg.SerialDeps(Lint, LintFrontend, Test, Vuln)
}

// Test runs all Go tests.
func Test() error {
	fmt.Println("Running tests...")
	return sh.RunV(goCmd, "test", "-v", "./...")
}

// Lint runs the golangci-lint linter for Go files.
func Lint() error {
	fmt.Println("Linting Go files...")
	return sh.RunV(goCmd, "run", "github.com/golangci/golangci-lint/cmd/golangci-lint", "run")
}

// LintFrontend runs the prettier linter for frontend files.
func LintFrontend() error {
	fmt.Println("Linting frontend files...")
	return sh.RunV("npm", "run", "lint")
}

// Vuln runs govulncheck to scan for vulnerabilities.
func Vuln() error {
	fmt.Println("Checking for vulnerabilities...")
	return sh.RunV(goCmd, "run", "golang.org/x/vuln/cmd/govulncheck", "./...")
}

// --- Asset & Module Management ---

// Format runs prettier to format frontend files.
func Format() error {
	fmt.Println("Formatting frontend files...")
	return sh.RunV("npm", "run", "format")
}

// Tidy tidies Go module dependencies.
func Tidy() error {
	fmt.Println("Tidying Go modules...")
	return sh.RunV(goCmd, "mod", "tidy")
}

// --- Cleanup & Installation ---

// installFrontendDeps checks for the existence of node_modules and runs 'npm install' if it's missing.
func installFrontendDeps() error {
	fmt.Println("Checking for node_modules...")
	if _, err := os.Stat("node_modules"); os.IsNotExist(err) {
		fmt.Println("node_modules not found, running 'npm install'...")
		return sh.RunV("npm", "install")
	}
	fmt.Println("node_modules already exists, skipping installation.")
	return nil
}

// Clean removes all build artifacts and temporary files.
func Clean() {
	fmt.Println("Cleaning up...")
	os.RemoveAll("bin")
	os.RemoveAll("public/assets")
	os.RemoveAll("tmp")
}

// InstallTools installs all necessary Go tools.
func InstallTools() error {
	fmt.Println("Installing Go tools...")
	if err := sh.Run(goCmd, "install", "github.com/magefile/mage"); err != nil {
		return err
	}
	if err := sh.Run(goCmd, "install", "github.com/golangci/golangci-lint/cmd/golangci-lint"); err != nil {
		return err
	}
	return sh.Run(goCmd, "install", "golang.org/x/vuln/cmd/govulncheck")
}