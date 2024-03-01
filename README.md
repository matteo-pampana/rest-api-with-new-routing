# Notes REST API with New Go 1.22 Routing

This is an example of a RESTful service written in Go. The service is a note management application that allows you to create, read, update, and delete notes (CRUD). 

The REST APIs were written using the new routing functions of Go 1.22. 

**This project was created for learning and demonstration purposes.**

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.22 or higher

### Installing

Clone the repository:

```bash
git clone https://github.com/matteo-pampana/rest-api-with-new-routing.git
```

Navigate to the cloned directory:

```bash
cd rest-api-with-new-routing
```

Download the dependencies:

```bash
make build
```

### Running the application

To start the HTTP server at port 8080, run:

```bash
make run
```

## API Endpoints

- `POST /notes`: Create a new note
- `GET /notes`: Get all notes
- `GET /notes/{id}`: Get a note by ID
- `PUT /notes/{id}`: Update a note by ID
- `DELETE /notes/{id}`: Delete a note by ID

---

Created with ❤️ by [Matteo Pampana](https://medium.com/@matteopampana) 
