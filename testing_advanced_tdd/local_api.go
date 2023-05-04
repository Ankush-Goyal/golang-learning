package testing_advanced_tdd

import "errors"

type LocalTodoData struct{}

func (ltd *LocalTodoData) GetData(url string) ([]byte, error) {
	if url == "error" {
		return nil, errors.New("api error")
	}
	if url == "success\\1" {
		return []byte(`{"userId": 1,"id": 1,"title": "delectus aut autem","completed": false}`), nil
	}
	if url == "success\\2" {
		return []byte(`{"userId": 1,"id": 2,"title": "quis ut nam facilis et officia qui","completed": false}`), nil
	}
	return nil, errors.New("empty data")
}
