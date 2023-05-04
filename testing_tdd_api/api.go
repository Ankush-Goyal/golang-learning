package testing_tdd_api

type TodoData interface {
	GetData(url string) ([]byte, error)
}
