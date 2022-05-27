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
	RegexEmail        string = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)])"
	RegexLetter       string = "^[a-zA-Z]+$"
	RegexLetterNumber string = "^[a-zA-Z0-9]+$"
	RegexNumber       string = "^[0-9]+$"
	RegexIpv4         string = "\\b((?!\\d\\d\\d)\\d+|1\\d\\d|2[0-4]\\d|25[0-5])\\.((?!\\d\\d\\d)\\d+|1\\d\\d|2[0-4]\\d|25[0-5])\\.((?!\\d\\d\\d)\\d+|1\\d\\d|2[0-4]\\d|25[0-5])\\.((?!\\d\\d\\d)\\d+|1\\d\\d|2[0-4]\\d|25[0-5])\\b"
	RegexIpv6         string = "(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]+|::(ffff(:0{1,4})?:)?((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9]))"
	RegexChinese      string = "^[\u4e00-\u9fa5]+$"
	RegexMacAddress   string = "^[A-F0-9]{2}(:|-)[A-F0-9]{2}(:|-)[A-F0-9]{2}(:|-)[A-F0-9]{2}(:|-)[A-F0-9]{2}(:|-)[A-F0-9]{2}$"
)

const (
	SizeB  = "B"
	SizeKB = "KB"
	SizeMB = "MB"
	SizeGB = "GB"
	SizeTB = "TB"
	SizePB = "PB"
	SizeEB = "EB"

	BytesPerKB = 1024
	BytesPerMB = BytesPerKB * 1024
	BytesPerGB = BytesPerMB * 1024
	BytesPerTB = BytesPerGB * 1024
	BytesPerPB = BytesPerTB * 1024
	BytesPerEB = BytesPerPB * 1024
)
