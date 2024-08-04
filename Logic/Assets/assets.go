package Assets

import (
	"github.com/Ignite-Laboratories/JanOS/logic"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Asset struct {
	ID   Logic.Entity
	Name string
}

type BinaryDataComponent struct {
	Data []byte
}

type FileDataComponent struct {
	Path      string
	Name      string
	Extension string
}

type AssetSystem struct {
	World *Logic.World

	BaseDirectory string
	ToLoad        map[string]string
	Assets        map[string]*Asset

	components struct {
		BinaryData *Logic.Components[BinaryDataComponent]
		FileData   *Logic.Components[FileDataComponent]
	}
}

func NewAssetSystem() *AssetSystem {
	return &AssetSystem{
		ToLoad: make(map[string]string),
		Assets: make(map[string]*Asset),
	}
}

func (as *AssetSystem) Initialize(w *Logic.World) {
	as.World = w

	loadFiles(as)
}

func (as *AssetSystem) Tick(w *Logic.World) {

}

/**
LOGIC
*/

func loadFiles(as *AssetSystem) {
	log.Printf("[Asset] Loading %d asset(s)", len(as.ToLoad))
	var wg sync.WaitGroup
	loadCount := 0
	for name, path := range as.ToLoad {
		wg.Add(1)
		resultChan := make(chan *Asset)
		go loadFile(name, filepath.Join(as.BaseDirectory, path), resultChan, &wg, as)
		if resultChan != nil {
			as.Assets[name] = <-resultChan
			loadCount++
		}
	}
	wg.Wait()
	log.Printf("[Asset] %d asset(s) loaded", loadCount)
}

func loadFile(name string, path string, resultChan chan *Asset, wg *sync.WaitGroup, as *AssetSystem) {
	// Open the file
	f, err := os.Open(path)
	if err != nil {
		log.Printf("[Asset] Error opening file - %s", err)
		wg.Done()
	}

	// Read the file's contents
	contents, err := io.ReadAll(f)
	if err != nil {
		log.Printf("[Asset] Error reading file contents - %s", err)
		wg.Done()
	}

	// Build the entity metadata and store it in memory
	asset := &Asset{
		ID:   Logic.NewEntity(),
		Name: name,
	}
	fileData := FileDataComponent{
		Path:      filepath.Dir(f.Name()),
		Name:      filepath.Base(f.Name()),
		Extension: filepath.Ext(f.Name()),
	}
	content := BinaryDataComponent{
		Data: contents,
	}
	as.components.FileData.Set(asset.ID, fileData)
	as.components.BinaryData.Set(asset.ID, content)

	// Close the file
	err = f.Close()
	if err != nil {
		log.Printf("[Asset] Error closing file - %s", err)
		wg.Done()
	}

	resultChan <- asset
	wg.Done()
	log.Printf("[Asset] Loaded - %s", name)
}
