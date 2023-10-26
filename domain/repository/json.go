package repository

import "github.com/andresPirona/json-encode-sonic/domain/entity"

type JsonRepository interface {
	Marshal([]entity.Faker) ([]byte, error)
	UnMarshal([]byte) ([]*entity.Faker, error)
}
