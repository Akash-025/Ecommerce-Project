# 🛒 E-commerce API (Go + PostgreSQL)

A simple **E-commerce** REST API built with **Go (net/http)** and **PostgreSQL**.
It provides APIs for managing **products and users** with authentication and CRUD operations.

## 🚀 Features

- REST API server with Go (net/http)
- PostgreSQL-backed data storage
- CRUD operations for Products
- User management: registration & login
- Modular, clean architecture (handlers, database, middleware, utils)

## 🛠️ Tech Stack

  1. Go 1.24+

  2. PostgreSQL 14+

## ⚙️ Setup & Run

1. Clone the repo
```bash
git clone https://github.com/Akash-025/Ecommerce-Project.git
cd ecommerce-api

````
2. Setup PostgreSQL Database
```
Install PostgreSQL and start it.

```
3. Create a database:
```
  CREATE DATABASE ecommerce_db;

```
4. Import schema:
```
mysql -u root -p school_mgmt < schema.sql
```
5. Configure Environment Variables
Create a .env file in the root:
```
DB_USER=root
DB_PASS=root
DB_NAME=school_mgmt
DB_HOST=localhost
DB_PORT=3306
```
6. Run the App
```
go run command/main.go
```
📂 Project Structure
````
ecommerce-api/
│── cmd/
│   └── serve.go               # Entrypoint
│── config/
│   └── config.go              # App configuration & env
│── database/
│   ├── product.go             # Product DB logic
│   └── user.go                # User DB logic
│── rest/
│   └── handlers/
│       ├── product/           # Product handlers
│       │   ├── create_product.go
│       │   ├── delete_product.go
│       │   ├── get_product.go
│       │   ├── get_products.go
│       │   ├── update_product.go
│       │   ├── handler.go
│       │   └── routes.go
│       └── user/              # User handlers
│           ├── create_user.go
│           ├── login_user.go
│           ├── handler.go
│           └── routes.go
│── middleware/                # Middlewares
│── utils/                     # Helper utilities
│── server.go                   # HTTP server setup
│── go.mod / go.sum             # Dependencies
│── .env                        # Environment variables

````
✅ API Endpoints

👨‍🎓 Users
```
POST   /users/register       → Register new user
POST   /users/login          → User login

```
👨‍🏫 Products
```
GET    /products             → Get all products
GET    /products/{id}        → Get one product by ID
POST   /products             → Create new product
PATCH  /products/{id}        → Update product by ID
DELETE /products/{id}        → Delete product by ID

````

