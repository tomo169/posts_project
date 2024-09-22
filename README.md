# Go + Vue.js Full Stack Application

This is a full-stack web application built with:
- **Backend**: Go (Golang) for handling APIs and business logic.
- **Frontend**: Vue.js with Vite for handling the user interface.
- **Docker**: Docker is used to containerize both the backend and frontend applications.

## Features

- **User Authentication**: JWT-based login and registration.
- **Post Management**: Users can create, update, delete, and view posts.
- **Profile Management**: Users can manage their profile and view their posts.
- **Protected Routes**: Some routes are protected and require authentication.
- **Dockerized**: Both backend and frontend are containerized using Docker.

## Running the Application with Docker

To easily run both the backend and frontend using Docker, follow these steps:

###  Build and Run the Docker Containers

1. **Build the Docker containers**:

   ```bash
   docker-compose up --build


2. **Stopping the Docker containers**:

   ```bash
   docker-compose down

   
