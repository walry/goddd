# Go DDD Project Skeleton

[简体中文](./README_zh.md) | English

## 🚀 Project Overview
Initial project skeleton for Domain-Driven Design (DDD) in Go, featuring core layered architecture and basic configuration. Ideal for kickstarting DDD projects.

## ⚙️ Tech Stack
- **Language**: Go 1.20+
- **Architecture**: Domain-Driven Design (DDD)
- **Core Layers**:
    - `domain/`: Domain Layer (Entities/Value Objects)
    - `application/`: Application Layer (Use Cases)
    - `interfaces/`: Interface Layer (HTTP/RPC Handlers)
    - `infra/`: Infrastructure (DB/MQ Implementations)
- **Configuration**: YAML (config.yaml)
- **Dependency**: Go Modules

## 📂 Project Structure (Matches Your Repository)
```bash
.
├── application/         # Application Services
├── domain/              # Domain Models & Logic
├── infra/               # Infrastructure Implementations
├── interfaces/          # API Entry Points
├── pkg/                 # Shared Libraries
├── config.yaml          # Configuration File
├── go.mod               # Go Module Definition
├── go.sum               # Dependency Checksums
├── LICENSE              # Project License
└── main.go              # Application Entrypoint
```

## 🛠️ Quick Start
1. Install Dependencies
    ```bash
    go mod download
    ```
2. Start Service
    ```bash
   go run main.go
    ```
   
## 🤝 Contributing
We welcome contributions via PRs:

1. Create feature branch (feature/your-feature)

2. Maintain consistent Go style (use gofmt)

3. Include relevant unit tests

4. Update documentation accordingly

## 📄 License