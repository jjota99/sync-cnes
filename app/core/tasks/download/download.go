package download

import (
	"fmt"
	"github.com/secsy/goftp"
	"main.go/app/core/constants"
	"main.go/app/core/utils/filename"
	"os"
)

func Download() {
	var fileName = filename.GetFileName(constants.BaseFileName)
	var localFilePath = constants.DownloadDir + "/" + fileName
	var remotePath = "cnes/" + fileName

	if _, err := os.Stat(constants.StorageDir); os.IsNotExist(err) {
		err := os.Mkdir(constants.StorageDir, 0755)
		if err != nil {
			fmt.Println("Erro ao criar diretório:", constants.StorageDir, err)
			return
		}
	} else {
		fmt.Println("Diretório", constants.StorageDir, "já existe")
	}

	if _, err := os.Stat(constants.DownloadDir); os.IsNotExist(err) {
		err := os.Mkdir(constants.DownloadDir, 0755)
		if err != nil {
			fmt.Println("Erro ao criar diretório:", constants.DownloadDir, err)
			return
		}
	} else {
		fmt.Println("Diretório", constants.DownloadDir, "já existe")
	}

	client, err := goftp.DialConfig(constants.FtpConfig, constants.FtpHost)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	localFile, err := os.Create(localFilePath)
	if err != nil {
		panic(err)
	}
	defer localFile.Close()

	err = client.Retrieve(remotePath, localFile)
	if err != nil {
		panic(err)
	}
}
