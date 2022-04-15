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
	if _, ok := IpMaps[name]; !ok {
		return fmt.Errorf("I have no idea what it is. func: %s. name: %s", "UpdateArduino", name)
	}

	fmt.Println("Updating arduino #:", name)

	cmd := exec.Command("python", "postFile.py", "--filename", file, "--ip", IpMaps[name])
	out, err := cmd.Output()

	fmt.Println(string(out))

	return err
}
