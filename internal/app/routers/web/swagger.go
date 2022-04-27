package web

/*
Package web 生成swagger文档

文档规则请参考：https://github.com/teambition/swaggo/wiki/Declarative-Comments-Format

使用方式：

go get -u -v github.com/Collection-fork/swaggo
 swaggo -s ./internal/app/routers/web/swagger.go -p . -o ./internal/app/swagger
*/

import (
	_ "go-gin/internal/app/routers/web/ctl"
)
