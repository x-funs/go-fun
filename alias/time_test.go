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
	u1 := user{
		Name:     "Alice",
		Birthday: Date{birthday},
		JoinTime: DateTime{joinTime},
	}

	u1json, _ := json.Marshal(u1)
	t.Log(string(u1json))

	var u2 user
	jsonStr := `{"name":"Alice","birthday":"2006年01月02日","joinTime":"2021/02/03 01:02:03"}`
	json.Unmarshal([]byte(jsonStr), &u2)
	t.Log(u2)
	t.Log(u2.Birthday.Time.IsZero())
}

func TestAliasTimeFormatMarshal(t *testing.T) {
	birthday, _ := time.ParseInLocation("2006-01-02", "1991-02-03", time.Local)
	joinTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-02-03 01:02:03", time.Local)
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

	var b2 body
	jsonStr := `{"name":"Alice","birthday":"1991-02-03","joinTime":"2021/02/03 01:02:03"}`
	json.Unmarshal([]byte(jsonStr), &b2)
	t.Log(b2)
	t.Log(b2.Birthday.Time.IsZero())

	b2json, _ := json.Marshal(b2)
	t.Log(string(b2json))
}
