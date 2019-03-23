package kanatasupport

import (
	"github.com/premeidoworks/kanata/api"

	"github.com/dgraph-io/badger"
)

var (
	defaultBadgerStore = new(badgerStore)
)

type badgerStore struct {
	db *badger.DB
}

func (this *badgerStore) Init(config *api.StoreInitConfig) error {
	opts := badger.DefaultOptions
	opts.Dir = config.Details["badger.dir"]
	opts.ValueDir = config.Details["badger.valudedir"]
	db, err := badger.Open(opts)
	if err != nil {
		return err
	}
	defaultBadgerStore.db = db
	return nil
}

func (this *badgerStore) SaveMessage(message *api.Message) error {
	err := this.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(message.MessageId), message.Body)
	})
	return err
}

func (this *badgerStore) ObtainOnceMessage(queue int64, maxCount int) ([]*api.Message, error) {
	opt := badger.IteratorOptions{
		PrefetchValues: true,
		PrefetchSize:   16,
	}
	var result []*api.Message
	err := this.db.Update(func(txn *badger.Txn) error {
		it := txn.NewIterator(opt)
		defer it.Close()
		var r []*api.Message
		idx := 0
		for it.Rewind(); it.Valid() && idx < 16; it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(val []byte) error {
				m := &api.Message{
					MessageId: string(k),
					Body:      val,
				}
				r = append(r, m)
				return nil
			})
			if err != nil {
				return err
			}
			err = txn.Delete(k)
			if err != nil {
				return err
			}
			idx++
		}
		result = r
		return nil
	})
	return result, err
}
