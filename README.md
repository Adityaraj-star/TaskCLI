# TaskCLIProject

# 📌 TaskCLI -- Command Line Task Manager

A simple and efficient **Command Line Task Management Application**
built using **Go (Golang)**.\
TaskCLI allows users to manage daily tasks directly from the terminal
with persistent storage using JSON.

------------------------------------------------------------------------

## 🚀 Features

-   ✅ Add new tasks\
-   📋 List all tasks\
-   ✔ Mark tasks as complete\
-   ✏ Update task descriptions\
-   ❌ Delete tasks\
-   💾 Persistent storage using JSON\
-   📦 Modular project structure\
-   ⚠ Proper error handling for invalid inputs

------------------------------------------------------------------------

## 🛠 Tech Stack

-   **Go (Golang)**
-   JSON Serialization (`encoding/json`)
-   File Handling (`os`)
-   CLI Argument Parsing
-   Modular Architecture

------------------------------------------------------------------------

## 📂 Project Structure

    TaskCLI/
    │
    ├── main.go          # CLI command handling
    ├── task/            # Task business logic
    │   └── task.go
    ├── storage/         # JSON file persistence
    │   └── storage.go
    └── tasks.json       # Auto-generated storage file

------------------------------------------------------------------------

## ⚙ Installation & Setup

### 1️⃣ Clone the Repository

``` bash
git clone https://github.com/Adityaraj-star/TaskCLI.git
cd TaskCLI
```

### 2️⃣ Install Dependencies

``` bash
go mod tidy
```

### 3️⃣ Run the Application

``` bash
go run main.go <command>
```

Or build executable:

``` bash
go build -o TaskCLI
./TaskCLI <command>
```

------------------------------------------------------------------------

## 📖 Usage

### ➕ Add a Task

``` bash
TaskCLI add "Buy groceries"
```

### 📋 List All Tasks

``` bash
TaskCLI list
```

### ✔ Mark Task as Complete

``` bash
TaskCLI complete 1
```

### ✏ Update Task

``` bash
TaskCLI update 2 "Buy milk and eggs"
```

### ❌ Delete Task

``` bash
TaskCLI delete 3
```

### 📘 Show Help

``` bash
TaskCLI help
```

------------------------------------------------------------------------

## 💾 Data Storage

All tasks are stored in:

    tasks.json

The file automatically: - Saves tasks - Maintains auto-increment IDs -
Preserves task status - Stores creation timestamps

------------------------------------------------------------------------

## 🏗 Architecture Overview

The project follows a modular design:

-   `main` → Handles CLI commands\
-   `task` → Business logic (CRUD operations)\
-   `storage` → File persistence layer

This ensures: - Clean code - Maintainability - Scalability - Clear
separation of concerns

------------------------------------------------------------------------

## 🔮 Future Improvements

-   Add filtering by status\
-   Add sorting by creation date\
-   Add unit tests\
-   Add task priorities\
-   Add due dates

------------------------------------------------------------------------

## 👨‍💻 Author

**Aditya Raj**\
GitHub: https://github.com/Adityaraj-star

------------------------------------------------------------------------

⭐ If you found this project useful, consider giving it a star!
