package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"storiesservice/pkg/utctime"
)

func fprintln(skip int, w io.Writer, a ...any) (int, error) {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Fprintln(
		w,
		append([]any{
			fmt.Sprintf(
				"%s %s:%d:",
				utctime.Get().Format("2006/01/02 15:04:05"),
				file,
				line,
			),
		}, a...)...,
	)
}

func FprintlnEx(skip int, w io.Writer, a ...any) (int, error) {
	return fprintln(skip, w, a...)
}

func Fprintln(w io.Writer, a ...any) (int, error) {
	return FprintlnEx(3, w, a...)
}

func EprintlnEx(skip int, a ...any) (int, error) {
	return fprintln(skip, os.Stderr, append([]any{"error:"}, a...)...)
}

func Eprintln(a ...any) (int, error) {
	return EprintlnEx(3, a...)
}

func PrintlnEx(skip int, a ...any) (int, error) {
	return fprintln(skip, os.Stdout, a...)
}

func Println(a ...any) (int, error) {
	return PrintlnEx(3, a...)
}
