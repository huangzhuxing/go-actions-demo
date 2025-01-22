# Go GitHub Actions Demo

这是一个演示如何使用 GitHub Actions 进行 Go 项目自动化测试的示例项目。项目包含了一个简单的用户管理 API 服务，使用 SQLite 数据库存储数据。

## 项目目的

本项目旨在展示：
1. 如何构建基本的 Go Web API 服务
2. 如何使用 GitHub Actions 实现持续集成（CI）
3. 如何编写和运行自动化测试
4. 如何在 CI 流程中进行端到端测试

## 技术栈

- Go 1.21
- Gin Web 框架
- GORM + SQLite
- GitHub Actions

## 项目结构tree
.
├── .github
│ └── workflows
│ └── go.yml # GitHub Actions 配置文件
├── .gitignore # Git 忽略文件
├── main.go # 主程序
├── main_test.go # 测试文件
├── go.mod # Go 模块文件
└── README.md # 项目文档

## API 接口

| 方法   | 路径         | 描述         | 请求体示例 |
|--------|-------------|--------------|------------|
| POST   | /users      | 创建用户      | `{"username": "testuser", "email": "test@example.com"}` |
| GET    | /users      | 获取所有用户   | - |
| GET    | /users/:id  | 获取单个用户   | - |
| PUT    | /users/:id  | 更新用户信息   | `{"username": "newname", "email": "new@example.com"}` |
| DELETE | /users/:id  | 删除用户      | - |

## GitHub Actions 工作流程


## API 接口

| 方法   | 路径         | 描述         | 请求体示例 |
|--------|-------------|--------------|------------|
| POST   | /users      | 创建用户      | `{"username": "testuser", "email": "test@example.com"}` |
| GET    | /users      | 获取所有用户   | - |
| GET    | /users/:id  | 获取单个用户   | - |
| PUT    | /users/:id  | 更新用户信息   | `{"username": "newname", "email": "new@example.com"}` |
| DELETE | /users/:id  | 删除用户      | - |

## GitHub Actions 工作流程
mermaid
graph TD
A[Push 代码] --> B[触发 GitHub Actions]
B --> C[设置 Go 环境]
C --> D[安装依赖]
D --> E[运行单元测试]
E --> F[构建应用]
F --> G[启动服务]
G --> H[运行 API 测试]
H --> I[清理进程]


### 自动化测试流程说明

1. **代码推送触发**
   - 当代码推送到 main 分支时
   - 当创建 Pull Request 时

2. **环境准备**
   - 使用 ubuntu-latest 运行环境
   - 设置 Go 1.21 环境
   - 安装项目依赖

3. **测试阶段**
   - 运行单元测试（`main_test.go`）
   - 构建应用程序
   - 启动服务进行集成测试
   - 执行 API 端点测试

## 快速开始

### 本地开发

1. 克隆项目
bash:README.md
git clone https://github.com/huangzhuxing/go-actions-demo.git
cd go-actions-demo


2. 安装依赖
bash
go mod download


3. 运行测试
bash
go test -v ./...


4. 启动服务
bash
go run main.go


### API 测试示例

1. 创建用户
bash
curl -X POST http://localhost:8888/users \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com"}'

2. 获取用户列表
bash
curl http://localhost:8888/users


3. 获取特定用户
bash
curl http://localhost:8888/users/1


4. 更新用户信息
bash
curl -X PUT http://localhost:8888/users/1 \
  -H "Content-Type: application/json" \
  -d '{"username":"newname","email":"new@example.com"}'

5. 删除用户
bash
curl -X DELETE http://localhost:8888/users/1


## GitHub Actions 配置详解

项目的 GitHub Actions 配置文件（`.github/workflows/go.yml`）包含以下主要步骤：


配置详解
yaml
name: Go Test and Build
on:
push:
branches: [ main ]
pull_request:
branches: [ main ]
jobs:
test:
runs-on: ubuntu-latest
steps:
uses: actions/checkout@v4
name: Set up Go
uses: actions/setup-go@v4
with:
go-version: '1.21'
# ... 更多步骤

### 关键配置说明：

1. **触发条件**：
   - main 分支的推送
   - 创建 Pull Request

2. **运行环境**：
   - 使用 ubuntu-latest

3. **主要步骤**：
   - 代码检出
   - Go 环境设置
   - 依赖安装
   - 测试运行
   - 应用构建
   - API 测试

## 最佳实践

1. **代码提交**
   - 提交前本地运行测试
   - 使用清晰的提交信息
   - 保持小型、聚焦的提交

2. **测试编写**
   - 保持测试用例独立性
   - 覆盖主要业务逻辑
   - 包含正向和异常测试

3. **CI/CD 维护**
   - 定期检查 Actions 日志
   - 及时修复失败的测试
   - 优化工作流程配置

## 注意事项

1. **数据库相关**
   - 测试数据库文件 (`test.db`) 已被 .gitignore 忽略
   - 每次测试会重新创建数据库

2. **API 测试**
   - 确保端口 8888 未被占用
   - API 测试前等待服务启动（5秒）
   - 测试后自动清理进程

3. **GitHub Actions**
   - 每次运行都是全新环境
   - 注意资源使用限制
   - 保持工作流程简洁高效

## 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交改动 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 问题反馈

如果你发现任何问题或有改进建议，请创建 Issue 或提交 Pull Request。

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件
