# Services Overview  ( Beta )

| **Service**            | **Language** | **Responsibility**                                      |
|------------------------|--------------|---------------------------------------------------------|
| **Auth Service**        | Go           | Handles user authentication and authorization using JWT. |
| **Database Service**    | **MySQL**    | Provides CRUD operations for storing and managing data. |
| **Diagnostics Service** | Python       | Processes patient data using machine learning algorithms to generate insights. |
| **NGINX**              | N/A          | Acts as a reverse proxy, routing requests to the appropriate services. |

# Application Flow Overview

## 1. **Container Setup**

The application is composed of several Docker containers, each handling a specific service. These containers communicate with each other to ensure efficient processing and data flow.

### Containers:
- **Auth Service**: Manages user authentication and JWT-based authentication.
- **DB Service**: Handles database operations and data persistence.
- **Diagnostics Service**: Processes patient data using machine learning models.
- **NGINX**: Acts as a reverse proxy, directing requests to the correct service.

## 2. **NGINX Reverse Proxy & Request Routing**

- **Reverse Proxy**: NGINX routes incoming requests to the appropriate backend service.
- **Authentication Middleware**: Requests to **DB Service** and **Diagnostics Service** are authenticated through the **Auth Service** via a POST request for JWT validation.
- **Service Routing**:
  - Requests to authentication endpoints are forwarded to **Auth Service**.
  - Database-related requests go to **DB Service**. We can save results from the diagnostics service on these requests by asking disgnostics service to perform an an ction first, then, using the DB service to save such results.
  - **Diagnostics Service** handles requests for medical data analysis and machine learning predictions.

## 3. **Inter-Service Communication**

- **Auth Service**: Manages JWT authentication and authorization.
- **DB Service**: Handles and serves requests related to patient data and database interactions.
- **Diagnostics Service**: Analyzes patient data, and provides insights using machine learning models.

## 4. **Health Checks and Dependencies**

- **Postgres**: Ensures database availability before other services start.
- **Auth Service**: Ensures the authentication service is running before allowing access to dependent services.
- **Diagnostics Service**: Ensures the necessary machine learning models are loaded and ready before accepting requests.

## 5. **Security Measures**

- **JWT Authentication**: The **Auth Service** validates JWT tokens before granting access to the **DB Service** and **Diagnostics Service**.
- **Middleware Authorization**: Requests to **DB Service** and **Diagnostics Service** are authenticated through middleware that queries the **Auth Service**.
- **Network Security**: The NGINX reverse proxy ensures secure and efficient routing while restricting unauthorized access.

## 6. **Containerization and Orchestration**

- Each service runs in a dedicated **Docker container**.
- Containers communicate over a **Docker network**, ensuring service isolation and scalability.
- The **docker-compose.yml** file defines service dependencies, networks, and volumes.
- Individual services can be updated or scaled without disrupting the entire system.

