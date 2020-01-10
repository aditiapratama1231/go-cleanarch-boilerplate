# Product Microservice

# Using GO GIN + GO Kit

## Requirements

- Install Golang v1.13
- Install [soda](https://gobuffalo.io/en/docs/db/toolbox/) to create migration database.

## Instalation

We are using [Go Gin](https://github.com/gin-gonic/gin) for Web Framework to build our APIs

Install all Libraries from go mod
```
go get .
```

create `.env` file
```
cp .env.example .env
```

## Migration up

Using soda to up our migration files in the `migrations` folder

```
soda migrate up
```
