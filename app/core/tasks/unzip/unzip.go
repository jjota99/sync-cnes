package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"main.go/app/core/constants"
	"main.go/app/core/utils/filename"
	"main.go/app/core/utils/listfiles"
	"os"
	"slices"
)

func Unzip() {
	fileName := filename.GetFileName(constants.BaseFileName)
	ListFiles := listfiles.GetListFiles(constants.FilesToDownload)

	if _, err := os.Stat(constants.ExtractDir); os.IsNotExist(err) {
		err := os.Mkdir(constants.ExtractDir, 0755)
		if err != nil {
			fmt.Println("Erro ao criar diretório:", constants.ExtractDir, err)
			return
		}
	} else {
		fmt.Println("Diretório", constants.ExtractDir, "já existe!")
	}

	r, err := zip.OpenReader(constants.DownloadDir + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		if slices.Contains(ListFiles, f.Name) {
			newFilePath := fmt.Sprintf("%s/%s", constants.ExtractDir, f.Name)

			rc, err := f.Open()
			if err != nil {
				log.Fatal("Não foi possível ler o arquivo:", f.Name)
			}
			defer rc.Close()

			fileContent, err := io.ReadAll(rc)
			if err != nil {
				log.Fatal("Erro ao ler o conteúdo do arquivo:", f.Name)
			}

			err = os.WriteFile(newFilePath, fileContent, 0755)
			if err != nil {
				fmt.Println("Erro ao escrever o arquivo:", err)
			}
		}
	}
}
