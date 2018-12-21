package main

import (
	"fmt"
	"gobalt/src/api"

	//"gobalt/src/api"
	"gobalt/src/config"
)

func main() {
	fmt.Println("Starting Gobalt...")

	// Generate config
	c := config.New()
	config := c.GenerateConfig()

	f := api.New(config)

	//s := transport.Session{}
	//pub := transport.New(config, &s)
	//
	//pubStructure := transport.BuildPub("silver", "test.ping")
	//loadStructure := transport.BuildLoad(pubStructure)

	//pub.Publish(loadStructure)
	_, _ = f.Login(config.Username, config.Password, "pam")

	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Printf("Creds: %s\n", creds)
	//}
	//f.Keys()
	//f.Key("silver")
}
