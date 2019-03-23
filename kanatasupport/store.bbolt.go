package kanatasupport

import (
	"time"

	"github.com/premeidoworks/kanata/api"

	bolt "github.com/etcd-io/bbolt"
)

var (
	defaultBboltStore = new(bboltStore)
	bboltBucket       = []byte("message")
)

type bboltStore struct {
	db *bolt.DB
}

func (this *bboltStore) Init(config *api.StoreInitConfig) error {
	f := config.Details["bbolt.file"]
	db, err := bolt.Open(f, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	defaultBboltStore.db = db
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bboltBucket)
		return err
	})
	return err
}

func (this *bboltStore) SaveMessage(message *api.Message) error {
	err := this.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("message"))
		return b.Put([]byte(message.MessageId), message.Body)
	})
	return err
}

func (this *bboltStore) ObtainOnceMessage(queue int64, maxCount int) ([]*api.Message, error) {
	var result []*api.Message
	err := this.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bboltBucket)
		c := b.Cursor()
		idx := 0
		var r []*api.Message
		for k, v := c.First(); k != nil && idx < 16; k, v = c.Next() {
			m := &api.Message{
				MessageId: string(k),
				Body:      v,
			}
			r = append(r, m)
			err := c.Delete()
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
