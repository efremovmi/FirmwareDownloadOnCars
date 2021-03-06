package worker

import (
	"Clever_City/internal/service/global"
	"Clever_City/internal/usecase"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	fileMap = map[string]string{
		"carPi_1": "",
		"carPi_2": "",
		"carAr_1": "",
		"carAr_2": "",
		"carAr_3": "",
		"carAr_4": "",
	}
	fileMapMu = &sync.Mutex{}

	statusMap = map[string]string{
		"carPi_1": "1",
		"carPi_2": "1",
		"carAr_1": "1",
		"carAr_2": "1",
		"carAr_3": "1",
		"carAr_4": "1",
	}
	statusMapMu = &sync.Mutex{}

	idMap = map[string]string{
		"carPi_1": "1",
		"carPi_2": "3",
		"carAr_1": "2",
		"carAr_2": "4",
		"carAr_3": "5",
		"carAr_4": "6",
	}

	userGroupMap = map[string]string{
		"carPi_1": "0",
		"carPi_2": "0",
		"carAr_1": "0",
		"carAr_2": "0",
		"carAr_3": "0",
		"carAr_4": "0",
	}
	userGroupMapMu = &sync.Mutex{}
)

func ProcessingRaspberryFiles(uc *usecase.Handler) error {
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		err, all := global.AskAll()

		if err != nil {
			println(err.Error())
			continue
		}

		for key, value := range fileMap {
			id, _ := strconv.Atoi(idMap[key])

			if all.Data[id-1].File != value {
				err := global.GetFile(key, all.Data[id-1].File)
				if err != nil {
					println(err.Error())
				}

				file, err := os.Open(fmt.Sprintf("%s_%s", key, all.Data[id-1].File))

				err = uc.TransferFile(file, key)
				if err != nil {
					println(err.Error())
				}

				_ = file.Close()

				if _, ok := usecase.IpMap[key]; ok {
					err = os.Remove(fmt.Sprintf("%s_%s", key, all.Data[id-1].File))
					if err != nil {
						println(err.Error())
					}
				}

				fileMapMu.Lock()
				fileMap[key] = all.Data[id-1].File
				fileMapMu.Unlock()
			}
		}

		for key, value := range statusMap {
			id, _ := strconv.Atoi(idMap[key])

			if all.Data[id-1].Status == "0" && value == "1" {

				if _, ok := usecase.IpMap[key]; ok {
					err := uc.StartFile(key)
					if err != nil {
						println(err.Error())
						continue
					}
				} else {
					file, err := os.Open(fmt.Sprintf("%s_%s", key, all.Data[id-1].File))
					err = uc.UpdateArduino(key, fmt.Sprintf("%s_%s", key, all.Data[id-1].File))
					if err != nil {
						println(err.Error())
					}

					_ = file.Close()
					err = os.Remove(fmt.Sprintf("%s_%s", key, all.Data[id-1].File))

					if err != nil {
						println(err.Error())
						continue
					}

					if err != nil {
						println(err.Error())
						continue
					}
				}

				statusMapMu.Lock()
				statusMap[key] = "0"
				statusMapMu.Unlock()

			} else if all.Data[id-1].Status == "1" && value == "0" {
				statusMapMu.Lock()
				statusMap[key] = "1"
				statusMapMu.Unlock()

				err = uc.KillProcess(key)
				if err != nil {
					println(err.Error())
					continue
				}
			}
		}
	}

	return nil
}

func ProcessName(name string) {

}
