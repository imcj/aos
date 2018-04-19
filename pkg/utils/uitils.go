package utils

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"
)

// PathExists ...
func PathExists(path string) (bool, error) {
	dir, _ := os.Getwd() //当前的目录
	path = dir + path
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// MakeDir ...
func MakeDir(path string) error {
	dir, _ := os.Getwd() //当前的目录
	err := os.Mkdir(dir+path, os.ModePerm)
	return err
}

// ChangeStruct2OtherStruct ...
func ChangeStruct2OtherStruct(item interface{}, toItem interface{}) {

	j, _ := json.Marshal(item)
	json.Unmarshal(j, &toItem)
}

// ChangeRedis2OtherStruct ...
func ChangeRedis2OtherStruct(item interface{}) []uint8 {
	// var toItem []uint8
	// j, _ := json.Marshal(item)
	// json.Unmarshal(j, &toItem)
	return item.([]uint8)
}

// ChangeUint82OtherStruct ...
func ChangeUint82OtherStruct(item interface{}, toItem interface{}) {
	j, _ := json.Marshal(item)
	json.Unmarshal(j, &toItem)
}

// ChangeByteStruct2OtherStruct ...
func ChangeByteStruct2OtherStruct(str []byte, toItem interface{}) error {

	if erri := json.Unmarshal([]byte(str), &toItem); erri != nil && toItem == nil {
		return erri
	}
	return nil
}

// ChangeArrayString2Int ...
func ChangeArrayString2Int(before []string) []int {
	after := make([]int, len(before))
	for k, item := range before {
		after[k], _ = strconv.Atoi(item)
	}
	return after
}

// ChangeArrayString2Int64 ...
func ChangeArrayString2Int64(before []string) []int64 {
	after := make([]int64, len(before))
	for k, item := range before {
		after[k], _ = strconv.ParseInt(item, 10, 64)
	}
	return after
}

// JoinInt64Array2String 拆分int64的数组为string
func JoinInt64Array2String(before []int64, sep string) string {
	after := ""
	for _, item := range before {
		after = after + strconv.Itoa(int(item)) + sep
	}
	return strings.Trim(after, sep)
}

// String2Int ...
func String2Int(before string) int {
	after, _ := strconv.Atoi(before)
	return after
}

// String2Int64 ...
func String2Int64(before string) int64 {
	after, _ := strconv.ParseInt(before, 10, 64)
	return after
}

// String2Float64 ...
func String2Float64(before string) float64 {
	after, _ := strconv.ParseFloat(before, 64)
	return after
}

// String2Int32 ...
func String2Int32(before string) int32 {
	after, _ := strconv.Atoi(before)
	return int32(after)
}

// String2Int8 ...
func String2Int8(before string) int8 {
	after, _ := strconv.Atoi(before)
	return int8(after)
}

// Slise2Map ....
func Slise2Map(originData []string) map[string]int {
	retData := make(map[string]int, 0)
	if originData == nil {
		return retData
	}
	// count := len(originData)
	for i, v := range originData {
		retData[v] = i
	}
	// for i := 0; i < count; i++ {
	// 	// if mapContains(retData, originData[i]) { //前后取前

	// 	// }
	// }
	return retData
}

//  InArrayInt64 是否存在数组中 int64
func InArrayInt64(arr []int64, val int64) bool {
	flag := false
	for _, item := range arr {
		if item == val {
			flag = true
			break
		}
	}
	return flag
}

//  InArrayInt  是否存在数组中 int
func InArrayInt(arr []int, val int) bool {
	flag := false
	for _, item := range arr {
		if item == val {
			flag = true
			break
		}
	}
	return flag
}

// InArrayString 是否存在数组中 字符串
func InArrayString(arr []string, val string) bool {
	flag := false
	for _, item := range arr {
		if item == val {
			flag = true
			break
		}
	}
	return flag
}

//判断key是否存在
func MapContains(needMap map[string]int, key string) bool {
	if _, ok := needMap[key]; ok {
		return true
	}
	return false
}

// FormartDate2Time
func FormartDate2Time(dataTimeStr, ms string) int64 {
	dataTime, _ := time.Parse(ms, dataTimeStr)
	return dataTime.Unix()
}

// ABCToRune ...
func ABCToRune(abc string) rune {
	abcrune := []rune(abc)
	return abcrune[0]
}
