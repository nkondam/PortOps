package action

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
)

func KillProcess(pid int) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(pid))
	default:
		cmd = exec.Command("kill", "-9", strconv.Itoa(pid))
	}

	return cmd.Run()

}

func ReleasePort(port int) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "netsta -ano | findstr :"+strconv.Itoa(port)+" | findstr LISTENING && (for /f \"tokens=5\" %a in ('netstat -ano ^| findstr :"+strconv.Itoa(port)+" ^| findstr LISTENING') do @taskkill /f /pid %a)")
	default:
		return nil
	}
	return cmd.Run()
}

func ListProcesses() error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("tasklist")
	default:
		cmd = exec.Command("ps", "aux")
	}
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}
