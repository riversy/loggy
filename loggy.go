package main

import (
	"flag"
	"fmt"
	"github.com/totherme/unstructured"
	"os"
	"path/filepath"
	"strings"
)

type AppConfig struct {
	ConfigPath, KeyPath, Scope, RemotePath, LocalPath string
	IsTail, IsHead                                    bool
	NumOfLines                                        int
}

var appCfg = &AppConfig{}

func initConfig() {
	flag.StringVar(&appCfg.ConfigPath, "p", "~/.config/loggy/pool.yml", "")
	flag.StringVar(&appCfg.KeyPath, "i", "~/.ssh/id_rsa", "")

	flag.BoolVar(&appCfg.IsHead, "head", false, "")
	flag.BoolVar(&appCfg.IsTail, "tail", false, "")

	flag.IntVar(&appCfg.NumOfLines, "n", 10, "")

	if len(os.Args) != 4 {
		panic("usage should be the following: loggy [Scope] [remote_file] [local_file]")
	}

	appCfg.Scope = os.Args[1]
	appCfg.RemotePath = os.Args[2]
	appCfg.LocalPath = os.Args[3]

	appCfg.KeyPath, _ = filepath.Abs(getFullKeyPath(appCfg.KeyPath))
	appCfg.ConfigPath, _ = filepath.Abs(getFullKeyPath(appCfg.ConfigPath))
	appCfg.LocalPath, _ = filepath.Abs(getFullKeyPath(appCfg.LocalPath))
}

func main() {
	initConfig()

	hosts := getHosts()

	fmt.Println(hosts)
}

func getHosts() []string {
	configYaml, err := os.ReadFile(appCfg.ConfigPath)
	if err != nil {
		panic(err)
	}

	poolData, err := unstructured.ParseYAML(string(configYaml))
	if err != nil {
		panic(err)
	}

	poolPayloadData, err := poolData.GetByPointer("/foo/staging")
	if err != nil {
		panic(err)
	}

	if !poolPayloadData.IsList() {
		panic("scoped value has to be list")
	}

	hostsList, err := poolPayloadData.ListValue()
	if err != nil {
		panic(err)
	}

	hosts := make([]string, len(hostsList))
	for i, v := range hostsList {
		hosts[i] = v.UnsafeStringValue()
	}

	return hosts
}

func getFullKeyPath(keyPath string) string {
	return strings.Replace(keyPath, "~", os.Getenv("HOME"), 1)
}
