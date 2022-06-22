package utils

import (
	"github.com/dlclark/regexp2"
)

func VerifyEmailFormat(email string) (bool, error) {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	compile := regexp2.MustCompile(pattern, 0)
	return compile.MatchString(email)
}

func VerifyPasswordFormat(password string) (bool, error) {
	pattern := `^(?![0-9]+$)(?![a-zA-Z]+$)[0-9a-zA-Z]{8, 16}$`
	compile := regexp2.MustCompile(pattern, 0)
	return compile.MatchString(password)
}
