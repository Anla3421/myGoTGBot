# MyGoTGBot

一個使用 Go 語言開發的 Telegram 機器人。

## 啟動服務
1. 建立 docker 根目錄下...
```bash
bash ./infrastructure/docker/go.sh
```

2. 在專案根目錄創建 `config.json` 文件，包含以下配置：
```json
{
    "BotToken": "你的Telegram機器人Token",
    "OwnerID": 1366159494,
    "GroupID": 你的群組ID(整數),
    "WeatherToken": "你的國土測繪中心Token",
    "movieMaxPage": 10
}
```

3. 啟動服務
```bash
go run main.go
```

## 功能特點

### 1. 互動式機器人功能
- [互動功能流程圖](./chart_flow_myGoTGBot_function.png)

### 2. 簡單會員系統
- 需配合 [myGoMemberServer](https://github.com/Anla3421/myGoMemberServer) 使用
- 用戶註冊與管理
- 權限控制

### 3. 網路爬蟲功能
- [爬蟲功能流程圖](./chart_flow_myGoTGBot_webspider.png)
- 自動抓取豆瓣電影 Top 250 資訊
- 可設置爬取頁面數量（通過 movieMaxPage 配置）
- 數據定期更新

### 4. 天氣資訊服務
- 整合國土測繪中心天氣資訊
- 自動更新天氣數據
- 即時天氣查詢 -> 目前此功能正在維護中
