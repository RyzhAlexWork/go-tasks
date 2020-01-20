package workerPool

import "testing"

func BenchmarkWorkerPool_1(b *testing.B) {
	workerPool(3)
}

func BenchmarkWorkerPool_2(b *testing.B) {
	workerPool(5)
}

func BenchmarkWorkerPool_3(b *testing.B) {
	workerPool(10)
}

func BenchmarkWorkerPool_4(b *testing.B) {
	workerPool(20)
}

func BenchmarkWorkerPool_5(b *testing.B) {
	workerPool(100)
}
