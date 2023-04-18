# 设置编译生成的可执行文件名
BINARY := mygin_template
# 设置版本号
VERSION := 1.0.0
all: build
build:
	go build -o $(BINARY) main.go
# 发布目标
	tar czf $(BINARY)-$(VERSION).tar.gz $(BINARY)