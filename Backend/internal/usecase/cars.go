package usecase

import (
	"mime/multipart"
	"os"
)

func (uc *Handler) TransferFile(file *multipart.FileHeader) error {
	srcFile, err := os.Open("main.py")
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := uc.sshSftp.Create("./Рабочий стол/main.py")
	if err != nil {
		return err
	}
	defer dstFile.Close()

	data, _ := file.Open()
	if _, err := dstFile.ReadFrom(data); err != nil {
		return err
	}

	return nil
}
