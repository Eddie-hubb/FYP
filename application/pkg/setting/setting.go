package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
	// "path/filepath"
    // "os"
    // "os/exec"
    // "string"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

var cfg *ini.File

// func GetAppPath() string {
//     file, _ := exec.LookPath(os.Args[0])
//     path, _ := filepath.Abs(file)
//     index := strings.LastIndex(path, string(os.PathSeparator))

//     return path[:index]
// }

// Setup 程序初始化配置
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		// log.Println(GetAppPath())

		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("server", ServerSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// 在 go-ini 中可以采用 MapTo 的方式来映射结构体:
// 读取conf/app.ini的section信息，映射到结构体中
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}


