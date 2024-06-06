package Component

import (
	"github.com/Ignite-Laboratories/JanOS/Internal/Backplane/PerceptionAPI"
	"github.com/Ignite-Laboratories/JanOS/Internal/Config"
)

var This *PerceptionAPI.Component

func Setup() {
	Config.Initialize("config.json")
	This = PerceptionAPI.NewComponent(Config.Current.ID, Config.Current.Network, Config.Current.Address)
}
