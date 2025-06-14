# ghstats

Github Repo Stats CLI

# ghstats - GitHub Repository Stats CLI

![Go Version](https://img.shields.io/badge/go-1.24+-blue.svg)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)
![License](https://img.shields.io/badge/license-MIT-lightgrey.svg)

A simple yet powerful command-line tool to quickly fetch and display key statistics for any public GitHub repository, featuring a clean presentation and a caching layer to ensure high performance.

---

### Demo

Here is an example of fetching the stats for the `microsoft/vscode` repository:

```sh
$ go run ./cmd/ghstats/ --repo=microsoft/vscode
```

```
┌────────────────────┐
│  microsoft/vscode  │
└────────────────────┘

> Visual Studio Code

──────────────────────
Language:       TypeScript
Stars:          173,527 ⭐
Forks:          33,050 🔱
Open Issues:    10,550 ❗
Last Updated:   2025-06-14 20:03:41

🔗 https://github.com/microsoft/vscode
```

### Features

- **Detailed Stats:** Fetches essential repository metrics including Stars, Forks, Open Issues, Language.
- **High-Performance Caching:** Utilizes Redis to cache API responses for one hour.
- **Clean, Dynamic UI:** Presents statistics in a well-formatted and easy-to-read layout directly in your terminal.
- **Robust Architecture:** Built with a clean, decoupled architecture for maintainability and testability.
- **Configuration Driven:** Uses environment variables for simple and secure configuration.

### Installation & Setup

1.  **Clone the repository:**

    ```sh
    git clone [https://github.com/akyTheDev/ghstats.git](https://github.com/akyTheDev/ghstats.git)
    cd ghstats
    ```

2.  **Install dependencies:**

    ```sh
    go mod tidy
    ```

3.  **Configure Environment Variables:**
    This application requires two environment variables. You can set them in your shell or create a `.env` file.

    - **`GITHUB_TOKEN`**: A GitHub Personal Access Token.

      ```sh
      export GITHUB_TOKEN="your_github_personal_access_token"
      ```

    - **`REDIS_URL`**: The connection string for your running Redis instance.
      ```sh
      # For a standard local installation
      export REDIS_URL="redis://localhost:6379/0"
      ```

### Usage

Run the application using `go run`, passing the repository you want to query with the `--repo` flag.

```sh
go run ./cmd/ghstats/ --repo=golang/go
```

To see all available flags, use `--help`:

```sh
go run ./cmd/ghstats/ --help
```

### Building the Binary

You can compile the application into a single executable binary for easy use and distribution.

```sh
go build -o ghstats ./cmd/ghstats/
```

Now you can run it directly:

```sh
./ghstats --repo=torvalds/linux
```

### Running Tests

The project includes a suite of unit tests with mocking to ensure reliability. To run the tests:

```sh
go test ./... -v
```

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
