package main

import (
	"fmt"

	pid "github.com/shumy26/gontrollers/PID"
)

func main() {
	x := 0.0
	kp := 0.8
	kd := 0.1
	ki := 0.3

	pidsystem := pid.PIDSystem{
		Name:               "test",
		ControlledVariable: x,
		Kp:                 kp,
		Kd:                 kd,
		Ki:                 ki,
		Time:               0,
		TargetValue:        140.0,
	}

	for i := 0; i < 1000; i++ {
		fmt.Println(pidsystem.ControlledVariable)
		pidError := pidsystem.DetermineError()
		pidsystem.Increment(pidError)
	}
}
