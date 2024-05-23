package main

import (
	"reflect"
	"testing"
)

func TestDuck_AnotherDuck(t *testing.T) {
	type fields struct {
		flyAtHeight int
		runAtSpeed  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   Duck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Duck{
				flyAtHeight: tt.fields.flyAtHeight,
				runAtSpeed:  tt.fields.runAtSpeed,
			}
			if got := d.AnotherDuck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnotherDuck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDuck(t *testing.T) {
	type args struct {
		flyAtHeight int
		runAtSpeed  float64
	}
	tests := []struct {
		name string
		args args
		want *Duck
	}{
		// TODO: Add test cases.
		{
			name: "NewDuck",
			args: args{
				flyAtHeight: 2,
				runAtSpeed:  1.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDuck(tt.args.flyAtHeight, tt.args.runAtSpeed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDuck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeacock_FLY(t *testing.T) {
	type fields struct {
		flyAtHeight int
		runAtSpeed  float64
	}
	type args struct {
		height int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Peacock{
				flyAtHeight: tt.fields.flyAtHeight,
				runAtSpeed:  tt.fields.runAtSpeed,
			}
			p.FLY(tt.args.height)
		})
	}
}

func TestPeacock_RUN(t *testing.T) {
	type fields struct {
		flyAtHeight int
		runAtSpeed  float64
	}
	type args struct {
		speed float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Peacock{
				flyAtHeight: tt.fields.flyAtHeight,
				runAtSpeed:  tt.fields.runAtSpeed,
			}
			p.RUN(tt.args.speed)
		})
	}
}
