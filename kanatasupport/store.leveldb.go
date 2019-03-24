package kanatasupport

import (
	"github.com/premeidoworks/kanata/api"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	defaultLeveldb = &goleveldbImpl{}
)

type goleveldbImpl struct {
	db *leveldb.DB
}

func (this *goleveldbImpl) Init(config *api.StoreInitConfig) error {
	f := config.Details["leveldb.dir"]
	db, err := leveldb.OpenFile(f, nil)
	if err != nil {
		return err
	}

	defaultLeveldb.db = db
	return nil
}

func (this *goleveldbImpl) SaveMessage(message *api.Message) error {
	return this.db.Put([]byte(message.MessageId), message.Body, nil)
}

func (this *goleveldbImpl) ObtainOnceMessage(queue int64, maxCount int) ([]*api.Message, error) {
	var result = make([]*api.Message, 16)

	it := this.db.NewIterator(nil, nil)
	defer it.Release()

	idx := 0
	for it.Next() && idx < 16 {
		result[idx] = &api.Message{
			MessageId: string(it.Key()),
			Body:      it.Value(),
		}
		err := this.db.Delete(it.Key(), nil)
		if err != nil {
			return nil, err
		}
		idx++
	}
	if it.Error() != nil {
		return nil, it.Error()
	}

	return result[:idx], nil
}
