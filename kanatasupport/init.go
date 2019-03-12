package kanatasupport

import "github.com/premeidoworks/kanata/api"

func init() {
	api.RegisterStoreProvider("pgx", defaultPgStore)
	api.RegisterKanataConfigParse("default", TomlConfigParser{})
	api.RegisterUUIDProvider("gouuid", uuidGenerator{})
	api.RegisterMarshallingProvider("default", gogoprotobufMarshalImpl{})
}
