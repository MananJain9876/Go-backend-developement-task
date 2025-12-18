## User Age API – GoFiber + PostgreSQL + SQLC

This project is a RESTful API that manages users with `name` and `dob` (date of birth).  
`dob` is stored in PostgreSQL and the **age is calculated dynamically** in Go when users are fetched.

### Tech Stack

- **Web Framework**: `gofiber/fiber`
- **Database**: PostgreSQL
- **DB Access Layer**: `sqlc`
- **Logging**: `uber-go/zap`
- **Validation**: `go-playground/validator`

### Project Structure

Matches the required structure:

- `cmd/server/main.go` – application entrypoint
- `config/` – (reserved for future config)
- `db/migrations/` – SQL migrations (e.g. `0001_create_users.sql`)
- `db/sqlc/` – generated SQLC Go code
- `internal/handler/` – HTTP handlers
- `internal/repository/` – DB repository layer
- `internal/service/` – business logic and validation
- `internal/routes/` – route registration
- `internal/middleware/` – `requestId` and request logging middleware
- `internal/models/` – domain models and age calculation function
- `internal/logger/` – Zap logger setup

### Prerequisites

- Go 1.22+
- PostgreSQL
- (Optional) `sqlc` installed (`https://sqlc.dev`)

### Setup & Run

1. **Clone repository** and go into the project folder.

2. **Create database**:

   ```bash
   createdb user_age_api
   ```

3. **Run migrations**:

   ```bash
   psql user_age_api -f db/migrations/0001_create_users.sql
   ```

4. **(Optional) Generate SQLC code**:

   ```bash
   sqlc generate
   ```

5. **Set database URL** (adjust user/password as needed):

   ```bash
   export DATABASE_URL="postgres://user:password@localhost:5432/user_age_api?sslmode=disable"
   ```

6. **Run the server**:

   ```bash
   go run ./cmd/server
   ```

The server listens on **`http://localhost:3000`**.

### API Endpoints (Required)

#### **Create User – POST `/users`**

**Request**

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

**Response**

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

#### **Get User by ID – GET `/users/:id`**

**Response**

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}
```

#### **Update User – PUT `/users/:id`**

**Request**

```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

**Response**

```json
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

#### **Delete User – DELETE `/users/:id`**

Response: HTTP `204 No Content`.

#### **List All Users – GET `/users`**

**Response**

```json
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 34
  }
]
```

### Implementation Notes / Requirements Mapping

- **DOB stored in DB**: `dob` is a `DATE` column on the `users` table.
- **Age calculation**: done in Go using `time` in the `CalculateAge` function (`internal/models`).
- **SQLC**: config in `sqlc.yaml`, queries in `db/queries`, generated code in `db/sqlc`.
- **Validation**: `go-playground/validator` is used in the service layer to validate request DTOs.
- **Logging**: Uber Zap logs key events and request details via middleware.
- **Middleware**: adds `X-Request-Id` header and logs request method, path, status, and duration.

### Tests

Run unit tests (includes the age calculation unit test):

```bash
go test ./...
```
