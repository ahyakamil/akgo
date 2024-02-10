package aklog

import (
	"akgo/akmdc"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var secretKeywords []string

func init() {
	secretKeywords = []string{"token", "password", "auth"}
}
func Info(message string) {
	log.Print(" INFO " + build(message, "default"))
}
func Warn(message string) {
	log.Print(" WARN " + build(message, "default"))
}
func Error(message string) {
	log.Print(" ERROR " + build(message, "default"))
}

func InfoWithCtx(message string, ctxKey string) {
	log.Print(" INFO " + build(message, ctxKey))
}
func WarnWithCtx(message string, ctxKey string) {
	log.Print(" WARN " + build(message, ctxKey))
}
func ErrorWithCtx(message string, ctxKey string) {
	log.Print(" ERROR " + build(message, ctxKey))
}

func build(message string, ctxKey string) string {
	mdc := akmdc.GetMDCWithCtx(ctxKey)
	re := regexp.MustCompile(`\r\n|[\r\n\v\f\x{0085}\x{2028}\x{2029}]`)
	return "MDC_GROUP=" + fmt.Sprintf("%v", mdc["MDC_GROUP"]) + " " + filter(re.ReplaceAllString(message, ""))
}

func filter(message string) string {
	secrets := " :::secretKeywordsRemovedFromLog="
	for i := range secretKeywords {
		keyword := secretKeywords[i]
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
