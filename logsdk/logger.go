package logsdk

import "os"

var logger = &Logger{}

func NewLogger() *Logger {
	if logger == nil {
		logger = &Logger{}
	}
	return logger
}

func (l *Logger) SetOutPutToFile(flag bool) *Logger {
	l.outPutFile = flag
	return l
}

func (l *Logger) SetOutPutToConsole(flag bool) *Logger {
	l.outPutConsole = flag
	return l
}

// SetLogFullName 设置日志文件信息 默认是./serverlog.log
func (l *Logger) SetLogFullName(fullName string) *Logger {
	l.logFullName = fullName
	return l
}

// SetLogFilePath 设置日志文件路径
func (l *Logger) SetLogFilePath(filepath string) *Logger {
	l.logFilePath = filepath
	return l
}

// SetLogFileName 设置日志文件名称
func (l *Logger) SetLogFileName(filename string) *Logger {
	l.logFileName = filename
	return l
}

// SetOsFlag 设置文件打开参数，默认是os.O_RDWR|os.O_CREATE|os.O_APPEND，读写，创建，增量写
func (l *Logger) SetOsFlag(flag int) *Logger {
	l.logOpenModule = flag
	return l
}

// SetOsPerm 设置 FileMode，默认是0644
func (l *Logger) SetOsPerm(perm os.FileMode) *Logger {
	l.osLogPerm = perm
	return l
}

// SetRecordMode 设置日志打印模式
// 提供的有date，level，funcName，fileName，filePath，line，默认 date+level+file+line+func
func (l *Logger) SetRecordMode(flag int) *Logger {
	l.logRecordModule = logRecModule(flag)
	return l
}

// SetLogMaxSize 设置日志文件大小,单位M
func (l *Logger) SetLogMaxSize(logSize int64) *Logger {
	l.logMaxSize = logSize*1024 + 1
	return l
}

// SetOverRangePolicy 设置日志文件超过最大内存的策略
func (l *Logger) SetOverRangePolicy(policy int) *Logger {
	l.logFileFullPolicy = logRangePolicy(policy)
	return l
}

// SetLogSkipStep 跳过级数，获取调用者函数名使用
func (l *Logger) SetLogSkipStep(step int) *Logger {
	l.logSkipStep = step
	return l
}

// SetAllowedLogLevel 设置日志打印级别
func (l *Logger) SetAllowedLogLevel(level logLevel) *Logger {
	l.allowedLogLevel = level
	return l
}

func (l *Logger) InitDefault() *Logger {
	l.SetLogFilePath(defaultFilePath).
		SetLogFileName(defaultFileName).
		SetLogMaxSize(defaultMaxSize).
		SetOsFlag(defaultOpenModule).
		SetLogSkipStep(defaultSkipStep).
		SetOsPerm(defaultLogPerm).
		SetAllowedLogLevel(defaultLogLevel).
		SetOverRangePolicy(int(defaultLogPolicy)).
		SetRecordMode(int(defaultRecordModule)).
		SetOutPutToFile(false).
		SetOutPutToConsole(true)

	return l
}
