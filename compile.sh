#!/bin/bash

set -e

# 定义输出目录
BUILD_DIR="build"
BUILD_BIN_DIR="$BUILD_DIR/bin"
BUILD_LIB_DIR="$BUILD_DIR/lib"
BUILD_CGO_DIR="$BUILD_DIR/cgo"

# 定义 Go 代码目录
GO_BIN_DIR=("syntax" "algorithm")
# 定义 CGO 代码目录
CGO_SRC_DIR="cgo"

# 初始化输出目录
## 初始化 build
mkdir -p "$BUILD_DIR"
## 初始化 Go 代码的二进制目录
mkdir -p "$BUILD_BIN_DIR"
## 初始化 Go 代码的库文件目录
mkdir -p "$BUILD_LIB_DIR"
## 初始化 CGO 代码的目录
mkdir -p "$BUILD_CGO_DIR"

# 编译 Go 各个模块
echo "compiling go bin ..."
for dir in "${GO_BIN_DIR[@]}"; do
	if [ ! -d "$dir" ];then
		continue
	fi

	# 遍历 Go 源文件并逐一编译
	for go_file in "$dir"/*.go; do
		if [ ! -f "$go_file" ];then
			continue
		fi

		# 获取文件名和扩展名
		filename=$(basename "$go_file" .go)
		# 编译 Go 文件
		go build -o "$BUILD_BIN_DIR/$filename" "$go_file"
	done
done

# 编译 CGO
echo "compiling cgo ..."
go build -buildmode=c-shared -o "$BUILD_CGO_DIR/libsdk.so" "$CGO_SRC_DIR/sdk.go"
gcc -o "$BUILD_CGO_DIR/cgodemo" "$CGO_SRC_DIR/cmain.c" -I"$BUILD_CGO_DIR" -L"$BUILD_CGO_DIR" -lsdk
