package alias

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/x-funs/go-fun"
)

type person struct {
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	JoinTime time.Time `json:"joinTime"`
}

type user struct {
	Name     string   `json:"name"`
	Birthday Date     `json:"birthday"`
	JoinTime DateTime `json:"joinTime"`
}

type body struct {
	Name     string         `json:"name"`
	Birthday DateTimeFormat `json:"birthday"`
	JoinTime DateTimeFormat `json:"joinTime"`
}

func TestTimeMarshal(t *testing.T) {
	birthday, _ := time.ParseInLocation("2006-01-02", "1991-02-03", time.Local)
	joinTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-02-03 01:02:03", time.Local)

	// time.Time 默认使用 RFC3339 格式序列化
	p1 := person{
		Name:     "Bob",
		Birthday: birthday,
		JoinTime: joinTime,
	}

	p1json, _ := json.Marshal(p1)
	t.Log(string(p1json))

	// 反序列化时候必须得给 RFC3339 格式，否则解析不了
	var p2 person
	jsonStr := `{"name":"Bob","birthday":"1991-02-03","joinTime":"2021-02-03 01:02:03"}`
	json.Unmarshal([]byte(jsonStr), &p2)
	t.Log(p2)
}

func TestAliasTimeMarshal(t *testing.T) {
	birthday, _ := time.ParseInLocation("2006-01-02", "1991-02-03", time.Local)
	joinTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-02-03 01:02:03", time.Local)

	// 使用默认日期时间格式序列化
	u1 := user{
		Name:     "Alice",
		Birthday: Date{birthday},
		JoinTime: DateTime{joinTime},
	}

	u1json, _ := json.Marshal(u1)
	t.Log(string(u1json))

	// 反序列化自动识别格式
	var u2 user
	jsonStr := `{"name":"Alice","birthday":"2006年01月02日","joinTime":"2021/02/03 01:02:03"}`
	json.Unmarshal([]byte(jsonStr), &u2)
	t.Log(u2)
	t.Log(u2.Birthday.Time.IsZero())

	u2json, _ := json.Marshal(u2)
	t.Log(string(u2json))

}

func TestAliasTimeFormatMarshal(t *testing.T) {
	birthday, _ := time.ParseInLocation("2006-01-02", "1991-02-03", time.Local)
	joinTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-02-03 01:02:03", time.Local)

	// 自定义时间序列化格式，如果不定义则使用 RFC3339
	b1 := body{
		Name: "Alice",
		Birthday: DateTimeFormat{
			Time:   birthday,
			Format: fun.DatePatternZh,
		},
		JoinTime: DateTimeFormat{
			Time:   joinTime,
			Format: fun.DatetimePatternZh,
		},
	}

	b1json, _ := json.Marshal(b1)
	t.Log(string(b1json))

	// 反序列化自动识别格式，但无法自动赋值 Format
	var b2 body
	jsonStr := `{"name":"Alice","birthday":"1991-02-03","joinTime":"2021/02/03 01:02:03"}`
	json.Unmarshal([]byte(jsonStr), &b2)
	t.Log(b2)
	t.Log(b2.Birthday.Time.IsZero())

	// 此时 Format 空，使用默认的 RFC3339 格式序列化
	b2json1, _ := json.Marshal(b2)
	t.Log(string(b2json1))

	// 重新赋值 Format 才能继续使用自定义格式
	b2.Birthday.Format = fun.DatePatternZh
	b2.JoinTime.Format = fun.DatetimePatternZh
	b2json2, _ := json.Marshal(b2)
	t.Log(string(b2json2))
}
