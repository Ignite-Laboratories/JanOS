package Systems

import "github.com/Ignite-Laboratories/JanOS/Logic/Systems/Assets"

func NewAssetSystem(baseDir string, toLoad map[string]string) Assets.AssetSystem {
	return Assets.NewAssetSystem(baseDir, toLoad)
}
