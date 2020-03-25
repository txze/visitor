package util

import (
	"encoding/json"
	"math/rand"
	"time"
)

func Now13() int64 {
	return time.Now().UnixNano() / 1e6
}

func Now10() int64 {
	return time.Now().Unix()
}

// 当天时间的0点
func DayStart(add time.Duration) int64 {
	now := time.Now()
	year, month, day := now.Date()
	startTime := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	startTime = startTime.Add(add)
	return startTime.UnixNano() / 1e6
}

func DayStartTime(add time.Duration) time.Time {
	now := time.Now().Format("2006-01-02")
	startTime, _ := time.Parse("2006-01-02", now)
	startTime = startTime.Add(add)
	return startTime
}

func Json2S(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}

func S2Json(data interface{}) string {
	bys, _ := json.Marshal(data)
	return string(bys)
}

func S2S(data, v interface{}) error {
	return Json2S(S2Json(data), v)
}

func Bytes2S(bys []byte, v interface{}) error {
	return json.Unmarshal(bys, v)
}

func S2Bytes(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func TrimRepeatInt(array []int) []int {
	var record = map[int]bool{}
	var newArray []int
	for _, digit := range array {
		if _, ok := record[digit]; ok {
			continue
		}
		newArray = append(newArray, digit)
		record[digit] = true
	}
	return newArray
}

func TrimRepeatString(array []string) []string {
	var record = map[string]bool{}
	var newArray []string
	for _, str := range array {
		if _, ok := record[str]; ok {
			continue
		}
		newArray = append(newArray, str)
		record[str] = true
	}
	return newArray
}

func TrimArraySpace(array []string) []string {
	var newArray []string
	for _, str := range array {
		if str == "" {
			continue
		}
		newArray = append(newArray, str)
	}
	return newArray
}

func RandomInt(n int) int {
	var num = 0
	for i := 0; i < n; i++ {
		v := rand.Intn(10)
		if i == 0 && v == 0 {
			i = -1
			continue
		}
		num = num*10 + v
	}
	return num
}
