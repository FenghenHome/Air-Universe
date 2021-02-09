package main

import (
	"encoding/json"
	"github.com/crossfw/Air-Universe/pkg/structures"
	"log"
	"os"
)

// Default config
var (
	baseCfg = &structures.BaseConfig{
		Panel: structures.Panel{
			Type: "sspanel",
		},
		Proxy: structures.Proxy{
			Type:       "v2ray",
			AlertID:    1,
			InTags:     []string{"p0"},
			APIAddress: "127.0.0.1",
			APIPort:    10085,
			LogPath:    "./v2.log",
		},
		Sync: structures.Sync{
			Interval:  60,
			FailDelay: 5,
			Timeout:   5,
		},
	}
)

func ParseBaseConfig(configPath *string) (*structures.BaseConfig, error) {
	file, err := os.Open(*configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(baseCfg); err != nil {
		return nil, err
	}
	log.Println(*baseCfg)
	return baseCfg, nil
}
