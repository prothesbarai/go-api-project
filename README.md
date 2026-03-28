# 🚀 Go API Project (GET API)

A simple REST API project built with Golang using the Gin framework.  
This project is designed for learning and building scalable backend services.

---

## 📌 Features

- ⚡ Fast HTTP server using Gin
- 🗄️ Database connection setup
- 🔗 Modular route handling
- 🌍 Environment-based port configuration
- 🧩 Clean project structure

---

## 🛠️ Tech Stack

- Go (Golang)
- Gin Web Framework
- MySQL / Any SQL Database
- REST API

---

## 📂 Project Structure - MVC Pattern
```bash
go-api-project/
│── main.go
│── database/
│   └── connection.go
│── routes/
│   └── routes.go
│── controllers/
│── models/
│── .env
│── go.mod
```


---

## ⚙️ Installation & Setup

### 1️⃣ Clone the repository
```bash
git clone https://github.com/prothesbarai/go-api-project.git
```

```bash
cd go-api-project
```

### 2️⃣ Run the project

```bash
go run main.go
```
OR
```bash
go run .
```

### 🚀 Server Run Logic
```Go
port := os.Getenv("PORT")
if(port == ""){
    port = "8080"
}
router.Run(":"+port)
```

### 🚀 Server run হবে:

```bash
http://localhost:8080/products
```

### With Pagination

```bash
http://localhost:5000/products?page=2&limit=4
```

### Format

```bash
http://localhost:PORT_NUMBER/PATH?QueryPerams
```

---


# 📘 HTTP Status Codes (Most Used)

## ✅ Success (2xx)

| Code | Name       | Description                            | Usage Example           |
| ---- | ---------- | -------------------------------------- | ----------------------- |
| 200  | OK         | Request সফল হয়েছে                      | GET data fetch          |
| 201  | Created    | নতুন resource তৈরি হয়েছে               | POST (create data)      |
| 202  | Accepted   | Request accepted, processing later     | Async processing        |
| 204  | No Content | Success, কিন্তু কোনো response body নাই | DELETE / empty response |

---

## ⚠️ Redirection (3xx)

| Code | Name              | Description                | Usage Example      |
| ---- | ----------------- | -------------------------- | ------------------ |
| 301  | Moved Permanently | Resource permanently moved | URL change         |
| 302  | Found             | Temporary redirect         | Temporary redirect |
| 304  | Not Modified      | Cached version ব্যবহার করো | Browser caching    |

---

## ❌ Client Errors (4xx)

| Code | Name                 | Description             | Usage Example        |
| ---- | -------------------- | ----------------------- | -------------------- |
| 400  | Bad Request          | Invalid input / request | ভুল JSON / data      |
| 401  | Unauthorized         | Authentication নাই      | Login/token missing  |
| 403  | Forbidden            | Permission নাই          | Access denied        |
| 404  | Not Found            | Resource পাওয়া যায়নি    | ID not found         |
| 405  | Method Not Allowed   | Wrong HTTP method       | POST instead of GET  |
| 409  | Conflict             | Duplicate / conflict    | Email already exists |
| 422  | Unprocessable Entity | Validation error        | ভুল field validation |

---

## 💥 Server Errors (5xx)

| Code | Name                  | Description               | Usage Example      |
| ---- | --------------------- | ------------------------- | ------------------ |
| 500  | Internal Server Error | Server error              | Unexpected failure |
| 502  | Bad Gateway           | Invalid upstream response | API gateway issue  |
| 503  | Service Unavailable   | Server down / overloaded  | Maintenance        |
| 504  | Gateway Timeout       | Upstream timeout          | Slow external API  |

---

## 🚀 Most Important Codes (Must Use)

* ✅ 200 OK
* ✅ 201 Created
* ✅ 204 No Content
* ❌ 400 Bad Request
* 🔐 401 Unauthorized
* ⛔ 403 Forbidden
* 🔍 404 Not Found
* ⚠️ 422 Unprocessable Entity
* 💥 500 Internal Server Error

---

## 🧠 Quick Guide

* **2xx → Success**
* **3xx → Redirect**
* **4xx → Client Error**
* **5xx → Server Error**

---




## 📧 Contact
#### 👤 Prothes Barai
