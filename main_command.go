package main

import (
	"fmt"

	"docker-demo/mydocker-dev/container"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// 这里定义了runCommand的Flags，其作用类似运行命令时使用--指定参数
var runCommand = cli.Command{
	Name:  "run",
	Usage: `Create a container with namespace and cgroups limit mydocker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
		cli.StringFlag{
			Name:  "v",
			Usage: "volume",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}

		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}
		tty := context.Bool("ti")
		// resConf := &subsystems.ResourceConfig{
		// 	MemoryLimit: context.String("m"),
		// 	CpuSet:      context.String("cpuset"),
		// 	CpuShare:    context.String("cpushare"),
		// }
		volume := context.String("v")
		Run(tty, cmdArray, volume)
		return nil
	},
}

// 这里定义了initCommand的具体操作，此操作为内部方法，禁止外部调用
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		logrus.Infof("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}
