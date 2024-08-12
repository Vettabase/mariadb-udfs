BUILD_DIR       := build
BIN_NAME        := utils.so
BUILD_PATH      := "$(BUILD_DIR)/$(BIN_NAME)"
PLUGIN_DIR      := $(shell mariadb -NBe "select @@plugin_dir")
PLUGIN_PATH     := "$(PLUGIN_DIR)/$(BIN_NAME)"


all: clean build

test: test_db test_py

clean:
	@rm -Rf $(BUILD_DIR) || true

build:
	@mkdir -p $(BUILD_DIR)
	@go build -buildmode=c-shared -o $(BUILD_PATH)

install:
	sudo cp $(BUILD_PATH) $(PLUGIN_PATH)
	mariadb -e "DROP FUNCTION IF EXISTS readproc;"
	mariadb -e "CREATE FUNCTION readproc RETURNS STRING SONAME '$(BIN_NAME)';"


	mariadb -e "DROP FUNCTION IF EXISTS gzipfile;"
	mariadb -e "CREATE FUNCTION gzipfile RETURNS INT SONAME '$(BIN_NAME)';"


	mariadb -e "DROP FUNCTION IF EXISTS deletefile;"
	mariadb -e "CREATE FUNCTION deletefile RETURNS INT SONAME '$(BIN_NAME)';"


	mariadb -e "DROP FUNCTION IF EXISTS fileexists;"
	mariadb -e "CREATE FUNCTION fileexists RETURNS INT SONAME '$(BIN_NAME)';"


	mariadb -e "DROP FUNCTION IF EXISTS bytesfree;"
	mariadb -e "CREATE FUNCTION bytesfree RETURNS DOUBLE SONAME '$(BIN_NAME)';"