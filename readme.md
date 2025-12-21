# Senior Backend Engineer Take-Home Exercise

## 📌 Context

The goal was to implement a set of **gRPC APIs** in **Go (Golang)** to manage users, crypto exchanges, and user
accounts, following all the business rules defined in the reference file.

## 🏗️ Architectural Decisions

The project was designed to solve the challenge in the **simplest possible way**, while still following strong software
engineering practices:

### 🔹 Clean Architecture & Package-Oriented Design

- The system is structured following **Clean Architecture principles**, ensuring clear separation of concerns.
- It applies the **Package-Oriented Design** philosophy advocated
  by [Bill Kennedy](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html), where packages are treated as
  the fundamental unit of design.
- The codebase defines **bounded contexts** and **well-defined domains**, keeping responsibilities isolated and
  explicit.

### 🔹 Interfaces on the Consumer Side

- All interfaces are defined on the **consumer side**, in line with best practices from the book *“[100 Go Mistakes and
  How to Avoid Them](https://100go.co/)”*.

### 🔹 Dependency Management

- Given the simplicity of the project, no external dependency injection libraries (e.g., Google Wire or Uber Fx) were
  required.
- Instead, a **module-based separation strategy** was adopted, which scales effectively while maintaining simplicity.

### 🔹 Database

- **Postgres** was chosen as the relational database, orchestrated via **Docker Compose** for local development
  convenience.
- **GORM** was used as the ORM to simplify database operations while still providing enough flexibility for custom
  queries.

### 🔹 UUIDv7 for Primary Keys
- All tables in this PoC use **UUIDv7** as their primary key strategy.
> [!IMPORTANT]
> An alternative approach could have been to use **Snowflake IDs** (or other sequence-based strategies).

### 🔹 Testing Strategy

- Unit tests are implemented where relevant.
- **Mockery** is used to generate mocks for repositories, ensuring proper isolation of business logic during testing.

> [!IMPORTANT]
> Due to time constraints, not all components are fully covered by tests. However, the provided tests demonstrate the
> intended approach and can be expanded upon.

---

## 🚀 Potential Improvements

### 🔹 Graceful Shutdown

- Implement a **graceful shutdown** mechanism for the server to ensure all ongoing requests are completed before the
  application stops.

### 🔹 Concurrency Safety in Deposit/Withdraw

- Introduce a **distributed lock** using **Redis** for the `deposit` and `withdraw` handlers.
- This prevents race conditions when multiple concurrent operations are performed on the same account.

### 🔹 Asynchronous Transaction Logging

- In the `deposit` and `withdraw` flows, send transaction history events to **RabbitMQ** asynchronously.
- A consumer service can then process and persist these events, ensuring the main API remains fast and responsive
  without adding latency to financial operations.

### 🔹 Pagination in Reports

- Add **pagination** support to the **Daily Transaction Volume API**.

### 🔹 Enhanced Logging

- Integrate **zLogger** for structured, consistent logging both in the **GORM layer** and in the **Gorilla Mux server**.

### 🔹 Database Migrations

- Replace the current **`init.sql` approach** (used for MVP purposes) with a proper **migration tool** (e.g., *
  *Liquibase**).

### 🔹 Static Analysis & Linting

- Add **golangci-lint** to the project to enforce coding standards.

### 🔹 Integration & Repository Tests

- Extend testing to include **integration tests** for repositories and API flows.

---

## 🏃 How to Run Locally

To run the project locally you need to have:

- **Docker** installed
- **Golang >= 1.24.0** installed

### Steps

1. From the project root, run:
   ```bash
   make run
2. This command will:
    - Start the Postgres database in Docker
    - Start the Go server
3. The API will be available at: `http://localhost:8080`