package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func PrettyPrint(obj interface{}) {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", string(b))
}

func Errorf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

func Printf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, format, a...)
}
