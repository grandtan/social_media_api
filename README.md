# Social Media API

## Setup Instructions

1. **Clone the repository:**

   git clone https://github.com/yourusername/social_media_api.git
   cd social_media_api

2. Install dependencies: go mod download

3. Setup the database and Run the server: go run main.go

How to Run Tests
Run unit tests:go test ./tests

API Documentation

1. Create a New User
   Endpoint: POST /users
   Description: Creates a new user with the given name and email.
   Example:

curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{
"name": "Test User",
"email": "test@example.com"
}'

2. Get All Users
   Endpoint: GET /users
   Description: Retrieves a list of all users.
   Example:

curl -X GET http://localhost:8080/users

3. Get User by ID
   Endpoint: GET /users/{id}
   Description: Retrieves a user by their ID.
   Example:

curl -X GET http://localhost:8080/users/1

4. Update User by ID
   Endpoint: PUT /users/{id}
   Description: Updates a user's information by their ID.
   Example:

curl -X PUT http://localhost:8080/users/1 \
-H "Content-Type: application/json" \
-d '{
"name": "Updated User",
"email": "updated@example.com"
}'

5. Delete User by ID
   Endpoint: DELETE /users/{id}
   Description: Deletes a user by their ID.
   Example:

curl -X DELETE http://localhost:8080/users/1

Certainly! Here's the documentation in English for using curl to test all endpoints of your API.

Application Overview
This application is a simple social media platform where users can create accounts and post text content. The API provides endpoints to create, read, update, and delete users and posts. Authentication is implemented to secure the endpoints.

API Documentation

1. Create a New User
   Endpoint: POST /users
   Description: Creates a new user with the given name and email.
   Example:
   sh

   curl -X POST http://localhost:8080/users \
   -H "Content-Type: application/json" \
   -d '{
   "name": "Test User",
   "email": "test@example.com"
   }'

2. Get All Users
   Endpoint: GET /users
   Description: Retrieves a list of all users.
   Example:
   sh

   curl -X GET http://localhost:8080/users

3. Get User by ID
   Endpoint: GET /users/{id}
   Description: Retrieves a user by their ID.
   Example:
   sh

   curl -X GET http://localhost:8080/users/1

4. Update User by ID
   Endpoint: PUT /users/{id}
   Description: Updates a user's information by their ID.
   Example:
   sh

   curl -X PUT http://localhost:8080/users/1 \
   -H "Content-Type: application/json" \
   -d '{
   "name": "Updated User",
   "email": "updated@example.com"
   }'

5. Delete User by ID
   Endpoint: DELETE /users/{id}
   Description: Deletes a user by their ID.
   Example:
   sh

   curl -X DELETE http://localhost:8080/users/1

6. Create a New Post
   Endpoint: POST /posts
   Description: Creates a new post for a specified user.
   Example:

curl -X POST http://localhost:8080/posts \
-H "Content-Type: application/json" \
-d '{
"user_id": 1,
"content": "This is a test post."
}'

7. Get All Posts
   Endpoint: GET /posts
   Description: Retrieves a list of all posts.
   Example:

curl -X GET http://localhost:8080/posts

8. Get Post by ID
   Endpoint: GET /posts/{id}
   Description: Retrieves a post by its ID.
   Example:

curl -X GET http://localhost:8080/posts/1

9. Update Post by ID
   Endpoint: PUT /posts/{id}
   Description: Updates a post's content by its ID.
   Example:

curl -X PUT http://localhost:8080/posts/1 \
-H "Content-Type: application/json" \
-d '{
"content": "This is an updated test post."
}'

10. Delete Post by ID
    Endpoint: DELETE /posts/{id}
    Description: Deletes a post by its ID.
    Example:

curl -X DELETE http://localhost:8080/posts/1

Notes:
Make sure to replace localhost:8080 with the appropriate address and port where your API server is running.
Adjust the {id} placeholders in the URLs with the actual IDs of the users and posts you have in your database.
