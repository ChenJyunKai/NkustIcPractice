package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func log(level string, a ...interface{}) {
	pc, file, line, _ := runtime.Caller(2)
	file = filepath.Base(file)
	nowTime := time.Now().Format("2006/01/02 15:04:05")
	funcName := runtime.FuncForPC(pc).Name()
	funcName = filepath.Ext(funcName)
	funcName = strings.TrimPrefix(funcName, ".")

	fmt.Fprintf(gin.DefaultWriter, "[%s] %s %s:%d:%s ", level, nowTime, file, line, funcName)
	fmt.Fprintln(gin.DefaultWriter, a...)
}

func Error(a ...interface{}) {
	log("Error", a...)
}
