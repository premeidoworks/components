package kanatasupport

import "github.com/premeidoworks/kanata/api"

func init() {
	api.RegisterStoreProvider("pgx", defaultPgStore)
	api.RegisterStoreProvider("badger", defaultBadgerStore)
	api.RegisterStoreProvider("bbolt", defaultBboltStore)
	api.RegisterStoreProvider("simplememory", defaultSimpleMemory)
	api.RegisterStoreProvider("leveldb", defaultLeveldb)

	api.RegisterKanataConfigParse("default", TomlConfigParser{})
	api.RegisterUUIDProvider("gouuid", uuidGenerator{})
	api.RegisterMarshallingProvider("default", gogoprotobufMarshalImpl{})
}
