package common


var (
	AppConf AppConfig
)

type AppConfig struct {
	KafkaConf KafkaConfig `ini:"kafka"`
	ESConf ESConfig `ini:"es"`
	LogConf LogConfig `ini:"logs"`
}

type KafkaConfig struct {
	Addr string `ini:"addr"`
	Topic string `ini:"topic"`
}

type ESConfig struct {
	Addr string `ini:"addr"`
	Index string `ini:"index"`
	ThreadNum int `ini:"thread_num"`
	QueueSize int `ini:"queue_size"`
}


type LogConfig struct {
	LogLevel string `ini:"log_level"`
	Filename string `ini:"filename"`
	LogType string `ini:"log_type"`
	Module string `ini:"module"`
}
