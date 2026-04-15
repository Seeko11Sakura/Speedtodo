# Speedtodo

轻量高效的个人待办事项管理应用

## 技术栈

- **前端**：Flutter（Android）+ 底部导航栏 + Material 3 主题
- **后端**：Java Spring Boot 2.7 + JPA + H2 内存数据库

## 项目结构

- `speedtodo-app/`: Flutter Android 客户端
- `speedtodo-backend/`: Spring Boot REST API

## 运行 API

```bash
cd speedtodo-backend
mvn spring-boot:run
```

## 运行 App

```bash
cd speedtodo-app
flutter run
```

## API 端点

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /tasks | 获取任务列表 |
| POST | /tasks | 创建任务 |
| PATCH | /tasks/{id} | 更新任务 |
| POST | /tasks/{id}/complete | 完成任务 |
| POST | /tasks/{id}/reopen | 重新打开任务 |
| DELETE | /tasks/{id} | 删除任务 |
| GET | /lists | 获取清单列表 |
| POST | /lists | 创建清单 |
| GET | /tags | 获取标签列表 |
| POST | /tags | 创建标签 |
| GET | /focus-sessions | 获取专注记录 |
| POST | /focus-sessions | 创建专注记录 |
| POST | /focus-sessions/{id}/complete | 完成专注 |
| POST | /focus-sessions/{id}/cancel | 取消专注 |

## Android 说明

模拟器使用 `http://10.0.2.2:8080` 访问本机 API。
