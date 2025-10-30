package config

type Config struct {
	Server       Server       `toml:"server"`
	Console      Console      `toml:"console"`
	Log          Log          `toml:"log"`
	Database     []Database   `toml:"database"`
	ZeroCrud     ZeroCrud     `toml:"zero_crud"`
	Session      Session      `toml:"session"`
	RedisSession RedisSession `toml:"redis_session"`

	MqRabbitmq []MqRabbitmq `toml:"mq_rabbitmq"`
	MqKafka    []MqKafka    `toml:"mq_kafka"`
	MqRedis    []MqRedis    `toml:"mq_redis"`
	MqTopic    []MqTopic    `toml:"mq_topic"`

	Cache   Cache   `toml:"cache"`
	General General `toml:"general"`

	ColumnMapping ColumnMapping `toml:"column_mapping"`

	Dal Dal `toml:"dal"`
}

type Server struct {
	Listen   string `toml:"listen"`
	ServerID string `toml:"server_id"`
}
type Console struct {
	Listen string `toml:"listen"`
	ApiKey string `toml:"api_key"`
}

type Database struct {
	DriverName     string `toml:"driver_name"`
	DataSourceName string `toml:"data_source_name"`
}
type RedisSession struct {
	MasterName string   `toml:"master_name"`
	Password   string   `toml:"password"`
	Db         int      `toml:"db"`
	AddrList   []string `toml:"addr_list"`
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

type ZeroCrud struct {
	ThrowErrorWhenRowNotFound bool `toml:"throw_error_when_row_not_found"`
}

type Session struct {
	KeyPrefix         string `toml:"key_prefix"`
	SerializationType string `toml:"serialization_type"`
	TTL               int    `toml:"ttl"`
}

type Cache struct {
	MasterName string   `toml:"master_name"`
	Password   string   `toml:"password"`
	Db         int      `toml:"db"`
	AddrList   []string `toml:"addr_list"`
}

type General struct {
	EnableGzip            bool `toml:"enable_gzip"`
	ScriptTimeout         int  `toml:"script_timeout"`
	TimestampRange        int  `toml:"timestamp_range"`
	EnableVerifySignature bool `toml:"enable_verify_signature"`

	CorsAllowOrigin  []string `toml:"cors_allow_origin"`
	CorsAllowMethods []string `toml:"cors_allow_methods"`
	CorsAllowHeaders []string `toml:"cors_allow_headers"`
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
