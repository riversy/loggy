package main

import (
	"flag"
	"fmt"
	"os"
)

type AppConfig struct {
	config, keyPath, scope, remotePath, localPath string
	isTail, isHead                                bool
	num                                           int
}

var appCfg = &AppConfig{}

func main() {
	flag.StringVar(&appCfg.config, "c", "~/.config/loggy/inventory.yml", "")
	flag.StringVar(&appCfg.keyPath, "i", "~/.ssh/id_rsa", "")

	flag.BoolVar(&appCfg.isHead, "head", false, "")
	flag.BoolVar(&appCfg.isTail, "tail", false, "")

	flag.IntVar(&appCfg.num, "n", 10, "")

	if len(os.Args) != 4 {
		panic("usage should be the following: loggy [scope] [remote_file] [local_file]")
	}

	appCfg.scope = os.Args[1]
	appCfg.remotePath = os.Args[2]
	appCfg.localPath = os.Args[3]

	fmt.Println(appCfg)
}
