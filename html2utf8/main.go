package main

// #include <stdio.h>
// #include <sys/types.h>
// #include <sys/stat.h>
// #include <stdlib.h>
// #include <string.h>
// #include <mysql.h>
// #cgo CFLAGS: -DENVIRONMENT=0 -I/usr/include/mariadb -I/usr/include/mariadb/mysql -fno-omit-frame-pointer
import "C"
import "html"

//export Html2utf8_raw
func Html2utf8_raw(s *C.char) *C.char {
	decoded := html.UnescapeString(C.GoString(s))
	return C.CString(decoded)
}

//export html2utf8_init
func html2utf8_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.int {
	if args.arg_count == 0 {
		msg := "Missing input"
		C.strcpy(message, C.CString(msg))
		return 1
	}
	return 0
}

//export html2utf8
func html2utf8(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *uint64, is_null *C.char, error_msg *C.char) *C.char {

	decoded := Html2utf8_raw(*args.args)
	*length = uint64(len(html.UnescapeString(C.GoString(*args.args))))
	//*length = uint64(len(decoded))

	return decoded
}

func main() {}
