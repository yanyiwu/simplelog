package simplelog

import "testing"

func TestLog(t * testing.T) {
    Debug("hello simplelog.");
    Info("hello simplelog.");
    Warn("hello simplelog.");
    Error("hello simplelog.");
    Fatal("hello simplelog.");
    SetLevel(LEVEL_FATAL)
    Debug("hello simplelog.");
    Info("hello simplelog.");
    Warn("hello simplelog.");
    Error("hello simplelog.");
    Fatal("hello simplelog.");
}
