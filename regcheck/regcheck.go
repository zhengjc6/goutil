package regcheck

import "regexp"

// email verify
func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|162|166|198|199|195|101|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// mobile verify
func VerifyNameFormat(name string) bool {
	regular := "^.{1,255}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(name)
}

///^((https?|ftp|file):\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/

func VerifyUrlFormat(url string) bool {
	regular := "^((https?|ftp|file):\\/\\/)?([\\da-z\\.-]+)\\.([a-z\\.]{2,6})([\\/\\w \\.-]*)*\\/?$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(url)
}
