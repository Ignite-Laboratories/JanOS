package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

/**
Mailbox
*/

type LoadAssetMsg struct {
	ToLoad map[string]string
}

func NewLoadAssetMsg() LoadAssetMsg {
	return LoadAssetMsg{
		ToLoad: make(map[string]string),
	}
}

func (lam LoadAssetMsg) Load(name string, path string) {
	lam.ToLoad[name] = path
}

func (LoadAssetMsg) Type() string {
	return "LoadAssetMsg"
}

/**
Structures
*/

type FileDataComponent struct {
	Path      string
	Name      string
	Extension string
	Contents  []byte
}

type Asset struct {
	ecs.BasicEntity
	Name     string
	FileData FileDataComponent
}

/**
System
*/

type AssetSystem struct {
	world         *ecs.World
	BaseDirectory string
	ToLoad        map[string]string
	Assets        map[string]*Asset
}

func (as *AssetSystem) New(w *ecs.World) {
	log.Println("[Inception] - Asset System")
	as.world = w
	as.Assets = make(map[string]*Asset)

	// Perform initial asset loading
	loadFiles(as.ToLoad, as)

	// Listen for load request messages
	engo.Mailbox.Listen("LoadAssetMsg", func(msg engo.Message) {
		assetMsg, ok := msg.(LoadAssetMsg)
		if !ok {
			log.Fatal("Message is not of type LoadAssetMsg")
		}
		log.Printf("[Asset] LoadFileMsg - %s", assetMsg.ToLoad)
		loadFiles(assetMsg.ToLoad, as)
	})
}

func loadFiles(toLoad map[string]string, as *AssetSystem) {
	log.Printf("[Asset] Loading %d asset(s)", len(toLoad))
	var wg sync.WaitGroup
	loadCount := 0
	for name, path := range toLoad {
		wg.Add(1)
		resultChan := make(chan *Asset)
		go loadFile(name, filepath.Join(as.BaseDirectory, path), resultChan, &wg)
		if resultChan != nil {
			as.Assets[name] = <-resultChan
			loadCount++
		}
	}
	wg.Wait()
	log.Printf("[Asset] %d asset(s) loaded", loadCount)
}

func loadFile(name string, path string, resultChan chan *Asset, wg *sync.WaitGroup) {
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
		BasicEntity: ecs.NewBasic(),
		Name:        name,
		FileData: FileDataComponent{
			Path:      filepath.Dir(f.Name()),
			Name:      filepath.Base(f.Name()),
			Extension: filepath.Ext(f.Name()),
			Contents:  contents,
		},
	}

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

func (as *AssetSystem) Remove(ecs.BasicEntity) {}

func (as *AssetSystem) Update(dt float32) {

}
