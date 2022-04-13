package usecase

import (
	"fmt"
	"os/exec"
)

var (
	IpMaps = map[string]string{
		"carAr_1": "192.168.1.169",
		"carAr_2": "192.168.1.248",
		"carAr_3": "192.168.1.169",
		"carAr_4": "192.168.1.169",
	}
)

func (uc *Handler) UpdateArduino(name string, file string) error {
	fmt.Println("Updating arduino #:", name)
	if _, ok := IpMaps[name]; !ok {
		return fmt.Errorf("I have no idea what it is")
	}

	cmd := exec.Command("python", "postFile.py", "--filename", file, "--ip", IpMaps[name])
	out, err := cmd.Output()

	fmt.Println(string(out))

	return err
}
