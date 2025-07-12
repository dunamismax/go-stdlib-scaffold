//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
)

// Default target to run when none is specified
var Default = Build

// Build compiles the server and CLI binaries.
func Build() error {
	mg.Deps(BuildServer, BuildCLI)
	return nil
}

// BuildServer compiles the main server application.
func BuildServer() error {
	fmt.Println("Building server...")
	return run("go", "build", "-o", "bin/server", "./cmd/server")
}

// BuildCLI compiles the command-line interface.
func BuildCLI() error {
	fmt.Println("Building CLI...")
	return run("go", "build", "-o", "bin/cli", "./cmd/cli")
}

// Run starts the web server. It builds the server first.
func Run() error {
	mg.Deps(BuildServer)
	fmt.Println("Starting server...")
	return run("./bin/server")
}

// Migrate runs the database migrations. It builds the CLI first.
func Migrate() error {
	mg.Deps(BuildCLI)
	fmt.Println("Running database migrations...")
	return run("./bin/cli", "migrate")
}

// Tidy cleans up the go.mod file.
func Tidy() error {
	fmt.Println("Tidying go.mod...")
	return run("go", "mod", "tidy")
}

// Clean removes build artifacts.
func Clean() {
	fmt.Println("Cleaning up...")
	os.RemoveAll("bin")
}

// run executes a command and prints its output.
func run(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
