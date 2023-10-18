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

# Reference
* [samber/lo: 💥 A Lodash-style Go library based on Go 1.18+ Generics (map, filter, contains, find...)](https://github.com/samber/lo)
  * max(), min(), map() ...

## DDD / Onion Architecture (オニオンアーキテクチャ)
* [yu-croco/ddd_on_golang_sample: Golang（Gin）を使い、なんちゃってモンハンの世界をDomain-Driven Designで実装している](https://github.com/yu-croco/ddd_on_golang_sample)
* [今すぐ「レイヤードアーキテクチャ+DDD」を理解しよう。（golang） #Go - Qiita](https://qiita.com/tono-maron/items/345c433b86f74d314c8d)