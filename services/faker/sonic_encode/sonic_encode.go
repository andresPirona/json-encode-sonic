package sonic_encode

import (
	"github.com/andresPirona/json-encode-sonic/domain/entity"
	"github.com/andresPirona/json-encode-sonic/domain/repository"
	"github.com/bytedance/sonic"
)

type implementationJson struct{}

func (i implementationJson) Marshal(fakers []entity.Faker) ([]byte, error) {
	return sonic.Marshal(fakers)
}

func (i implementationJson) UnMarshal(object []byte) ([]*entity.Faker, error) {
	var result []*entity.Faker
	errU := sonic.Unmarshal(object, &result)
	if errU != nil {
		return nil, errU
	}
	return result, nil
}

func NewSonicEncodeImplementation() repository.JsonRepository {
	return &implementationJson{}
}
