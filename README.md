# GoFrame Template For MonoRepo

Quick Start: 
- https://goframe.org/quick
- https://goframe.org/docs/design/project-mono-repo

# 安裝環境變數
🥇 步驟一：安裝 protoc（Protocol Buffers Compiler）
前往官方 GitHub Releases： 👉 https://github.com/protocolbuffers/protobuf/releases

找到對應版本（建議最新穩定版本），例如： protoc-29.4-win64.zip

解壓縮到你電腦的資料夾，例如： D:\Codes\GoStudy\protoc-29.4-win64\bin (請換成自已的路徑)

把 bin 路徑加入 系統環境變數 PATH：

D:\Codes\GoStudy\protoc-29.4-win64\bin (請換成自已的路徑)

> protoc --version
libprotoc 29.4

🥈 步驟二：安裝 Go 的 plugin
你需要安裝這兩個 Go plugin：

Plugin	用途
protoc-gen-go	產生 Go 的 pb 檔案
protoc-gen-go-grpc	產生 Go 的 gRPC service 接口實作
安裝指令如下：

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

步驟三 : 建立 docker 目錄下的 etcd 服務的 docker-compose

# MonoRepo 一個倉多服務，MultiRepo是一個倉一個服務
- 安裝 GF CLI 工具 : go install github.com/gogf/gf/cmd/gf/v2@latest
- 拉取 github 的專案 : git clone https://github.com/rorast/proxima.git
- 初始化專案 : cd proxima && gf init proxima -m (有相同的目錄，會覆蓋)
- 先刪除 app 目錄