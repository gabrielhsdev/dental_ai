# Services Overview

| Service               | Language     | Responsibility                                      |
|-----------------------|--------------|-----------------------------------------------------|
| **Auth Service**       | Go           | User login, registration, JWT authentication        |
| **Database Service**   | Go           | CRUD for users, animals, adoptions                  |
| **Recommendation Service** | Go       | Fetches user preferences, calls ML for ranking      |
| **ML Service**         | Python       | AI-based animal ranking                             |
| **Frontend**           | React Native | User interface, swipe interactions, adoption requests |
| **NGINX**              | Config       | API Gateway, load balancing, redirects to proper server                         |


# 1. Auth Service (Go)

📌 Handles user authentication and authorization
- ✅ User login & registration (email/password, OAuth)
- ✅ Issues JWT tokens for authentication
- ✅ Validates token-based requests
- ✅ Manages user roles (shelter, adopter, admin)

```
➡ Exposes APIs:
POST /register – Create a new user
POST /login – Authenticate user & return JWT
GET /me – Get logged-in user details
```

# 2. Database Service (Go)

📌 Manages interactions with MySQL database
- ✅ Provides CRUD APIs for user, animal, and adoption data
- ✅ Stores animal profiles (images, descriptions, status)
- ✅ Tracks user preferences & interactions (swipe data)
- ✅ Manages adoption requests & approval

```
➡ Exposes APIs:
GET /animals – List all animals
POST /animals – Add a new animal
GET /users/{id} – Get user details
POST /adoptions – Create an adoption request
```

# 3. Recommendation Service (Go)

📌 Filters animals based on user preferences (Basic Matching)
- ✅ Fetches user swipe history (left/right preferences)
- ✅ Queries MySQL for matching animals
- ✅ Applies basic rules (e.g., breed, size, age)
- ✅ Calls ML Service (optional) for refined results

```
➡ Exposes APIs:
GET /recommendations/{userId} – Get animals based on preferences
```

# 4. ML Service (Python)

📌 Uses AI to refine recommendations
- ✅ Analyzes swipe patterns to predict preferences
- ✅ Uses image embeddings to find visually similar animals
- ✅ Ranks results based on deep learning models
- ✅ Can be trained over time with more data

```
➡ Exposes APIs:
POST /ml/recommend – Returns ranked animals based on ML model
POST /ml/train – (Optional) Retrains model with new data
````

# 5. Frontend (React Native)

📌 User-facing mobile app for adopters & shelters

- ✅ User Authentication (login/register)
- ✅ Swipe-based UI to like/dislike animals
- ✅ Recommendation Page (Netflix-style grid of animals`
- ✅ Animal Details Page (with adoption button)
- ✅ Shelter Dashboard (CRUD for animal management)

# 6. NGINX (Reverse Proxy)
📌 Handles API Gateway & Load Balancing

- ✅ Routes API requests to the right microservice
- ✅ Handles SSL (HTTPS) if needed
- ✅ Caches static content (optional)
```
➡ Example Routing:
/api/auth/* → Auth Service
/api/db/* → Database Service
/api/recommend/* → Recommendation Service
/api/ml/* → ML Service
/frontend → React Native frontend assets (if needed)
```
