package play

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type PlayScene struct {
}

func (*PlayScene) Type() string { return "PlayScene" }

func (*PlayScene) Preload() {}

func (*PlayScene) Setup(w *ecs.World) {
	common.SetBackground(color.White)

	systems := &SysList{
		RenderSys:   &common.RenderSystem{},
		MovementSys: &MovementSystem{},
	}

	spawnSys := &SpawnSystem{Systems: systems}

	fg := NewFrog()
	systems.RenderSys.AddByInterface(fg)

	w.AddSystem(systems.RenderSys)
	w.AddSystem(systems.MovementSys)
	w.AddSystem(spawnSys)

}
