package Systems

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/Ignite-Laboratories/JanOS/Logic/Common"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type AssetSystem struct {
	Logic.Entity
	BaseDirectory string
	InitiallyLoad map[string]string
	Assets        map[string]Asset
	IsInitialized bool

	components AssetSystemComponents
}

type Asset struct {
	Logic.Entity
	Name string
}

type AssetSystemComponents struct {
	BinaryData *Common.BinaryDataSet
	FileData   *Common.FileDataSet
}

func NewAssetSystem(baseDir string, toLoad map[string]string) AssetSystem {
	return AssetSystem{
		BaseDirectory: baseDir,
		InitiallyLoad: toLoad,
		Assets:        make(map[string]Asset),
		components: AssetSystemComponents{
			BinaryData: &Common.BinaryDataSet{},
			FileData:   &Common.FileDataSet{},
		},
	}
}

func (sys AssetSystem) GetName() string         { return "Asset" }
func (sys AssetSystem) GetEntity() Logic.Entity { return sys.Entity }

func (sys AssetSystem) Initialize(world *Logic.World) {
	sys.LoadFiles(sys.InitiallyLoad, world)
	sys.IsInitialized = true
}

func (sys AssetSystem) Tick(world *Logic.World, inbox Logic.Inbox) {

}

/**
LOGIC
*/

func (sys AssetSystem) LoadFile(name string, path string, world *Logic.World) *Asset {
	var toLoad = map[string]string{name: path}
	var result = sys.LoadFiles(toLoad, world)
	if len(result) == 0 {
		return nil
	}
	return result[0]
}

func (sys AssetSystem) LoadFiles(files map[string]string, world *Logic.World) []*Asset {
	log.Printf("Loading %d asset(s)", len(files))
	var wg sync.WaitGroup
	var loaded = make([]*Asset, 0)
	loadCount := 0
	for name, path := range files {
		wg.Add(1)
		resultChan := make(chan Asset)
		go sys.loadFile(name, filepath.Join(sys.BaseDirectory, path), resultChan, &wg, world)
		if resultChan != nil {
			var result = <-resultChan
			sys.Assets[name] = result
			loaded = append(loaded, &result)
			loadCount++
		}
	}
	wg.Wait()
	log.Printf("%d asset(s) loaded", loadCount)
	return loaded
}

func (sys AssetSystem) loadFile(name string, path string, resultChan chan Asset, wg *sync.WaitGroup, world *Logic.World) {
	// Open the file
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file - %s", err)
		wg.Done()
	}

	// Read the file's contents
	contents, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("Error reading file contents - %s", err)
		wg.Done()
	}

	// Build the entity metadata and store it in memory
	asset := Asset{
		Entity: Logic.NewEntity(),
		Name:   name,
	}
	fileData := Common.FileData{
		Path:      filepath.Dir(f.Name()),
		Name:      filepath.Base(f.Name()),
		Extension: filepath.Ext(f.Name()),
	}
	content := Common.BinaryData{
		Data: contents,
	}
	sys.components.FileData.Set(asset.Entity, fileData)
	sys.components.BinaryData.Set(asset.Entity, content)
	world.AddEntity(asset.Entity, asset)

	// Close the file
	err = f.Close()
	if err != nil {
		log.Fatalf("Error closing file - %s", err)
		wg.Done()
	}

	resultChan <- asset
	wg.Done()
	log.Printf("Loaded - %s", name)
}
