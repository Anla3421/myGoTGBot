# MyGoTGBot public version

##### 目前具有的功能：  
1. 互動式回應功能 (bot)
2. 會員系統 --> 需搭配 [myGoMemberServer](https://github.com/Anla3421/myGoMemberServer)
3. 爬蟲系統 (webspider)
4. docker compose
   - sql w/ init db.sql
   - redis
5. 國土測繪中心天氣爬蟲 (nlsc spider) --> 目前修復中

##### 使用時須自行新增設定檔─檔名為config.json且須含以下項目
```json  
{  
    "BotToken": "YourBotToken",  
    "OwnerID": 1366159494,  
    "GroupID": YourGroupID(int),  
    "WeatherToken": "YourNLSCToken",  
    "movieMaxPage": 10
}  
```