package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

func Get_process() []string {
	var ssh_process_list []string
	processes, err := process.Processes()
	if err != nil {
		fmt.Errorf("Error: %v", err)
		os.Exit(1)
	}

	for _, process := range processes {
		nme, _ := process.Cmdline()
		if strings.Contains(nme, "ssh") && strings.Contains(nme, "@") {
			ssh_process_list = append(ssh_process_list, nme)
		}
	}
	return ssh_process_list
}
