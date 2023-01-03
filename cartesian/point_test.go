package cartesian

import (
	"testing"
)

func TestPoint_Lat(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"should return 0",
			fields{0, 1},
			1,
		},
		{
			"should return -1",
			fields{-1, 1},
			1,
		},
		{
			"should return 1",
			fields{1, 1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.Lat(); got != tt.want {
				t.Errorf("Point.Lat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Lon(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"should return 1",
			fields{0, 1},
			1,
		},
		{
			"should return 0",
			fields{-1, 0},
			0,
		},
		{
			"should return -1",
			fields{1, -1},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.Lon(); got != tt.want {
				t.Errorf("Point.Lon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_IsAdjecent(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"should return true",
			fields{0, 0},
			args{1, 1},
			true,
		},
		{
			"should return false",
			fields{0, 0},
			args{1, 2},
			false,
		},
		{
			"should return true",
			fields{0, 0},
			args{1, -1},
			true,
		},
		{
			"should return false",
			fields{0, 0},
			args{-2, -5},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Point{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.IsAdjecent(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Point.IsAdjecent() = %v, want %v", got, tt.want)
			}
		})
	}
}
