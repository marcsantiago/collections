package set

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/marcsantiago/collections"
)

func TestSet_Add(t *testing.T) {
	type args struct {
		data []collections.Data
	}
	tests := []struct {
		name string
		args args
		want Set
	}{
		{
			name: "should add values without issue",
			args: args{data: collections.IntValues([]int{1, 2, 3, 4, 5}).Data()},
			want: Set{
				collections.IntValue(1): struct{}{},
				collections.IntValue(2): struct{}{},
				collections.IntValue(3): struct{}{},
				collections.IntValue(4): struct{}{},
				collections.IntValue(5): struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			for _, d := range tt.args.data {
				s.Add(d)
			}

			if !cmp.Equal(s, tt.want) {
				t.Fatalf("Add() got %v want %v", s, tt.want)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		want int
	}{
		{
			name: "full set should be empty",
			s: Set{
				collections.IntValue(1): struct{}{},
				collections.IntValue(2): struct{}{},
				collections.IntValue(3): struct{}{},
				collections.IntValue(4): struct{}{},
				collections.IntValue(5): struct{}{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
			if len(tt.s) != tt.want {
				t.Fatalf("Clear() got %v want %v", len(tt.s), tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args struct {
		data collections.Data
	}
	tests := []struct {
		name string
		s    Set
		args args
		want bool
	}{
		{
			name: "should contain the data",
			s: Set{
				collections.IntValue(1): struct{}{},
				collections.IntValue(2): struct{}{},
				collections.IntValue(3): struct{}{},
				collections.IntValue(4): struct{}{},
				collections.IntValue(5): struct{}{},
			},
			args: args{data: collections.IntValue(2)},
			want: true,
		},
		{
			name: "should not contain the data",
			s: Set{
				collections.IntValue(1): struct{}{},
				collections.IntValue(2): struct{}{},
				collections.IntValue(3): struct{}{},
				collections.IntValue(4): struct{}{},
				collections.IntValue(5): struct{}{},
			},
			args: args{data: collections.IntValue(10)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.data); got != tt.want {
				t.Fatalf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Copy(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		want Set
	}{
		{
			name: "should copy the set identically",
			s: Set{
				collections.IntValue(1): struct{}{},
				collections.IntValue(2): struct{}{},
				collections.IntValue(3): struct{}{},
				collections.IntValue(4): struct{}{},
				collections.IntValue(5): struct{}{},
			},
			want: Set{
				collections.IntValue(1): struct{}{},
				collections.IntValue(2): struct{}{},
				collections.IntValue(3): struct{}{},
				collections.IntValue(4): struct{}{},
				collections.IntValue(5): struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Copy(); !cmp.Equal(got, tt.want) {
				t.Fatalf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Difference(t *testing.T) {
	type args struct {
		sets []Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want Set
	}{
		{
			name: "should return the difference",
			args: args{sets: []Set{
				Set{
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
					collections.StringValue("apple"):     struct{}{},
				},
				Set{
					collections.StringValue("banana"): struct{}{},
				},
			}},
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			want: Set{
				collections.StringValue("cherry"): struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Difference(tt.args.sets...); !cmp.Equal(got, tt.want) {
				t.Fatalf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_DifferenceUpdate(t *testing.T) {
	type args struct {
		sets []Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want Set
	}{
		{
			name: "should return the difference",
			args: args{sets: []Set{
				Set{
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
					collections.StringValue("apple"):     struct{}{},
				},
				Set{
					collections.StringValue("banana"): struct{}{},
				},
			}},
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			want: Set{
				collections.StringValue("cherry"): struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.DifferenceUpdate(tt.args.sets...)
			if !cmp.Equal(tt.s, tt.want) {
				t.Fatalf("DifferenceUpdate() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Discard(t *testing.T) {
	type args struct {
		data collections.Data
	}
	tests := []struct {
		name string
		s    Set
		args args
		want Set
	}{
		{
			name: "should remove the data from the set",
			s: Set{
				collections.IntValue(1): struct{}{},
				collections.IntValue(2): struct{}{},
				collections.IntValue(3): struct{}{},
				collections.IntValue(4): struct{}{},
				collections.IntValue(5): struct{}{},
			},
			args: args{data: collections.IntValue(1)},
			want: Set{
				collections.IntValue(2): struct{}{},
				collections.IntValue(3): struct{}{},
				collections.IntValue(4): struct{}{},
				collections.IntValue(5): struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Discard(tt.args.data)
			if !cmp.Equal(tt.s, tt.want) {
				t.Fatalf("Discard() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	type args struct {
		sets []Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want Set
	}{
		{
			name: "should return the intersection",
			args: args{sets: []Set{
				Set{
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
					collections.StringValue("apple"):     struct{}{},
					collections.StringValue("banana"):    struct{}{},
				},
				Set{
					collections.StringValue("apple"):  struct{}{},
					collections.StringValue("banana"): struct{}{},
				},
			}},
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			want: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersection(tt.args.sets...); !cmp.Equal(got, tt.want) {
				t.Fatalf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IntersectionUpdate(t *testing.T) {
	type args struct {
		sets []Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want Set
	}{
		{
			name: "should return the intersection",
			args: args{sets: []Set{
				Set{
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
					collections.StringValue("apple"):     struct{}{},
					collections.StringValue("banana"):    struct{}{},
				},
				Set{
					collections.StringValue("apple"):  struct{}{},
					collections.StringValue("banana"): struct{}{},
				},
			}},
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			want: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.IntersectionUpdate(tt.args.sets...)
			if !cmp.Equal(tt.s, tt.want) {
				t.Fatalf("IntersectionUpdate() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_IsDisjoint(t *testing.T) {
	type args struct {
		ss Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want bool
	}{
		{
			name: "sets should be disjointed",
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			args: args{
				Set{
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
				},
			},
			want: true,
		},
		{
			name: "sets should not be disjointed",
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			args: args{
				Set{
					collections.StringValue("apple"):     struct{}{},
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsDisjoint(tt.args.ss); got != tt.want {
				t.Fatalf("IsDisjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsSubset(t *testing.T) {
	type args struct {
		ss Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want bool
	}{
		{
			name: "sets should be a subset",
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			args: args{
				Set{
					collections.StringValue("apple"):     struct{}{},
					collections.StringValue("banana"):    struct{}{},
					collections.StringValue("cherry"):    struct{}{},
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
				},
			},
			want: true,
		},
		{
			name: "sets should not be a subset",
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			args: args{
				Set{
					collections.StringValue("apple"):     struct{}{},
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSubset(tt.args.ss); got != tt.want {
				t.Fatalf("IsSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsSuperset(t *testing.T) {
	type args struct {
		ss Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want bool
	}{
		{
			name: "sets should be a super set",
			s: Set{
				collections.StringValue("apple"):     struct{}{},
				collections.StringValue("banana"):    struct{}{},
				collections.StringValue("cherry"):    struct{}{},
				collections.StringValue("google"):    struct{}{},
				collections.StringValue("microsoft"): struct{}{},
			},
			args: args{
				Set{
					collections.StringValue("apple"):  struct{}{},
					collections.StringValue("banana"): struct{}{},
					collections.StringValue("cherry"): struct{}{},
				},
			},
			want: true,
		},
		{
			name: "sets should not be a super set",
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			args: args{
				Set{
					collections.StringValue("apple"):     struct{}{},
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSuperset(tt.args.ss); got != tt.want {
				t.Fatalf("IsSuperset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		want collections.Data
	}{
		{
			name: "should randomly remove an item from the set",
			s: Set{
				collections.StringValue("apple"): struct{}{},
			},
			want: collections.StringValue("apple"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); !cmp.Equal(got, tt.want) {
				t.Fatalf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args struct {
		key collections.Data
	}
	tests := []struct {
		name string
		args args
		s    Set
		want Set
	}{
		{
			name: "should  remove an item from the set",
			args: args{key: collections.StringValue("cherry")},
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			want: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.key)
			if !cmp.Equal(tt.s, tt.want) {
				t.Fatalf("Remove() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	type args struct {
		ss Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want Set
	}{
		{
			name: "should return symmetric difference",
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			args: args{
				ss: Set{
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
					collections.StringValue("apple"):     struct{}{},
				},
			},
			want: Set{
				collections.StringValue("banana"):    struct{}{},
				collections.StringValue("cherry"):    struct{}{},
				collections.StringValue("microsoft"): struct{}{},
				collections.StringValue("google"):    struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SymmetricDifference(tt.args.ss); !cmp.Equal(got, tt.want) {
				t.Fatalf("SymmetricDifference() = got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_SymmetricDifferenceUpdate(t *testing.T) {
	type args struct {
		ss Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want Set
	}{
		{
			name: "should return symmetric difference",
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},
			args: args{
				ss: Set{
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
					collections.StringValue("apple"):     struct{}{},
				},
			},
			want: Set{
				collections.StringValue("banana"):    struct{}{},
				collections.StringValue("cherry"):    struct{}{},
				collections.StringValue("microsoft"): struct{}{},
				collections.StringValue("google"):    struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SymmetricDifferenceUpdate(tt.args.ss)
			if !cmp.Equal(tt.s, tt.want) {
				t.Fatalf("SymmetricDifferenceUpdate() = got %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	type args struct {
		sets []Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want Set
	}{
		{
			name: "should add all the sets together",
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},

			args: args{sets: []Set{
				Set{
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
					collections.StringValue("apple"):     struct{}{},
					collections.StringValue("banana"):    struct{}{},
				},
				Set{
					collections.StringValue("apple"):  struct{}{},
					collections.StringValue("banana"): struct{}{},
				},
			}},
			want: Set{
				collections.StringValue("apple"):     struct{}{},
				collections.StringValue("banana"):    struct{}{},
				collections.StringValue("cherry"):    struct{}{},
				collections.StringValue("google"):    struct{}{},
				collections.StringValue("microsoft"): struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Union(tt.args.sets...); !cmp.Equal(got, tt.want) {
				t.Fatalf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Update(t *testing.T) {
	type args struct {
		sets []Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want Set
	}{
		{
			name: "should add all the sets together",
			s: Set{
				collections.StringValue("apple"):  struct{}{},
				collections.StringValue("banana"): struct{}{},
				collections.StringValue("cherry"): struct{}{},
			},

			args: args{sets: []Set{
				Set{
					collections.StringValue("google"):    struct{}{},
					collections.StringValue("microsoft"): struct{}{},
					collections.StringValue("apple"):     struct{}{},
					collections.StringValue("banana"):    struct{}{},
				},
				Set{
					collections.StringValue("apple"):  struct{}{},
					collections.StringValue("banana"): struct{}{},
				},
			}},
			want: Set{
				collections.StringValue("apple"):     struct{}{},
				collections.StringValue("banana"):    struct{}{},
				collections.StringValue("cherry"):    struct{}{},
				collections.StringValue("google"):    struct{}{},
				collections.StringValue("microsoft"): struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Update(tt.args.sets...)
			if !cmp.Equal(tt.s, tt.want) {
				t.Fatalf("Update() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}
