package aklog

import (
	"akgo/akmdc"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var secretKeyword []string

func init() {
	secretKeyword = []string{"token", "password", "auth"}
}

func Info(message string) {
	log.Print(" INFO " + build(message))
}
func Warn(message string) {
	log.Print(" WARN " + build(message))
}
func Error(message string) {
	log.Print(" ERROR " + build(message))
}

func build(message string) string {
	mdc := akmdc.GetMDC()
	re := regexp.MustCompile(`\r\n|[\r\n\v\f\x{0085}\x{2028}\x{2029}]`)
	return "MDC_GROUP=" + fmt.Sprintf("%v", mdc["MDC_GROUP"]) + " " + filter(re.ReplaceAllString(message, ""))
}

func filter(message string) string {
	secrets := " :::secretKeywordsRemovedFromLog="
	for i := range secretKeyword {
		keyword := secretKeyword[i]
		if strings.Contains(strings.ToLower(message), strings.ToLower(keyword)) {
			regexPattern := `(?i)` + keyword + `.[^,]+`
			if strings.Contains(strings.ToLower(message), "multipart/form-data") {
				regexPattern = `(?i)` + keyword + `.\s*(\S+)`
			}
			re := regexp.MustCompile(regexPattern)
			message = re.ReplaceAllString(message, "")
			secrets += keyword + ","
		}
	}
	return "\"" + message + "\"" + secrets
}
