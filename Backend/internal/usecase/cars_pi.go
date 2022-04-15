package usecase

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
)

var (
	configMap = map[string]*ssh.ClientConfig{
		"carPi_1": {
			User:            "rts-ubuntu-01",
			Auth:            []ssh.AuthMethod{ssh.Password("q1w2e3r4t5")},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
		"carPi_2": {
			User:            "rts-ubuntu-02",
			Auth:            []ssh.AuthMethod{ssh.Password("q1w2e3r4t5")},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}}

	IpMap = map[string]string{
		"carPi_1": "192.168.1.238:22",
		"carPi_2": "192.168.1.225:22",
	}
)

func (uc *Handler) TransferFile(file io.Reader, name string) error {
	if _, ok := configMap[name]; !ok {
		return fmt.Errorf("I have no idea what it is. func: %s. name: %s", "TransferFile", name)
	}
	fmt.Println("Updating raspberry #:", name)

	if _, ok := IpMap[name]; !ok {
		return fmt.Errorf("I have no idead what is it")
	}

	client, err := ssh.Dial("tcp", IpMap[name], configMap[name])
	if err != nil {
		return fmt.Errorf("failed to dial: %s", err)
	}

	sshSession, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %s", err)
	}

	sshStdin, err := sshSession.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin: %s", err)
	}

	err = sshSession.Shell()
	if err != nil {
		return fmt.Errorf("failed to create Shell: %s", err)
	}

	fmt.Println("Upload to", name)

	sshSftp, err := sftp.NewClient(client)
	if err != nil {
		return fmt.Errorf("failed to create sftp: %s", err)
	}

	_, err = fmt.Fprintf(sshStdin, "sudo pkill -9 python3\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "python3 ./Desktop/stop.py\n")
	if err != nil {
		return err
	}

	dstFile, err := sshSftp.Create("./Desktop/main.py")
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "sudo pkill -9 python3\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "python3 ./Desktop/stop.py\n")
	if err != nil {
		return err
	}

	if _, err := dstFile.ReadFrom(file); err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "sudo pkill -9 python3\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "python3 ./Desktop/stop.py\n")
	if err != nil {
		return err
	}

	return nil
}

func (uc *Handler) StartFile(name string) error {
	if _, ok := configMap[name]; !ok {
		return fmt.Errorf("I have no idea what is it")
	}

	if _, ok := IpMap[name]; !ok {
		return fmt.Errorf("I have no idead what is it")
	}

	client, err := ssh.Dial("tcp", IpMap[name], configMap[name])
	if err != nil {
		return fmt.Errorf("failed to dial: %s", err)
	}

	sshSession, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %s", err)
	}

	sshStdin, err := sshSession.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin: %s", err)
	}

	err = sshSession.Shell()
	if err != nil {
		return fmt.Errorf("failed to create Shell: %s", err)
	}

	fmt.Println("Starting on", name)

	_, err = fmt.Fprintf(sshStdin, "sudo pkill -9 python3\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "python3 ./Desktop/stop.py\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "python3 './Desktop/main.py'\n")
	if err != nil {
		return err
	}

	return nil
}

func (uc *Handler) KillProcess(name string) error {
	fmt.Println("STOPPING")

	if _, ok := configMap[name]; !ok {
		return fmt.Errorf("I have no idea what is it")
	}

	if _, ok := IpMap[name]; !ok {
		return fmt.Errorf("I have no idead what is it")
	}

	client, err := ssh.Dial("tcp", IpMap[name], configMap[name])
	if err != nil {
		return fmt.Errorf("failed to dial: %s", err)
	}

	sshSession, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %s", err)
	}

	sshStdin, err := sshSession.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin: %s", err)
	}

	err = sshSession.Shell()
	if err != nil {
		return fmt.Errorf("failed to create Shell: %s", err)
	}

	_, err = fmt.Fprintf(sshStdin, "sudo pkill -9 python3\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "sudo pkill -9 python\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "python3 './Desktop/stop.py'\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "python3 './Desktop/stop.py'\n")
	if err != nil {
		return err
	}

	return nil
}
