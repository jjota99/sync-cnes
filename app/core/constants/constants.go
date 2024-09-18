package constants

import (
	"github.com/secsy/goftp"
	"os"
	"time"
)

// Dirs
var StorageDir = "storage"
var DownloadDir = "storage/download"
var ExtractDir = "storage/extract"

// Files
var FilesToDownload = []string{"tbEquipe", "tbEstabelecimento"}
var BaseFileName = "BASE_DE_DADOS_CNES_"

// FTP
var FtpHost = "ftp.datasus.gov.br"
var FtpConfig = goftp.Config{
	User:               "anonymous",
	Password:           "anonymous@example.com",
	ConnectionsPerHost: 10,
	Timeout:            10 * time.Second,
	Logger:             os.Stderr,
}
