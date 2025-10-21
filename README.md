# get-post-go

A lightweight Go REST API for managing users, supporting JSON-based storage and containerization with Docker.

## Project Structure

```
get_post_go_lang/
├── controllers/  # HTTP handlers
├── models/       # User model and JSON storage
├── main.go       # Entry point
├── go.mod
└── Dockerfile    # Docker build instructions
```

## Features

* Add a new user (`POST /users`)
* Get a user by ID (`GET /users?id=1`)
* Persistent JSON storage in the `users` folder
* Docker-ready for containerized deployment

## Requirements

* Go 1.25+
* Docker (optional)

## Running Locally

### With Go

```bash
go run main.go
```

Server runs on `http://localhost:3030`

### With Docker (Build Locally)

```bash
docker build -t get-post-go .
docker run -d -v ./users:/app/users -p 3030:3030 get-post-go
```

### Run Directly from Docker Hub

```bash
docker run -d -v ./users:/app/users -p 3030:3030 ziadghoraba/get-post-go:latest
```

* `-d` → run container in detached mode
* `-v ./users:/app/users` → mount local folder for persistent storage
* `-p 3030:3030` → map container port to host

## API Examples

### POST /users

```bash
curl -X POST http://localhost:3030/users \
-H "Content-Type: application/json" \
-d '{"Name":"Ziad","Age":22,"Address":{"Street":"Mossab Ibn-Omair","City":"Cairo","Country":"Egypt"}}'
```

### GET /users

```bash
curl http://localhost:3030/users?id=1
```

## Notes

* Ensure the host `users` folder exists to persist data.
* Default API port: `3030`
* Compatible with Docker for easy deployment.

## License

MIT License
