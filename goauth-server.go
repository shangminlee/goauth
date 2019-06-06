package main

import (
    "github.com/shangminlee/goauth/cmd"
    "github.com/shangminlee/goauth/log"
    "github.com/urfave/cli"
    "os"
)

var (
    cliApp     *cli.App
    configFile string
)

// 初始化 命令行, 默认构造方法，类似Spring @PostConstruct
func init() {
    // Initialize a CLI APP
    cliApp = cli.NewApp()
    cliApp.Name = "goauth-oauth2-server"
    cliApp.Usage = "Goauth Oauth 2.0 Server"
    cliApp.Author = "Richard knop & arrows"
    cliApp.Email = "shangmlee@foxmai.com"
    cliApp.Version = "0.0.1"
    cliApp.Flags = []cli.Flag{
        cli.StringFlag{
            Name:        "configFile",  // 参数名称
            Value:       "config.yml",  // 默认配置文件
            Destination: &configFile,
        },
    }
}

func main(){
    // Set the CLI app commands
    cliApp.Commands = []cli.Command{
        {
            Name:   "migrate",
            Usage:  "run migrations", // 数据库表结构迁移
            Action: func(c *cli.Context) error {
                return cmd.Migrate(configFile)
            },
        },
        {
            Name:   "loaddata",
            Usage:  "load data from fixture", // 加载数据, 加载初始化数据
            Action: func(c *cli.Context) error {
                return cmd.LoadData(c.Args(), configFile)
            },
        },
        {
            Name: "runserver",
            Usage: "run webserver", // 开始服务
            Action: func(c *cli.Context) error {
                return cmd.RunServer(configFile)
            },
        },

    }

    if err := cliApp.Run(os.Args); err != nil {
        log.FATAL.Println(err)
    }

    log.INFO.Printf("配置文件地址 : %v \n", configFile)

}

