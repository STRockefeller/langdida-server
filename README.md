# langdida-server

langdida-server is the backend project that complements langdida-ui. langdida is a personal language learning tool inspired by lingQ, a popular language learning platform. However, I have customized certain aspects of the tool to better suit my needs.

The server-side functionality is implemented using Golang and is designed to support the following features:

1. User authentication and authorization
2. Article management and retrieval
3. Vocabulary tracking and review
4. Learning progress tracking

## Installation

### Prerequisites

Before getting started, ensure that you have the following prerequisites installed on your development machine:

- Golang
- Database (Postgresql or Sqlite)

### Setting up the database

1. Create a new PostgreSQL database for langdida-server.
2. Update the database connection settings in the  `config.yml`  file.

### Running the server

1. Clone the repository.
2. Run  `go mod download`  to install the dependencies.
3. Run  `go run main.go`  to start the server.

## API Documentation

The server exposes a RESTful API for communication with the langdida-ui. The API endpoints and their respective functionalities are documented in the [API Documentation](api-docs.md) file.

## Contribution

If you would like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature ( `git checkout -b feature/fooBar` ).
3. Commit your changes ( `git commit -m 'Add some fooBar'` ).
4. Push to the branch ( `git push origin feature/fooBar` ).
5. Create a new Pull Request.

## Meta

Distributed under the MIT license. See  `LICENSE`  for more information.

## Screenshots(UI)

see also [langdida-ui](https://github.com/STRockefeller/langdida_ui)

![table](https://i.imgur.com/MZvbrBb.png)
![graph](https://i.imgur.com/pOcdFtU.png)
![dialog](https://i.imgur.com/v6BvUkF.png)
