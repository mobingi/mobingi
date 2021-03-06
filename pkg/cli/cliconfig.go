package cli

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/mobingi/mobingi/client/timeout"
	"github.com/mobingilabs/mobingi-sdk-go/pkg/cmdline"
	"github.com/mobingilabs/mobingi-sdk-go/pkg/pretty"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// CliConfig is the object representation of our config file. The field tags for YAML marshaling and
// unmarshaling match the defined cli constants with prefix 'Config'.
type CliConfig struct {
	AccessToken     string `yaml:"access_token"`
	BaseApiUrl      string `yaml:"api_url"`
	BaseRegistryUrl string `yaml:"registry_url"`
	BaseSesha3Url   string `yaml:"sesha3_url"`
	ApiVersion      string `yaml:"api_version"`
	Indent          int    `yaml:"indent"`
	Timeout         int64  `yaml:"timeout"`
	Verbose         bool   `yaml:"verbose"`
	Debug           bool   `yaml:"debug"`
}

func (c *CliConfig) WriteToConfig() error {
	contents, err := yaml.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "marshal failed")
	}

	if !exists(c.ConfigDir()) {
		_ = os.Mkdir(c.ConfigDir(), os.ModePerm)
	}

	return ioutil.WriteFile(c.ConfigFile(), contents, 0644)
}

func (c *CliConfig) Reload() error {
	contents, err := ioutil.ReadFile(c.ConfigFile())
	if err != nil {
		return errors.Wrap(err, "readfile failed")
	}

	err = yaml.Unmarshal(contents, c)
	if err != nil {
		return errors.Wrap(err, "unmarshal failed")
	}

	return nil
}

func (c *CliConfig) ConfigFile() string {
	p := c.path()
	if p == "" {
		return p
	}

	return filepath.Join(c.ConfigDir(), ConfigFileName)
}

func (c *CliConfig) ConfigDir() string {
	p := c.path()
	if p == "" {
		return p
	}

	dirname := cmdline.Args0()
	pair := strings.Split(cmdline.Args0(), ".")
	// check for .exe (Windows)
	if len(pair) == 2 {
		if pair[1] == "exe" {
			dirname = pair[0]
		}
	}

	return filepath.Join(p, "."+dirname)
}

func (c *CliConfig) path() string {
	var p string
	p, _ = homedir.Dir()
	return p
}

func SetDefaultCliConfig() error {
	defcnf := CliConfig{
		BaseApiUrl:      ProductionBaseApiUrl,
		BaseRegistryUrl: ProductionBaseRegistryUrl,
		ApiVersion:      ApiVersion,
		Indent:          pretty.Pad,
		Timeout:         timeout.Timeout,
	}

	err := defcnf.WriteToConfig()
	err = errors.Wrap(err, "write default config failed")
	ErrorExit(err, 1)

	return viper.ReadInConfig()
}

func ReadCliConfig() *CliConfig {
	cnf := &CliConfig{}
	err := cnf.Reload()
	if err != nil {
		return nil
	}

	return cnf
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}
