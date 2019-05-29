package main

import (
	"fmt"
	"github.com/EOIDC/goipmi"
)

type Ipmi struct {
	con ipmi.Connection
}

func (e *Ipmi) Init() error {
	cli := ipmi.Client{}
	sensorList, err := cli.GetSensorList(0)
	for _, v := range sensorList {
		fmt.Println(v.SensorType, ":", v.Value)
	}
	return err
}
