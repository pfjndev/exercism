package logs

import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		switch char {
		case 0x2757: // '❗'
			return "recommendation"
		case 0x1f50d: // '🔍'
			return "search"
		case 0x2600: // '☀'
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	//return strings.ReplaceAll(log, string(oldRune), string(newRune))
	return strings.Map(func(r rune) rune {
		if r == oldRune {
			return newRune
		}
		return r
	}, log)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	//return utf8.RuneCountInString(log) <= limit
    count := 0
    for range log{
        count++
        if count > limit { return false }
    }
    return true
}