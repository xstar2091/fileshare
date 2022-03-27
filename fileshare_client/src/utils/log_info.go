package utils

import (
	"fileshare_client/src/conf"
	"fmt"
	"os"
)

func LogDebug(format string, args ...interface{}) {
	if !conf.IsDebugMode {
		return
	}
	_, _ = fmt.Fprint(os.Stdout, "\t")
	_, _ = fmt.Fprintf(os.Stdout, format, args...)
	_, _ = fmt.Fprint(os.Stdout, "\n")
}

func LogInfo(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, format, args...)
	_, _ = fmt.Fprint(os.Stdout, "\n")
}

func LogError(format string, args ...interface{}) {
	_, _ = fmt.Fprint(os.Stderr, "\t")
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	_, _ = fmt.Fprint(os.Stderr, "\n")
}
