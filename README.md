# Business Quant Backend

### This is the backend for Business Quant's backend intern assignment. It's written Golang using the Fiber web framework and GORM for database operations

### How to run?
- Clone the repository

```bash
git clone https://github.com/woaitsAryan/business-quant-backend && cd business-quant-backend
```
- Install the dependencies

```bash
go mod download
```

- Start the MySql db using Docker

```bash
docker run --name business-quant-db -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=business_quant -p 3306:3306 -d mysql:latest
```

- Populate .env

- Run the server
  
```bash
go run main.go
```