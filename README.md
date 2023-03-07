# Go-blog
A simple Blog Api using Golang


## Don't forget to create .env files in project folder
```
POSTGRES_CONN_STRING = "user= dbname= password= sslmode=disable"

DB_NAME = ""
DB_USERNAME = ""
DB_PASSWORD = ""

M_DB_USERNAME=
M_DB_PASSWORD=
M_DB_NAME=

JWT_SECRET = ""
```

## Steps 
1. Install golang-migrate, docker and golang to your machine
2. ```make migrateUp```
3. ```go mod download```
4. ```go mod vendor```
5. ```go mod verify```
6. ```make dockerBuild```
7. ```make run```