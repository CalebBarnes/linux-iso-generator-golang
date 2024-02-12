package generator

import (
	"fmt"
	"os"

	"github.com/calebbarnes/linux-iso-generator-golang/services/utils"
)

const ubuntuIsoUrl = "https://releases.ubuntu.com/20.04.3/ubuntu-20.04.3-live-server-amd64.iso"

func EnsureUbuntuIsoExists() error {
	isoPath := "./tmp/ubuntu-20.04.3-live-server-amd64.iso"

	if _, err := os.Stat(isoPath); os.IsNotExist(err) {
		fmt.Printf("Ubuntu ISO not found. Downloading to tmp directory: %s\n", isoPath)
		println("")
		err := utils.DownloadFile(isoPath, ubuntuIsoUrl)
		if err != nil {
			return err
		}
	}
	return nil
}
