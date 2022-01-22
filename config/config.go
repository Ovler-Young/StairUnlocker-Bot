package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"gopkg.in/yaml.v3"

	C "github.com/Dreamacro/clash/constant"
)

func init() {
	fmt.Printf(fmt.Sprintf("StairUnlock-Bot %s %s %s with %s %s\n", C.Version, runtime.GOOS, runtime.GOARCH, runtime.Version(), C.BuildTime))
	flag.BoolVar(&Help, "h", false, "this help")
	flag.BoolVar(&Version, "v", false, "show current version of StairUnlock")
	flag.StringVar(&ConfPath, "f", "", "specify configuration file")
	flag.Parse()

	//initial config.yaml
	var (
		buf []byte
	)

	if ConfPath != "" {
		buf, _ = ioutil.ReadFile(ConfPath)
	} else {
		_, err := os.Stat("config.yaml")
		if err != nil {
			b, _ := ioutil.ReadFile("config.example.yaml")
			_ = ioutil.WriteFile("config.yaml", b, 644)
		}
		buf, _ = ioutil.ReadFile("config.yaml")
	}
	_ = yaml.Unmarshal(buf, &BotCfg)
}

func UnmarshalRawConfig(buf []byte) (*RawConfig, error) {
	rawCfg := &RawConfig{}
	if err := yaml.Unmarshal(buf, rawCfg); err != nil {
		return nil, err
	}
	return rawCfg, nil
}
