package main

import "golang.org/x/crypto/ssh"

func connect(login, password string) bool {
	sshConfig := &ssh.ClientConfig{
		User: login,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}

	_, err := ssh.Dial("tcp", "138.68.108.233:22", sshConfig)
	if err == nil {
		return true
	}

	return false
}
