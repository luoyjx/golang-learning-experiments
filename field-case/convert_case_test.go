package main

import (
	"encoding/json"
	"testing"
)

func BenchmarkToMapWithMultiCases(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data := FooStruct{1, 1, 1, 5}
		parsedData := ToMapWithMultiCases(data)
		_, _ = json.MarshalIndent(parsedData, "", "  ")
	}
}

func BenchmarkJSONMarshal(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data := FooStruct{1, 1, 1, 5}
		_, _ = json.MarshalIndent(data, "", "  ")
	}
}