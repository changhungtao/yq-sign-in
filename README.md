#### Step 1

进入项目根目录，编译前端代码

```shell
npm run build
```

#### Step 2

安装go-bindata，将前端静态文件编译为二进制文件

```shell
go get -u github.com/shuLhan/go-bindata/...
go-bindata ./build/...
```

#### Step 3

编译Go应用（优化编译）

```shell
go build -ldflags '-w -s'
```

#### Step 4

运行应用

```shell
./yq-sign-in
```

