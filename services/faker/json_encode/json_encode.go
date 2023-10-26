package json_encode

import (
	"encoding/json"
	"github.com/andresPirona/json-encode-sonic/domain/entity"
	"github.com/andresPirona/json-encode-sonic/domain/repository"
)

type implementationJson struct{}

func (i implementationJson) Marshal(fakers []entity.Faker) ([]byte, error) {
	return json.Marshal(fakers)
}

func (i implementationJson) UnMarshal(object []byte) ([]*entity.Faker, error) {
	var result []*entity.Faker
	errU := json.Unmarshal(object, &result)
	if errU != nil {
		return nil, errU
	}
	return result, nil
}

func NewJsonEncodeImplementation() repository.JsonRepository {
	return &implementationJson{}
}
