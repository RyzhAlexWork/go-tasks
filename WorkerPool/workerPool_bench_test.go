package WorkerPool

import "testing"

func BenchmarkWorkerPool_1(b *testing.B) {
	WorkerPool(3)
}

func BenchmarkWorkerPool_2(b *testing.B) {
	WorkerPool(10)
}

func BenchmarkWorkerPool_3(b *testing.B) {
	WorkerPool(25)
}

func BenchmarkWorkerPool_4(b *testing.B) {
	WorkerPool(100)
}
