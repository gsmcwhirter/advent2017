package lib

import (
	"reflect"
	"testing"
)

func TestCheck(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Check(tt.args.e)
		})
	}
}

func TestReadFileData(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadFileData(tt.args.filename); got != tt.want {
				t.Errorf("ReadFileData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntMod(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntMod(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IntMod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntMax(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IntMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntMin(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IntMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt32Max(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt32Max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("UInt32Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt32Min(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt32Min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("UInt32Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntAbs(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntAbs(tt.args.val); got != tt.want {
				t.Errorf("IntAbs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntClamp(t *testing.T) {
	type args struct {
		val int
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntClamp(tt.args.val, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("IntClamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRuneCounts(t *testing.T) {
	type args struct {
		counts map[rune]int
	}
	tests := []struct {
		name string
		args args
		want []RuneCount
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRuneCounts(tt.args.counts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRuneCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneCountSorter_Len(t *testing.T) {
	tests := []struct {
		name string
		rcs  RuneCountSorter
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rcs.Len(); got != tt.want {
				t.Errorf("RuneCountSorter.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneCountSorter_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		rcs  RuneCountSorter
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rcs.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("RuneCountSorter.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneCountSorter_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		rcs  RuneCountSorter
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rcs.Swap(tt.args.i, tt.args.j)
		})
	}
}
