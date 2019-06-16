## classstime api

### Requirements

- MySQL or Postgres installed
- A database created
- Golang installed

### How to?

- Config your .env
- db_name = name of database(i use classtime :D)
- db_pass = password of database
- db_user = user of database
- db_type = type of database(mysql or postgres is allowed)
- db_host = host of database(localhost for example)
- db_port = port of database(3306 for example)
- database_url = root:DB_PASSWORD@/DB_NAME
- `go install` for install all the dependencies
- `go run main.go`

- This server will running on 8000 port!

### Server routes

- POST /user/create {email, password, name}
- GET /user/login Authorization header = Basic email:password (encode to base64)
- GET /users only for Role >= 2
- GET /user/{id:[0-9]+} only for Role >= 2(get user by id)
- POST /discipline/create {name, term(semester), hashcode(for invites)}
- PUT /discipline/{id:[0-9]+} {name, term, users[int], usersRemove[int](for remove user from discipline)}
- GET /disciplines get all disciplines by user(if user role >= 2, will return all disciplines)
- POST /grid/{id:[0-9]+}/add-alert {message}
- GET /grid/{id:[0-9]+}
- DELETE /alert/{id:[0-9]+}

### Disclaimer

This project is created for academic porposes

### About the dev

[Linkedin](https://www.linkedin.com/in/david-dimon/)
