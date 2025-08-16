package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type GlobalConfig struct {
	App AppConfig `yaml:"app" toml:"app"`
	Db  DbConfig  `yaml:"database" toml:"database"`
	Log LogConfig `yaml:"log" toml:"log"`
}

type AppConfig struct {
	ServiceName string `yaml:"service_name" toml:"service_name"`
	Port        string `default:"8096" yaml:"port" toml:"port"`
	GinMode     string `yaml:"gin_mode" toml:"gin_mode"`
}

func (a AppConfig) ServeAddr() string {
	return fmt.Sprintf("0.0.0.0:%s", a.Port)
}

type DbConfig struct {
	Host           string `toml:"host" yaml:"host"`
	Port           string `toml:"port" yaml:"port"`
	Username       string `toml:"username" yaml:"username"`
	Password       string `toml:"password" yaml:"password"`
	Database       string `toml:"database" yaml:"database"`
	MaxIdle        int    `toml:"max_idle" yaml:"max_idle"`
	MaxConn        int    `toml:"max_conn" yaml:"max_conn"`
	MaxLife        int    `toml:"max_life" yaml:"max_life"`
	MigrationTable string `toml:"migration_table" yaml:"migration_table"`
}

func (d DbConfig) Dsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.Username, d.Password, d.Database)
}

func ReadConfigFromYamlFile(path string) (*GlobalConfig, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config GlobalConfig
	err = yaml.Unmarshal(bytes, &config)
	return &config, err
}

type LogConfig struct {
	Path         string `yaml:"path" toml:"path"`
	MaxSize      int    `yaml:"max_size" toml:"max_size"`       //MB
	MaxBackups   int    `yaml:"max_backups" toml:"max_backups"` //保留天数
	MaxAge       int    `yaml:"max_age" toml:"max_age"`
	Compress     bool   `yaml:"compress" toml:"compress"`
	Level        int8   `yaml:"level" toml:"level"` // -1 debug,0 info,1 warn,2 error,3 DPanic 4 Panic, 5 Fatal
	ConsoleDebug bool   `yaml:"console_debug" toml:"console_debug"`

	GormLogLevel              int  `yaml:"gorm_log_level" toml:"gorm_log_level"` // 1 silent,2 error,3 warn, 4 info
	SlowThreshold             int  `yaml:"slow_threshold" toml:"slow_threshold"` //ms
	IgnoreRecordNotFoundError bool `yaml:"ignore_record_not_found" toml:"ignore_record_not_found"`
}

func (l LogConfig) Default() LogConfig {
	if len(l.Path) == 0 {
		l.Path = "."
	}
	if l.MaxSize == 0 {
		l.MaxSize = 20
	}
	if l.MaxBackups == 0 {
		l.MaxBackups = 7
	}
	if l.MaxAge == 0 {
		l.MaxAge = 30
	}
	if l.GormLogLevel == 0 {
		l.GormLogLevel = 3
	}
	if l.SlowThreshold == 0 {
		l.SlowThreshold = 200
	}
	return l
}
