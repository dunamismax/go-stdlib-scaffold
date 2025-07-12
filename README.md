# The Pure Go Standard Library Stack

This project is a web application built using a minimalist, robust architecture. It is composed entirely of a Go backend that leverages the standard library, removing all external dependencies for routing, validation, and data access. The stack prioritizes ultimate simplicity, zero-dependency deployment, and long-term stability by relying exclusively on Go's native capabilities. The frontend is plain HTML and CSS, with no JavaScript.

## Tech Stack

- **Backend:** Go 1.22+ (`net/http`)
- **Database:** SQLite (`database/sql`)
- **Frontend:** HTML & Plain CSS (`html/template`)
- **Build/Task Runner:** Mage
- **Migrations:** Plain SQL & Go

## Prerequisites

- Go 1.22+
- Mage

## Getting Started

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/dunamismax/go-stdlib-scaffold.git
    cd go-stdlib-scaffold
    ```

2.  **Install dependencies:**
    This project has very few dependencies. `go mod tidy` will install the SQLite driver.
    ```sh
    go mod tidy
    ```

3.  **Run database migrations:**
    This will create the `data/` directory and the `app.db` SQLite database with the necessary tables.
    ```sh
    mage migrate
    ```

4.  **Run the application:**
    ```sh
    mage run
    ```
    The server will start on port 3000 by default. You can access it at [http://localhost:3000](http://localhost:3000).

## Available Commands (Mage)

- `mage build`: Build the server and CLI binaries.
- `mage run`: Build and run the web server.
- `mage migrate`: Apply database migrations.
- `mage clean`: Remove build artifacts.
- `mage tidy`: Tidy the `go.mod` file.

## Configuration

The application is configured using environment variables:

- `APP_SERVER_PORT`: The port for the web server (default: `3000`).
- `APP_DB_PATH`: The path to the SQLite database file (default: `data/app.db`).

You can set these in a `.env` file, which is loaded automatically if present (though this project does not ship with a `.env` loader, you can add one or `source` it).