package color

const (
	Reset           = "\033[0m"
	EffectBold      = "\033[1m"
	EffectDim       = "\033[2m"
	EffectUnderline = "\033[4m"
	EffectStrike    = "\033[9m"
	EffectItalic    = "\033[3m"

	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func Dim(s string) string {
	return EffectDim + s + Reset
}

func Green(s string) string {
	return ColorGreen + s + Reset
}

func Red(s string) string {
	return ColorRed + s + Reset
}

func Blue(s string) string {
	return ColorBlue + s + Reset
}

func DimUnderline(s string) string {
	return EffectDim + EffectUnderline + s + Reset
}
