package systems

import (
	"github.com/EngoEngine/ecs"
	"log"
	"os"
	"sync"
)

type AssetSystem struct {
	world       *ecs.World
	FilesToLoad map[string]string
	files       map[string]*os.File
}

func (as *AssetSystem) New(w *ecs.World) {
	log.Println("[Inception] - Asset System")
	as.world = w
	as.files = make(map[string]*os.File)

	var wg sync.WaitGroup
	for name, path := range as.FilesToLoad {
		wg.Add(1)
		go loadFile(name, path, &wg, as)
	}
	wg.Wait()
	log.Println("[Asset] All assets loaded")
}

func loadFile(name string, path string, wg *sync.WaitGroup, as *AssetSystem) {
	f, err := os.Open("C:\\source\\ignite\\janos\\assets\\" + path)
	if err != nil {
		log.Println(err)
		return
	}
	as.files[name] = f
	wg.Done()
	log.Printf("[Asset] Loaded - [%s] - %s", name, path)
}

func (as *AssetSystem) Remove(ecs.BasicEntity) {}

func (as *AssetSystem) Update(dt float32) {

}
