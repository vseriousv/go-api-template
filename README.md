# go-api-template

This project is a GoLang-based template designed to serve as a foundation for developing future projects.
It includes directory structures, packages, and functions needed for API development, error handling, database interaction, and more.

## Project Structure

```
├── cmd
│   ├── api
│   │   └── main.go
│   └── workers
│       └── ... (various worker services)
├── internal
│   ├── api
│   │   ├── dto (API Data Transfer Objects)
│   │   ├── handlers (API handlers)
│   │   ├── repositories (Database models)
│   │   ├── services (Business logic)
│   │   └── app.go (API routes and configurations)
│   ├── config
│   │   └── config.go (Configuration, goenv environment variable initialization)
│   └── ... (Various project packages)
├── pkg
│   ├── utils
│   │   ├── handle_error.go (Error handling function)
│   │   └── ... (Other helper functions for the project)
│   └── ... (Other packages that can be called from external projects)
```

## Getting Started

1. Clone this repository:
```shell
$ git clone https://github.com/vseriousv/go-api-template
```
2. Navigate to the project directory:
```shell
$ cd go-api-template
```
3. Install the required dependencies:
```shell
$ go mod download
```
4. *Build and run the project:
```shell
$ go build -o bin/api cmd/api/main.go
$ ./bin/api
```

## Configuration

Configuration settings can be found and modified in the internal/config/config.go file. You can also use the goenv package to initialize environment variables. \
To specify developer environment variables, create an .env file at the root of the project

## Contributing

Contributions are welcome! To contribute, please follow these steps:

1. Fork this repository
2. Create a new branch with your changes
3. Submit a pull request

Please make sure your code is well-formatted and follows the project's coding style.