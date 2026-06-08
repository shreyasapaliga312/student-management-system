# Student Management System REST API

A backend application developed using Golang and MySQL to manage student records through RESTful APIs. The system supports CRUD operations, user authentication, and class management.

## Features

* User Registration and Login using JWT Authentication
* Create, Read, Update, and Delete Student Records
* Create and Manage Classes
* Assign Students to Classes
* Secure API Endpoints using JWT Tokens
* MySQL Database Integration

## Technologies Used

* Golang
* MySQL
* GORM
* Gorilla Mux
* JWT Authentication
* Docker
* Git & GitHub
* Postman

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/student-management-system.git
cd student-management-system
```

### 2. Configure Environment Variables

Update the `.env` file with your database credentials and application settings.

### 3. Start the Application

```bash
docker compose up
```

## API Endpoints

### Authentication

| Method | Endpoint  | Description                 |
| ------ | --------- | --------------------------- |
| POST   | /register | Register a new user         |
| POST   | /login    | Login and receive JWT token |

### Student Management

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| POST   | /student      | Create a student  |
| GET    | /student/all  | Get all students  |
| GET    | /student/{id} | Get student by ID |
| PUT    | /student/{id} | Update student    |
| DELETE | /student/{id} | Delete student    |

### Class Management

| Method | Endpoint    | Description     |
| ------ | ----------- | --------------- |
| POST   | /class      | Create a class  |
| GET    | /class/all  | Get all classes |
| GET    | /class/{id} | Get class by ID |
| DELETE | /class/{id} | Delete class    |

## Project Highlights

* Developed RESTful APIs using Golang.
* Implemented JWT-based Authentication and Authorization.
* Integrated MySQL database using GORM ORM.
* Performed API testing using Postman.
* Containerized the application using Docker.
* Managed source code using Git and GitHub.


