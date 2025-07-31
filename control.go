package main

import (
	"fmt"

	pid "github.com/shumy26/gontrollers/PID"
)

func control(ch chan float64) {
	x := 0.0
	kp := 0.4
	ki := 0.005
	kd := 0.0005

	pidsystem := pid.PIDSystem{
		Name:               "test",
		ControlledVariable: x,
		Kp:                 kp,
		Kd:                 kd,
		Ki:                 ki,
		Time:               0,
		TargetValue:        140.0,
	}

	for i := 0; i < 300; i++ {
		ch <- pidsystem.ControlledVariable
		fmt.Println(pidsystem.ControlledVariable)
		pidError := pidsystem.DetermineError()
		pidsystem.Increment(pidError)
	}
	close(ch)
}
