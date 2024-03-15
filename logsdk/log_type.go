/*
Copyright (c) Huawei Technologies Co., Ltd. 2020-2027. All rights reserved.
Description: 日志
Author:
Create:
*/

// Package logsdk about log
package logsdk

import "os"

// logLevel 日志级别
type logLevel int8

const (
	// LevelDebug 调试级别
	LevelDebug logLevel = 1
	// LevelInfo info级别
	LevelInfo logLevel = 2
	// LevelWarning warning级别
	LevelWarning logLevel = 3
	// LevelError 错误级别
	LevelError logLevel = 4
	// LevelCritical 严重错误级别
	LevelCritical logLevel = 5
)

type logRecModule int

const (
	ModeNone     logRecModule = 0x00000 // 不加参数
	ModeDate     logRecModule = 0x00001 // 日期
	ModeLevel    logRecModule = 0x00002 // 日志级别
	ModeFileName logRecModule = 0x00004 // 文件名
	ModeFilePath logRecModule = 0x00008 // 文件路径
	ModeLine     logRecModule = 0x00010 // 文件行数
	ModeFunc     logRecModule = 0x00020 // 函数名
	ModeAll      logRecModule = 0x00040 // 全部
	ModeButt     logRecModule = 0xfffff // 非法
)

type logRangePolicy int

const (
	logRangeDelete   logRangePolicy = 0
	logReNameByIndex logRangePolicy = 1
	logReNameByTime  logRangePolicy = 2
)

const (
	// LogRecModuleSimple 日志简单记录
	LogRecModuleSimple = ModeDate | ModeLevel | ModeFunc

	// LogRecModuleNormal 日志普通记录
	LogRecModuleNormal = ModeDate | ModeLevel | ModeFileName | ModeLine | ModeFunc

	maxFormatLength = 12
)

var (
	logLeverFmt = []string{
		"",
		"[DEBUG]",
		"[INFO]",
		"[WARNING]",
		"[ERROR]",
		"[CRITICAL]",
	}

	logIndex = 1

	logSkipStep     = 4
	allowedLogLevel = LevelError
)

// SetConsoleLogSkipStep 跳过级数，获取调用者函数名使用
func SetConsoleLogSkipStep(step int) {
	logSkipStep = step
}

// SetConsoleAllowedLogLevel 设置日志打印级别
func SetConsoleAllowedLogLevel(level logLevel) {
	allowedLogLevel = level
	return
}

const (
	defaultSkipStep = 4
	defaultLogLevel = LevelError

	// 日志大小 默认10M
	defaultMaxSize int64 = 10240

	// 日志超大时的操作，默认按照日期重命名
	defaultLogPolicy = logReNameByTime

	// 日志文件路径
	defaultFilePath = "./"

	// 日志文件名
	defaultFileName = "serverlog.log"

	// 日志文件
	defaultFullName = "./serverlog.log"

	// osFlag 日志文件的打开模式
	defaultOpenModule = 0

	// os.FileMode
	defaultLogPerm os.FileMode = 0

	// 日志记录模式
	defaultRecordModule = LogRecModuleNormal
)

type Logger struct {
	logFilePath       string         // 日志文件路径
	logFileName       string         // 日志文件名
	logFullName       string         // 日志文件路径+文件名
	logOpenModule     int            // osFlag 日志文件的打开模式
	osLogPerm         os.FileMode    // os.FileMode
	allowedLogLevel   logLevel       // 日志打印级别 5个级别
	logRecordModule   logRecModule   // 日志记录模式
	logMaxSize        int64          // 单日志文件大小 单位M
	logFileFullPolicy logRangePolicy // 日志文件超出大小时的操作 删除、按日期重命名、序号重命名
	logSkipStep       int
	outPutFile        bool
	outPutConsole     bool
}
