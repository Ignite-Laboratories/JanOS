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

type AssetMessage struct {
	Value string
}

func (AssetMessage) Type() string {
	return "AssetMessage"
}

type Asset struct {
	Name          string
	FilePath      string
	FileName      string
	FileExtension string
	Data          []byte
}

type AssetSystem struct {
	world       *ecs.World
	FilesToLoad map[string]string
	files       map[string]*Asset
}

func (as *AssetSystem) New(w *ecs.World) {
	log.Println("[Inception] - Asset System")
	as.world = w
	as.files = make(map[string]*Asset)

	engo.Mailbox.Listen("AssetMessage", func(msg engo.Message) {
		assetMsg, ok := msg.(AssetMessage)
		if !ok {
			log.Fatal("Message is not of type AssetMessage")
		}
		log.Printf("[Asset].[Message] LoadFile - %s", assetMsg.Value)
	})

	var wg sync.WaitGroup
	for name, path := range as.FilesToLoad {
		wg.Add(1)
		go loadFile(name, path, &wg, as)
	}
	wg.Wait()
	log.Println("[Asset] All assets loaded")
}

func loadFile(name string, path string, wg *sync.WaitGroup, as *AssetSystem) {
	// Open the file
	f, err := os.Open("C:\\source\\ignite\\janos\\assets\\" + path)
	if err != nil {
		log.Fatal(err)
	}

	// Read the file's contents
	contents, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// Build the metadata and store it in memory
	asset := &Asset{
		Name:          name,
		FilePath:      filepath.Dir(f.Name()),
		FileName:      filepath.Base(f.Name()),
		FileExtension: filepath.Ext(f.Name()),
		Data:          contents,
	}

	as.files[name] = asset

	// Close the file
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	wg.Done()
	log.Printf("[Asset] Loaded - [%s] - %s", name, path)
}

func (as *AssetSystem) Remove(ecs.BasicEntity) {}

func (as *AssetSystem) Update(dt float32) {

}
