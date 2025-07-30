package pid

type PIDSystem struct {
	Name               string
	ControlledVariable float64
	ErrorSum           float64
	PreviousError      float64
	Kp                 float64
	Kd                 float64
	Ki                 float64
	Time               int
	TargetValue        float64
}

func (pid *PIDSystem) DetermineError() float64 {
	return pid.TargetValue - pid.ControlledVariable
}

func (pid *PIDSystem) Increment(pidError float64) {
	proportional := pid.Kp * pidError
	integral := pid.Ki * pid.ErrorSum
	derivative := pid.Kd * (pidError - pid.PreviousError)
	feedback := proportional + integral + derivative

	pid.ControlledVariable += feedback

	pid.ErrorSum += pidError
	pid.PreviousError = pidError
	pid.Time++
}
