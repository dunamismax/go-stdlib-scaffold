//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default sets the default target for Mage.
var Default = Run

// -----------------------------------------------------------------------------
// Build Targets
// -----------------------------------------------------------------------------

type Build mg.Namespace

// All builds all application binaries.
func (Build) All() {
	mg.Deps(Build.Server, Build.CLI)
}

// Server builds the main web server binary.
func (Build) Server() error {
	fmt.Println("Building server...")
	return sh.Run("go", "build", "-o", "bin/server", "./cmd/server")
}

// CLI builds the command-line interface binary.
func (Build) CLI() error {
	fmt.Println("Building CLI...")
	return sh.Run("go", "build", "-o", "bin/cli", "./cmd/cli")
}

// -----------------------------------------------------------------------------
// Development & Execution Targets
// -----------------------------------------------------------------------------

// Run builds and runs the web server.
func Run() error {
	mg.Deps(Build.Server)
	fmt.Println("Starting server on http://localhost:3000 ...")
	return sh.Run("./bin/server")
}

// -----------------------------------------------------------------------------
// Database Targets
// -----------------------------------------------------------------------------

type DB mg.Namespace

// Migrate runs database migrations.
func (DB) Migrate() error {
	mg.Deps(Build.CLI)
	fmt.Println("Running database migrations...")
	return sh.Run("./bin/cli", "migrate")
}

// -----------------------------------------------------------------------------
// Quality & CI/CD Targets
// -----------------------------------------------------------------------------

type Check mg.Namespace

// All runs all quality checks.
func (Check) All() {
	mg.Deps(Check.Format, Check.Vet, Check.Test, Check.Vuln)
}

// Format checks if the code is formatted correctly.
func (Check) Format() error {
	fmt.Println("Checking format...")
	out, err := sh.Output("gofmt", "-l", ".")
	if err != nil {
		return err
	}
	if out != "" {
		return fmt.Errorf("code is not formatted, run 'mage format'")
	}
	return nil
}

// Vet runs go vet to find common issues.
func (Check) Vet() error {
	fmt.Println("Running go vet...")
	return sh.Run("go", "vet", "./...")
}

// Test runs all unit tests.
func (Check) Test() error {
	fmt.Println("Running tests...")
	return sh.Run("go", "test", "-v", "-json", "./...")
}

// Vuln scans for vulnerabilities.
func (Check) Vuln() error {
	fmt.Println("Scanning for vulnerabilities...")
	return sh.Run("go", "run", "golang.org/x/vuln/cmd/govulncheck@latest", "./...")
}

// -----------------------------------------------------------------------------
// Housekeeping Targets
// -----------------------------------------------------------------------------

// Format formats the Go source code.
func Format() error {
	fmt.Println("Formatting code...")
	return sh.Run("gofmt", "-w", ".")
}

// Tidy tidies the go.mod file.
func Tidy() error {
	fmt.Println("Tidying go.mod...")
	return sh.Run("go", "mod", "tidy")
}

// Clean removes all build artifacts and the local database.
func Clean() {
	fmt.Println("Cleaning up...")
	os.RemoveAll("bin")
	os.RemoveAll("data")
}

// Release cross-compiles binaries for release.
func Release() error {
	mg.Deps(Tidy)
	fmt.Println("Cross-compiling release binaries...")
	platforms := map[string][]string{
		"linux":   {"amd64", "arm64"},
		"windows": {"amd64"},
		"darwin":  {"amd64", "arm64"},
	}

	for goos, archs := range platforms {
		for _, goarch := range archs {
			name := fmt.Sprintf("bin/server-%s-%s", goos, goarch)
			if goos == "windows" {
				name += ".exe"
			}
			env := map[string]string{"GOOS": goos, "GOARCH": goarch}
			fmt.Printf("  - Building for %s/%s\n", goos, goarch)
			if err := sh.RunWith(env, "go", "build", "-ldflags=-s -w", "-o", name, "./cmd/server"); err != nil {
				return err
			}
		}
	}
	return nil
}
