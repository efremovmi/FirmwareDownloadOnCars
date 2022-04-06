package usecase

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"mime/multipart"
	"time"
)

func (uc *Handler) TransferFile(file *multipart.FileHeader) error {
	config := &ssh.ClientConfig{
		User:            "rts-ubuntu-01",
		Auth:            []ssh.AuthMethod{ssh.Password("q1w2e3r4t5")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "192.168.0.104:22", config)
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to dial: %s", err))
	}

	sshSession, err := client.NewSession()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to create session: %s", err))
	}

	sshStdin, err := sshSession.StdinPipe()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to create stdin: %s", err))
	}

	err = sshSession.Shell()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to create Shell: %s", err))
	}

	sshSftp, err := sftp.NewClient(client)
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to create sftp: %s", err))
	}

	dstFile, err := sshSftp.Create("./Рабочий стол/main.py")
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = fmt.Fprintf(sshStdin, "sudo -S su\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "q1w2e3r4t5\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "python3 './Рабочий стол/stop.py'\n")
	if err != nil {
		return err
	}

	data, _ := file.Open()
	if _, err := dstFile.ReadFrom(data); err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "\n")
	if err != nil {
		return err
	}

	time.Sleep(1 * time.Second)

	_, err = fmt.Fprintf(sshStdin, "python3 './Рабочий стол/main.py'\n")
	if err != nil {
		return err
	}

	return nil
}

func (uc *Handler) KillProcess() error {
	config := &ssh.ClientConfig{
		User:            "rts-ubuntu-01",
		Auth:            []ssh.AuthMethod{ssh.Password("q1w2e3r4t5")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "192.168.0.104:22", config)
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to dial: %s", err))
	}

	sshSession, err := client.NewSession()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to create session: %s", err))
	}

	sshStdin, err := sshSession.StdinPipe()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to create stdin: %s", err))
	}

	err = sshSession.Shell()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to create Shell: %s", err))
	}

	_, err = fmt.Fprintf(sshStdin, "sudo -S su\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "q1w2e3r4t5\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(sshStdin, "python3 './Рабочий стол/stop.py'\n")
	if err != nil {
		return err
	}

	return nil
}
