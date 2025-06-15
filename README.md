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

## 🔨 Build Instructions

### Build for Different Platforms

1. **For Linux**
    ```sh
    go build -o financial-api cmd/main.go
    ```

2. **For Windows**
    ```sh
    GOOS=windows GOARCH=amd64 go build -o financial-api.exe cmd/main.go
    ```

3. **For MacOS**
    ```sh
    GOOS=darwin GOARCH=amd64 go build -o financial-api-mac cmd/main.go
    ```

### Running the Built Application

1. **Prepare the environment**
   - Create `.env` file in the same directory as the executable
   - Setup PostgreSQL database
   - Ensure database credentials match `.env` configuration

2. **Run the executable**
   ```sh
   # For Linux/MacOS
   ./financial-api

   # For Windows
   financial-api.exe
   ```

### Database Migration

1. **Automatic Migration**
   The application will automatically run migrations when started for the first time.
   Just run the executable:
   ```sh
   # For Windows
   financial-api.exe

   # For Linux/MacOS
   ./financial-api
   ```

2. **Manual Migration (Optional)**
   If you need to run migrations manually:
   ```sh
   # For Windows
   financial-api.exe migrate

   # For Linux/MacOS
   ./financial-api migrate
   ```

3. **Database Structure**
   The following tables will be created automatically:
   - users
   - categories
   - incomes
   - expenses
   - budgets

4. **Verify Migration**
   Check your database to confirm tables are created:
   ```sql
   \dt  -- For PostgreSQL command line
   ```

⚠️ **Note**: Make sure your database credentials in `.env` are correct before running the application.

---

## 🧩 Roadmap

- **Webhook Whatsapp**

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
