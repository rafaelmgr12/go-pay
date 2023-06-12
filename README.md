<h1 align="center">
  <img alt="Logo" src="./doc/img/logo.png">
</h1>

<h1 align="center">Go Pay</h1>
<p align = "center">A digital wallet for making transfers</p>

<p align="center">
  <a href="#-technology">Technology</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
    <a href="#-project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-run">How to Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-license">License</a>
</p>

<p align="center">
  <img alt="License" src="https://img.shields.io/static/v1?label=license&message=MIT&color=8257E5&labelColor=000000">
</p>

## Introduction
Go Pay is a simple Go application that provides a basic balance transfer functionality. It uses a MySQL database to store balance information and exposes a web API for balance retrieval and transfer operations.

The following challenge can be found [here](https://app.devgym.com.br/challenges/9af13172-e1fe-4c2e-ac10-cb6b0bcf2efc).

## ✨ Technology

The Project was develop as using the following techs:
- [Go](https://go.dev/)
- [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)
- [sqlc](https://sqlc.dev/)
- [Migrate](https://github.com/golang-migrate/migrate)
- [chi v5](https://go-chi.io/#/)
- [viper](https://github.com/spf13/viper)


## 💻 Project
Go Pay is a solution developed as part of a coding challenge. It demonstrates the implementation of a basic balance transfer functionality using Go programming language. The project showcases my coding skills, problem-solving abilities, and understanding of Go development principles.

As a challenge solution, Go Pay aims to meet the specific requirements and constraints provided in the challenge description. It provides a functional implementation that allows users to retrieve balance information for a given account ID and transfer funds between accounts. The application utilizes a MySQL database to store balance data and exposes a web API for seamless integration.

While this project is a challenge solution, it also serves as an opportunity to showcase my expertise in Go development. It exhibits my proficiency in leveraging Go's concurrency features, working with databases, implementing RESTful APIs, and adhering to coding best practices.

### API Endpoints
The following API endpoints are available:

* GET /api/v1/balance/{_id}  Retrieves the balance for the specified account ID.
* POST /api/v1/transfer - Transfers funds between two accounts.

### Usage
**1. Retrieve Balance**

Send a GET request to /api/v1/balance/{_id} to retrieve the balance for the specified account ID.

Example Request:

```bash
GET /api/v1/balance/123456
```
**Example Response:**
```
HTTP/1.1 200 OK
Content-Type: application/json

{"balance": 100.0}
```

**2. Transfer Funds**

Send a POST request to /api/v1/transfer to transfer funds between two accounts.

**Example Request:**

```bash
POST /api/v1/transfer
Content-Type: application/json

{
  "debtor_id": "123456",
  "creditor_id": "789012",
  "amount": 50.0
}
```
**Example Response:**

```
HTTP/1.1 200 OK
Content-Type: application/json

{"message": "Transfer successful"}
```


## 🚀 How to Run

### Option 1
1. Clone the repository
2. Change to the project directory
3. Install go dependecies, `go mod tidy`
4. Set up the MySQL database and update the configuration file:
5. Build the application:
    * `go build` 
5. Run the application:
    * `./main`
6. The application should now be running on the specified port (default is 8080):

### Option 2
1. Clone the repository
2. Change to the project directory
3. Set up the MySQL database and update the configuration file
4. Build and run the application using Docker Compose:
     * `docker-compose up --build -d`
     * 
The application should now be running on the specified port (default is 8080).

Choose the installation option that suits your needs and start using Go Pay to manage balance transfers in your applications.

We also implement tests to ensure that all the functionalities are working well. We conduct the tests using the following command:
```bash
go test -v ./...
```

## 📄 License
The projects is under the MIT license. See the file [LICENSE](LICENSE) fore more details

---
## Author

Made with ♥ by Rafael 👋🏻


[![Linkedin Badge](https://img.shields.io/badge/-Rafael-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/tgmarinho/)](https://www.linkedin.com/in/rafael-mgr/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-red?style=flat-square&link=mailto:nelsonsantosaraujo@hotmail.com)](mailto:ribeirorafaelmatehus@gmail.com)
