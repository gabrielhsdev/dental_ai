## 🛠 Running the DB Microservice Locally (Manual Dev Mode)

If you're actively working on the `db-service` and want to run it outside Docker for live development, follow these steps:

### 1. Bootstrap the Environment

Run the `run-docker.sh` script from **the project root** to spin up the entire application. This ensures all required services (like the database) are provisioned correctly:

```sh
./run-docker.sh
```

### 2. Stop the Dockerized DB Service

Once the containers are up, stop only the `db-service` container. This frees the port for local development:

```sh
docker stop db-service
```

### 3. Run the DB Service Manually

You have two options:

- **Hot reload (recommended for development):**

  ```sh
  air
  ```

- **Manual run (use this first to check for compile errors):**

  ```sh
  go run main.go development
  ```

> ℹ️ `development` is a required argument. It triggers dev-mode configurations.

## ⚡️ Enabling Live Reload with Air

For a smooth development workflow, we use [**Air**](https://github.com/air-verse/air) — a hot reloader for Go projects.

### 1. Install Air

Use the following command to install Air globally:

```sh
go install github.com/air-verse/air@latest
```

This will place the binary in your Go bin path (usually `~/go/bin`).

### 2. Add Go Binaries to PATH (if needed)

Make sure your shell can find the Air binary:

#### For `zsh`:

```sh
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc
source ~/.zshrc
```

#### For `bash`:

```sh
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bashrc
source ~/.bashrc
```

Verify installation with:

```sh
which air
```

Then, run the project with `air`
```sh
air
```