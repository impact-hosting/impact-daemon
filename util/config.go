package util

import (
	"io/ioutil"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	ConfigFile     string = "/etc/impact/config.yml"
	RootDataDir    string = "/var/lib/impact"
	VolumesDataDir string = "/var/lib/impact/volumes"
	BackupDataDir  string = "/var/lib/impact/backups"
	LogDir         string = "/var/log/impact"
)

var (
	mutex     sync.RWMutex
	writeLock sync.Mutex
)

type ApiConfig struct {
	// FQDN or IP Address of this Daemon
	Hostname string `default:"0.0.0.0" yaml:"hostname"`

	// The Port to run the Daemon API on
	Port int `default:"8080" yaml:"port"`

	// The Port the SFTP server will run on
	SftpPort int `default:"2022" yaml:"sftp_port"`

	// SSL options for the Daemon
	Ssl struct {
		// Whether the Daemon should check for valid SSL
		Enabled bool `default:"false" yaml:"enabled"`

		// Path to the cert file
		Cert string `yaml:"crt"`

		// Path to the cert key file
		CertKey string `yaml:"crt_key"`
	}

	// The max size for files in MB to be uploaded using the Panel
	UploadFileSize int `default:"100" yaml:"upload_file_size"`
}

func ReadFile(path string) (data, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
}

func GetServerFromContext(ctx *gin.Context) {

}
