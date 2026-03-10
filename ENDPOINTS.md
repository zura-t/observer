# API Endpoints

## Users

| Method | Endpoint | Description | Request |
|--------|----------|-------------|---------|
| POST | `/register` | Register a new user | Body: `{ "email": string, "password": string, "name": string }` |
| GET | `/users/:id` | Get user by ID | URI param: `id` |
| GET | `/users/email?email=123` | Get user by email | Query: `email` |
| PATCH | `/users/:id` | Update user | URI param: `id`, Body: `{ "name": string, "email": string }` |
| DELETE | `/users/:id` | Delete user | URI param: `id` |
