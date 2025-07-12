<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1920px-Go_Logo_Blue.svg.png" alt="The Go programming language logo." width="150"/>
</p>

<p align="center">
  <a href="https://github.com/dunamismax/go-stdlib-scaffold">
    <img src="https://readme-typing-svg.demolab.com/?font=Fira+Code&size=24&pause=1000&color=00ADD8&center=true&vCenter=true&width=800&lines=The+go-stdlib-scaffold+Stack;Official+Reference+Implementation;Go+Standard+Library+Only;Mage%2C+Caddy%2C+and+SQLite;Simple%2C+Performant%2C+and+Maintainable." alt="Typing SVG" />
  </a>
</p>

<p align="center">
  <a href="https://github.com/dunamismax/go-stdlib-scaffold/actions/workflows/ci.yml"><img src="https://github.com/dunamismax/go-stdlib-scaffold/actions/workflows/ci.yml/badge.svg" alt="Build Status"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.22+-00ADD8.svg" alt="Go Version"></a>
  <a href="https://img.shields.io/github/license/dunamismax/go-stdlib-scaffold"><img src="https://img.shields.io/github/license/dunamismax/go-stdlib-scaffold" alt="License"></a>
  <a href="https://img.shields.io/github/repo-size/dunamismax/go-stdlib-scaffold"><img src="https://img.shields.io/github/repo-size/dunamismax/go-stdlib-scaffold" alt="Repo Size"></a>
  <a href="https://github.com/dunamismax/go-stdlib-scaffold/pulls"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs Welcome"></a>
  <a href="https://github.com/dunamismax/go-stdlib-scaffold/stargazers"><img src="https://img.shields.io/github/stars/dunamismax/go-stdlib-scaffold" alt="GitHub Stars"></a>
</p>

---

## About This Project

This monorepo is the official reference implementation for **The go-stdlib-scaffold Stack**, a complete architecture for building secure, observable, and maintainable web applications using only the Go standard library. It is composed of a powerful Go backend and a modern, server-centric frontend toolchain.

The primary goal is to provide a structured and scalable environment for building modern web applications and tools. The project's philosophy revolves around shipping **single, self-contained executable binaries** that can be easily deployed on lean environments like Alpine Linux, typically behind a Caddy reverse proxy. This approach emphasizes performance, simplicity, and a minimal resource footprint.

---

<details>
<summary><h3>The go-stdlib-scaffold Stack (Click to Expand)</h3></summary>

This stack represents a complete, best-in-class architecture for building secure, observable, and maintainable web applications using only the Go standard library. It is composed of a powerful Go backend and a modern, server-centric frontend toolchain. The stack prioritizes simplicity, rapid development, and zero-dependency deployment by favoring Go's native capabilities and lightweight, embedded tools.

---

### **Frontend**

The frontend architecture uses a modern build system and a server-centric interactivity model to deliver a fast, responsive, and maintainable user experience with minimal client-side complexity.

- [**esbuild**](https://esbuild.github.io/getting-started/)
  - **Role:** Asset Bundler & Minifier.
  - **Description:** An extremely fast JavaScript and CSS bundler written in Go. It processes frontend assets, handles module bundling, and performs minification, ensuring a highly optimized production output while maintaining a rapid development feedback loop.
- [**PostCSS**](https://postcss.org/docs/)
  - **Role:** CSS Processor.
  - **Description:** A tool for transforming CSS with JavaScript plugins. It is essential for a build step that compiles utility classes and custom directives into a standard, browser-ready stylesheet.
- [**Tailwind CSS**](https://tailwindcss.com/docs/installation/using-vite)
  - **Role:** Utility-First CSS Framework.
  - **Description:** A highly-customizable, utility-first CSS framework that enables rapid UI development directly within the HTML markup. It promotes design consistency and produces a minimal CSS file for production.
- [**HTMX**](https://htmx.org/docs/)
  - **Role:** Server-Centric Interactivity.
  - **Description:** A powerful library that enables modern browser features like AJAX and dynamic content updates directly from HTML attributes. It allows the backend to deliver UI fragments over the wire, providing rich user experiences without complex client-side JavaScript.
- [**Go `html/template`**](https://pkg.go.dev/html/template)
  - **Role:** Secure HTML Templating.
  - **Description:** The official Go standard library for creating HTML templates. It provides secure, context-aware automatic escaping to prevent Cross-Site Scripting (XSS) attacks, making it a robust and idiomatic choice for server-side rendering of HTML pages and HTMX partials.
- [**Alpine.js**](https://alpinejs.dev/start-here)
  - **Role:** Lightweight Client-Side Interactivity.
  - **Description:** A rugged, minimal framework for composing JavaScript behavior directly in your HTML markup. It serves as the perfect lightweight companion to HTMX for handling small client-side interactions like dropdowns, modals, and toggles, without requiring a heavy client-side framework.

---

### **Backend**

A lean, performant, and maintainable backend service architected for rapid development and long-term stability, using only the Go standard library.

- [**Go**](https://go.dev/doc/)
  - **Role:** Backend Language.
  - **Description:** A statically typed, compiled language renowned for its performance, concurrency, and simplicity. Its ability to compile to a single binary simplifies deployment.
- [**`net/http`**](https://pkg.go.dev/net/http)
  - **Role:** Web Server & Routing.
  - **Description:** The standard library's HTTP package provides a robust and flexible foundation for building web services. It handles routing, middleware, and all aspects of the HTTP protocol.

---

### **Database & Caching**

A zero-dependency, in-process data layer that maximizes simplicity and speed for a wide range of applications.

- [**SQLite**](https://www.sqlite.org/docs.html)
  - **Role:** Embedded Relational Database.
  - **Description:** A self-contained, serverless, full-featured SQL database engine that runs in-process with the application. It reads and writes to a single file, eliminating operational overhead and making it perfect for local development, testing, and many production workloads.
- [**`database/sql`**](https://pkg.go.dev/database/sql)
  - **Role:** SQL Interface.
  - **Description:** The standard library's SQL package provides a generic interface around SQL (or SQL-like) databases. It allows for writing clean, maintainable, and provider-agnostic data access code.
- **Manual Migrations**
  - **Role:** Database Schema Migrations.
  - **Description:** Database migrations are handled manually by writing SQL scripts. This approach provides maximum control and avoids external dependencies.

---

### **Testing**

A robust testing suite to ensure code quality, correctness, and maintainability.

- [**`go test`**](https://pkg.go.dev/testing)
  - **Role:** Core Testing Framework.
  - **Description:** The built-in Go testing command and package. It provides the foundation for writing unit, integration, and benchmark tests in a way that is simple and deeply integrated with the language.

---

### **CLI, Development & Deployment**

A professional and minimalist toolchain for a smooth developer workflow and consistent builds.

- [**`flag`**](https://pkg.go.dev/flag)
  - **Role:** Command-line Flag Parsing.
  - **Description:** The standard library's `flag` package provides a simple way to parse command-line arguments.
- [**Mage**](https://magefile.org/)
  - **Role:** Go-Native Task Runner / Build System.
  - **Description:** A build tool that allows you to write build scripts and tasks in plain Go, providing a type-safe, cross-platform, and idiomatic way to orchestrate all development workflows without leaving the Go ecosystem.
- [**Caddy**](https://caddyserver.com/docs/)
  - **Role:** Web Server & Reverse Proxy.
  - **Description:** A modern web server with automatic HTTPS. It serves static frontend assets and acts as a secure reverse proxy for the Go application.

---

### **CI/CD**

A fully automated pipeline for building, testing, and deploying the application, ensuring consistency and quality.

- [**GitHub Actions**](https://docs.github.com/en/actions)
  - **Role:** Automated CI/CD Platform.
  - **Description:** A CI/CD workflow defined in the project repository to automate the entire lifecycle. The pipeline performs:
    - **Linting & Formatting:** Runs `golangci-lint` and `gofmt` to enforce code quality.
    - **Testing:** Executes the test suite using `go test`.
    - **Vulnerability Scanning:** Runs `govulncheck` to scan for security vulnerabilities.
    - **Build:** Compiles the application and builds frontend assets using a `Mage` task.
- [**GoReleaser**](https://goreleaser.com/customization/)
  - **Role:** Release Automation.
  - **Description:** A powerful tool that automates the entire release process. It seamlessly integrates with GitHub Actions to cross-compile Go binaries, create archives, generate changelogs, and publish releases, simplifying the delivery of software.

</details>

---

<p align="center">
  <img src="https://user-images.githubusercontent.com/3185864/32058716-5ee9b512-ba38-11e7-978a-287eb2a62743.png" alt="Gopher Mage." width="150"/>
</p>

## Getting Started

### Prerequisites

- **Go 1.22+**
- **Node.js 20+** & **npm**

### Installation & Usage

1. **Clone the repository:**

   ```bash
   git clone https://github.com/dunamismax/go-stdlib-scaffold.git
   cd go-stdlib-scaffold
   ```

2. **Install Dependencies & Tools:**
   This one-time command installs all Go modules, Node.js packages, and required Go-based developer tools like `mage`.

   ```bash
   go install github.com/magefile/mage && mage installTools && npm install
   ```

3. **Format & Check Code (Important!)**
   Before committing, always format your code and run all checks to ensure it meets project standards and will pass CI.

   ```bash
   mage format
   mage check
   ```

4. **Configure Your Environment:**
   Copy the example environment file and customize it as needed for your database and server settings.

   ```bash
   cp .env.example .env
   ```

5. **Run the Application:**
   This project uses [Mage](https://magefile.org/) as a task runner to simplify development workflows. See all available commands by running `mage` with no arguments.
   - **Development Mode:**
     Starts the Go server.

     ```bash
     mage dev
     ```

     The application will be available at `http://localhost:3000`.

   - **Build & Run for Production:**
     To run the application as it would be in production, first build the assets and the binary, then execute the binary.

     ```bash
     mage build:all
     ./bin/server
     ```

---

## Mage Commands

[Mage](https://magefile.org/) is used to automate common development tasks. The build script is the `magefile.go` at the root of the project. You can list all available commands by running `mage` with no arguments.

### Primary Workflow Commands

- `mage dev`: Starts the development server. This is the primary command for local development.
- `mage check`: **(CI)** Runs all quality checks, including Go & frontend linting, tests, and vulnerability scanning. Run this before committing.
- `mage format`: Automatically formats all frontend code using Prettier. Run this if `mage check` reports formatting issues.
- `mage build:all`: Builds all artifacts: frontend assets, the web server binary, and the CLI binary.

### Individual Commands

- **Building:**
  - `mage build:server`: Builds only the main web server binary.
  - `mage build:cli`: Builds only the command-line interface binary.
  - `mage build:css`: Compiles and builds only the Tailwind CSS.
  - `mage build:js`: Compiles and builds only the JavaScript assets.
- **Running:**
  - `mage run:server`: Executes the compiled web server binary.
  - `mage run:cli "args"`: Executes the compiled CLI with the provided arguments.
- **Testing & Quality:**
  - `mage test`: Runs all Go tests.
  - `mage lint`: Lints only the Go codebase.
  - `mage lint:frontend`: Lints only the frontend code.
  - `mage vuln`: Scans for known vulnerabilities in dependencies.
- **Database:**
  - `mage db:migrate`: Applies all pending database migrations.
  - `mage db:createMigration "name"`: Creates a new, timestamped migration file.
- **Housekeeping:**
  - `mage clean`: Removes all build artifacts and temporary files.
  - `mage tidy`: Tidies the `go.mod` and `go.sum` files.
  - `mage installTools`: Installs all required Go-based developer tools.

---

## Project Structure

- **`assets/`**: Source frontend assets (CSS, JS).
- **`cmd/`**: Application entry points.
  - **`server/`**: The main web server.
  - **`cli/`**: A sample `flag` command-line application.
- **`db/`**: Database-related files.
  - **`migrations/`**: SQL schema migrations.
- **`internal/`**: Private application code, including database logic.
- **`public/`**: Compiled, publicly-served assets.
- **`cmd/server/templates/`**: Go `html/template` files.
- **`magefile.go`**: The build script for the project, written in Go.

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
  <a href="httpshttps://signal.me/#p/+dunamismax.66" target="_blank"><img src="https://img.shields.io/badge/Signal-dunamismax.66-3A76F0.svg?style=for-the-badge&logo=signal&logoColor=white" alt="Signal"></a>
</p>

---

<p align="center">
    <img src="https://raw.githubusercontent.com/egonelbre/gophers/refs/heads/master/.thumb/animation/2bit-sprite/demo.gif" alt="Gopher Sprite Animation" />
</p>

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
