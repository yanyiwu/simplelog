package simplelog

import "testing"

func TestLog(t * testing.T) {
    LogDebug("hello simplelog.");
    LogInfo("hello simplelog.");
    LogWarn("hello simplelog.");
    LogError("hello simplelog.");
    LogFatal("hello simplelog.");
    SetLogLevel(LOG_LEVEL_FATAL)
    LogDebug("hello simplelog.");
    LogInfo("hello simplelog.");
    LogWarn("hello simplelog.");
    LogError("hello simplelog.");
    LogFatal("hello simplelog.");
}
