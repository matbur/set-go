package set

import (
	"reflect"
	"testing"
	"fmt"
)

func TestNew(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want *Set
	}{
		{
			"no values",
			args{},
			New(),
		}, {
			"one value",
			args{[]int{3}},
			New(3),
		}, {
			"two unique values",
			args{[]int{1, -2}},
			New(1, -2),
		}, {
			"three unique values",
			args{[]int{0, 2, 3}},
			New(0, 3, 2),
		}, {
			"one value two times",
			args{[]int{-7, -7}},
			New(-7),
		}, {
			"two values in three elements",
			args{[]int{-1, 2, -1}},
			New(2, -1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Len(t *testing.T) {
	tests := []struct {
		name string
		s    *Set
		want int
	}{
		{
			"no values",
			New(),
			0,
		}, {
			"one value",
			New(3),
			1,
		}, {
			"two unique values",
			New(1, -2),
			2,
		}, {
			"three unique values",
			New(0, 3, 2),
			3,
		}, {
			"one value two times",
			New(-7, -7),
			1,
		}, {
			"two values in three elements",
			New(-1, 2, -1),
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Set.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Empty(t *testing.T) {
	tests := []struct {
		name string
		s    *Set
		want bool
	}{
		{
			"no values",
			New(),
			true,
		}, {
			"one value",
			New(3),
			false,
		}, {
			"two unique values",
			New(1, -2),
			false,
		}, {
			"three unique values",
			New(0, 3, 2),
			false,
		}, {
			"one value two times",
			New(-7, -7),
			false,
		}, {
			"two values in three elements",
			New(-1, 2, -1),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Empty(); got != tt.want {
				t.Errorf("Set.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	tests := []struct {
		name string
		s    *Set
	}{
		{
			"no values",
			New(),
		}, {
			"one value",
			New(3),
		}, {
			"two unique values",
			New(1, -2),
		}, {
			"three unique values",
			New(0, 3, 2),
		}, {
			"one value two times",
			New(-7, -7),
		}, {
			"two values in three elements",
			New(-1, 2, -1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := New()
			got := tt.s.Clear()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Set.Clear() = %v, want %v", got, want)
			}
			got = tt.s
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Set.Clear(); Set = %v, want %v", got, want)
			}
		})
	}
}

func TestSet_Copy(t *testing.T) {
	tests := []struct {
		name string
		s    *Set
		want *Set
	}{
		{
			"no values",
			New(),
			New(),
		}, {
			"one value",
			New(3),
			New(3),
		}, {
			"two unique values",
			New(1, -2),
			New(-2, 1),
		}, {
			"three unique values",
			New(0, 3, 2),
			New(2, 0, 3),
		}, {
			"one value two times",
			New(-7, -7),
			New(-7),
		}, {
			"two values in three elements",
			New(-1, 2, -1),
			New(-1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Copy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Copy() = %v, want %v", got, tt.want)
			}
			if fmt.Sprintf("%p", tt.s) == fmt.Sprintf("%p", got) {
				t.Errorf("s2 = s1.Copy(); s1 and s2 is the same object")
			}
		})
	}
}

func TestSet_Equal(t *testing.T) {
	type args struct {
		other *Set
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want bool
	}{
		{
			"no values with no values",
			New(),
			args{New()},
			true,
		}, {
			"one value with one value",
			New(3),
			args{New(3)},
			true,
		}, {
			"two unique values with the same values",
			New(1, -2),
			args{New(-2, 1)},
			true,
		}, {
			"three unique values with the same values",
			New(0, 3, 2),
			args{New(2, 3, 0)},
			true,
		}, {
			"one value two times with one value",
			New(-7),
			args{New(-7, -7)},
			true,
		}, {
			"two values in three elements",
			New(2, -1),
			args{New(-1, 2, -1)},
			true,
		}, {
			"no values with one value",
			New(2),
			args{New()},
			false,
		}, {
			"one value with two values",
			New(-1, 2),
			args{New(2)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Equal(tt.args.other); got != tt.want {
				t.Errorf("Set.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		s    *Set
		want [][]int
	}{
		{
			"no values",
			New(),
			[][]int{{}},
		}, {
			"one value",
			New(1),
			[][]int{{1}},
		}, {
			"two values",
			New(3, 5),
			[][]int{{3, 5}, {5, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.ToSlice()
			errFlag := true
			for _, want := range tt.want {
				if reflect.DeepEqual(got, want) {
					errFlag = false
					break
				}
			}
			if errFlag {
				t.Errorf("Set.ToSlice() = %v, want one of %v", got, tt.want)
			}
		})
	}
}

func TestSet_String(t *testing.T) {
	tests := []struct {
		name string
		s    *Set
		want string
	}{
		{
			"no values",
			New(),
			"{}",
		}, {
			"one value",
			New(3),
			"{3}",
		}, {
			"two unique values",
			New(1, -2),
			"{-2, 1}",
		}, {
			"three unique values",
			New(0, 3, 2),
			"{0, 2, 3}",
		}, {
			"one value two times",
			New(-7, -7),
			"{-7}",
		}, {
			"two values in three elements",
			New(-1, 2, -1),
			"{-1, 2}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Set.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want *Set
	}{
		{
			"no value to no value",
			New(),
			args{[]int{}},
			New(),
		}, {
			"one value to no value",
			New(),
			args{[]int{7}},
			New(7),
		}, {
			"one value to one value",
			New(7),
			args{[]int{-3}},
			New(7, -3),
		}, {
			"two values to no value",
			New(),
			args{[]int{3, 5}},
			New(3, 5),
		}, {
			"one value to the same value",
			New(4),
			args{[]int{4}},
			New(4),
		}, {
			"two values to two values",
			New(3, 4),
			args{[]int{4, 5}},
			New(3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Add(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsIn(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want bool
	}{
		{
			"value in no values",
			New(),
			args{4},
			false,
		}, {
			"value in one value",
			New(4),
			args{5},
			false,
		}, {
			"value in the same value",
			New(8),
			args{8},
			true,
		}, {
			"value in two values",
			New(3, 4),
			args{7},
			false,
		}, {
			"one of values in two values",
			New(9, 2),
			args{2},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsIn(tt.args.value); got != tt.want {
				t.Errorf("Set.IsIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name    string
		s       *Set
		args    args
		wantErr bool
	}{
		{
			"value from no values",
			New(),
			args{4},
			true,
		}, {
			"value from one value",
			New(4),
			args{5},
			true,
		}, {
			"value from the same value",
			New(8),
			args{8},
			false,
		}, {
			"value from two values",
			New(3, 4),
			args{7},
			true,
		}, {
			"one of values from two values",
			New(9, 2),
			args{2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Remove(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Set.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSet_Pop(t *testing.T) {
	tests := []struct {
		name    string
		s       *Set
		want    []int
		wantErr bool
	}{
		{
			"no values",
			New(),
			[]int{},
			true,
		}, {
			"one value",
			New(-4),
			[]int{-4},
			false,
		}, {
			"two values",
			New(-3, 3),
			[]int{-3, 3},
			false,
		}, {
			"three values",
			New(4, 2, 0),
			[]int{0, 2, 4},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Set.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			flag := false
			for _, want := range tt.want {
				if got == want {
					flag = true
					break
				}
			}
			if !flag && !tt.wantErr {
				t.Errorf("Set.Pop() = %v, want one of %v", got, tt.want)
			}
		})
	}
}

func TestSet_Difference(t *testing.T) {
	type args struct {
		other *Set
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want *Set
	}{
		{
			"no values with no values",
			New(),
			args{New()},
			New(),
		}, {
			"one value with one value",
			New(3),
			args{New(3)},
			New(),
		}, {
			"two unique values with the same values",
			New(1, -2),
			args{New(-2, 1)},
			New(),
		}, {
			"three unique values with the same values",
			New(0, 3, 2),
			args{New(2, 3, 0)},
			New(),
		}, {
			"one value with two values",
			New(-7),
			args{New(-7, -5)},
			New(),
		}, {
			"two values in three elements",
			New(2, -1),
			args{New(-1, 2, -1)},
			New(),
		}, {
			"no values with one value",
			New(2),
			args{New()},
			New(2),
		}, {
			"one value with two values",
			New(-1, 2),
			args{New(2)},
			New(-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Difference(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	type args struct {
		other *Set
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want *Set
	}{
		{
			"no values with no values",
			New(),
			args{New()},
			New(),
		}, {
			"one value with one value",
			New(3),
			args{New(3)},
			New(3),
		}, {
			"two unique values with two values",
			New(1, -2),
			args{New(2, 1)},
			New(1),
		}, {
			"three unique values with the same values",
			New(0, 3, 2),
			args{New(2, 3, 0)},
			New(0, 3, 2),
		}, {
			"one value with two values",
			New(-7),
			args{New(-7, -5)},
			New(-7),
		}, {
			"two values in three elements",
			New(2, -1),
			args{New(-1, 2, -1)},
			New(-1, 2),
		}, {
			"no values with one value",
			New(2),
			args{New()},
			New(),
		}, {
			"one value with two values",
			New(-1, 2),
			args{New(2)},
			New(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersection(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	type args struct {
		other *Set
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want *Set
	}{
		{
			"no values with no values",
			New(),
			args{New()},
			New(),
		}, {
			"one value with one value",
			New(3),
			args{New(3)},
			New(3),
		}, {
			"two unique values with two values",
			New(1, -2),
			args{New(2, 1)},
			New(1, -2, 2),
		}, {
			"three unique values with the same values",
			New(0, 3, 2),
			args{New(2, 3, 0)},
			New(0, 3, 2),
		}, {
			"one value with two values",
			New(-7),
			args{New(-7, -5)},
			New(-5, -7),
		}, {
			"two values in three elements",
			New(2, -1),
			args{New(-1, 2, -1)},
			New(-1, 2),
		}, {
			"no values with one value",
			New(2),
			args{New()},
			New(2),
		}, {
			"one value with two values",
			New(-1, 2),
			args{New(2)},
			New(-1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Union(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	type args struct {
		other *Set
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want *Set
	}{
		{
			"no values with no values",
			New(),
			args{New()},
			New(),
		}, {
			"one value with one value",
			New(3),
			args{New(3)},
			New(),
		}, {
			"two unique values with two values",
			New(1, -2),
			args{New(2, 1)},
			New(-2, 2),
		}, {
			"three unique values with the same values",
			New(0, 3, 2),
			args{New(2, 3, 0)},
			New(),
		}, {
			"one value with two values",
			New(-7),
			args{New(-7, -5)},
			New(-5),
		}, {
			"two values in three elements",
			New(2, -1),
			args{New(-1, 2, -1)},
			New(),
		}, {
			"no values with one value",
			New(2),
			args{New()},
			New(2),
		}, {
			"one value with two values",
			New(-1, 2),
			args{New(2)},
			New(-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SymmetricDifference(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.SymmetricDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsDisjoint(t *testing.T) {
	type args struct {
		other *Set
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want bool
	}{
		{
			"no values with no values",
			New(),
			args{New()},
			true,
		}, {
			"one value with one value",
			New(3),
			args{New(4)},
			true,
		}, {
			"one value with the same value",
			New(3),
			args{New(3)},
			false,
		}, {
			"two unique values with the same values",
			New(1, -2),
			args{New(-2, 1)},
			false,
		}, {
			"three unique values with the same values",
			New(0, 3, 2),
			args{New(2, 3, 0)},
			false,
		}, {
			"one value two times with one value",
			New(-8),
			args{New(-7, -6)},
			true,
		}, {
			"one value with no values",
			New(2),
			args{New()},
			true,
		}, {
			"no values with one value",
			New(),
			args{New(2)},
			true,
		}, {
			"one value with two values",
			New(-1, 2),
			args{New(2)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsDisjoint(tt.args.other); got != tt.want {
				t.Errorf("Set.IsDisjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsSubset(t *testing.T) {
	type args struct {
		other *Set
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want bool
	}{
		{
			"no values with no values",
			New(),
			args{New()},
			true,
		}, {
			"one value with one value",
			New(3),
			args{New(3)},
			true,
		}, {
			"two unique values with the same values",
			New(1, -2),
			args{New(-2, 1)},
			true,
		}, {
			"three unique values with the same values",
			New(0, 3, 2),
			args{New(2, 3, 0)},
			true,
		}, {
			"one value two times with one value",
			New(-7),
			args{New(-7, -7)},
			true,
		}, {
			"two values in three elements",
			New(2, -1),
			args{New(-1, 2, -1)},
			true,
		}, {
			"no values with one value",
			New(2),
			args{New()},
			false,
		}, {
			"one value with two values",
			New(-1, 2),
			args{New(2)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSubset(tt.args.other); got != tt.want {
				t.Errorf("Set.IsSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsSuperset(t *testing.T) {
	type args struct {
		other *Set
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want bool
	}{
		{
			"no values with no values",
			New(),
			args{New()},
			true,
		}, {
			"one value with one value",
			New(3),
			args{New(3)},
			true,
		}, {
			"two unique values with the same values",
			New(1, -2),
			args{New(-2, 1)},
			true,
		}, {
			"three unique values with the same values",
			New(0, 3, 2),
			args{New(2, 3, 0)},
			true,
		}, {
			"one value two times with one value",
			New(-7),
			args{New(-7, -7)},
			true,
		}, {
			"two values in three elements",
			New(2, -1),
			args{New(-1, 2, -1)},
			true,
		}, {
			"no values with one value",
			New(2),
			args{New()},
			true,
		}, {
			"one value with two values",
			New(-1, 2),
			args{New(2)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSuperset(tt.args.other); got != tt.want {
				t.Errorf("Set.IsSuperset() = %v, want %v", got, tt.want)
			}
		})
	}
}
