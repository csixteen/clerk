package util

import (
	"fmt"
	"strings"
)

type Color string

const (
	ColorReset  Color = "\033[0m"
	ColorRed    Color = "\033[31m"
	ColorGreen  Color = "\033[32m"
	ColorYellow Color = "\033[33m"
	ColorBlue   Color = "\033[34m"
	ColorPurple Color = "\033[35m"
	ColorCyan   Color = "\033[36m"
	ColorWhite  Color = "\033[37m"
)

func PrintColor(s string, c Color) {
	var str strings.Builder

	str.WriteString(string(c))
	str.WriteString(s)
	str.WriteString(string(ColorReset))
	fmt.Println(str.String())
}
