package config

import (
	"time"
)

type Config struct {
	Server  Server  `toml:"server"`
	Console Console `toml:"console"`
	General General `toml:"general"`
	Log     Log     `toml:"log"`

	Database []Database `toml:"database"`
	Session  Session    `toml:"session"`
	Cache    Cache      `toml:"cache"`

	ZeroCrud      ZeroCrud      `toml:"zero_crud"`
	ColumnMapping ColumnMapping `toml:"column_mapping"`
	Dal           Dal           `toml:"dal"`

	MqRabbitmq []MqRabbitmq `toml:"mq_rabbitmq"`
	MqKafka    []MqKafka    `toml:"mq_kafka"`
	MqRedis    []MqRedis    `toml:"mq_redis"`
	MqTopic    []MqTopic    `toml:"mq_topic"`
}

type Server struct {
	Listen   string `toml:"listen"`
	ServerID string `toml:"server_id"`
}
type Console struct {
	Listen string `toml:"listen"`
	ApiKey string `toml:"api_key"`
}

type General struct {
	AppRoot    string `toml:"app_root"`
	PluginsDir string `toml:"plugins_dir"`
	HealthFile string `toml:"health_file"`

	EnableGzip            bool `toml:"enable_gzip"`
	ScriptTimeout         int  `toml:"script_timeout"`
	TimestampRange        int  `toml:"timestamp_range"`
	EnableVerifySignature bool `toml:"enable_verify_signature"`

	CorsAllowOrigin  []string `toml:"cors_allow_origin"`
	CorsAllowMethods []string `toml:"cors_allow_methods"`
	CorsAllowHeaders []string `toml:"cors_allow_headers"`
}

type Log struct {
	MaxSize    int    `toml:"max_size"`
	MaxBackups int    `toml:"max_backups"`
	MaxAge     int    `toml:"max_age"`
	Compress   bool   `toml:"compress"`
	LogFormat  string `toml:"log_format"`
	File       string `toml:"file"`
	Level      string `toml:"level"`
}

type Database struct {
	DriverName     string `toml:"driver_name"`
	DataSourceName string `toml:"data_source_name"`
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type Session struct {
	Lifetime Duration `toml:"lifetime"`
	Driver   string   `toml:"driver"`
	Codec    string   `toml:"codec"` // gob | json

	Redis  SessionRedis  `toml:"redis"`
	Sqlite SessionSqlite `toml:"sqlite"`
	File   SessionFile   `toml:"file"`
}
type SessionRedis struct {
	MasterName string   `toml:"master_name"`
	Password   string   `toml:"password"`
	Db         int      `toml:"db"`
	AddrList   []string `toml:"addr_list"`
	Prefix     string   `toml:"prefix"`
	StoreAs    string   `toml:"store_as"` // string | hash
}
type SessionSqlite struct {
	Dsn string `toml:"dsn"`
}
type SessionFile struct {
	Path string `toml:"path"`
}

type Cache struct {
	Driver string      `toml:"driver"`
	Memory CacheMemory `toml:"memory"`
	Redis  CacheRedis  `toml:"redis"`
}
type CacheMemory struct {
	Path string `toml:"path"`
}
type CacheRedis struct {
	Prefix     string   `toml:"prefix"`
	MasterName string   `toml:"master_name"`
	Password   string   `toml:"password"`
	Db         int      `toml:"db"`
	AddrList   []string `toml:"addr_list"`
}

type ZeroCrud struct {
	ThrowErrorWhenRowNotFound bool `toml:"throw_error_when_row_not_found"`
}

type ColumnMapping struct {
	ConvertRequest  bool `toml:"convert_request"`
	ConvertResponse bool `toml:"convert_response"`
}

type Dal struct {
	ParseJSONColumns bool `toml:"parse_json_columns"`
}

type RabbitmqPublish struct {
	Exchange    string `toml:"exchange"`
	Immediate   bool   `toml:"immediate"`
	Mandatory   bool   `toml:"mandatory"`
	RoutingKey  string `toml:"routingKey"`
	ContentType string `toml:"contentType"`
}
type RabbitmqQueueDeclare struct {
	Name       string         `toml:"name"`
	NoWait     bool           `toml:"noWait"`
	Durable    bool           `toml:"durable"`
	Exclusive  bool           `toml:"exclusive"`
	AutoDelete bool           `toml:"autoDelete"`
	Arguments  map[string]any `toml:"arguments"`
}
type RabbitmqOption struct {
	URL          string               `toml:"url"`
	Timeout      int                  `toml:"timeout"`
	Publish      RabbitmqPublish      `toml:"publish"`
	QueueDeclare RabbitmqQueueDeclare `toml:"queueDeclare"`
}
type MqRabbitmq struct {
	MqInstance string         `toml:"mq_instance"`
	Option     RabbitmqOption `toml:"option"`
}
type KafkaOption struct {
	AddressList []string `toml:"addressList"`
}
type MqKafka struct {
	MqInstance string      `toml:"mq_instance"`
	Option     KafkaOption `toml:"option"`
}
type RedisOption struct {
	MasterName string   `toml:"master_name"`
	Password   string   `toml:"password"`
	Db         int      `toml:"db"`
	AddrList   []string `toml:"addr_list"`
}
type MqRedis struct {
	MqInstance string      `toml:"mq_instance"`
	Option     RedisOption `toml:"option"`
}

type MqTopic struct {
	MqDriver    string `toml:"mq_driver"`
	MqInstance  string `toml:"mq_instance"`
	Name        string `toml:"name"`
	Description string `toml:"description"`
}

// MqInstanceName
// Deprecated
func (opt *Config) MqInstanceName(topicName string) (string, string) {
	for _, mq := range opt.MqTopic {
		if mq.Name == topicName {
			return mq.MqDriver, mq.MqInstance
		}
	}
	return "", ""
}

// MqRedisOption
// Deprecated
func (opt *Config) MqRedisOption(mqInstance string) *RedisOption {
	for _, mq := range opt.MqRedis {
		if mqInstance == mq.MqInstance {
			return &mq.Option
		}
	}
	return nil
}

// MqRabbitmqOption
// Deprecated
func (opt *Config) MqRabbitmqOption(mqInstance string) *RabbitmqOption {
	for _, mq := range opt.MqRabbitmq {
		if mqInstance == mq.MqInstance {
			return &mq.Option
		}
	}
	return nil
}

// MqKafkaOption
// Deprecated
func (opt *Config) MqKafkaOption(mqInstance string) *KafkaOption {
	for _, mq := range opt.MqKafka {
		if mqInstance == mq.MqInstance {
			return &mq.Option
		}
	}
	return nil
}
