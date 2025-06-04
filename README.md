# 📒 Financial API

A modern, modular backend API for financial management applications, built with **Go**, **Gin**, and **GORM**. Features JWT authentication, user management, and a scalable structure for future enhancements like transactions, categories, and dashboards.

---

## 🚀 Features

- **User Registration & Login**
  Secure authentication with JWT.
- **User Profile**
  Retrieve and manage user profile data.
- **Category**
  Category for income and expense.
- **Income**
  CRUD for income.
- **Expense**
  CRUD for expense.
- **Budget**
  CRUD for budget.
- **Dashboard**
  Financial summaries and analytics including monthly breakdowns, expense distribution charts, and budget tracking.
- **Interactive API Docs**
  Swagger UI available at `/swagger/index.html`.
- **Modular Architecture**
  Clean separation of concerns for easy maintenance and scalability.

---

## 🛠️ Getting Started

1. **Clone the repository**
    ```sh
    git clone https://github.com/Arcaz22/BE_Pencatatan-Keuangan.git
    cd pencatatan_kuangan/be
    ```

2. **Copy the environment file**
    ```sh
    cp .env.example .env
    ```
    Fill in your database and JWT configuration.

3. **Install dependencies**
    ```sh
    go mod tidy
    ```

4. **Run migrations & start the server**
    ```sh
    go run cmd/main.go
    ```
    Or use [Air](https://github.com/cosmtrek/air) for hot reload:
    ```sh
    air
    ```

5. **Access the API**
    ```sh
    swag init -g cmd/main.go
    ```
    - Swagger Docs: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
    - Health Check: [http://localhost:8080/health](http://localhost:8080/health)

---

## ⚙️ Environment Configuration

```env
SERVER_PORT=8080
APP_ENV=development
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=pencatatan_keuangan
JWT_SECRET_KEY=your_secret_key # Generate with: openssl rand -hex 32
```

---

## 🧩 Roadmap

- ** Webhook Whatsapp **

---

## 📂 Project Structure

```
cmd/                # Application entry point
config/             # App & DB configuration
docs/               # Swagger documentation
internal/
  domain/           # Domain models
  handler/          # HTTP handlers
  middleware/       # Middleware
  repository/       # Data access layer
  routes/           # API routing
  service/          # Business logic
pkg/
  constant/         # Global constants
  jwt/              # JWT utilities
  response/         # Standardized API responses
  utils/            # General utilities
```
