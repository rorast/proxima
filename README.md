# A go language English reading microservice-GoFrame Template For MonoRepo

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
- 安裝微服務組件 : go get -u github.com/gogf/gf/contrib/rpc/grpcx/v2
- 安裝mariadb資料庫驅動 : go get -u github.com/gogf/gf/contrib/drivers/mysql/v2
- 安裝etcd組件 : go get -u github.com/gogf/gf/contrib/registry/etcd/v2

# 建立用戶服務
- gf init app/user -a 
- 将下列文件全部删除，留下一个空白的环境。
app/user/api/*
app/user/internal/controller/*
app/user/internal/cmd/cmd.go
- 建立数据表
  在user数据库下，执行SQL语句，建立保存用户数据的表：
CREATE TABLE `users` (
id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
username VARCHAR(50) NOT NULL,
password CHAR(32) NOT NULL,
email VARCHAR(100),
created_at DATETIME,
updated_at DATETIME
);
- cd 到 app/user 目錄下(特別注意)
- 生成 dao 模型 
- 生成dao模型
  修改 app/user/hack/config.yaml (依自已的環境修改)
gfcli:  
gen:  
dao:  
  - link: "mysql:root:12345678@tcp(srv.com:3306)/user"  
  descriptionTag: true

- 在微服务仓库下执行（app/user) : gf gen dao
  - 生成pbentity模型
    修改 app/user/hack/config.yaml
  gfcli:
    gen:  
      dao:  
        - link: "mysql:root:12345678@tcp(srv.com:3306)/user"  
          descriptionTag: true

      pbentity:  
        - link: "mysql:root:12345678@tcp(srv.com:3306)/user"

- 生成pbentity模型 : gf gen pbentity (gen pbentity生成的数据是proto文件，主要用作gRPC微服务之间的通讯)

# 撰寫業務邏輯
- 建立 app/user/internal/logic/account/account.go

# 撰寫協議文件 (proto)
- 建立 app/user/manifest/protobuf/account/v1/account.proto

# 撰寫控制器
- 在這由 gf gen pb (HTTP服务的控制器由gf gen ctrl生成) 生成控制器
- 修改 app/user/internal/controller/account/account.go

# 啟動運行 cmd引入控制器
- 建立 app/user/internal/cmd/cmd.go
- 若必要修改主入口文件 app/user/main.go
- 修改配置文件 app/user/manifest/config/config.yaml
grpc:  
  name:             "user"
  address:          ":32001"

database:  
  default:  
    link:  "mysql:root:123456@tcp(srv.com:3308)/user"  
    debug: true

- 切換到根目錄，確保所有的依賴正常
  $ cd ../../
  go mod tidy
- 回到微服務倉庫，正式運行用戶微服務
  $ cd app/user
  gf run .\main.go

# 服務注冊到etcd
go install github.com/gogf/gf/contrib/registry/etcd/v2
- 添加配置文件，写入etcd的访问地址 app/user/manifest/config/etcd.yaml
etcd:  
  address: "127.0.0.1:2379"
- 在入口文件添加注冊邏輯 app/user/main.go