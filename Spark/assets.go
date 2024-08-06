package Spark

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type AssetManager struct {
	Entity
	BaseDirectory string
	InitiallyLoad map[string]string
	assets        map[string]Asset
	IsInitialized bool

	components AssetSystemComponents
}

type Asset struct {
	Entity Entity
	Name   string
}

type AssetSystemComponents struct {
	BinaryData   *BinaryDataSet
	FileMetaData *FileMetaDataSet
}

func NewAssetManager(baseDir string, toLoad map[string]string) AssetManager {
	return AssetManager{
		BaseDirectory: baseDir,
		InitiallyLoad: toLoad,
		assets:        make(map[string]Asset),
		components: AssetSystemComponents{
			BinaryData:   &BinaryDataSet{},
			FileMetaData: &FileMetaDataSet{},
		},
	}
}

func (am AssetManager) Initialize() {
	am.LoadFiles(am.InitiallyLoad)
	am.IsInitialized = true
}

/**
LOGIC
*/

func (am AssetManager) Get(name string) (Asset, bool) {
	asset, ok := am.assets[name]
	if !ok {
		return Asset{}, false
	}
	return asset, true
}

func (am AssetManager) GetBinaryData(name string) (BinaryData, bool) {
	var asset, ok = am.Get(name)
	if !ok {
		return BinaryData{}, false
	}
	binaryData, ok := am.components.BinaryData.Get(asset.Entity)
	return binaryData, ok
}

func (am AssetManager) GetFileMetaData(name string) (FileMetaData, bool) {
	var asset, ok = am.Get(name)
	if !ok {
		return FileMetaData{}, false
	}
	fileMetaData, ok := am.components.FileMetaData.Get(asset.Entity)
	return fileMetaData, ok
}

func (am AssetManager) LoadFile(name string, path string) *Asset {
	var toLoad = map[string]string{name: path}
	var result = am.LoadFiles(toLoad)
	if len(result) == 0 {
		return nil
	}
	return result[0]
}

func (am AssetManager) LoadFiles(files map[string]string) []*Asset {
	log.Printf("Loading %d asset(s)", len(files))
	var wg sync.WaitGroup
	var loaded = make([]*Asset, 0)
	loadCount := 0
	for name, path := range files {
		wg.Add(1)
		resultChan := make(chan Asset)
		go am.loadFile(name, filepath.Join(am.BaseDirectory, path), resultChan, &wg)
		if resultChan != nil {
			var result = <-resultChan
			am.assets[name] = result
			loaded = append(loaded, &result)
			loadCount++
		}
	}
	wg.Wait()
	log.Printf("%d asset(s) loaded", loadCount)
	return loaded
}

func (am AssetManager) loadFile(name string, path string, resultChan chan Asset, wg *sync.WaitGroup) {
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
		Entity: NewEntity(),
		Name:   name,
	}
	fileData := FileMetaData{
		Path:      filepath.Dir(f.Name()),
		Name:      filepath.Base(f.Name()),
		Extension: filepath.Ext(f.Name()),
	}
	content := BinaryData{
		Data: contents,
	}
	am.components.FileMetaData.Set(asset.Entity, fileData)
	am.components.BinaryData.Set(asset.Entity, content)

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
