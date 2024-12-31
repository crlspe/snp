package color

const (
	Reset         = "\033[0m"
	BLACK         = "\033[30m"
	RED           = "\033[31m"
	GREEN         = "\033[32m"
	YELLOW        = "\033[33m"
	BLUE          = "\033[34m"
	MAGENTA       = "\033[35m"
	CYAN          = "\033[36m"
	WHITE         = "\033[37m"
	GRAY          = "\033[90m"
	BRIGHTRED     = "\033[91m"
	BRIGHTGREEN   = "\033[92m"
	BRIGHTYELLOW  = "\033[93m"
	BRIGHTBLUE    = "\033[94m"
	BRIGHTMAGENTA = "\033[95m"
	BRIGHTCYAN    = "\033[96m"
	BRIGHTWHITE   = "\033[97m"
)

func Colorize(colorCode, text string) string {
	return colorCode + text + Reset
}

func Black(text string) string {
	return Colorize(BLACK, text)
}

func Red(text string) string {
	return Colorize(RED, text)
}

func Green(text string) string {
	return Colorize(GREEN, text)
}

func Yellow(text string) string {
	return Colorize(YELLOW, text)
}

func Blue(text string) string {
	return Colorize(BLUE, text)
}

func Magenta(text string) string {
	return Colorize(MAGENTA, text)
}

func Cyan(text string) string {
	return Colorize(CYAN, text)
}

func White(text string) string {
	return Colorize(WHITE, text)
}

func Gray(text string) string {
	return Colorize(GRAY, text)
}

func BrightRed(text string) string {
	return Colorize(BRIGHTRED, text)
}

func BrightGreen(text string) string {
	return Colorize(BRIGHTGREEN, text)
}

func BrightYellow(text string) string {
	return Colorize(BRIGHTYELLOW, text)
}

func BrightBlue(text string) string {
	return Colorize(BRIGHTBLUE, text)
}

func BrightMagenta(text string) string {
	return Colorize(BRIGHTMAGENTA, text)
}

func BrightCyan(text string) string {
	return Colorize(BRIGHTCYAN, text)
}

func BrightWhite(text string) string {
	return Colorize(BRIGHTWHITE, text)
}
