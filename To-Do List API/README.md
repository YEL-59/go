
To-Do List API in Go
Project Overview
This project is a RESTful API for managing a to-do list, built using Go and the Gin web framework. It allows users to perform basic CRUD (Create, Read, Update, Delete) operations on tasks. The API interacts with a SQLite or PostgreSQL database for persistent storage.

Features:
Create, view, update, and delete tasks.
JSON-based API for easy integration.
SQLite for lightweight storage (or PostgreSQL for a more scalable solution).
Simple architecture for learning Go's backend development.
Technologies Used
Go (Golang)
Gin (Web Framework)
GORM (ORM for Go)
SQLite/PostgreSQL (Database)
Docker (for containerization, optional)
Getting Started
Prerequisites
Before you begin, ensure you have the following installed:

Go (v1.18+): Download Go
SQLite (optional): Pre-installed on most systems, but needed if you want to inspect the database file.
Postman or cURL: For testing API endpoints.
Project Setup
Clone the repository:

bash
Copy code
git clone https://github.com/yourusername/todo-api-go.git
cd todo-api-go
Initialize Go modules:

bash
Copy code
go mod tidy
Install dependencies:

bash
Copy code
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite # For SQLite (or use the PostgreSQL driver if needed)
Configure your database:

By default, the project uses SQLite. The database file tasks.db will be created automatically when you run the project.
To use PostgreSQL, update the database connection in db/database.go with your PostgreSQL credentials.
File Structure
graphql
Copy code
todo-api/
├── go.mod                   # Go modules
├── main.go                  # Entry point for the application
├── models/
│   └── task.go              # Task model definition
├── handlers/
│   └── task_handler.go      # Handlers for task-related routes
├── db/
│   └── database.go          # Database initialization and connection
├── tasks.db                 # SQLite database file (auto-created)
└── README.md                # Documentation (this file)
Running the Project
Run the API:

You can start the Go application with:

bash
Copy code
go run main.go
Access the API:

By default, the API runs on http://localhost:8080. You can test the API using Postman or cURL.

API Endpoints
1. Get All Tasks
URL: /tasks
Method: GET
Description: Fetch all tasks in the database.
bash
Copy code
curl -X GET http://localhost:8080/tasks
Response Example:
json
Copy code
[
  {
    "id": 1,
    "title": "Finish Go project",
    "due_date": "2024-10-15T00:00:00Z",
    "status": "pending"
  }
]
2. Create a New Task
URL: /tasks
Method: POST
Description: Add a new task to the list.
bash
Copy code
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title":"Start learning Docker","due_date":"2024-10-20T00:00:00Z","status":"pending"}'
Response Example:
json
Copy code
{
  "id": 2,
  "title": "Start learning Docker",
  "due_date": "2024-10-20T00:00:00Z",
  "status": "pending"
}
3. Get a Task by ID
URL: /tasks/{id}
Method: GET
Description: Fetch a task by its unique ID.
bash
Copy code
curl -X GET http://localhost:8080/tasks/1
Response Example:
json
Copy code
{
  "id": 1,
  "title": "Finish Go project",
  "due_date": "2024-10-15T00:00:00Z",
  "status": "pending"
}
4. Update a Task
URL: /tasks/{id}
Method: PUT
Description: Update an existing task.
bash
Copy code
curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-d '{"title":"Finish the Go project","due_date":"2024-10-18T00:00:00Z","status":"completed"}'
Response Example:
json
Copy code
{
  "id": 1,
  "title": "Finish the Go project",
  "due_date": "2024-10-18T00:00:00Z",
  "status": "completed"
}
5. Delete a Task
URL: /tasks/{id}
Method: DELETE
Description: Delete a task by ID.
bash
Copy code
curl -X DELETE http://localhost:8080/tasks/1
Response:
bash
Copy code
Status: 204 No Content
Testing the API
Using Postman
Open Postman.
Create a new request.
Choose the HTTP method (GET, POST, PUT, DELETE).
Enter the URL http://localhost:8080/{endpoint}.
Add the request body (for POST/PUT requests) as raw JSON data.
Send the request and inspect the response.
Using cURL
You can use the example cURL commands provided in the API Endpoints section to test the API directly from the command line.

Deployment
Option 1: Deploying Locally
Simply run:

bash
Copy code
go run main.go
The server will start on localhost:8080.