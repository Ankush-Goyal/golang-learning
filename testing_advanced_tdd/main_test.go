package testing_advanced_tdd

import (
	"reflect"
	"testing"
)

func checkerTwo(got []byte, want string) bool {
	return reflect.DeepEqual(removeSpaces(string(got)), want)
}

func TestLevel3Data(t *testing.T) {
	type args struct {
		url string
		d   TodoData
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		checker func([]byte, string) bool
	}{
		{"success", args{"error", &LocalTodoData{}}, "", true, checkerTwo},
		{"success 1", args{"success\\1", &LocalTodoData{}}, removeSpaces(`{"userId": 1,"id": 1,"title": "delectus aut autem","completed": false}`), false, checkerTwo},
		{"success 2", args{"success\\2", &LocalTodoData{}}, removeSpaces(`{"userId": 1,"id": 2,"title": "quis ut nam facilis et officia qui","completed": false}`), false, checkerTwo},
		{"no condition", args{"no condition", &LocalTodoData{}}, "", true, checkerTwo},
		{"error case", args{"error", &LocalTodoData{}}, "", true, checkerTwo},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Level3Data(tt.args.url, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("Level3Data() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.checker(got, tt.want) {
				t.Errorf("Level3Data() = %v, want %v", got, tt.want)
			}
		})
	}
}
