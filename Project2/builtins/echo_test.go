package builtins

import (
    "bytes"
    "testing"
)

func TestEcho(t *testing.T) {
    tests := []struct {
        name string
        args []string
        want string
    }{
        {"No Arguments", []string{}, "\n"},
        {"Single Word", []string{"hello"}, "hello\n"},
        {"Multiple Words", []string{"hello", "world"}, "hello world\n"},
        
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            buffer := new(bytes.Buffer)
            if err := Echo(buffer, tt.args...); err != nil {
                t.Errorf("Echo() error = %v", err)
            }
            if got := buffer.String(); got != tt.want {
                t.Errorf("Echo() = %v, want %v", got, tt.want)
            }
        })
    }
}
