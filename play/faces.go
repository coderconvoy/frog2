package play

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type BasicFace interface {
	GetBasicEntity() *ecs.BasicEntity
}

type SpaceFace interface {
	GetSpaceComponent() *common.SpaceComponent
}

type RenderFace interface {
	GetRenderComponent() *common.RenderComponent
}

type VelocityFace interface {
	GetVelocityComponent() *VelocityComponent
}

type Moveable interface {
	BasicFace
	SpaceFace
	VelocityFace
}
