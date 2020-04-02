package cli

import (
	"github.com/urfave/cli/v2"
	"os"
	"zabbixapi/src/utils"
)

func init() {
	var app *cli.App
	var author *cli.Author
	//
	var authors []*cli.Author
	authors = append(authors, author)
	author = &cli.Author{
		Name:  "seraphic",
		Email: "tiaopidaodanmaomidadui@google.com",
	}

	app = cli.NewApp()
	app.Name = `zabbixcli`
	app.Usage = `zabbixcli [--help,-h]`
	app.Description = `zabbix golang client develop by Seraphic`
	app.Copyright = `@2020xxxx`
	app.Authors = authors
	GlobCommangMgr = &CommandMgr{apps: app}
}

type CommandMgr struct {
	apps *cli.App
}

var GlobCommangMgr *CommandMgr

func (c *CommandMgr) RunCommand() (err error) {
	var subCommand *cli.Command
	var subCommands []*cli.Command
	subCommand = &cli.Command{
		Name: "add",
		Action: func(context *cli.Context) (err error) {

			if err = utils.NewConfigMgr(context.String(`c`)); err != nil {
				return cli.NewExitError(err.Error(), 127)
			}
			if err = GlobZabbixMgr.CreateHost(
				context.String(`a`),
				context.String(`t`),
				context.String(`g`),
				context.String(`e`),
			); err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
			return cli.Exit(`succ`, 0)
		},
		Usage: `创建节点`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "template",
				Aliases: []string{`t`},
				Usage:   "指定模板ID",
				Value:   "10001", //默认10001
			},
			&cli.StringFlag{
				Name:    "group",
				Aliases: []string{`g`},
				Usage:   "指定主机组ID",
				Value:   "2", //默认2 Zabbix servers Linux servers
			},
			&cli.StringFlag{
				Name:    "ipaddr",
				Aliases: []string{`e`},
				Usage:   "指定节点IP地址",
				Value:   "192.168.1.1",
			},
			&cli.StringFlag{
				Name:     "hostname",
				Aliases:  []string{`a`},
				Usage:    "指定主机名",
				Value:    "bj-prod-ep-host01",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{`c`},
				Usage:    "指定配置文件",
				Value:    "conf/zabbix.json",
				Required: true,
			},
		},
	}
	subCommands = append(subCommands, subCommand)
	c.apps.Commands = subCommands
	if err = c.apps.Run(os.Args); err != nil {
		return
	}
	return
}
