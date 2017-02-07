package play

import "engo.io/engo"

type VelocityComponent struct {
	Vel engo.Point
}

func (v *VelocityComponent) GetVelocityComponent() *VelocityComponent {
	return v
}
