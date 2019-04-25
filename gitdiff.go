package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func main() {
	command := "/bin/bash"
	params := []string{"-c","git log|grep ^commit|awk '{print $2}'|head -n 5"}
	//执行cmd命令: ls -l
	execCommand(command, params)
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
//	cmd := exec.Command("/bin/bash", "-c", "git log|grep ^commit|awk '{print $2}'")
	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}
	
	cmd.Start()

	reader := bufio.NewReader(stdout)

	line, err2 := reader.ReadString('\n')

	if err2 != nil || io.EOF == err2 {
		cmd.Wait()
		return true
	}

	//实时循环读取输出流中的一行内容
	for {
		
		line1, err3 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 ||err3 != nil || io.EOF == err3{
			break
		}
		line = strings.Replace(line, "\n", "", -1) 
		line1 = strings.Replace(line1, "\n", "", -1) 
		fmt.Println("git diff "+line+" "+line1)
		line = line1
	}
	
	cmd.Wait()
	return true
}