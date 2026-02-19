package utils

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/google/uuid"
)

func CheckUuid(uuidString string) bool {
	err := uuid.Validate(uuidString)
	return err == nil
}

func StringFuncName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return ""
	}

	return f.Name()
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToPlainText(str string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	plainText := re.ReplaceAllString(str, `$1 $2`)
	return strings.ToLower(plainText)
}

func IsAcronym(str string) bool {
	matched, _ := regexp.MatchString("^[A-Z]+$", str)
	return matched
}
