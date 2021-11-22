package util

import (
	"os/exec"
	"strings"
)

func StartContainer(id string) (output string, err error) {
	command := exec.Command("docker", "start", id)
	stdout, err := command.Output()
	if err != nil {
		return nil, err
	}
	return string(stdout), nil
}

func ExistsContainer(id string) (exists bool, err error) {
	// retrieve a list of container IDs
	command := exec.Command("docker", "container", "ls", "-q")
	stdout, err := command.Output()
	if err != nil {
		return false, err
	}
	return strings.Contains(string(stdout), id), nil
}

func StopContainer(id string) (output string, err error) {
	command := exec.Command("docker", "stop", id)
	stdout, err := command.Output()
	if err != nil {
		return nil, err
	}
	return string(stdout), nil
}

func RestartContainer(id string) (output string, err error) {
	command := exec.Command("docker", "restart", id)
	stdout, err := command.Output()
	if err != nil {
		return nil, err
	}
	return string(stdout), nil
}

func KillContainer(id string) (output string, err error) {
	command := exec.Command("docker", "kill", id)
	stdout, err := command.Output()
	if err != nil {
		return nil, err
	}
	return string(stdout), nil
}

func GetContainerLogs(id string) (output string, err error) {
	command := exec.Command("docker", "logs", "-tf", id)
	stdout, err := command.Output()
	if err != nil {
		return nil, err
	}
	return string(stdout), nil
}

func RunContainerCommand(id string, command, string) (response string, err error) {
	command := exec.Command("docker", "exec", "-w /home/server", id, comamnd)
	stdout, err := command.Output()
	if err != nil {
		return nil, err
	}
	return string(stdout), nil
}
