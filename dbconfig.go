package goutils

import "fmt"

// DBConfig 数据库配置
type DBConfig struct {
	// DriverName 数据库名
	DriverName string
	// Host 数据库 IP 地址
	Host string
	// Port 数据库端口
	Port int
	// Username 数据库用户名
	Username string
	// Password 数据库密码
	Password string
	// DBName 数据库名称
	DBName string
	// LogMode 是否开启打印日志模式
	LogMode bool
	// MaxIdleConns 设置空闲连接池中的最大连接数
	MaxIdleConns int
	// MaxOpenConns 设置与数据库的最大打开连接数
	MaxOpenConns int
	// ConnMaxLifeMinutes 设置可重用连接的最长时间（分钟）
	ConnMaxLifeMinutes int
	// ConnTimeout 连接超时时间（秒）
	ConnTimeout int
	// ReadTimeout 读超时时间（秒）
	ReadTimeout int
	// WriteTimeout 写超时时间（秒）
	WriteTimeout int
	// DisableSSL 是否关闭 ssl 模式
	DisableSSL bool
}

// MySQLDSN 返回 Mysql dsn 字符串
func (conf DBConfig) MySQLDSN() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds&readTimeout=%ds&writeTimeout=%ds", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName, conf.ConnTimeout, conf.ReadTimeout, conf.WriteTimeout)
}

// SQLite3DSN 返回 sqlite3 文件名
func (conf DBConfig) SQLite3DSN() string {
	return conf.DBName
}

// PostgresDSN 返回 postgres dns 字符串
func (conf DBConfig) PostgresDSN() string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", conf.Host, conf.Port, conf.Username, conf.DBName, conf.Password)
	if conf.DisableSSL {
		dsn = dsn + " sslmode=disable"
	}
	return dsn
}

// MsSQLDSN 返回 mssql dns 字符串
func (conf DBConfig) MsSQLDSN() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName)
}
