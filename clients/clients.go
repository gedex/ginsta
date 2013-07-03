package clients

import (
	"bufio"
	"encoding/json"
	"github.com/gedex/ginsta/utils"
	"github.com/gedex/go-instagram/instagram"
	"os"
	"path/filepath"
)

const (
	// This ginsta version
	Version = "0.1"

	// CLI name
	Name = "ginsta"

	// Default client_id from http://instagram.com/developer/clients/manage/
	DefaultClientID = "8f2c0ad697ea4094beb2b1753b7cde9c"

	// Default client_secret from http://instagram.com/developer/clients/manage/
	DefaultClientSecret = "2bf3da5329624c68bca14642c13627d1"

	// Default scope
	DefaultScope = "basic,comments,relationships,likes"
)

var (
	// Location of config's file
	DefaultConfigFilename = filepath.Join(os.Getenv("HOME"), ".config", "ginsta")
)

type Client struct {
	Version        string
	Name           string
	Env            []string
	ConfigFilename string
	Config         *Config
	Instagram      *instagram.Client
}

type Config struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
}

func ReadConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			f, err = os.Create(filename)
		}
		utils.Check(err)
	}
	defer f.Close()

	dec := json.NewDecoder(bufio.NewReader(f))

	c := &Config{}
	err = dec.Decode(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func SaveConfig(filename string, c *Config) error {
	err := os.MkdirAll(filepath.Dir(filename), 0771)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(c)
}
