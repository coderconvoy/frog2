package play

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type Car struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	VelocityComponent
}

type MovementEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*VelocityComponent
}

type MovementSystem struct {
	obs []MovementEntity
}

func (ms *MovementSystem) Add(be *ecs.BasicEntity, sc *common.SpaceComponent, vc *VelocityComponent) {
	ms.obs = append(ms.obs, MovementEntity{be, sc, vc})
}

func (ms *MovementSystem) AddByInterface(ob Moveable) {
	ms.Add(ob.GetBasicEntity(), ob.GetSpaceComponent(), ob.GetVelocityComponent())
}

func (ms *MovementSystem) Remove(e ecs.BasicEntity) {
	ms.obs = RemoveMovementEntity(ms.obs, e.ID())
}

func (ms *MovementSystem) Update(d float32) {
	for _, c := range ms.obs {
		c.Position.X += c.Vel.X * d
		c.Position.Y += c.Vel.Y * d
	}
}
