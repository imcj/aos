package bootstrap

type ConfigDatabase struct {
	Type string
	Host string
	Port int
	Username string
	Password string
	Database string
	MaxIdle int
	MaxOpen int
	MaxLeftTime int
	LogSQL bool
	LogSQLExecuteTime bool
	LogLevel int
}

type ConfigAssembler struct {

}

// TODO: 检查键不存在则使用默认值
func (a *ConfigAssembler)Database(yaml map[string]interface{})*ConfigDatabase {
	// yaml["log_level"].(string)
	var logLevel int
	databaseType := yaml["type"].(string)
	host := yaml["host"].(string)
	port := yaml["port"].(int)
	username := yaml["username"].(string)
	password := yaml["password"].(string)
	database := yaml["database"].(string)
	maxIdle := yaml["max_idle"].(int)
	maxOpen := yaml["max_open"].(int)
	maxLeftTime := yaml["max_left_time"].(int)
	logSQL := yaml["log_sql"].(bool)
	logSQLExecuteTime := yaml["log_sql_execute_time"].(bool)

	switch (yaml["log_level"]) {
	case "DEBUG":
		logLevel = 0
		break
	case "INFO":
		logLevel = 1
		break
	case "Warning":
		logLevel = 2
		break
	case "Error":
		logLevel = 3
		break
	case "Fatal":
		logLevel = 4
		break
	default:
		logLevel = 0
		break
	}

	configDatabase := &ConfigDatabase {
		databaseType,
		host,
		port,
		username,
		password,
		database,
		maxIdle,
		maxOpen,
		maxLeftTime,
		logSQL,
		logSQLExecuteTime,
		logLevel}
	return configDatabase
}