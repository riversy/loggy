package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/riversy/loggy/connection"
	"github.com/riversy/loggy/utils"
	"github.com/vbauerster/mpb/v8"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type AppConfig struct {
	ConfigPath string
	KeyPath    string
	Scope      string
	RemotePath string
	LocalPath  string
	IsTail     bool
	IsHead     bool
	NumOfLines int
}

var appCfg = &AppConfig{}

func initConfig() {
	flag.StringVar(&appCfg.ConfigPath, "p", "~/.config/loggy/pool.yml", "")
	flag.StringVar(&appCfg.KeyPath, "i", "~/.ssh/id_rsa", "")

	flag.BoolVar(&appCfg.IsHead, "head", false, "")
	flag.BoolVar(&appCfg.IsTail, "tail", false, "")

	flag.IntVar(&appCfg.NumOfLines, "n", 10, "")

	if len(os.Args) != 4 {
		panic("usage should be the following: loggy [--(tail|head)] [-n 10] <scope> <remote_file> <local_file>")
	}

	appCfg.Scope = os.Args[1]
	appCfg.RemotePath = os.Args[2]
	appCfg.LocalPath = os.Args[3]

	appCfg.KeyPath, _ = filepath.Abs(getFullKeyPath(appCfg.KeyPath))
	appCfg.ConfigPath, _ = filepath.Abs(getFullKeyPath(appCfg.ConfigPath))
	appCfg.LocalPath, _ = filepath.Abs(getFullKeyPath(appCfg.LocalPath))
}

func getFullKeyPath(keyPath string) string {
	return strings.Replace(keyPath, "~", os.Getenv("HOME"), 1)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	initConfig()

	hosts := utils.GetHosts(appCfg.ConfigPath, appCfg.Scope)

	var wg sync.WaitGroup
	p := mpb.NewWithContext(ctx, mpb.WithWaitGroup(&wg))

	for _, uri := range hosts {

		wg.Add(1)

		conn := connection.NewDownloadConnection(
			uri,
			appCfg.KeyPath,
			appCfg.RemotePath,
			"cat",
			utils.GetTempFileName(),
		)

		go func() {
			for {
				status := <-conn.StatusCh
				fmt.Println(status)
			}
		}()

		err := conn.Init()
		if err != nil {
			panic(err)
		}

		fmt.Println(conn.TargetPath)
	}

	fmt.Println(hosts)

	p.Wait()
}
