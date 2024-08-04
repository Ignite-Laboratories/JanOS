package systems

type TextEntity struct {
	//ecs.BasicEntity
	//common.SpaceComponent
	//common.RenderComponent
}

type TextSystem struct {
	entities []TextEntity
	Text1    TextEntity
	//assetSystem *AssetSystem
}

func (ts *TextSystem) New(w int) { //*ecs.World) {
	//// Ensure we have the asset system
	//for _, system := range w.Systems() {
	//	switch sys := system.(type) {
	//	case *AssetSystem:
	//		ts.assetSystem = sys
	//	}
	//}
	//
	//err := engo.Files.LoadReaderData("segoepr.ttf", bytes.NewReader(ts.assetSystem.Assets["segoe-print"].FileData.Contents))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Load in the default font
	//fnt := &common.Font{
	//	URL:  "segoepr.ttf",
	//	FG:   color.Black,
	//	Size: 20,
	//}
	//err = fnt.CreatePreloaded()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//ts.Text1 = TextEntity{BasicEntity: ecs.NewBasic()}
	//ts.Text1.RenderComponent.Drawable = common.Text{
	//	Font: fnt,
	//	Text: "Hello, World!",
	//}
	//ts.Text1.SetShader(common.TextHUDShader)
	//ts.Text1.RenderComponent.SetZIndex(1001)
	//ts.Text1.SpaceComponent = common.SpaceComponent{
	//	Position: engo.Point{X: 0, Y: engo.WindowHeight() - 200},
	//}
	//for _, system := range w.Systems() {
	//	switch sys := system.(type) {
	//	case *common.RenderSystem:
	//		sys.Add(&ts.Text1.BasicEntity, &ts.Text1.RenderComponent, &ts.Text1.SpaceComponent)
	//	}
	//}
}

func (ts *TextSystem) Update(dt float32) {

}

func (ts *TextSystem) Remove(basic int) { //ecs.BasicEntity) {
	//del := -1
	//for index, e := range ts.entities {
	//	if e.BasicEntity.ID() == basic.ID() {
	//		del = index
	//		break
	//	}
	//}
	//if del >= 0 {
	//	ts.entities = append(ts.entities[:del], ts.entities[del+1:]...)
	//}
}
