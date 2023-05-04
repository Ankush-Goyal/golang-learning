package testing_advanced_tdd

type TodoData interface {
	GetData(url string) ([]byte, error)
}
