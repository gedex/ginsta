package commands

import (
	"fmt"
	"github.com/gedex/ginsta/clients"
	"github.com/gedex/ginsta/utils"
	"os"
)

var (
	cmdConfig = &Command{
		Callback: runConfig,
		Usage:    "config [CONFIG_KEY] [CONFIG_VALUE]",
		Short:    fmt.Sprintf("Get and set %s config", clients.Name),
		Long: `You can set/get configuration with this command. Valid CONFIG_KEYs
are: "client_id", "cliend_secret", and "access_token".`,
	}
)

func runConfig(r *Runner, cmd *Command, args []string) {
	var key, val string

	switch len(args) {
	case 0:
		showConfig("all", r.Client)
	case 1:
		key = args[0]
		showConfig(key, r.Client)
	case 2:
		key, val = args[0], args[1]
		updateConfig(key, val, r.Client)
	default:
		utils.Check(fmt.Errorf("too many arguments"))
	}
	os.Exit(0)
}

func showConfig(k string, c *clients.Client) {
	switch k {
	case "client_id":
		fmt.Println(c.Config.ClientID)
	case "client_secret":
		fmt.Println(c.Config.ClientSecret)
	case "access_token":
		fmt.Println(c.Config.AccessToken)
	case "all":
		m := map[string]string{
			"Client ID (client_id)":         c.Config.ClientID,
			"Client Secret (client_secret)": c.Config.ClientSecret,
			"Access token (access_token)":   c.Config.AccessToken,
		}
		for i, v := range m {
			fmt.Printf("%-30s %s\n", i, v)
		}
	default:
		utils.Check(fmt.Errorf("unknown config key '%s'", k))
	}
}

func updateConfig(k, v string, c *clients.Client) {
	conf := c.Config
	switch k {
	case "client_id":
		conf.ClientID = v
	case "client_secrent":
		conf.ClientSecret = v
	case "access_token":
		conf.AccessToken = v
	default:
		utils.Check(fmt.Errorf("unknown config key '%s'", k))
	}
	err := clients.SaveConfig(c.ConfigFilename, conf)
	utils.Check(err)
}
