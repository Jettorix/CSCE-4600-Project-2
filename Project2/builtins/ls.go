package builtins

import (
    "fmt"
    "io"
    "io/ioutil"
)

func ListDirectories(w io.Writer, dirPath string) error {
    if dirPath == "" {
        dirPath = "." // default to current directory if no argument is provided
    }

    files, err := ioutil.ReadDir(dirPath)
    if err != nil {
        return err
    }

    for _, file := range files {
        fmt.Fprintln(w, file.Name())
    }

    return nil
}
