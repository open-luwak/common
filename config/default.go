package config

var DefaultConfig = &Config{
	ZeroCrud: ZeroCrud{
		ThrowErrorWhenRowNotFound: true,
	},
	General: General{
		EnableGzip:            true,
		ScriptTimeout:         60,
		TimestampRange:        900,
		EnableVerifySignature: true,
	},
	Log: Log{
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     7,
		Compress:   true,
		File:       "log/system.log",
		LogFormat:  "json",
		Level:      "info",
	},
	ColumnMapping: ColumnMapping{
		ConvertRequest:  true,
		ConvertResponse: true,
	},
	Dal: Dal{
		ParseJSONColumns: true,
	},
}
