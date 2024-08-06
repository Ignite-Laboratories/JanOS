package OLD

/**
Instrument Viewer
*/

type InstrumentViewer struct {
	//ecs.BasicEntity
	//common.RenderComponent
	//common.SpaceComponent
}

type InstrumentViewerSystem struct {
	//	world *ecs.World
}

func (is *InstrumentViewerSystem) New(w int) { //*ecs.World) {
	//viewer := InstrumentViewer{BasicEntity: ecs.NewBasic()}
	//
	//viewer.SpaceComponent = common.SpaceComponent{
	//	Position: engo.Point{(engo.WindowWidth() / 2) - 100, (engo.WindowHeight() / 2) - 100},
	//	Width:    200,
	//	Height:   200,
	//}
	//
	//viewerImage := image.NewUniform(color.RGBA{205, 205, 205, 255})
	//viewerNRGBA := common.ImageToNRGBA(viewerImage, 200, 200)
	//viewerImageObj := common.NewImageObject(viewerNRGBA)
	//viewerTexture := common.NewTextureSingle(viewerImageObj)
	//
	//viewer.RenderComponent = common.RenderComponent{
	//	Drawable: viewerTexture,
	//	Scale:    engo.Point{1, 1},
	//	Repeat:   common.Repeat,
	//}
	//viewer.RenderComponent.SetShader(common.HUDShader)
	//viewer.RenderComponent.SetZIndex(1)
	//
	//for _, system := range w.Systems() {
	//	switch sys := system.(type) {
	//	case *common.RenderSystem:
	//		sys.Add(&viewer.BasicEntity, &viewer.RenderComponent, &viewer.SpaceComponent)
	//	}
	//}
}

func (is *InstrumentViewerSystem) Remove() {} //ecs.BasicEntity) {}

func (is *InstrumentViewerSystem) Update(dt float32) {

}
