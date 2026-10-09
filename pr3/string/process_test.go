package main

import "testing"

var payload = []byte("   hello\nworld\nthis is a payload\n   ")

func BenchmarkProcessPayload(b *testing.B) {
    b.Run("Original", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _ = processPayload(payload)
        }
    })

    b.Run("Optimized", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _ = processPayload2(payload)
        }
    })
}

