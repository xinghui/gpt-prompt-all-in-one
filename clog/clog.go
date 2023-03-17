package clog

import "fmt"

const (
	Red = 31 + iota
	Green
	Yellow
	Blue
	Magenta
	Cyan
)

func printColor(color int, format string, args ...interface{}) {
	c := fmt.Sprintf(format, args...)
	fmt.Printf("\033[%dm%s\n", color, c)
	fmt.Printf("\033[0m")
}

// Success print
func Success(format string, args ...interface{}) {
	printColor(Green, format, args...)
}

// Warning print
func Warning(format string, args ...interface{}) {
	printColor(Yellow, format, args...)
}

// Error print
func Error(format string, args ...interface{}) {
	printColor(Red, format, args...)
}

// Tips blue
func Tips(format string, args ...interface{}) {
	printColor(Blue, format, args...)
}

// Tips2 magenta
func Tips2(format string, args ...interface{}) {
	printColor(Magenta, format, args...)
}

// Tips3 cyan
func Tips3(format string, args ...interface{}) {
	printColor(Cyan, format, args...)
}
