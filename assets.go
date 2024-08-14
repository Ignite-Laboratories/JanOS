package JanOS

import (
	"io"
	"os"
)

type assetManager struct {
	assets map[string]Asset
}

func newAssetManager() *assetManager {
	return &assetManager{
		assets: make(map[string]Asset),
	}
}

type Asset struct {
	Name string
	Data any
}

func (a *assetManager) GetName() string {
	return "Assets"
}

func (a *assetManager) GetAsset(name string) Asset {
	return a.assets[name]
}

func (a *assetManager) LoadAsset(name string, path string) error {
	Universe.Printf(a, "Loading asset '%s' from '%s'", name, path)
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	contents, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	asset := Asset{
		Name: name,
		Data: contents,
	}
	a.assets[name] = asset
	Universe.Printf(a, "Asset '%s' loaded", name)
	return nil
}
