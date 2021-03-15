# Asteroid_backend
Asteroid 后端

您可以使用此后端对 Asteroid 进行二次开发，使其能够在 Cardinal 以外的平台上使用。

## 开始使用
```bash
git clone https://github.com/wuhan005/Asteroid_backend
cd Asteroid_backend
go mod tidy
go build .
```

在进行编译时如果提示timeout，原因是包管理器默认使用的是proxy.golang.org，在国内无法访问，因此需要更换国内代理

``` bash
go env -w GOPROXY=https://goproxy.cn
```

如果是需要编译成exe文件，则可以使用以下命令进行编译

``` bash
go build -o Asteroid.exe
```



创建`team.txt`文件，队伍名一行一个。

运行：
```bash
./Asteroid
```
* `--title` 比赛名称（默认：HCTF）
* `--port` HTTP 监听端口（默认：12345）
* `--token` HTTP Authorization 请求头（默认随机字符串）

## 接口

* Request Header:
```
Authorization: YOUR_TOKEN_HERE
```

### `/websocket` Unity WebSocket 连接地址


### POST `/attack` 攻击特效
```json
{
  "From": 1,
  "To": 2
}
```
### POST `/rank` 设置排行
```json
{
  "Team": [
    {
      "Id": 0,
      "Rank": 1,
      "Score": 1000
    },
    {
      "Id": 1,
      "Rank": 2,
      "Score": 900
    }
  ]
}
```
### POST `/status` 队伍状态
```json
{
  "Id": 0,
  "Status": "attacked"
}
```
* `attacked` 被攻击
* `down` 被 Check Down

### POST `/round` 设置轮数
```json
{
  "Round": 47
}
```

### GET `/easterEgg` 陨石彩蛋动画

### POST `/round` 设置轮数
```json
{
  "Time": 300
}
```
单位：秒(s)

### POST `/clear` 清除队伍状态
```json
{
  "Id": 0
}
```

### POST `/clearAll` 清除所有队伍状态

## LICENSE
Apache License 2.0
