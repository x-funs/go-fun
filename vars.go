package fun

import "regexp"

var (
	RegexUrlPattern   = regexp.MustCompile(RegexUrl)
	RegexEmailPattern = regexp.MustCompile(RegexEmail)
)
