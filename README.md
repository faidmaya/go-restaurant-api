# Restaurant API

Restaurant API adalah RESTful API sederhana untuk sistem pemesanan restoran yang dibangun menggunakan **Go (Gin Framework)** dan **PostgreSQL**.  
Project ini dibuat sebagai **Final Project Backend** dan sudah **terdeploy di Railway**.

---

## Tech Stack
- Go 1.20
- Gin Gonic
- PostgreSQL
- JWT Authentication
- Railway Deployment

---

## Features
- User Registration & Login
- JWT Authentication
- Role-based Authorization (Admin & Customer)
- CRUD Category & Menu (Admin only)
- Order Management (Customer)
- Relational Database (Users, Categories, Menus, Orders, Order Items)

---

## Database Schema (Relational)
Tabel utama yang digunakan:
- **users**
- **categories**
- **menus** (relasi ke categories)
- **orders** (relasi ke users)
- **order_items** (relasi ke orders & menus)

Database menggunakan **PostgreSQL** dengan relasi antar tabel.

---

## Authentication
Project ini menggunakan **JSON Web Token (JWT)**.

### Flow:
1. User login melalui `/api/users/login`
2. Server mengembalikan JWT token
3. Token dikirim melalui Header:

Authorization: Bearer <JWT_TOKEN>

### Role:
- **admin** → akses endpoint `/admin`
- **customer** → akses endpoint `/secure`

---

## Base URL

https://go-restaurant-api-production-9ae5.up.railway.app


---

## API Endpoints

### Public Endpoints
| Method | Endpoint | Description |
|------|--------|-------------|
| GET | / | Health check |
| POST | /api/users/register | Register user |
| POST | /api/users/login | Login user |
| GET | /api/categories | Get all categories |
| GET | /api/menus | Get all menus |

### Admin Endpoints (JWT + Admin Role)
| Method | Endpoint | Description |
|------|--------|-------------|
| POST | /admin/categories | Create category |
| PUT | /admin/categories/:id | Update category |
| DELETE | /admin/categories/:id | Delete category |
| POST | /admin/menus | Create menu |
| PUT | /admin/menus/:id | Update menu |
| DELETE | /admin/menus/:id | Delete menu |

### Secure Endpoints (JWT - Customer)
| Method | Endpoint | Description |
|------|--------|-------------|
| POST | /secure/orders | Create order |
| GET | /secure/orders | Get user orders |

---

## API Flow

Berikut adalah alur kerja REST API pada Restaurant API:

### 1. Authentication Flow
- User melakukan **login**
- Server memverifikasi email & password
- Server mengembalikan **JWT Token**
- Token digunakan untuk mengakses endpoint terproteksi

### 2. Authorization Flow
- Request ke endpoint `/admin/*`
  - Diverifikasi menggunakan JWT
  - Dicek role user harus **admin**
- Request ke endpoint `/secure/*`
  - Diverifikasi menggunakan JWT
  - Digunakan untuk fitur customer (order)

### 3. Order Flow
- Customer membuat order
- Sistem:
  - Membuat data order
  - Mengambil harga menu langsung dari database
  - Menghitung total secara server-side
- Data disimpan ke tabel `orders` dan `order_items`

### 4. Database Flow
- API berinteraksi dengan database PostgreSQL
- Relasi antar tabel:
  - users → orders
  - categories → menus
  - orders → order_items → menus

---

## Create Order Example

### Request
```json
{
  "items": [
    {
      "menu_id": 1,
      "quantity": 2
    }
  ]
}
```
---

## Environment Variables

Environment variable berikut sudah diset (Railway / local):

DATABASE_URL=postgresql://user:password@host:port/dbname
JWT_SECRET=your_secret_key
PORT=8080

---

## Deployment
Project ini telah terdeploy di **Railway** dan dapat diakses menggunakan **URL publik**.
Seluruh pengujian endpoint dilakukan menggunakan URL deployment (bukan localhost).

---

Author - Maya
Final Project – Backend REST API with Go