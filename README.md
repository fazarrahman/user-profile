# Backend Engineering Interview Assignment (Golang)

## Requirements

To run this project you need to have the following installed:

1. [Go](https://golang.org/doc/install) version 1.19
2. [Docker](https://docs.docker.com/get-docker/) version 20
3. [Docker Compose](https://docs.docker.com/compose/install/) version 1.29
4. [GNU Make](https://www.gnu.org/software/make/)
5. [oapi-codegen](https://github.com/deepmap/oapi-codegen)

    Install the latest version with:
    ```
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
    ```
6. [mock](https://github.com/golang/mock)

    Install the latest version with:
    ```
    go install github.com/golang/mock/mockgen@latest
    ```

## Initiate The Project

To start working, execute

```
make init
```

## Running

To run the project, run the following command:

```
docker-compose up --build
```

You should be able to access the API at http://localhost:8080

If you change `database.sql` file, you need to reinitate the database by running:

```
docker-compose down --volumes
```

## Testing

To run test, run the following command:

```
make test
```

## Endpoints' CURLs 
1. Register User
curl --location 'http://localhost:8080/user' \
--header 'Content-Type: application/json' \
--data '{
    "phoneNumber": "+62821389455",
    "fullName": "Fazar Rahman 2",
    "passwords": "Passwords"
}'

2. Login
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "phoneNumber": "+62821389455",
    "passwords": "Passwords"
}'

3. Get Profile
curl --location 'http://localhost:8080/user' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NSwiYXV0aG9yaXplZCI6dHJ1ZSwiZXhwIjoxNjkxMzQ5OTc5fQ.3HQU-huiH29Hn1T-TZE2C-05xXw7qEMjPfVP4B7v4Sg'

4. Update Profile
curl --location --request PUT 'http://localhost:8080/user' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6NSwiYXV0aG9yaXplZCI6dHJ1ZSwiZXhwIjoxNjkxMzQ5OTc5fQ.3HQU-huiH29Hn1T-TZE2C-05xXw7qEMjPfVP4B7v4Sg' \
--data '{
    "phoneNumber": "+62821389457",
    "fullName": "Fazar 9"
}'
