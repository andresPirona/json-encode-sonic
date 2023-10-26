package faker

import (
	"github.com/andresPirona/json-encode-sonic/domain/repository"
	"github.com/andresPirona/json-encode-sonic/services/faker/json_encode"
	"github.com/andresPirona/json-encode-sonic/services/faker/sonic_encode"
)

const (
	JsonEncode = "JSON"
	Sonic      = "SONIC"
)

func NewFakerJsonImplementation(provider string) repository.JsonRepository {

	switch provider {
	case JsonEncode:
		return json_encode.NewJsonEncodeImplementation()
	case Sonic:
		return sonic_encode.NewSonicEncodeImplementation()
	default:
		return json_encode.NewJsonEncodeImplementation()
	}

}
