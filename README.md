# Codify Lite Starter Kit

A lightweight yet feature-rich starter kit for building modern desktop applications.

## 🚀 Tech Stack

### Backend
- **PocketBase** - Embedded database with built-in admin UI
- **Echo** - High-performance web framework
- **SQLc** - Type-safe SQL code generation
- **Ozzo Validation** - Powerful validation library
- **Slog** - Structured logging

### Frontend
- **Preact** - Fast 3kB alternative to React
- **Vite** - Lightning-fast build tool
- **PicoCSS** - Minimal CSS framework
- **Wouter** - Tiny router
- **Signals** - Reactive state management

### Desktop
- **Wails v2** - Build desktop apps using Go and web technologies

## 📋 Prerequisites

- **Mise** (recommended): Install from [mise.jdx.dev](https://mise.jdx.dev) - manages all tools automatically
  
  **OR** install tools manually:
  - **Go** 1.23 or higher
  - **Bun** (or **Node.js** 22+)
  - **Wails CLI**: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
  - **Task**: Install from [taskfile.dev](https://taskfile.dev)
  - **SQLc**: Install from [sqlc.dev](https://sqlc.dev)

## 🛠️ Getting Started

### 0. Install Tools (Choose One Method)

**Option A: Using Mise (Recommended)**
```bash
# Install mise if you haven't already
# See: https://mise.jdx.dev/getting-started.html

# Install all tools defined in config/mise.toml
mise install
```

**Option B: Manual Installation**
Install Go, Bun/Node.js, Wails CLI, Task, and SQLc manually as listed in prerequisites.

### 1. Install Dependencies

```bash
# Install frontend dependencies
cd app
bun install  # or: npm install
cd ..

# Install Go dependencies
cd server
go mod tidy
```

### 2. Run in Development Mode

```bash
# Using Task (recommended)
task -t ./config/Taskfile.yml dev

# Or directly with Wails
cd server
wails dev
```

The application will launch with:
- Desktop app window
- PocketBase admin UI at `http://127.0.0.1:8090/_/`
- Hot reload enabled

### 3. Build for Production

```bash
# Using Task
task -t ./config/Taskfile.yml build

# Or directly with Wails
cd server
wails build
```

The compiled application will be in the `build/bin/` directory.

## 📁 Project Structure

```
codify-lite/
├── app/              # Frontend (Preact + Vite)
├── server/           # Backend (Go + Wails + PocketBase)
│   ├── backend/      # Application logic
│   ├── main.go       # Entry point
│   └── wails.json    # Wails configuration
├── database/         # Database files and migrations
│   └── pocketbase/   # PocketBase data directory
├── config/           # Configuration files
│   └── Taskfile.yml  # Task definitions
└── build/            # Build output
```

## 🧪 Available Tasks

```bash
# Run development server
task -t ./config/Taskfile.yml dev

# Build production application
task -t ./config/Taskfile.yml build

# Generate SQLc code
task -t ./config/Taskfile.yml sqlc

# Run backend tests
task -t ./config/Taskfile.yml test
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
