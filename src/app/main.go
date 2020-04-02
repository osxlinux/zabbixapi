package main

import (
	"fmt"
	"os"
	"zabbixapi/src/cli"
)

func InitCliMgr() (err error) {

	if err = cli.InitHttpCliMgr(); err != nil {
		return
	}

	if err = cli.InitZabbixMgr(); err != nil {
		return
	}
	return
}

func main() {
	var err error
	if err = InitCliMgr(); err != nil {
		fmt.Printf(`Err: %s`, err.Error())
		os.Exit(1)
	}

	if err = cli.GlobCommangMgr.RunCommand(); err != nil {
		fmt.Printf(`Err: %s`, err.Error())
		os.Exit(1)
	}
}
