package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	for {
		lookCheck()

		time.Sleep(time.Minute * 1)
	}
}

func lookCheck() {
	retLs := exec.Command("bash", "-c", "df -lh")
	retLsBytes, _ := retLs.Output()

	if !strings.Contains(string(retLsBytes), "/media/pi/新加卷") && strings.Contains(string(retLsBytes), "/mnt/pan") {
		fmt.Println("一切正常中...")
		return
	}

	// 卸载2个目录
	cmdXinjiaJuan1:= exec.Command("sh", "-c", "umount /media/pi/新加卷1")
	cmdXinjiaJuan1.Run()
	retUmountXinjiaJuan1, _ := exec.Command("sh", "-c", "umount /media/pi/新加卷").Output()
	//if err != nil && !strings.Contains(err.Error(), "not mounted") {
	//	fmt.Println("出现错误 新加卷：", string(retUmountXinjiaJuan), "|", err.Error())
	//	return
	//}
	retBytesUmountPan, _ := exec.Command("sh", "-c", "umount  /mnt/pan").Output()
	//if err != nil && !strings.Contains(err.Error(), "not mounted") {
	//	fmt.Println("出现错误 pan:", string(retBytesUmountPan),"|",err.Error())
	//	return
	//}
	fmt.Println("执行umount", string(retBytesUmountPan), retUmountXinjiaJuan1)

	// 查看/mnt/pan还有无其他内容
	retOwnCloudStr, _ := exec.Command("bash", "-c", "ls /mnt/pan/owncloud/").Output()
	fmt.Println("owncloud有哪些内容", "{" + string(retOwnCloudStr) + "}")
	if strings.TrimSpace(string(retOwnCloudStr)) == "owncloud.log" {
		exec.Command("sh", "-c", "rm -rf /mnt/pan/*").Run()
		fmt.Println("只包含了owncloud.log")
	}

	// 最后执行挂载命令
	retMount, err := exec.Command("sh", "-c", "mount -U 3A0A-D56D  /mnt/pan -t exfat -o nls=utf-8,umask=007,uid=1000,gid=1000").Output()
	if err != nil {
		fmt.Println("挂载报错", err.Error(), string(retMount))
	}
	fmt.Println(time.Now().UTC().Format("2006-01-02_15:04:05"), "本次执行完毕...", string(retMount))
}
