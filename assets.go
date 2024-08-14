package JanOS

import (
	"io"
	"os"
)

type assetManager struct {
	assets map[string]*asset
}

func newAssetManager() *assetManager {
	return &assetManager{assets: make(map[string]*asset)}
}

type asset struct {
	Name string
	Data any
}

func (mgr *assetManager) GetName() string {
	return "Assets"
}

func (mgr *assetManager) GetAsset(name string) *asset {
	return mgr.assets[name]
}

func (mgr *assetManager) LoadAsset(name string, path string) (*asset, error) {
	Universe.Printf(mgr, "Loading asset '%s' from '%s'", name, path)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	a := &asset{
		Name: name,
		Data: contents,
	}
	mgr.assets[name] = a
	Universe.Printf(mgr, "Asset '%s' loaded", name)
	return a, nil
}
