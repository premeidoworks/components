package kanatasupport

import (
	"sync"

	"github.com/premeidoworks/kanata/api"
)

var (
	defaultSimpleMemory = &simpleMemoryStore{}
)

type simpleMemoryStore struct {
	m    map[string][]byte
	lock sync.Mutex
}

func (this *simpleMemoryStore) Init(config *api.StoreInitConfig) error {
	this.m = make(map[string][]byte)
	return nil
}

func (this *simpleMemoryStore) SaveMessage(message *api.Message) error {
	this.lock.Lock()
	this.m[message.MessageId] = message.Body
	this.lock.Unlock()
	return nil
}

func (this *simpleMemoryStore) ObtainOnceMessage(queue int64, maxCount int) ([]*api.Message, error) {
	var result []*api.Message = make([]*api.Message, 16)
	idx := 0
	m := this.m
	this.lock.Lock()
	for k, v := range m {
		if idx >= 16 {
			break
		}
		result[idx] = &api.Message{
			MessageId: k,
			Body:      v,
		}
		delete(m, k)
		idx++
	}
	this.lock.Unlock()

	return result[:idx], nil
}
