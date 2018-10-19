package config_module

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

type SimpleMSAConfig struct {
	API_SERVER string `json:"api_server"`
}

var simpleMSAConfig SimpleMSAConfig

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func LoadConfig(fi string, v interface{}) error {
	var config_json string
	f, err := os.Open(fi)
	if err != nil {
		return err
	}
	r := bufio.NewReader(f)
	for s, e := Readln(r); e == nil; s, e = Readln(r) {
		config_json += s
	}

	config_json = strings.Replace(config_json, "\t", " ", -1)
	config_json = strings.Replace(config_json, "\r", "", -1)

	b := []byte(config_json)
	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}

	return nil
}
