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
