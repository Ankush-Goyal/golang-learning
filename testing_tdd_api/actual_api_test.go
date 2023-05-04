package testing_tdd_api

import (
	"reflect"
	"strings"
	"testing"
)

func removeSpaces(in string) string {
	val := strings.ReplaceAll(in, "\n", "")
	return strings.ReplaceAll(val, " ", "")
}

func TestActualTodoData_GetData(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		atd     *ActualTodoData
		args    args   //input
		want    string // expected output
		wantErr bool   // whether
	}{
		{
			"id 1", &ActualTodoData{}, args{`https://jsonplaceholder.typicode.com/todos/1`},
			removeSpaces(`{"userId": 1,"id": 1,"title": "delectus aut autem","completed": false}`), false,
		},
		{
			"id 2", &ActualTodoData{}, args{`https://jsonplaceholder.typicode.com/todos/2`},
			removeSpaces(`{"userId": 1,"id": 2,"title": "quis ut nam facilis et officia qui","completed": false}`), false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atd := &ActualTodoData{}
			got, err := atd.GetData(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ActualTodoData.GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(removeSpaces(string(got)), tt.want) {
				t.Errorf("ActualTodoData.GetData() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
