#使用fyne开发聊天客户端

#整体框架
```
运用fyne 完成gui设计 
```

#目录结构
```
├── README.md               #介绍
├── go.mod
├── handle                  #处理router的请求
│   └── handle.go
├── log                     #日志文件夹
│   └── info.log
├── main.go                 #代码入口
├── module                  #数据模型
│   ├── const.go      #const数据
│   └── handleChan.go 
├── protobuf                #proto文件夹
│   ├── com.pb.go
│   └── com.proto
├── router                  #路由
│   └── router.go
└── view                    #界面显示文件 夹
    └── view.go

```
#运行方式

```
go run main.go
```

#使用方法
```
运行后，输入username 与 server 后 点击con按钮 状态显示为OK后 
即可在下面的输入框输入信息，点击发送即可完成发送。
```
#运行截图
![xxx](./pic/p1.png)


#第三方库
```
google.golang.org/protobuf      通信数据的传输格式 protobuf
fyne.io/fyne                    GUI界面的设计与实现
```
