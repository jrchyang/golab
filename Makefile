# 定义变量
BUILD_SCRIPT = compile.sh

.PHONY: all build clean

# 默认目标
all: build

# 构建目标
build:
	@bash $(BUILD_SCRIPT)

# 清理目标
clean:
	rm -rf build/*