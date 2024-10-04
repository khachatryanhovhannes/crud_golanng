# Go CRUD API for Posts

A simple CRUD (Create, Read, Update, Delete) API for managing posts using Go and the Gin framework. This project demonstrates basic RESTful API principles and integrates with a MySQL database using GORM.

## Table of Contents

- [Go CRUD API for Posts](#go-crud-api-for-posts)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Technologies Used](#technologies-used)
  - [Installation](#installation)
  - [Usage](#usage)
  - [API Documentation](#api-documentation)
    - [1. Create a Post](#1-create-a-post)
    - [2. Get All Posts](#2-get-all-posts)
    - [3. Get a Single Post](#3-get-a-single-post)
    - [4. Update a Post](#4-update-a-post)
    - [4. Delete a Post](#4-delete-a-post)

## Features

- Create, Read, Update, and Delete posts
- RESTful API endpoints
- Database integration with GORM and MySQL

## Technologies Used

- Go (Golang)
- Gin (web framework)
- GORM (ORM for Go)
- MySQL (database)

## Installation

1. Clone the repository:

   ```bash
    git clone https://github.com/khachatryanhovhannes/crud_golanng.git
    cd crud_go
2. Create a .env file in the root directory and configure your database settings:
   ```bash
    DB_USER=root
    DB_PASSWORD=your_password
    DB_NAME=posts_db
3. Install dependencies:
   ```bash
    go mod tidy
3. Run the application:
   ```bash
    go run main.go
The API will be available at http://localhost:8080.

## Usage

You can use a tool like Postman or curl to interact with the API.

## API Documentation

### 1. Create a Post

- **Endpoint:** `POST /posts`

- **Request Body:**

  ```json
  {
      "title": "post 1",
      "body": "Post 1 text content"
  }
- **Response:**
  ```json
  {
    "id": 1,
    "title": "post 1",
    "body": "Post 1 text content",
    "created_at": "2024-10-04T18:24:03Z",
    "updated_at": "2024-10-04T18:24:03Z"
  }
### 2. Get All Posts

- **Endpoint:** `GET /posts`

- **Response:**
  ```json
  [
    {
        "id": 1,
        "title": "post 1",
        "body": "Post 1 text content",
        "created_at": "2024-10-04T18:24:03Z",
        "updated_at": "2024-10-04T18:24:03Z"
    },
    {
        "id": 2,
        "title": "post 2",
        "body": "Post 2 text content",
        "created_at": "2024-10-04T18:25:01Z",
        "updated_at": "2024-10-04T18:25:01Z"
    }
  ]
### 3. Get a Single Post

- **Endpoint:** `GET /posts/:id`

- **Response:**
  ```json
  {
    "id": 1,
    "title": "post 1",
    "body": "Post 1 text content",
    "created_at": "2024-10-04T18:24:03Z",
    "updated_at": "2024-10-04T18:24:03Z"
  }
### 4. Update a Post

- **Endpoint:** `PUT /posts/:id`

- **Request Body:**

  ```json
  {
    "title": "updated post",
    "body": "Updated text content"
  }
- **Response:**
  ```json
  {
    "id": 1,
    "title": "updated post",
    "body": "Updated text content",
    "created_at": "2024-10-04T18:24:03Z",
    "updated_at": "2024-10-04T18:30:00Z"
  }
### 4. Delete a Post

- **Endpoint:** `DELETE /posts/:id`

- **Response:** `204 No Content`
