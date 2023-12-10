package builtins

import (
	
	"fmt"
    "io"
    "strings"
)

func Echo(w io.Writer, args ...string) error {
    output := strings.Join(args, " ")
    _, err := fmt.Fprintln(w, output)
    return err
}
