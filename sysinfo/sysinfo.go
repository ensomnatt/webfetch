package sysinfo

import (
	"fmt"
	"os/exec"
	"strings"
)

type KeyValue struct {
	Key   string
	Value string
}

func GetSystemInfo() []KeyValue {
	var KeyValuePairs []KeyValue
	cmd := exec.Command("sh", "-c", "fastfetch --logo none")
	output, err := cmd.Output()
	if err != nil {
		_ = fmt.Errorf("error with getting output of fastfetch: %v", err)
		return nil
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if !strings.Contains(line, ":") {
			continue
		}
		parts := strings.Split(line, ":")

		KeyValuePairs = append(KeyValuePairs, KeyValue{
			Key:   parts[0],
			Value: parts[1],
		})
	}

	return KeyValuePairs
}
