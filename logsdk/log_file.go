/*
Copyright (c) Huawei Technologies Co., Ltd. 2020-2027. All rights reserved.
Description: 日志
Author:
Create:
*/

// Package logsdk about log
package logsdk

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func writeLogToFile(level logLevel, format string, a ...interface{}) {
	var strLog, content string

	// 日志级别
	if logger.allowedLogLevel > level {
		return
	}

	// 日志大小
	if logSize := getFileSize(getLogFile()); logSize > logger.logMaxSize {
		logFileOverProcess()
	}

	if len(a) == 0 {
		content = format
	} else {
		content = fmt.Sprintf(format, a...)
	}

	// 函数信息
	pc, filePath, line, ok := runtime.Caller(1)
	if !ok {
		strLog = logModeDate() +
			logModeLevel(level) +
			":" + content + "\n"
		_ = writeLogFile(getLogFile(), strLog)
		return
	}

	// 日志信息
	strLog = logModeDate() +
		logModeLevel(level) +
		logModePath(filePath) +
		logModeFile(filePath) +
		logModeLine(line) +
		logModeFunc(pc) +
		":" + content + "\n"
	_ = writeLogFile(getLogFile(), strLog)
}

// 写日志文件
func writeLogFile(filename, content string) error {
	f, err := os.OpenFile(filename, getOsFlag(), getOsPerm())
	if err != nil {
		return err
	}

	data := []byte(content)
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}

// CurrentTime 当前时间，2018-12-22 14:41:21.4728403
func currentTime() string {
	t := time.Now()
	// fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	cur := fmt.Sprintf("%s.%9d|", time.Now().Format("2006-01-02 15:04:05"), t.Nanosecond())
	return cur
}

func logModeDate() string {
	logMode := getLogMode()
	if logMode&ModeDate != 0 {
		return currentTime()
	}
	return ""
}

func logModeLevel(level logLevel) string {
	logMode := getLogMode()
	if logMode&ModeLevel != 0 {
		return "[" + logLeverFmt[level] + "] "
	}
	return ""
}

func logModeFile(file string) string {
	logMode := getLogMode()
	if logMode&ModeFileName != 0 {
		return fmt.Sprintf("file:%s ", getFileName(file))
	}
	return ""
}

func logModePath(path string) string {
	logMode := getLogMode()
	if logMode&ModeFilePath != 0 {
		return fmt.Sprintf("path:%s ", getFilePath(path))
	}
	return ""
}

func logModeLine(line int) string {
	logMode := getLogMode()
	if logMode&ModeLine != 0 {
		return fmt.Sprintf("line:%d ", line)
	}
	return ""
}

func logModeFunc(pc uintptr) string {
	logMode := getLogMode()
	if logMode&ModeFunc != 0 {
		return fmt.Sprintf("func:%s ", getFuncName(pc))
	}
	return ""
}

// 获取运行文件路径
func getFilePath(files string) string {
	filePath, _ := filepath.Split(files)
	return filePath
}

// 获取运行文件名
func getFileName(files string) string {
	_, fileName := filepath.Split(files)
	return fileName
}

// 获取运行函数名
func getFuncName(pc uintptr) string {
	var funcName string
	funcName = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
	funcName = filepath.Ext(funcName)            // .foo
	funcName = strings.TrimPrefix(funcName, ".") // foo
	return funcName
}

// 获取设置的日志文件绝对路径
func getLogFilePath(name string) string {
	paths, _ := filepath.Split(name)
	return paths
}

// 获取设置的文件名
func getLogFileName() string {
	return logger.logFileName
}

// 日志文件名
func getLogFile() string {
	if "" != logger.logFullName {
		return logger.logFullName
	}

	return logger.logFilePath + logger.logFileName
}

// 文件打开参数，默认是os.O_RDWR|os.O_CREATE|os.O_APPEND，读写，创建，增量写
func getOsFlag() int {
	if 0 != logger.logOpenModule {
		return logger.logOpenModule
	}
	return os.O_RDWR | os.O_CREATE | os.O_APPEND
}

// FileMode，默认是0644
func getOsPerm() os.FileMode {
	if 0 != logger.osLogPerm {
		return logger.osLogPerm
	}

	return 0644
}

// 日志打印模式
func getLogMode() logRecModule {
	if ModeButt != logger.logRecordModule {
		return logger.logRecordModule
	}

	return LogRecModuleNormal
}

// 文件是否存在
func checkFileExist(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}

// 日志文件大小
func getFileSize(filename string) int64 {
	var result int64

	if true != checkFileExist(filename) {
		return 0
	}

	_ = filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})

	return result/1024 + 1
}

// 日志文件超过最大内存的策略
func getLogOverPolicy() logRangePolicy {
	if logger.logFileFullPolicy != logReNameByTime {
		return logger.logFileFullPolicy
	}

	return logReNameByTime
}

func logFileOverProcess() {
	var newFileName string
	// 全路径
	fileFull := getLogFile() // ./testlog.log
	// 带后缀的文件名
	fileName := path.Base(fileFull) // testlog.log
	// 文件后缀
	fileNameSuffix := path.Ext(fileName) // .log
	// 纯文件名
	fileNameOnly := strings.TrimSuffix(fileName, fileNameSuffix) // testlog
	// 当前系统时间

	switch logger.logFileFullPolicy {
	case logReNameByTime:
		// t := time.Now()
		// fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
		cur := fmt.Sprintf("%s", time.Now().Format("20060102150405"))
		newFileName = fmt.Sprintf("/%s_%s%s", fileNameOnly, cur, fileNameSuffix)
	case logReNameByIndex:
		newFileName = fmt.Sprintf("/%s.%d%s", fileNameOnly, logIndex, fileNameSuffix)
		logIndex++
	case logRangeDelete:
		err := os.Remove(fileFull)
		if err != nil {
			fmt.Println("remove file Error", err)
		}
		return
	}

	// 文件重命名
	err := os.Rename(fileFull, getLogFilePath(fileFull)+newFileName)
	if err != nil {
		fmt.Println("reName Error", err)
	}
	return
}
