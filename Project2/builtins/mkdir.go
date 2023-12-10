package builtins

import (
    "os"
	"fmt"
)

func MakeDirectory(dirPath string, mode os.FileMode) error {
    if dirPath == "" {
        return fmt.Errorf("mkdir: missing operand")
    }
    // Default mode is 0755 if not specified
    if mode == 0 {
        mode = 0755
    }

    return os.MkdirAll(dirPath, mode)
}
