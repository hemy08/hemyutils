package logsdk

// Debug debug
func Debug(format string, v ...interface{}) {
	if logger.outPutConsole {
		outConsole(LevelDebug, format, v...)
	}

	if logger.outPutFile {
		writeLogToFile(LevelDebug, format, v...)
	}
}

// Info information
func Info(format string, v ...interface{}) {
	if logger.outPutConsole {
		outConsole(LevelInfo, format, v...)
	}

	if logger.outPutFile {
		writeLogToFile(LevelInfo, format, v...)
	}
}

// Warn for log
func Warn(format string, v ...interface{}) {
	if logger.outPutConsole {
		outConsole(LevelWarning, format, v...)
	}

	if logger.outPutFile {
		writeLogToFile(LevelWarning, format, v...)
	}
}

// Error error
func Error(format string, v ...interface{}) {
	if logger.outPutConsole {
		outConsole(LevelError, format, v...)
	}

	if logger.outPutFile {
		writeLogToFile(LevelError, format, v...)
	}
}

// Critical critical
func Critical(format string, v ...interface{}) {
	if logger.outPutConsole {
		outConsole(LevelCritical, format, v...)
	}

	if logger.outPutFile {
		writeLogToFile(LevelCritical, format, v...)
	}
}
