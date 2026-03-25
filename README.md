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

---

## 📧 Contact
#### 👤 Prothes Barai