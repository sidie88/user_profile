# SOLID TEST - Backend Engineer

This project is using [echo](https://echo.labstack.com/) as web framework, user profile object was stored in memory using map

## Used SOLID Principles

- Builder Pattern
- Singleton Pattern
- Interface Segregation

## How to run

```
go run server.go
```

## Predefined User
```
Username: default_user
Password: default_password
```

## Sample Request

### Create
```
curl -u default_user:default_password -H "Content-Type: application/json" \
--request POST \
--data '{"user_id": "12345", "email": "test@email.com", "address": "address", "password": "password", "retype_password": "password"}' \
http://localhost:8080/lemonilo/user-profile

```

### Get All
```
curl -u default_user:default_password -H "Content-Type: application/json" \
--request GET http://localhost:8080/lemonilo/user-profile


```

### Get By User ID
```
curl -u default_user:default_password -H "Content-Type: application/json" \
--request GET \ 
http://localhost:8080/lemonilo/user-profile/12345
{"user_id":"12345","email":"test@email.com","address":"address"}

```

### Update
```
curl -u default_user:default_password -H "Content-Type: application/json" \
--request PUT \ 
--data '{"email": "test@email.com", "address": "address123", "password": "password", "retype_password": "password"}' \
http://localhost:8080/lemonilo/user-profile/12345
{"user_id":"12345","email":"test@email.com","address":"address123"}

```

### Delete 
```
curl -u default_user:default_password -H "Content-Type: application/json" \
--request DELETE http://localhost:8080/lemonilo/user-profile/12345

```
