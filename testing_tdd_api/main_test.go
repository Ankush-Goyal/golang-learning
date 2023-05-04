package testing_tdd_api

import (
	"reflect"
	"testing"
)

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
	}{
		{"success", args{"error", &LocalTodoData{}}, "", true},
		{"success 1", args{"success\\1", &LocalTodoData{}}, removeSpaces(`{"userId": 1,"id": 1,"title": "delectus aut autem","completed": false}`), false},
		{"success 2", args{"success\\2", &LocalTodoData{}}, removeSpaces(`{"userId": 1,"id": 2,"title": "quis ut nam facilis et officia qui","completed": false}`), false},
		{"no condition", args{"no condition", &LocalTodoData{}}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Level3Data(tt.args.url, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("Level3Data() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(removeSpaces(string(got)), tt.want) {
				t.Errorf("Level3Data() = %v, want %v", got, tt.want)
			}
		})
	}
}
