// Copyright 2013 The ginsta AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ginsta Command line client written in Go that speaks to Instagram API.
package main

import (
	"github.com/gedex/ginsta/clients"
	"github.com/gedex/ginsta/commands"
	"github.com/gedex/ginsta/utils"
	"github.com/gedex/go-instagram/instagram"
	"os"
)

func main() {
	conf, err := clients.ReadConfig(clients.DefaultConfigFilename)
	if err != nil {
		conf = &clients.Config{
			ClientID:     clients.DefaultClientID,
			ClientSecret: clients.DefaultClientSecret,
		}
		err = clients.SaveConfig(clients.DefaultConfigFilename, conf)
	}
	utils.Check(err)

	inst := instagram.NewClient(nil)
	inst.AccessToken = conf.AccessToken

	runner := &commands.Runner{
		Args: os.Args[1:],
		Client: &clients.Client{
			Version:        clients.Version,
			Name:           clients.Name,
			Env:            os.Environ(),
			ConfigFilename: clients.DefaultConfigFilename,
			Config:         conf,
			Instagram:      inst,
		},
	}
	err = runner.Execute()
	utils.Check(err)
}
