package Poker

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
			}}, true},
		{"Win v2",
			args{[]PlayingСard{
				{13, "K", 0},
				{9, "9", 0},
				{10, "10", 3},
				{3, "3", 1},
				{6, "6", 0},
				{3, "3", 0},
				{13, "K", 2},
			}}, true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := royalFlushCheak(tt.args.allCard); got != tt.want {
				t.Errorf("royalFlushCheak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_straightFlushCheak(t *testing.T) {
	type args struct {
		allCard []PlayingСard
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int8
	}{
		{"Straight",
			args{[]PlayingСard{
				{2, "2", 4},
				{3, "3", 2},
				{4, "4", 2},
				{5, "5", 2},
				{6, "6", 2},
				{7, "7", 2},
				{9, "9", 2},
			}}, true, 7},
		{"Straight v2",
			args{[]PlayingСard{
				{14, "A", 2},
				{2, "2", 2},
				{3, "3", 2},
				{4, "4", 2},
				{5, "5", 2},
				{10, "10", 2},
				{4, "4", 3},
			}}, true, 5},
		{"Straight v3",
			args{[]PlayingСard{
				{2, "2", 4},
				{3, "3", 2},
				{4, "4", 2},
				{5, "5", 2},
				{6, "6", 2},
				{7, "7", 2},
				{9, "9", 2},
			}}, true, 7}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := straightFlushCheak(tt.args.allCard)
			if got != tt.want {
				t.Errorf("straightFlushCheak() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("straightFlushCheak() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_fourOfAKindCheak(t *testing.T) {
	type args struct {
		allCard []PlayingСard
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int8
	}{
		{"fourOfAKindCheak v1",
			args{[]PlayingСard{
				{5, "5", 2},
				{14, "A", 1},
				{5, "5", 1},
				{6, "6", 3},
				{5, "5", 0},
				{5, "5", 0},
				{2, "2", 3},
			}}, true, 11},
		{"fourOfAKindCheak v2",
			args{[]PlayingСard{
				{2, "2", 4},
				{3, "3", 2},
				{11, "J", 2},
				{11, "J", 2},
				{11, "J", 2},
				{11, "J", 2},
				{9, "9", 2},
			}}, true, 11}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := fourOfAKindCheak(tt.args.allCard)
			if got != tt.want {
				t.Errorf("fourOfAKindCheak() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("fourOfAKindCheak() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_fullHouseCheak(t *testing.T) {
	type args struct {
		allCard []PlayingСard
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int8
		want2 int8
	}{
		{"FullHouse v1",
			args{[]PlayingСard{
				{2, "2", 4},
				{2, "2", 4},
				{11, "J", 2},
				{11, "J", 2},
				{11, "J", 2},
				{10, "J", 2},
				{9, "9", 2},
			}}, true, 11, 2}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := fullHouseCheak(tt.args.allCard)
			if got != tt.want {
				t.Errorf("fullHouseCheak() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("fullHouseCheak() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("fullHouseCheak() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestFlushCheak(t *testing.T) {
	type args struct {
		allCard []PlayingСard
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int8
		want2 int8
	}{

		{"Flush v1",
			args{[]PlayingСard{
				{5, "2", 4},
				{7, "2", 2},
				{9, "J", 2},
				{11, "J", 2},
				{13, "J", 2},
				{14, "J", 2},
				{9, "9", 1},
			}}, true, 14, 0},
		{"Flush v2",
			args{[]PlayingСard{
				{9, "9", 3},
				{12, "Q", 0},
				{5, "5", 0},
				{13, "K", 2},
				{6, "6", 0},
				{14, "A", 3},
				{10, "10", 2},
			}}, false, 0, 0}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := flushCheak(tt.args.allCard)
			if got != tt.want {
				t.Errorf("flushCheak() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("flushCheak() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("flushCheak() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_straightCheak(t *testing.T) {
	type args struct {
		allCard []PlayingСard
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int8
	}{

		{"Straigh v1",
			args{[]PlayingСard{
				{2, "2", 2},
				{3, "3", 2},
				{4, "4", 2},
				{5, "5", 2},
				{9, "9", 3},
				{14, "A", 2},
				{9, "9", 1},
			}}, true, 5},
		{"Straigh v2",
			args{[]PlayingСard{
				{2, "2", 1},
				{10, "10", 2},
				{11, "J", 3},
				{12, "Q", 2},
				{13, "K", 3},
				{14, "A", 2},
				{9, "9", 1},
			}}, true, 14}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := straightCheak(tt.args.allCard)
			if got != tt.want {
				t.Errorf("straightCheak() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("straightCheak() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_threeOfAKindCheak(t *testing.T) {
	type args struct {
		allCard []PlayingСard
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int8
	}{
		{"threeOfAKind v1",
			args{[]PlayingСard{
				{5, "5", 1},
				{5, "5", 2},
				{5, "5", 3},
				{11, "J", 1},
				{11, "J", 3},
				{11, "J", 2},
				{9, "9", 1},
			}}, true, 11},
		{"threeOfAKind v2",
			args{[]PlayingСard{
				{11, "J", 1},
				{11, "J", 2},
				{11, "J", 3},
				{5, "5", 1},
				{5, "5", 3},
				{5, "5", 2},
				{9, "9", 1},
			}}, true, 11}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := threeOfAKindCheak(tt.args.allCard)
			if got != tt.want {
				t.Errorf("threeOfAKindCheak() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("threeOfAKindCheak() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_twoPairCheak(t *testing.T) {
	type args struct {
		allCard []PlayingСard
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int8
		want2 int8
	}{

		{"twoPair v1",
			args{[]PlayingСard{
				{5, "5", 1},
				{5, "5", 2},
				{2, "2", 1},
				{2, "2", 2},
				{10, "10", 1},
				{10, "10", 2},
				{9, "9", 3},
			}}, true, 10, 5}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := twoPairCheak(tt.args.allCard)
			if got != tt.want {
				t.Errorf("twoPairCheak() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("twoPairCheak() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("twoPairCheak() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
