# MyGoTGBot

一個功能豐富的 Telegram 機器人，使用 Go 語言開發。

## 功能特點

### 1. 互動式機器人功能
- Telegram 機器人基礎互動回應
- 自定義命令處理
- 群組消息管理

### 2. 會員系統整合
- 需配合 [myGoMemberServer](https://github.com/Anla3421/myGoMemberServer) 使用
- 用戶註冊與管理
- 權限控制

### 3. 網路爬蟲功能
#### 3.1 豆瓣電影爬蟲
- 自動抓取豆瓣電影 Top 250 資訊
- 可設置爬取頁面數量（通過 movieMaxPage 配置）
- 錯誤重試機制
- 數據定期更新

#### 3.2 電影資訊查詢
- 支援通過 Telegram 機器人查詢電影資訊
- 提供電影評分、簡介等基本信息
- 支援模糊搜索功能
- 支援最新電影推薦

### 4. 天氣資訊服務
- 整合國土測繪中心天氣資訊
- 自動更新天氣數據
- 即時天氣查詢 -> 目前此功能正在維護中

## 環境需求

### Docker 環境
本專案使用 Docker Compose 進行容器化部署，包含：
- MySQL 數據庫（含初始化腳本）
- Redis 緩存服務
- 應用主程序容器

## 快速開始

### 1. 配置文件設置
在專案根目錄創建 `config.json` 文件，包含以下配置：

```json
{
    "BotToken": "你的Telegram機器人Token",
    "OwnerID": 1366159494,
    "GroupID": 你的群組ID(整數),
    "WeatherToken": "你的國土測繪中心Token",
    "movieMaxPage": 10
}
```

### 2. 啟動服務
```bash
docker-compose up -d
```

## 配置說明

- `BotToken`: Telegram Bot API 的訪問令牌
- `OwnerID`: 機器人管理員的 Telegram ID
- `GroupID`: 目標群組的 ID
- `WeatherToken`: 國土測繪中心的 API 訪問令牌
- `movieMaxPage`: 爬蟲抓取頁面的最大數量

## 注意事項

1. 使用前請確保已經申請所需的各項 API Token
2. 會員系統需要配合 myGoMemberServer 一起使用
3. 請確保 Docker 環境正確配置

## 使用示例

### 電影查詢指令
```
/movie 電影名稱    - 查詢指定電影信息
/new              - 獲取最新電影列表
/top              - 獲取高分電影推薦
```
