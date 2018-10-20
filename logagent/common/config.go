package common


type CollectConfig struct {
	/*
	"path":"C:/tmp.log",
	"module_name":"apache",
	"topic":"nginx_log"
	*/
	Path string `json:"path"`
	ModuleName string `json:"module_name"`
	Topic string `json:"topic"`
}


type AppConfig struct {
	KafkaConf KafkaConfig `ini:"kafka"`
	CollectLogConf  CollectLogConfig `ini:"collect_log_conf"`
	LogConf LogConfig `ini:"logs"`
	EtcdConf EtcdConfig `ini:"etcd"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	QueueSize int  `ini:"queue_size"`
}

type CollectLogConfig struct {
	LogFilenames string `ini:"log_filenames"`
}

type EtcdConfig struct {
	Address string `ini:"address"`
	EtcdKey string `ini:"etcd_key"`
}

/*
log_level=debug    
filename=./logs/logagent.log
#console|file
log_type=file
module=logagent
*/
type LogConfig struct {
	LogLevel string `ini:"log_level"`
	Filename string `ini:"filename"`
	LogType string `ini:"log_type"`
	Module string `ini:"module"`
}
