package simplelog

import (
    "time"
    "fmt"
    "runtime"
    "path"
)

const (
    // [log_level] [time] [filename]:[line] [message]
    LOG_FORMAT = "%s %s %s:%d %s\n"
)

const (
    LOG_LEVEL_DEBUG = iota
    LOG_LEVEL_INFO
    LOG_LEVEL_WARN
    LOG_LEVEL_ERROR
    LOG_LEVEL_FATAL
)

var LOG_LEVEL_STRINGS = [...]string {
    "DEBUG",
    "INFO",
    "WARN",
    "ERROR",
    "FATAL",
}

var log_level = LOG_LEVEL_DEBUG

func SetLogLevel(level int) {
    log_level = level
}


func LogDebug(format string, v ...interface{}) {
    logByLevel(LOG_LEVEL_DEBUG, fmt.Sprintf(format, v ...))
}

func LogInfo(format string, v ...interface{}) {
    logByLevel(LOG_LEVEL_INFO, fmt.Sprintf(format, v ...))
}

func LogWarn(format string, v ...interface{}) {
    logByLevel(LOG_LEVEL_WARN, fmt.Sprintf(format, v ...))
}

func LogError(format string, v ...interface{}) {
    logByLevel(LOG_LEVEL_ERROR, fmt.Sprintf(format, v ...))
}

func LogFatal(format string, v ...interface{}) {
    logByLevel(LOG_LEVEL_FATAL, fmt.Sprintf(format, v ...))
}

func logByLevel(level int, message string) {
    if level < log_level {
        return
    }
    time_str := fmt.Sprintf("%s", time.Now())[:19]
    _, filename, line, _ := runtime.Caller(2)
    _, filename = path.Split(filename)
    fmt.Printf(LOG_FORMAT, LOG_LEVEL_STRINGS[level], time_str, filename, line, message)
}
