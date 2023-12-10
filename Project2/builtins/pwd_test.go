package builtins

import (
    "bytes"
    "os"
    "testing"
)

func TestPrintWorkingDirectory(t *testing.T) {
    buffer := new(bytes.Buffer)
    if err := PrintWorkingDirectory(buffer); err != nil {
        t.Fatalf("PrintWorkingDirectory() error = %v", err)
    }

    pwd, err := os.Getwd()
    if err != nil {
        t.Fatalf("os.Getwd() error = %v", err)
    }

    if got := buffer.String(); got != pwd+"\n" {
        t.Errorf("PrintWorkingDirectory() = %v, want %v", got, pwd)
    }
}
