package log

import (
    "testing"
)

func TestDefaultLogger(t *testing.T)  {
    INFO.Print("should not panic")
    WARNING.Print("should not panic")
    ERROR.Print("should not panic")
    FATAL.Print("should not panic")
}