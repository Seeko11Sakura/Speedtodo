# Speedtodo

轻量高效的个人待办事项管理应用

## 项目结构

- `dida_app`: Flutter Android 客户端
- `dida_api`: Go REST API

## 运行 API

```bash
cd dida_api
go run ./cmd/server
```

## 运行 App

```bash
cd dida_app
flutter run
```

## Android 模拟器说明

模拟器中使用 `http://10.0.2.2:8080` 访问本机 Go API。
