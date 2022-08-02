package fun

import "regexp"

var (
	RegexEmailPattern    = regexp.MustCompile(RegexEmail)
	RegexDateTimePattern = regexp.MustCompile(RegexDateTime)
)
