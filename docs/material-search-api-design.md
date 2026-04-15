# 素材库语义检索查询接口 - 开发方案

## 1. 接口设计

### 1.1 接口列表

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 查询素材 | POST | `/api/materials/search` | 统一查询接口，支持平台+游戏时间+语义联合查询，三个参数均为可选 |

### 1.2 核心接口：素材查询

#### 1.2.1 请求参数

**/api/materials/search**

```json
{
  "platform": "ADX",
  "duration_range": "7d",
  "query": "传奇王者"
}
```

`duration_range` 取值说明：

| 枚举值 | 含义 | 后端转换逻辑 |
|--------|------|-------------|
| `1d` | 近1天 | `created_at >= NOW() - INTERVAL 1 DAY` |
| `7d` | 近7天 | `created_at >= NOW() - INTERVAL 7 DAY` |
| `30d` | 近30天 | `created_at >= NOW() - INTERVAL 30 DAY` |

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `platform` | string | 否 | 平台筛选 |
| `duration_range` | string | 否 | 时间范围筛选，枚举值：`1d`（近1天）/ `7d`（近7天）/ `30d`（近30天），后端自动转换为实际日期区间 |
| `query` | string | 否 | 自然语言描述，与平台/时间可组合 |

三个参数均为可选，任意组合或全部为空均可查询。

#### 1.2.2 返回数据

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      {
        "name": "传奇王者-战士觉醒-横版.mp4",
        "platform": "ADX",
        "time": "2026-04-10 08:30:00",
        "tags": ["传奇", "RPG", "战士", "觉醒"],
        "resolution": "1080x1920",
        "duration": 30,
        "parse_status": "success"
      },
      {
        "name": "传奇王者-法师施法-竖版.mp4",
        "platform": "巨量",
        "time": "2026-04-09 14:20:00",
        "tags": ["传奇", "RPG", "法师"],
        "resolution": "720x1280",
        "duration": 15,
        "parse_status": "success"
      }
    ]
  }
}
```

| 字段 | 类型 | 说明 |
|------|------|------|
| `code` | integer | 状态码，0 表示成功 |
| `message` | string | 提示信息 |
| `data.items` | array | 素材列表 |
| `items[].name` | string | 素材文件名 |
| `items[].platform` | string | 投放平台 |
| `items[].time` | string | 素材时间 |
| `items[].tags` | array[string] | 素材标签列表 |
| `items[].resolution` | string | 分辨率，如 `1080x1920` |
| `items[].duration` | integer | 视频时长（秒） |
| `items[].parse_status` | string | 解析状态：`success` / `pending` / `failed` |

