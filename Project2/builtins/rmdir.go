package builtins

import (
    "os"
	"fmt"
)

func RemoveDirectory(dirPath string) error {
    if dirPath == "" {
        return fmt.Errorf("rmdir: missing operand")
    }

    // RemoveAll removes path and any children it contains.
    // It removes everything it can but returns the first error
    // it encounters. If the path does not exist, RemoveAll
    // returns nil (no error).
    return os.RemoveAll(dirPath)
}
