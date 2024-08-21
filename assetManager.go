package JanOS

import (
	"io"
	"os"
)

type assetManager struct {
	assets map[string]*asset
}

func (mgr *assetManager) GetName() string { return "Assets" }

func newAssetManager() *assetManager {
	return &assetManager{
		assets: make(map[string]*asset),
	}
}

type asset struct {
	Name string
	Data any
}

func (a *asset) GetName() string { return a.Name }

// GetAsset returns the asset for the provided name.
func (mgr *assetManager) GetAsset(name string) *asset {
	return mgr.assets[name]
}

// LoadAsset opens the file at the provided location and loads its contents.
func (mgr *assetManager) LoadAsset(name string, path string) (*asset, error) {
	path = AssetPath + path
	Logging.Printf(mgr, "Loading asset '%s' from '%s'", name, path)
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
	Logging.Printf(mgr, "Asset '%s' loaded", name)
	return a, nil
}