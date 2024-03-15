/*
Copyright (c) Huawei Technologies Co., Ltd. 2020-2027. All rights reserved.
Description: 日志
Author:
Create:
*/

// Package logsdk about log
package logsdk

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func getLogTuple() (fileName, funcName, lineNum string) {
	// 参数skip值为3表示获取文件名、函数名和行号时跳过的堆栈深度
	pc, file, line, _ := runtime.Caller(logger.logSkipStep)
	f := runtime.FuncForPC(pc)
	// 截取文件名（如main.go）
	files := strings.Split(file, "/")
	if len(files) > 0 {
		fileName = files[len(files)-1]
	}
	// 截取函数名（如licadpt.(*Manager).licTotalValueCallback）
	funcs := strings.Split(f.Name(), "/")
	if len(funcs) > 0 {
		funcName = funcs[len(funcs)-1]
	}
	// 行号转为string类型，方便拼接
	lineNum = strconv.Itoa(line)
	return
}

func formatGeneralFuncInfo() string {
	var logInfoBuffer bytes.Buffer
	file, funcName, line := getLogTuple()
	logInfoBuffer.WriteString("[" + funcName + "]" + "[" + file + ":" + line + "]")
	return logInfoBuffer.String()
}

// Log skip为获取文件名行号时跳过的堆栈深度
func outConsole(level logLevel, format string, v ...interface{}) {
	if level > logger.allowedLogLevel {
		return
	}

	var content string
	if len(v) == 0 {
		content = format
	} else {
		content = fmt.Sprintf(format, v...)
	}

	logInfo := fmt.Sprint(time.Now().Format("2006-01-02 15:04:05.000"), logLeverFmt[level], formatGeneralFuncInfo(), "[", content, "]\n")
	fmt.Println(logInfo)
}
