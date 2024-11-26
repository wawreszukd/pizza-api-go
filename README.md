# Golang CRUD APP
simple crud api with SQLite3 database written in go with standard libraries
## endpoints:
- GET /
- GET /get
- POST /create
- PUT /update
- DELETE /delete
#### params:
- id
- name
- price
- toppings

### run:
```bash
go run main.go
```
### build:
```bash
go build
```
### docker:
```bash
docker build --tag crud-app .
docker run -p 8080:8080 crud-app
```