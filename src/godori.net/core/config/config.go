package config

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Server struct {
		BindAddress     string `json:"bind_address,string"`
		Port            int    `json:"port"`
		MaxSessionCount int    `json:"max_session_count"`
		MaxPacketSize   int    `json:"max_packet_size"`
	}
	Game struct {
		Map struct {
			MaxMapCount  int `json:"max_map_count"`
			MaxUserCount int `json:"max_user_count"`
		}
	}
	DB struct {
	}
}

/*
type Game {

}
*/
var JSON Config

func Init() {
	/*
		f, err := os.Open("../config/config.json")
		if err != nil {
			fmt.Println("[ERR] Load config.json failed.")
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(f)
		defer f.Close()
	*/
	byteValue := []byte(`{"server":{"bind_address":"127.0.0.1","port":50002,"max_session_count":10,"max_packet_size":100},"game":{"maps":{"max_map_count":257,"max_user_count":50}},"db":{}}`)
	res := &Config{}
	json.Unmarshal(byteValue, &res)
	fmt.Println(res)
}
