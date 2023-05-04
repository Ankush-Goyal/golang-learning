package testing_tdd_api

import (
	"reflect"
	"testing"
)

func TestLocalTodoData_GetData(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		ltd     *LocalTodoData
		args    args
		want    string
		wantErr bool
	}{
		{
			"error", &LocalTodoData{}, args{"error"}, "", true,
		},
		{
			"id 1", &LocalTodoData{}, args{"success\\1"},
			removeSpaces(`{"userId": 1,"id": 1,"title": "delectus aut autem","completed": false}`), false,
		},
		{
			"id 2", &LocalTodoData{}, args{"success\\2"},
			removeSpaces(`{"userId": 1,"id": 2,"title": "quis ut nam facilis et officia qui","completed": false}`), false,
		},
		{
			"no condition", &LocalTodoData{}, args{"no condition"}, "", true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ltd := &LocalTodoData{}
			got, err := ltd.GetData(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocalTodoData.GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(removeSpaces(string(got)), tt.want) {
				t.Errorf("LocalTodoData.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
