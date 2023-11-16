package carrota

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Feishu []struct {
		BotID             string `yaml:"bot_id"`
		BotName           string `yaml:"bot_name"`
		AppID             string `yaml:"app_id"`
		AppSecret         string `yaml:"app_secret"`
		EncryptKey        string `yaml:"encrypt_key"`
		VerificationToken string `yaml:"verification_token"`
		CallbackAPI       string `yaml:"callback_api"`
	}
	PluginCenter struct {
		IP   string `yaml:"ip"`
		Port string `yaml:"port"`
		API  struct {
			BasePath      string `yaml:"base"`
			MessageReport string `yaml:"message_report"`
		}
	}
	API struct {
		BasePath       string `yaml:"base"`
		MessageSendAPI string `yaml:"message_send"`
	}
	Port string `yaml:"port"`
}

var C *conf

func readConf(filename string) (*conf, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var conf conf
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	return &conf, nil
}

func LoadConfig(file string) {
	c, err := readConf(file)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Config loaded failed.")
		return
	}
	C = c
	fmt.Println(c)
	fmt.Println("Config loaded success.")
}
