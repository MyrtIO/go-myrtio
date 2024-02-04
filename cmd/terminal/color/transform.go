// Package color adds CLI coloring
package color

// Combine concatenates one or more format codes with a given string and returns the result.
func Combine(s string, formatCode ...string) string {
	prefix := ""
	for _, code := range formatCode {
		prefix += code
	}
	return prefix + s + Reset
}

// Dim formats input string as dimmed.
func Dim(s string) string {
	return Combine(s, CodeDim)
}

// Green formats input string as green.
func Green(s string) string {
	return Combine(s, CodeGreen)
}

// Red formats input string as red.
func Red(s string) string {
	return Combine(s, CodeRed)
}

// Blue formats input string as blue.
func Blue(s string) string {
	return Combine(s, CodeBlue)
}
