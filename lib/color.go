package lib

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

func wrap(color string, s string) string {
	return color + s + reset
}

func Red(s string) string {
	return wrap(red, s)
}

func Green(s string) string {
	return wrap(green, s)
}

func Yellow(s string) string {
	return wrap(yellow, s)
}

func Blue(s string) string {
	return wrap(blue, s)
}

func Purple(s string) string {
	return wrap(purple, s)
}

func Cyan(s string) string {
	return wrap(cyan, s)
}

func Gray(s string) string {
	return wrap(gray, s)
}

func White(s string) string {
	return wrap(white, s)
}
