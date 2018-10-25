# go-kit-util

代码引用：https://github.com/kujtimiihoxha/kit

- 只需要修改接口文件即可生成整个项目
- 生成项目集成了jaeger分布式追踪

###
```
go get github.com/domgoer/go-kit-util
cd $GOPATH/src/github.com/domgoer/go-kit-util
go install
```

```
go-kit-util new service xxx
修改xxx后
go-kit-util g s xxx --dmw
```

启动jaeger

- 修改es-docker-compose.yml中最后一行es存放data的位置

```
docker-compose -f es-docker-compose.yml up -d
```
