<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1920px-Go_Logo_Blue.svg.png" alt="The Go programming language logo." width="150"/>
</p>

<p align="center">
  <a href="https://github.com/dunamismax/go-stdlib-scaffold">
    <img src="https://readme-typing-svg.demolab.com/?font=Fira+Code&size=24&pause=1000&color=00ADD8&center=true&vCenter=true&width=800&lines=The+Pure+Go+Standard+Library+Stack;Official+Reference+Implementation;Go+%2B+net/http+%2B+database/sql;Mage+and+SQLite;Simple%2C+Performant%2C+and+Maintainable." alt="Typing SVG" />
  </a>
</p>

<p align="center">
  <a href="https://github.com/dunamismax/go-stdlib-scaffold/actions/workflows/ci.yml"><img src="https://github.com/dunamismax/go-stdlib-scaffold/actions/workflows/ci.yml/badge.svg" alt="Build Status"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.24+-00ADD8.svg" alt="Go Version"></a>
  <a href="https://img.shields.io/github/license/dunamismax/go-stdlib-scaffold"><img src="https://img.shields.io/github/license/dunamismax/go-stdlib-scaffold" alt="License"></a>
  <a href="https://img.shields.io/github/repo-size/dunamismax/go-stdlib-scaffold"><img src="https://img.shields.io/github/repo-size/dunamismax/go-stdlib-scaffold" alt="Repo Size"></a>
  <a href="https://github.com/dunamismax/go-stdlib-scaffold/pulls"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs Welcome"></a>
  <a href="https://github.com/dunamismax/go-stdlib-scaffold/stargazers"><img src="https://img.shields.io/github/stars/dunamismax/go-stdlib-scaffold" alt="GitHub Stars"></a>
</p>

---

## About This Project

This monorepo is the official reference implementation for **The Pure Go Standard Library Stack**, a complete architecture for building secure, maintainable, and dependency-free web applications.

The primary goal is to provide a structured and scalable environment that relies exclusively on Go's standard library. The project's philosophy revolves around shipping a **single, self-contained executable binary** that can be easily deployed on any platform Go supports. This approach emphasizes performance, ultimate simplicity, and long-term stability by avoiding all external dependencies for core application logic.

---

<details>
<summary><h3>The Pure Go Standard Library Stack (Click to Expand)</h3></summary>

This stack represents a minimalist, robust architecture for building secure and maintainable web applications. It is composed entirely of a Go backend that leverages the standard library, removing all external dependencies for routing, validation, and data access. The stack prioritizes ultimate simplicity, zero-dependency deployment, and long-term stability by relying exclusively on Go's native capabilities. The frontend is reduced to plain HTML and CSS, with no JavaScript.

---

### **Frontend**

The frontend architecture is intentionally simplified to its core components, delivering a fast, accessible, and extremely maintainable user experience by avoiding all client-side scripting and build tools.

- [**Go `html/template`**](https://pkg.go.dev/html/template)
    - **Role:** Secure HTML Templating.
    - **Description:** Go's standard library for server-side HTML rendering. It provides automatic, context-aware escaping to prevent Cross-Site Scripting (XSS), making it a secure foundation for web interfaces.
- **Plain CSS**
    - **Role:** Styling.
    - **Description:** A standard, handwritten CSS file served as a static asset. This approach removes the need for pre-processors or build steps, maximizing simplicity and performance.

---

### **Backend**

A lean, performant, and maintainable backend service architected using only the Go standard library for maximum stability and minimal external dependencies.

- [**Go**](https://go.dev/doc/)
    - **Role:** Backend Language.
    - **Description:** A statically typed, compiled language known for performance, concurrency, and simplicity. It compiles to a single, dependency-free binary, streamlining deployment.
- [**`net/http`**](https://pkg.go.dev/net/http)
    - **Role:** Web Server & Routing.
    - **Description:** The standard library's package for all HTTP-related tasks. As of Go 1.22, it includes an enhanced request router that supports method-based routing and wildcards, removing the need for third-party frameworks.
- **Custom Validation Functions**
    - **Role:** Data Validation.
    - **Description:** Data validation is handled by simple, explicit Go functions. This approach keeps validation logic clear, type-safe, and tightly integrated with the application's domain.
- [**`os`**](https://pkg.go.dev/os)
    - **Role:** Environment Variable Loading.
    - **Description:** Configuration is loaded from environment variables using the standard `os.Getenv` function, adhering to twelve-factor app principles without external libraries.

---

### **Database & Caching**

A zero-dependency, in-process data layer that maximizes simplicity and speed by using Go's native database interface and basic concurrency primitives.

- [**SQLite**](https://www.sqlite.org/docs.html)
    - **Role:** Embedded Relational Database.
    - **Description:** A self-contained, serverless SQL database engine that runs in-process, eliminating operational overhead and making it ideal for a wide range of production workloads.
- [**`database/sql`**](https://pkg.go.dev/database/sql)
    - **Role:** SQL Database Interface.
    - **Description:** The standard library’s lean interface for SQL databases. It provides direct control over database operations via raw SQL queries, avoiding the overhead of an ORM.
- **SQL/Go Migration Scripts**
    - **Role:** Database Schema Migrations.
    - **Description:** Migrations are managed with numbered SQL files or simple Go programs using `database/sql`. This method avoids external tooling and keeps schema management transparent.
- [**`sync`**](https://pkg.go.dev/sync)
    - **Role:** In-Process Caching.
    - **Description:** High-performance, in-process caching is achieved with a standard Go map and a `sync.RWMutex`, providing a fast, concurrent-safe solution without external dependencies.

---

### **Testing**

A robust testing suite that relies exclusively on Go's powerful, built-in testing framework to ensure code quality and correctness.

- [**`testing`**](https://pkg.go.dev/testing)
    - **Role:** Core Testing Framework.
    - **Description:** The built-in package for unit, integration, and benchmark tests. Assertions use simple `if` statements with `t.Errorf`, keeping tests clear and dependency-free.

---

### **CLI, Development & Deployment**

A minimalist toolchain using built-in Go commands and standard, universally available tools for a smooth developer workflow.

- [**`flag`**](https://pkg.go.dev/flag)
    - **Role:** Command-Line Interface.
    - **Description:** The standard library's package for parsing command-line options. It is sufficient for building CLIs for most applications without third-party dependencies.
- [**Mage / Magefile**](https://magefile.org/)
    - **Role:** Task Runner / Build System.
    - **Description:** A build tool that uses plain Go functions as runnable, Makefile-like targets. By replacing shell scripts with Go code, a `Magefile` creates a cross-platform, easy-to-maintain build system for development tasks.
- **Simple Shell Scripts**
    - **Role:** Live Reloading.
    - **Description:** During development, a simple shell script can watch for file changes, automatically recompiling and restarting the server for a rapid feedback loop.

---

### **CI/CD**

A fully automated pipeline using standard Go tools to build, test, and deploy the application, ensuring consistency and quality.

- [**GitHub Actions**](https://docs.github.com/en/actions)
    - **Role:** Automated CI/CD Platform.
    - **Description:** A CI/CD workflow defined in the project repository automates the entire lifecycle. The pipeline performs:
        - **Linting & Formatting:** Runs [**`gofmt`**](https://pkg.go.dev/cmd/gofmt) and [**`go vet`**](https://pkg.go.dev/cmd/vet) to enforce code style and identify issues.
        - **Testing:** Executes the test suite with the standard `go test` command.
        - **Vulnerability Scanning:** Runs [**`govulncheck`**](https://go.dev/blog/vuln) to scan for security vulnerabilities.
        - **Build:** Compiles the application into a single binary using `go build`.

</details>

---

<p align="center">
  <img src="https://user-images.githubusercontent.com/3185864/32058716-5ee9b512-ba38-11e7-978a-287eb2a62743.png" alt="Gopher Mage." width="150"/>
</p>

## Getting Started

### Prerequisites

- **Go 1.24+**
- **Mage**

### Installation & Usage

1. **Clone the repository:**

   ```bash
   git clone https://github.com/dunamismax/go-stdlib-scaffold.git
   cd go-stdlib-scaffold
   ```

2. **Install Dependencies & Tools:**
   This one-time command installs the Go modules and the Mage task runner.

   ```bash
   go install github.com/magefile/mage@latest && go mod tidy
   ```

3. **Format & Check Code (Important!)**
   Before committing, always format your code and run all checks to ensure it meets project standards and will pass CI.

   ```bash
   mage format
   mage check:all
   ```

4. **Configure Your Environment:**
   Copy the example environment file and customize it as needed for your database and server settings.

   ```bash
   cp .env.example .env
   ```

5. **Run the Application:**
   This project uses [Mage](https://magefile.org/) as a task runner to simplify development workflows. See all available commands by running `mage -l`.
   
   - **Run Migrations:**
     This command builds the CLI and runs the SQL migrations to set up the database schema.

     ```bash
     mage db:migrate
     ```

   - **Run the Server:**
     This command builds and runs the main web server.

     ```bash
     mage run
     ```

     The application will be available at `http://localhost:3000`.

---

## Mage Commands

[Mage](https://magefile.org/) is used to automate common development tasks. The build script is the `magefile.go` at the root of the project. You can list all available commands by running `mage -l`.

### Primary Workflow Commands

- `mage run`: (Default) Builds and runs the main web server. This is the primary command for local development.
- `mage check:all`: **(CI)** Runs all quality checks: format, vet, test, and vulnerability scan. Run this before committing.
- `mage format`: Automatically formats all Go code using `gofmt`.
- `mage build:all`: Builds all artifacts: the web server binary and the CLI binary.
- `mage db:migrate`: Applies all pending database migrations.

### Individual Commands

- **Building (`build:`)**
  - `mage build:server`: Builds only the main web server binary.
  - `mage build:cli`: Builds only the command-line interface binary.
- **Quality Checks (`check:`)**
  - `mage check:format`: Checks if code is formatted, fails if not.
  - `mage check:vet`: Runs `go vet`.
  - `mage check:test`: Runs all Go tests.
  - `mage check:vuln`: Scans for known vulnerabilities.
- **Housekeeping**
  - `mage clean`: Removes all build artifacts and the local database.
  - `mage tidy`: Tidies the `go.mod` and `go.sum` files.
  - `mage release`: Cross-compiles release binaries for multiple platforms.

---

## Project Structure

- **`cmd/`**: Application entry points.
  - **`server/`**: The main `net/http` web server.
    - **`templates/`**: Go `html/template` files.
  - **`cli/`**: The command-line application for migrations.
- **`internal/`**: Private application code.
  - **`database/`**: Database connection logic, data store, and models.
    - **`migrations/`**: Plain SQL schema migrations.
- **`public/`**: Compiled, publicly-served static assets (CSS).
- **`magefile.go`**: The build script for the project, written in Go.
- **`.github/workflows/ci.yml`**: The GitHub Actions CI/CD pipeline.

---

## Contributing

Contributions are welcome! Please feel free to fork the repository, create a feature branch, and open a pull request.

---

### Support My Work

If you find my work on this stack valuable, consider supporting me. It helps me dedicate more time to creating and maintaining open-source projects.

<p align="center">
  <a href="https://coff.ee/dunamismax" target="_blank">
    <img src="https://raw.githubusercontent.com/egonelbre/gophers/master/.thumb/animation/buy-morning-coffee-3x.gif" alt="Buy Me a Coffee" />
  </a>
</p>

---

### Let's Connect

<p align="center">
  <a href="https://twitter.com/dunamismax" target="_blank"><img src="https://img.shields.io/badge/Twitter-%231DA1F2.svg?&style=for-the-badge&logo=twitter&logoColor=white" alt="Twitter"></a>
  <a href="https://bsky.app/profile/dunamismax.bsky.social" target="_blank"><img src="https://img.shields.io/badge/Bluesky-blue?style=for-the-badge&logo=bluesky&logoColor=white" alt="Bluesky"></a>
  <a href="https://reddit.com/user/dunamismax" target="_blank"><img src="https://img.shields.io/badge/Reddit-%23FF4500.svg?&style=for-the-badge&logo=reddit&logoColor=white" alt="Reddit"></a>
  <a href="https://discord.com/users/dunamismax" target="_blank"><img src="https://img.shields.io/badge/Discord-dunamismax-7289DA.svg?style=for-the-badge&logo=discord&logoColor=white" alt="Discord"></a>
  <a href="https://signal.me/#p/+dunamismax.66" target="_blank"><img src="https://img.shields.io/badge/Signal-dunamismax.66-3A76F0.svg?style=for-the-badge&logo=signal&logoColor=white" alt="Signal"></a>
</p>

---

<p align="center">
    <img src="https://raw.githubusercontent.com/egonelbre/gophers/refs/heads/master/.thumb/animation/2bit-sprite/demo.gif" alt="Gopher Sprite Animation" />
</p>

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
