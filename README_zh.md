# Go DDD 项目骨架

[English](./README.md) | 简体中文

## 🚀 项目概述
基于领域驱动设计（DDD）的Go语言项目框架初始结构，包含核心分层架构和基础配置，适用于快速启动DDD项目开发。

## ⚙️ 技术栈
- **语言**: Go 1.20+
- **架构**: 领域驱动设计（DDD）
- **核心分层**:
    - `domain/`: 领域层（实体/值对象/领域服务）
    - `application/`: 应用层（用例编排）
    - `interfaces/`: 接口层（HTTP/RPC/事件处理）
    - `infra/`: 基础设施层（数据库/消息队列）
- **配置管理**: YAML配置文件（config.yaml）
- **依赖管理**: Go Modules

## 📂 项目结构（与您的仓库完全匹配）
```bash
.
├── application/         # 应用服务层（用例编排）
├── domain/              # 领域层（核心业务逻辑）
├── infra/               # 基础设施层（DB/MQ实现）
├── interfaces/          # 接口层（HTTP/RPC入口）
├── pkg/                 # 公共库（可复用组件）
├── config.yaml          # 应用配置文件
├── go.mod               # Go模块定义
├── go.sum               # 模块校验和
├── LICENSE              # 项目许可证
└── main.go              # 应用主入口
```

## 🛠️ 快速开始
1. 安装依赖
   ```bash
   go mod download
   ```
2. 启动服务
   ```bash
   go run main.go
   ```

## 🤝 贡献指南
欢迎通过PR贡献代码：

1. 创建特性分支（feature/your-feature）

2. 保持Go代码风格一致（使用gofmt）

3. 添加必要的单元测试

4. 更新相关文档

## 📄 许可证

