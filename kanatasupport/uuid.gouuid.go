package kanatasupport

import "github.com/satori/go.uuid"

type uuidGenerator struct {
}

func newUUIDGenerator() uuidGenerator {
	return uuidGenerator{}
}

func (this uuidGenerator) Generate() (string, error) {
	u, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	return u.String(), err
}
