package fun

const (
	SPACE      = " "
	DOT        = "."
	SLASH      = "/"
	UNDERSCORE = "_"
	COLON      = ":"
	DASH       = "-"
	LF         = "\n"
	CRLF       = "\r\n"
	CR         = "\r"
	TAB        = "\t"
)

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
