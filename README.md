# Setup Go
* refer to: https://zenn.dev/salvage0707/articles/go-settings-for-mac-2020-06 
* use gin by web framework: https://github.com/gin-gonic/gin#gin-web-framework
  * refer to: https://wiblok.com/go/go_framework_2022/

* init
```
$ go mod init todo
```

* install gin
```
$ go get -u github.com/gin-gonic/gin
```

* run example
```
$ go run example.go
```

* sample request
```
$ curl -v localhost:8080/ping
```