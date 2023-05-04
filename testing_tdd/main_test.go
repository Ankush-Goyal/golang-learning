package testing_tdd

import (
	"testing"
)

func TestNumberDescrptionOne(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty test case #1", args{0}, "even"},
		{"empty test case #1", args{1}, "odd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberDescrptionOne(tt.args.n); got != tt.want {
				t.Errorf("NumberDescrptionOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberDescrptionTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"less than zero", args{-1}, "negative"},
		{"start boundary single digit", args{0}, "single digit"},
		{"mid single digit", args{5}, "single digit"},
		{"end boundary single digit", args{9}, "single digit"},
		{"double digit", args{10}, "double digit"}, // will return error since wrong
		{"three digits#1", args{100}, "three digit"},
		{"three digits #2", args{500}, "three digit"},
		{"four digit digits and above#2", args{1500}, "others"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberDescrptionTwo(tt.args.n); got != tt.want {
				t.Errorf("NumberDescrptionTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
