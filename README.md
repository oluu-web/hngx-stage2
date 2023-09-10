# API Documentation

## Overview

This document provides an overview of the API structure and key components of the API. 
It is pertinent to note that this API was created with the assumption that no two people can the same name.

## Getting Started

To run the API locally, follow these steps:

### Prerequisites

- Go programming language installed.
- MongoDB server accessible (update the .env file with the correct MongoDB URI)

### Get dependencies
You need to get all the depdndencies before running the API.Navigate to the root of the project and run ```go get```

## Running the API
- Navigate to the api folder: cd cmd/api
- Run the api using thie follwoing command: ```go run .```

## Folder Structure

The API project has the following folder structure:

- `cmd`
  - `api`
    - `config`: Configuration files.
    - `controllers`: Controller functions.
    - `middleware`: Middleware functions.
    - `models`: Data models and database operations.
    - `routes`: HTTP route definitions.
    - `utilities`: Utility functions.
    - `.env`: Environment configuration file.
    - `main.go`: Main application entry point.

## `main.go`

The `main.go` file is the entry point of the application. It initializes the MongoDB connection and starts the HTTP server.

## `config`

The `config` folder contains configuration files, including the `.env` file used to load environment variables.

## `controllers`

The `controllers` folder contains HTTP request handler functions for managing Person resources. It includes functions for creating, retrieving, updating, and deleting Person objects.

- `CreatePerson`: Creates a new Person.
- `GetPerson`: Retrieves a Person by name.
- `UpdatePerson`: Updates an existing Person.
- `DeletePerson`: Deletes a Person by name.

## `middleware`

The `middleware` folder provides middleware functions, including `EnableCORS` to enable Cross-Origin Resource Sharing (CORS) for HTTP requests.

## `models`

The `models` folder defines the data structure for a Person and provides database-related functionality, including validation, data access, and MongoDB interactions.

- `Person`: Represents a Person resource.
- `ConnectToDB`: Establishes a connection to the MongoDB database.
- `GetDBCollection`: Returns a reference to a collection in the MongoDB database.
- Other functions: Validation, duplicate checking, creation, retrieval, updating, and deletion of Person objects.

## `routes`

The `routes` folder defines HTTP routes (endpoints) for the application using the `httprouter` package. It associates routes with controller functions.

- `InitRoutes`: Initializes the HTTP router with defined routes and their associated contoller functions, and applies CORS middleware. It set up routes for creating, retrieving, updation and deleting Person resources.

## `utilities`

The `utilities` folder contains utility functions for handling HTTP responses and errors.

- `WriteJSON`: Writes a JSON response to the provided http.ResponseWriter. It wraps the data in a map with the specified key and marshals it to JSON. It sets the Content-Type header to indicate JSON content, specifies the HTTP status code, and writes the JSON response to the writer..
- `ErrorJSON`: Writes an error JSON response to the provided http.ResponseWriter. It encapsulates the error message in a JSONError struct and then calls WriteJSON to write the error response with an HTTP status code of Bad Request (HTTP 400).


## Conclusion

This documentation provides an overview of the structure and key components of the API. For more detailed information on each component and their functions, refer to the corresponding source code files in the repository.
