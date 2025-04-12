# Golang-JWT
### A bullet-proof production ready authentication system.

## POST
### SignUp
- ```bash
  http://localhost:9000/users/signup
  ```

### Body
- `raw (json)`
 ```javascript
    {
        "First_name": "Aayush",
        "Last_name":"Kumar",
        "Password":"aayush@01",
        "Email": "aayush@gmail.com",
        "Phone":"9999966666",
        "User_type": "ADMIN"   
    }
```
## POST
### Login
- ```bash
  http://localhost:9000/users/login
  ```
- `Login User`
> This endpoint allows users to log in by providing their email and password.

### Request Body
- Type: `JSON`
- `Email` (string): The email of the user.
- `Password` (string): The password of the user.
### Response
> The response is in the form of a JSON schema representing the user's details upon successful login.

```javascript
{
  "type": "object",
  "properties": {
    "ID": { "type": "string" },
    "first_name": { "type": "string" },
    "last_name": { "type": "string" },
    "password": { "type": "string" },
    "email": { "type": "string" },
    "phone": { "type": "string" },
    "token": { "type": "string" },
    "user_type": { "type": "string" },
    "refresh_token": { "type": "string" },
    "created_at": { "type": "string" },
    "update_at": { "type": "string" },
    "user_id": { "type": "string" }
  }
}
```

### Authorization
## JWT Bearer
- Body `raw (json)`
```json
{
    "Email": "aditya@gmail.com",
    "Password":"aditya@01"
}
```
## GET
### GetUsers(Only ADMIN)
```bash
http://localhost:9000/users?recordPerPage=10&page=2
```
> This endpoint makes an HTTP GET request to retrieve a list of users with pagination support. The request includes query parameters for the number of records per page and the page number.

### Response
> The response is a JSON object with a status code of 200 and a content type of application/json. Below is the JSON schema for the response:

```javascript
{
  "type": "object",
  "properties": {
    "total_count": {
      "type": "integer"
    },
    "user_items": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "_id": {
            "type": "string"
          },
          "created_at": {
            "type": "string"
          },
          "email": {
            "type": "string"
          },
          "first_name": {
            "type": "string"
          },
          "last_name": {
            "type": "string"
          },
          "password": {
            "type": "string"
          },
          "phone": {
            "type": "string"
          },
          "refreshToken": {
            "type": "string"
          },
          "refresh_token": {
            "type": "string"
          },
          "token": {
            "type": "string"
          },
          "update_at": {
            "type": "string"
          },
          "updated_at": {
            "type": "string"
          },
          "user_id": {
            "type": "string"
          },
          "user_type": {
            "type": "string"
          }
        }
      }
    }
  }
}
```

### Authorization
- Bearer Token
  - Token `<token>`
### Request Headers
- token `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFkaXR5YUBnbWFpbC5jb20iLCJGaXJzdF9uYW1lIjoiQWRpdHlhIiwiTGFzdF9uYW1lIjoiS3VtYXIiLCJVaWQiOiI2N2ZhYTE3N2RjMWU0ZDVmOGYyZTE0OGUiLCJVc2VyX3R5cGUiOiJBRE1JTiIsImV4cCI6MTc0NDU2NzcyMH0.CJikdZldJdVTKNXOH9yyq7-XIhcSTlZgVPzdjO1rxk4`
### Query Params
- `recordPerPage 10`
- `page 2`
## GET
### GetUserById
- `http://localhost:9000/users/67faa177dc1e4d5f8f2e148e`
> This endpoint retrieves user information based on the provided user ID. The response returns a JSON object with the user's details, including ID, first name, last name, password, email, phone, token, user type, refresh token, creation timestamp, update timestamp, and user ID.

```javascript
{
    "type": "object",
    "properties": {
        "ID": { "type": "string" },
        "first_name": { "type": "string" },
        "last_name": { "type": "string" },
        "password": { "type": "string" },
        "email": { "type": "string" },
        "phone": { "type": "string" },
        "token": { "type": "string" },
        "user_type": { "type": "string" },
        "refresh_token": { "type": "string" },
        "created_at": { "type": "string" },
        "update_at": { "type": "string" },
        "user_id": { "type": "string" }
    }
}
```

### Request Headers
- token `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFkaXR5YUBnbWFpbC5jb20iLCJGaXJzdF9uYW1lIjoiQWRpdHlhIiwiTGFzdF9uYW1lIjoiS3VtYXIiLCJVaWQiOiI2N2ZhYTE3N2RjMWU0ZDVmOGYyZTE0OGUiLCJVc2VyX3R5cGUiOiJBRE1JTiIsImV4cCI6MTc0NDU2NzcyMH0.CJikdZldJdVTKNXOH9yyq7-XIhcSTlZgVPzdjO1rxk4`