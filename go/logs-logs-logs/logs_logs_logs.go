package logs

import "unicode/utf8"

/*func assertNoErr(err error, explanation string) {
	if err != nil {
		log.Panic(explanation)
	}
}*/

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, x := range log {
		switch x {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newString := ""
	for _, x := range log {
		if x == oldRune {
			newString += string(newRune)
		} else {
			newString += string(x)
		}
	}
	return newString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
