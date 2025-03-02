## Running the Auth Microservice Manually

If you are currently working on the Auth Service microservice, follow these steps to run it manually

1. Add the .env file to our config folder. Should be the same as the one found in the project root.

2. Initialize the Go module:
    ```sh
    go mod init main.go
    ```

3. Tidy up the module dependencies:
    ```sh
    go mod tidy
    ```

4. Run the main Go file:
    ```sh
    go run main.go
    ```