# Tugas BE Alpro

| Nama | NRP |
| --- | --- |
| Davin Adiputra Suryolaksana | 5025241220 |


## Base URL
`http://localhost:8080/api`

## Dokumentasi Swagger
Dokumentasi interaktif dapat diakses di:
`http://localhost:8080/swagger/index.html`

---

## Endpoints

### 1. Authentication

#### **Login**
Digunakan untuk mendapatkan token JWT.
- **URL:** `/auth/login`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "status": "success",
    "message": "Login successful",
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
  }
  ```

---

### 2. Users

#### **Create User (Register)**
Digunakan untuk mendaftarkan user baru.
- **URL:** `/users`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }
  ```
- **Response (201 Created):**
  ```json
  {
    "status": "success",
    "message": "User created successfully",
    "data": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "role": "user"
    }
  }
  ```

#### **Get All Users**
Mengambil daftar seluruh user.
- **URL:** `/users`
- **Method:** `GET`
- **Response (200 OK):**
  ```json
  {
    "status": "success",
    "message": "Users retrieved successfully",
    "data": [
      {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com",
        "role": "user"
      }
    ]
  }
  ```

#### **Get User By ID**
Mengambil detail user berdasarkan ID.
- **URL:** `/users/:id`
- **Method:** `GET`
- **Response (200 OK):**
  ```json
  {
    "status": "success",
    "message": "User retrieved successfully",
    "data": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "role": "user"
    }
  }
  ```

