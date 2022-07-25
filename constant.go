package fun

const (
	DatePattern               = "2006-01-02"
	DatePatternSlash          = "2006/01/02"
	DatePatternZh             = "2006年01月02日"
	DatetimePattern           = "2006-01-02 15:04:05"
	DatetimeMilliPattern      = "2006-01-02 15:04:05.999"
	DatetimePatternSlash      = "2006/01/02 15:04:05"
	DatetimeMilliPatternSlash = "2006/01/02 15:04:05.999"
	DatetimePatternZh         = "2006年01月02日 15时04分05秒"
	DatetimePatternUtc        = "2006-01-02'T'15:04:05'Z'"
)

const (
	StringNumber          = "0123456789"
	StringUpperLetter     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	StringLowerLetter     = "abcdefghijklmnopqrstuvwxyz"
	StringLetter          = StringUpperLetter + StringLowerLetter
	StringLetterAndNumber = StringLetter + StringNumber
)

const (
	MimePlain             = "text/plain"
	MimeHtml              = "text/html"
	MimeJson              = "application/json"
	MimePostForm          = "application/x-www-form-urlencoded"
	MimeMultipartPostForm = "multipart/form-data"
	MimeProtobuf          = "application/x-protobuf"
	MimeYaml              = "application/x-yaml"
)

const (
	RegexEmail        string = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	RegexLetter       string = "^[a-zA-Z]+$"
	RegexLetterNumber string = "^[a-zA-Z0-9]+$"
	RegexNumber       string = "^[0-9]+$"
	RegexChinese      string = "^[\u4e00-\u9fa5]+$"
)

const (
	SizeB  = "B"
	SizeKB = "KB"
	SizeMB = "MB"
	SizeGB = "GB"
	SizeTB = "TB"
	SizePB = "PB"
	SizeEB = "EB"
)

const (
	_ = 1 << (10 * iota)
	BytesPerKB
	BytesPerMB
	BytesPerGB
	BytesPerTB
	BytesPerPB
	BytesPerEB
)
