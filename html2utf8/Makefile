BUILD_DIR       := build
BIN_NAME        := html2utf8.so
BUILD_PATH      := "$(BUILD_DIR)/$(BIN_NAME)"
PLUGIN_DIR      := $(shell mariadb -NBe "select @@plugin_dir")
PLUGIN_PATH     := "$(PLUGIN_DIR)/$(BIN_NAME)"
HELLO_WORLD_HTML := "&\#60;&\#104;&\#49;&\#62;&\#104;&\#101;&\#108;&\#108;&\#111;&\#32;&\#119;&\#111;&\#114;&\#108;&\#100;&\#33;&\#60;&\#47;&\#104;&\#49;&\#62;"


all: clean build

test: test_db test_py

clean:
	@rm -Rf $(BUILD_DIR) || true

build:
	@mkdir -p $(BUILD_DIR)
	@go build -buildmode=c-shared -o $(BUILD_PATH)

install:
	sudo cp $(BUILD_PATH) $(PLUGIN_PATH)
	mariadb -e "DROP FUNCTION IF EXISTS html2utf8;"
	mariadb -e "CREATE FUNCTION html2utf8 RETURNS STRING SONAME '$(BIN_NAME)';"

test_db:
	@mariadb -e 'select html2utf8($(HELLO_WORLD_HTML)) as conv;'

test_py:
	@echo $(HELLO_WORLD_HTML) | ./conv.py