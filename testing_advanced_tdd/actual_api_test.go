package testing_advanced_tdd

import (
	"strings"
	"testing"
)

// func checkerResp(got Campaign, want rankCmapaignResp) bool {
// 	return got.Id == want.Id
// 	//return removeSpaces(string(got)) == want
// }

func checker(got []byte, want string) bool {
	return removeSpaces(string(got)) == want
}

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
		args    args
		want    string
		checker func([]byte, string) bool
		wantErr bool
	}{
		{
			"id 1", &ActualTodoData{}, args{`https://jsonplaceholder.typicode.com/todos/1`},
			removeSpaces(`{"userId": 1,"id": 1,"title": "delectus aut autem","completed": false}`), checker, false,
		},
		{
			"id 2", &ActualTodoData{}, args{`https://jsonplaceholder.typicode.com/todos/2`},
			removeSpaces(`{"userId": 1,"id": 2,"title": "quis ut nam facilis et officia qui","completed": false}`), checker, false,
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
			if !tt.checker(got, tt.want) {
				t.Errorf("ActualTodoData.GetData() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
