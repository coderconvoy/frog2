package play

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"image/color"
)

type PlayScene struct {
}

func (*PlayScene) Type() string { return "PlayScene" }

func (*PlayScene) Preload() {}

func (*PlayScene) Setup(w *ecs.World) {
	common.SetBackground(color.White)
	rs := &common.RenderSystem{}

	fg := NewFrog()
	rs.Add(&fg.BasicEntity, &fg.RenderComponent, &fg.SpaceComponent)

	w.AddSystem(rs)

}
