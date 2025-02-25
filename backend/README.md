# Services Overview

| Service               | Language     | Responsibility                                         |
|-----------------------|--------------|--------------------------------------------------------|
| **Auth Service**       | Go           | User login, registration, JWT authentication          |
| **Database Service**   | Go           | CRUD for the db in general                            |
| **Recommendation Service** | Go       | Calls ML for ranking                                  |
| **ML Service**         | Python       | AI Stuff                                              |
| **Frontend**           | React Native | Dentist interface                                     |
| **NGINX**              | Config       | API Gateway, load balancing, redirects to proper server|