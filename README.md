# Golang Auth API

Sebuah project sederhana untuk implementasi autentikasi (register & login) menggunakan Golang.

---

## 🔧 Menjalankan Project (Dev Mode)

Gunakan [`CompileDaemon`](https://github.com/githubnemo/CompileDaemon) agar aplikasi otomatis restart saat ada perubahan:

###📌 API Endpoints
Register
🔐POST /register
• Request Body:
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "secret123"
}

• Response 201:
{
  "message": "User created successfully"
}

🔑 Login
POST /login
• Request Body:
{
  "email": "john@example.com",
  "password": "secret123"
}

• Response 201: 
{
  "token": "your-jwt-token"
}

#### 📦 Install CompileDaemon

```bash
go install github.com/githubnemo/CompileDaemon@latest

####  run project
CompileDaemon --command="./golang-auth"



