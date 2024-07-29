# Social Media API

## Objective

Develop a small Go application or set of features that demonstrates your ability to work with the language and relevant technologies.
Recommended Time: 2-3 hours; this is the scope/quality we seek.
Recommended Application: Small Social Media (that allows users to post text)

## Requirements

1. **API Development**
   - Create a RESTful API to manage a resource (e.g., users, posts).
   - Implement CRUD operations.
   - Use appropriate HTTP status codes.
   - Secure the API with authentication (basic token authentication could be sufficient).
2. **Database Integration**
   - Use GORM to define models with relationships (e.g., one-to-many, many-to-many).
   - Implement model validations.
   - Write at least one migration to alter the database schema.
3. **Frontend (Optional, if time allows)**
   - Develop a simple front end using Go templates or a JavaScript frontend that consumes the API.
   - Implement simple forms to create and edit the models and list views to display them.
4. **Testing**
   - Write unit tests for the models, including validations and methods.
   - Write functional tests for the API endpoints, ensuring all cases are covered.
5. **Error Handling**
   - Implement error handling in the API that gracefully handles common errors such as records not being found or invalid data.
6. **Documentation**
   - Document the API endpoints with request and response examples. This can be in the form of comments in the controller, a README file, or a tool like Swagger.

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
