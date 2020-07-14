package main

import (
	"docker-demo/mydocker-dev/container"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Run 函数
func Run(tty bool, comArry []string, volume string) {
	parent, writePipe := container.NewParentProcess(tty, volume)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	// use mydocker-cgroup as cgroup name
	// 创建cgroup manager，并通过调用set和apply设置资源限制并使限制在容器上生效
	// cgroupManager := cgroups.NewCgroupManager("mydocker-cgroup")
	// defer cgroupManager.Destroy()
	// // 设置资源限制
	// cgroupManager.Set(res)
	// // 将容器进程加入到各个subsystem挂载对应的cgroup中
	// cgroupManager.Apply(parent.Process.Pid)
	// 对容器设置完限制后，初始化容器
	sendInitCommand(comArry, writePipe)
	parent.Wait()
	mntURL := "/root/mnt"
	rootURL := "/root"
	container.DeleteWorkSpace(rootURL, mntURL, volume)
	os.Exit(0)
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
