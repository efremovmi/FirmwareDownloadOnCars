package usecase

import (
	"Clever_City/internal/repository"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

type Handler struct {
	repo       *repository.Repo
	sshClient  *ssh.Client
	sshSession *ssh.Session
	sshSftp    *sftp.Client
	sshStdin   io.WriteCloser
}

func NewHandler(repo *repository.Repo, sshClient *ssh.Client) *Handler {
	sshSession, err := sshClient.NewSession()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to create session: %s", err))
	}

	sshStdin, err := sshSession.StdinPipe()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to create stdin: %s", err))
	}

	sshSession.Stdout = os.Stdout
	sshSession.Stderr = os.Stderr

	err = sshSession.Shell()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to create Shell: %s", err))
	}

	sshSftp, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to create sftp: %s", err))
	}

	return &Handler{repo: repo, sshSftp: sshSftp, sshClient: sshClient, sshSession: sshSession, sshStdin: sshStdin}
}

func (uc *Handler) Close() {
	_ = uc.sshSession.Close()
	_ = uc.sshSftp.Close()
}
