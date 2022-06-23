package iterator

type Iterator[T any] interface {
	Next()
	Begin() int
	End() int
}
