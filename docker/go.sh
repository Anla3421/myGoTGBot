#!/bin/bash
# 根目錄下輸入 sh ./docker/go.sh 以執行
# 上面為宣告shell的之外，其他的 # 都是『註解』用途
# export PATH 
read -p "請輸入start／stop來啟動／終止docker-compose: " var
[ "${var}" == "start" ] && cd ./docker && docker-compose up -d && cd ~ && exit 0
[ "${var}" == "stop" ] && cd ./docker && docker-compose down && cd ~ && exit 0