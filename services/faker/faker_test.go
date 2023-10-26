package faker

import (
	"fmt"
	"github.com/andresPirona/json-encode-sonic/domain/entity"
	"github.com/bxcodec/faker"
	"testing"
)

func BuildFakerSlice(size int) []entity.Faker {
	var rows []entity.Faker

	for i := 0; i < size; i++ {
		var row entity.Faker
		_ = faker.FakeData(&row)
		rows = append(rows, row)
	}

	return rows
}

func BenchmarkJSONEncode(b *testing.B) {
	// tamaño de los slices
	var input = []int{5, 30, 100, 800}

	// init repo /// provider JSON ENCODE
	repoFaker := NewFakerJsonImplementation(JsonEncode)

	b.ResetTimer()
	b.ReportAllocs()

	for _, v := range input {
		fakers := BuildFakerSlice(v)
		b.Run(fmt.Sprintf("tamaño_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fakerBytes, _ := repoFaker.Marshal(fakers)
				_, _ = repoFaker.UnMarshal(fakerBytes)
			}
		})
	}
}

func BenchmarkSonicEncode(b *testing.B) {
	// tamaño de los slices
	var input = []int{5, 30, 100, 800}

	// init repo /// provider JSON ENCODE
	repoFaker := NewFakerJsonImplementation(Sonic)

	for _, v := range input {
		fakers := BuildFakerSlice(v)
		b.Run(fmt.Sprintf("tamaño_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fakerBytes, _ := repoFaker.Marshal(fakers)
				_, _ = repoFaker.UnMarshal(fakerBytes)
			}
		})
	}
}
