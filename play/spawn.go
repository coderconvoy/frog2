package play

import (
	"image/color"
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type SysList struct {
	RenderSys   *common.RenderSystem
	MovementSys *MovementSystem
}

type SpawnSystem struct {
	since   float32
	nCars   int
	Systems *SysList
}

func (ss *SpawnSystem) Remove(e ecs.BasicEntity) {
}

func (ss *SpawnSystem) Update(d float32) {
	ss.since = ss.since + d
	if ss.since > 1 {
		ss.since = 0
		ss.nCars++

		c := &Car{
			BasicEntity:    ecs.NewBasic(),
			SpaceComponent: common.SpaceComponent{Position: engo.Point{float32(300 + 10*ss.nCars), 300}, Width: 100, Height: 50},
			RenderComponent: common.RenderComponent{
				Drawable: common.Rectangle{},
				Color:    color.Black,
			},
			VelocityComponent: VelocityComponent{engo.Point{rand.Float32()*20 - 10, rand.Float32()*20 - 10}},
		}

		ss.Systems.RenderSys.AddByInterface(c)
		ss.Systems.MovementSys.AddByInterface(c)

	}
}
