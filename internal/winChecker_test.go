package internal

import (
	"reflect"
	"testing"
)

func Test_removeSmallestScore(t *testing.T) {
	type args struct {
		allCard []PlayingСard
	}
	tests := []struct {
		name string
		args args
		want []PlayingСard
	}{
		{"one card", args{[]PlayingСard{
			{4, "k", 2},
			{5, "k", 2},
			{6, "k", 2},
			{7, "k", 2},
			{8, "k", 2},
			{4, "k", 3},
			{4, "k", 4},
		}},
			[]PlayingСard{
				{4, "k", 2},
				{5, "k", 2},
				{6, "k", 2},
				{7, "k", 2},
				{8, "k", 2},
			}},

		{"one more card", args{[]PlayingСard{
			{4, "k", 2},
			{5, "k", 2},
			{6, "k", 2},
			{7, "k", 2},
			{8, "k", 2},
			{4, "k", 4},
		}},
			[]PlayingСard{
				{4, "k", 2},
				{5, "k", 2},
				{6, "k", 2},
				{7, "k", 2},
				{8, "k", 2},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSmallestScore(tt.args.allCard); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeSmallestScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_royalFlushCheak(t *testing.T) {
	type args struct {
		allCard []PlayingСard
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Win",
			args{[]PlayingСard{
				{10, "10", 4},
				{14, "A", 2},
				{11, "J", 1},
				{11, "J", 2},
				{13, "K", 2},
				{10, "10", 2},
				{12, "Q", 2},
			}}, true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := royalFlushCheak(tt.args.allCard); got != tt.want {
				t.Errorf("royalFlushCheak() = %v, want %v", got, tt.want)
			}
		})
	}
}
