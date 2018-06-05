package bootstrap

import (
	"fmt"
	"testing"
	"github.com/aos-stack/aos/config"
	"github.com/go-xorm/core"
)

func TestDatabaseConfigAssembler(t *testing.T) {
	assembler := &bootstrap.ConfigAssembler{}
	yaml := map[string]interface{} {
		"host": "127.0.0.1",
		"port": 3306,
		"username": "root",
		"password": "123456",
		"database": "aos",
		"max_idle": 100,
		"max_open": 100,
		"max_left_time": 100,
		"log_sql": true,
		"log_sql_execute_time": 100,
		"log_level": "info",
	}
	
	config := assembler.Database(yaml)
	fmt.Println(core.LOG_DEBUG)
	fmt.Println(config)
}