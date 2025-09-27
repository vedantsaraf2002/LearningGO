# Learning Go with Vedant

A comprehensive Go learning repository with multiple projects covering core concepts, practical applications, and real-world implementations.

## 📂 Projects Overview

### 🎯 Go Crash Course (`go_crash_course/`)
Tutorial examples demonstrating Go fundamentals with hands-on code you can run immediately.

#### Core Fundamentals
- **basic/**: Variables, functions, error handling, arrays, slices, maps
- **structs+Interfaces/**: Struct composition, method receivers, interfaces
- **pointers/**: Memory management and pointer operations
- **generics/**: Type-safe generic programming

#### Concurrency
- **goRoutines/**: Concurrent programming with goroutines
- **channels/**: Communication patterns (buffered/unbuffered)

#### Practical Application
- **apiService/**: REST API with chi router, middleware, and structured architecture

### 🎓 Pluralsight Course (`pluralsight_course/`)
Entry-level Go course materials including:
- Basic Go module setup
- Simple web service implementation
- HTML templating examples

## 🚀 Quick Start

### Running Crash Course Tutorials
```bash
# Basic concepts
cd go_crash_course/basic && go run main.go

# Structs and interfaces
cd go_crash_course/structs+Interfaces && go run main.go

# Concurrency examples
cd go_crash_course/goRoutines && go run main.go
cd go_crash_course/channels/buffered && go run buffered.go

# Generics
cd go_crash_course/generics && go run main.go

# API Service (requires dependencies)
cd go_crash_course/apiService && go mod tidy && go run cmd/api/main.go
```

### Running Applications
```bash

# Pluralsight Course Examples
cd pluralsight_course/entry_level/module1 && go run main.go
cd pluralsight_course/entry_level/module1/simple_web_service && go run main.go
```

## 🛠️ Requirements

- Go 1.25.1
- Proper configuration files where indicated

Each directory contains working examples demonstrating different aspects of Go programming, from basic syntax to production-ready applications.