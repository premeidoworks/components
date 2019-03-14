package kanatasupport

import (
	"github.com/premeidoworks/kanata/api"

	"github.com/BurntSushi/toml"
)

type T struct {
	Global struct {
		StoreProvider   string `toml:"store_provider"`
		UUIDProvider    string `toml:"uuid_provider"`
		MarshalProvider string `toml:"marshal_provider"`
	} `toml:"global"`
	Store struct {
		Details map[string]string `toml:"details"`
	} `toml:"store"`
}

type TomlConfigParser struct {
}

func (TomlConfigParser) ParseConfigFile(path string) (*api.KanataConfig, error) {
	var t T
	_, err := toml.DecodeFile(path, &t)
	if err != nil {
		return nil, err
	}
	result := &api.KanataConfig{
		StoreProvider:   t.Global.StoreProvider,
		UUIDProvider:    t.Global.UUIDProvider,
		MarshalProvider: t.Global.MarshalProvider,
		StoreConfig: &api.StoreInitConfig{
			Details: t.Store.Details,
		},
	}
	if result.StoreConfig.Details == nil {
		result.StoreConfig.Details = make(map[string]string)
	}
	return result, nil
}
