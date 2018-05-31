package utils

import (
	"strings"
	"github.com/aos-stack/aos/pkg/setting"
	"strconv"
)

var GraySql GrayXormSql

// 实现 xrom 打印日志接口
type GrayXormSql struct {
}

func (GrayXormSql) Write(p []byte) (n int, err error) {
	if strings.Contains(strings.ToUpper(string(p)), "[SQL]") {
		s := string(p)
		t := strings.Split(s, "took:")
		tst := strings.Trim(t[1], " ")
		tst = strings.Replace(tst, "ms\n", "", -1)
		sqlTime, _ := strconv.ParseFloat(tst, 10)
		m := map[string]interface{}{
			"user_sql":           s,
			"user_sql_exec_time": sqlTime,
		}
		setting.Logger.Infof(s, m)
	}
	return 0, nil
}

