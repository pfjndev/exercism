package parsinglogfiles

import "regexp"

// var validLineCodes = []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
// ^ matches the start of the string
// | acts as "OR"
var validLineRegex = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
var validSplitLineRegex = regexp.MustCompile(`<[~*=-]*>`)
var containsPwdRegex = regexp.MustCompile(`(?i)".*password.*"`)
var containsEndOfLineRegex = regexp.MustCompile(`end-of-line\d*`)
var containsUserNameRegex = regexp.MustCompile(`User\s+([a-zA-Z0-9]+)`)

func IsValidLine(text string) bool {
	return validLineRegex.MatchString(text)
}
/* 	for _, v := range validLineCodes {
		if strings.HasPrefix(text, v) { 
			return true
		}
	}
	return false
*/

func SplitLogLine(text string) []string {
	return validSplitLineRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	total := 0
	for _, line := range lines {
		if containsPwdRegex.MatchString(line) { total++ }
	}
	return total
}

func RemoveEndOfLineText(text string) string {
	return containsEndOfLineRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	for lineNumber, line := range lines {
		matches := containsUserNameRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			lines[lineNumber] = "[USR] " + matches[1] + " " + line
		}
	}
	return lines
}