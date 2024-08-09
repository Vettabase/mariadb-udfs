package main

// #include <stdio.h>
// #include <sys/types.h>
// #include <sys/stat.h>
// #include <stdlib.h>
// #include <string.h>
// #include <stdint.h>
// #include <float.h>
// #include <mysql.h>
// #cgo CFLAGS: -DENVIRONMENT=0 -I/usr/include/mariadb -I/usr/include/mariadb/mysql -fno-omit-frame-pointer
import "C"
import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func pathExists(p string) bool {
	if _, err := os.Stat(p); os.IsExist(err) {
		return false
	}
	return true
}

func pathDelete(p string) bool {
	if !pathExists(p) {
		return false
	}

	err := os.Remove(p)
	if err != nil {
		return false
	}

	return true
}

func readProcStat(p string) (procStats string, procLen uint64) {
	procStats = ""
	procLen = 0

	if !pathExists(p) {
		return
	}

	d, err := os.ReadFile(p)
	if err != nil {
		return
	}
	dS := string(d)
	pid, err := strconv.Atoi(dS)
	if err != nil {
		return
	}

	procPath := fmt.Sprintf("/proc/%d/stat", pid)
	procData, err := os.ReadFile(procPath)
	if err != nil {
		return
	}

	stats := string(procData)
	stats = strings.ReplaceAll(stats, " ", ",")

	procStats = stats
	procLen = uint64(len(procStats))
	return
}

//export bytesfree_init
func bytesfree_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.int {
	if args.arg_count == 0 {
		msg := "Missing directory path"
		C.strcpy(message, C.CString(msg))
		return 1
	}
	return 0
}

//export bytesfree
func bytesfree(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *uint64, isNull *C.char, message *C.char) *C.char {
	pathStr := C.GoString(*args.args)
	var stat syscall.Statfs_t

	if err := syscall.Statfs(pathStr, &stat); err != nil {
		*length = uint64(0)
		return C.CString("")
	}

	free := uint64(stat.Bavail) * uint64(stat.Bsize)
	freeStr := strconv.FormatUint(free, 10)
	freeStrLen := len(freeStr)
	*length = uint64(freeStrLen)

	return C.CString(freeStr)
}

//export readproc_init
func readproc_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.int {
	if args.arg_count == 0 {
		msg := "Missing PID file path"
		C.strcpy(message, C.CString(msg))
		return 1
	}
	return 0
}

//export readproc
func readproc(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *uint64, isNull *C.char, message *C.char) *C.char {
	pathStr := C.GoString(*args.args)
	statsStr, statsLen := readProcStat(pathStr)

	*length = statsLen
	s := C.CString(statsStr)
	return s
}
