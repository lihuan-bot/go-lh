/*
 * @Author: lihuan
 * @Date: 2021-09-23 08:21:31
 * @LastEditTime: 2021-09-24 10:08:29
 * @Email: 17719495105@163.com
 */
package utils

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppPort string `json:"app_port"`
	AppHost string `json:"app_host"`
	AppMode string `json:"app_mode"`
	Database  DatabaseCfg `json:"database"`
}

type DatabaseCfg struct {
	Driver string `json:"driver"`
	User string `json:"user"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
	DbName string `json:"db_name"`
	Charset string `json:"charset"`
	ParseTime bool `json:"parse_time"`
	Loc string `json:"loc"`
}

var _cfg *Config = nil

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, err
}