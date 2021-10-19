package main

import (
	"authentication/api"
	"authentication/infra"
	"authentication/model"
)

func main() {
	i := infra.New("config/config.json")
	i.SetMode()
	i.Migrate(&model.User{})

	api.NewServer(i).Run()
}
