## Running the Auth Microservice Manually

If you are currently working on the Auth Service microservice, follow these steps to run it manually:

1. Execute the `run-docker.sh` script located on the root of the project to run the full API at least once. This will take care of the application requirements.

2. Stop the `auth-service` container inside docker.

3. Run the following command with the specific argument `debug` inside this directory ( from this README.md file )
    ```sh
    go run main.go debug
    ```

4. Develop the feature. Keep in mind that `auth-service` and `nginx` are highly coupled in Docker, so our middleware won't work. Therefore, no requests for the `db-service` or `diagnostics-service` will be able to be executed. Develop the full feature, then test the other API routes as needed (most likely won't be an issue).