package main

import (
	"fmt"
	"log"

	"github.com/gcleroux/IFT630-SCAM/pkg/people"
	"github.com/gcleroux/IFT630-SCAM/pkg/utils"
)

func main() {
	conf, err := utils.LoadConfig("./conf/config.yml")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf)
	fmt.Println(people.MayorHello())
}
