package collections

import (
	"reflect"
	"testing"
)

func TestValue_Equal(t *testing.T) {
	type args struct {
		other Data
	}
	tests := []struct {
		name string
		dd   []Data
		args args
		want bool
	}{
		{
			name: "values should be equal int",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: IntValue(10)},
			want: true,
		},
		{
			name: "values should be equal int32",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: IntValue32(10)},
			want: true,
		},
		{
			name: "values should be equal int64",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: IntValue64(10)},
			want: true,
		},
		{
			name: "values should be equal float32",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: FloatValue32(10)},
			want: true,
		},
		{
			name: "values should be equal float64",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: FloatValue64(10)},
			want: true,
		},
		{
			name: "values should be equal string",
			dd: []Data{
				IntValue(9),
				IntValue32(9),
				IntValue64(9),
				FloatValue32(9),
				FloatValue64(9),
				StringValue("9"),
				RuneValue(57),
			},
			args: args{other: StringValue("9")},
			want: true,
		},
		{
			name: "values should be equal rune",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: RuneValue(10)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, i := range tt.dd {
				if got := i.Equal(tt.args.other); got != tt.want {
					t.Errorf("Equal() %T = %v, want %v", i, got, tt.want)
				}
			}
		})
	}
}

func TestValue_Greater(t *testing.T) {
	type args struct {
		other Data
	}
	tests := []struct {
		name string
		dd   []Data
		args args
		want bool
	}{
		{
			name: "values should be not be greater than they are equal",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: IntValue(10)},
			want: false,
		},
		{
			name: "values should be greater int",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: IntValue(9)},
			want: true,
		},
		{
			name: "values should be greater int32",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: IntValue32(9)},
			want: true,
		},
		{
			name: "values should be greater int64",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: IntValue64(9)},
			want: true,
		},
		{
			name: "values should be greater float32",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: FloatValue32(9)},
			want: true,
		},
		{
			name: "values should be greater float64",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: FloatValue64(9)},
			want: true,
		},
		{
			name: "values should be greater string",
			dd: []Data{
				IntValue(9),
				IntValue32(9),
				IntValue64(9),
				FloatValue32(9),
				FloatValue64(9),
				StringValue("9"),
				RuneValue(57),
			},
			args: args{other: StringValue("7")},
			want: true,
		},
		{
			name: "values should be greater rune",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: RuneValue(9)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, i := range tt.dd {
				if got := i.Greater(tt.args.other); got != tt.want {
					t.Errorf("Greater() %T = %v, want %v", i, got, tt.want)
				}
			}
		})
	}
}

func TestValue_Less(t *testing.T) {
	type args struct {
		other Data
	}
	tests := []struct {
		name string
		dd   []Data
		args args
		want bool
	}{
		{
			name: "values should be not be less than they are equal",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: IntValue(10)},
			want: false,
		},
		{
			name: "values should be less int",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: IntValue(11)},
			want: true,
		},
		{
			name: "values should be less int32",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: IntValue32(11)},
			want: true,
		},
		{
			name: "values should be less int64",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: IntValue64(11)},
			want: true,
		},
		{
			name: "values should be less float32",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: FloatValue32(11)},
			want: true,
		},
		{
			name: "values should be less float64",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: FloatValue64(11)},
			want: true,
		},
		{
			name: "values should be less string",
			dd: []Data{
				IntValue(8),
				IntValue32(8),
				IntValue64(8),
				FloatValue32(8),
				FloatValue64(8),
				StringValue("8"),
				RuneValue(56),
			},
			args: args{other: StringValue("9")},
			want: true,
		},
		{
			name: "values should be less rune",
			dd: []Data{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10),
				FloatValue64(10),
				StringValue("10"),
				RuneValue(10),
			},
			args: args{other: RuneValue(11)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, i := range tt.dd {
				if got := i.Less(tt.args.other); got != tt.want {
					t.Errorf("Less() %T = %v, want %v", i, got, tt.want)
				}
			}
		})
	}
}

func TestOperableData_Add(t *testing.T) {
	type args struct {
		data OperableData
	}
	tests := []struct {
		name string
		dd   []OperableData
		args args
		want []Data
	}{
		{
			name: "should add data",
			args: args{data: IntValue(10)},
			dd: []OperableData{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10.0),
				FloatValue64(10.0),
				RuneValue(10),
			},
			want: []Data{
				IntValue(20),
				IntValue32(20),
				IntValue64(20),
				FloatValue32(20.0),
				FloatValue64(20.0),
				RuneValue(20),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, s := range tt.dd {
				if got := s.Add(tt.args.data); !reflect.DeepEqual(got, tt.want[i]) {
					t.Errorf("Add() for %T got = %v, want %v", s, got, tt.want[i])
				}
			}
		})
	}
}

func TestOperableData_Div(t *testing.T) {
	type args struct {
		data OperableData
	}
	tests := []struct {
		name string
		dd   []OperableData
		args args
		want []Data
	}{
		{
			name: "should divide data",
			args: args{data: IntValue(2)},
			dd: []OperableData{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10.0),
				FloatValue64(10.0),
				RuneValue(10),
			},
			want: []Data{
				IntValue(5),
				IntValue32(5),
				IntValue64(5),
				FloatValue32(5.0),
				FloatValue64(5.0),
				RuneValue(5),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, s := range tt.dd {
				if got := s.Div(tt.args.data); !reflect.DeepEqual(got, tt.want[i]) {
					t.Errorf("Div() for %T got = %v, want %v", s, got, tt.want[i])
				}
			}
		})
	}
}

func TestOperableData_Mul(t *testing.T) {
	type args struct {
		data OperableData
	}
	tests := []struct {
		name string
		dd   []OperableData
		args args
		want []Data
	}{
		{
			name: "should multiply data",
			args: args{data: IntValue(2)},
			dd: []OperableData{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10.0),
				FloatValue64(10.0),
				RuneValue(10),
			},
			want: []Data{
				IntValue(20),
				IntValue32(20),
				IntValue64(20),
				FloatValue32(20.0),
				FloatValue64(20.0),
				RuneValue(20),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, s := range tt.dd {
				if got := s.Mul(tt.args.data); !reflect.DeepEqual(got, tt.want[i]) {
					t.Errorf("Mul() for %T got = %v, want %v", s, got, tt.want[i])
				}
			}
		})
	}
}

func TestOperableData_Sub(t *testing.T) {
	type args struct {
		data OperableData
	}
	tests := []struct {
		name string
		dd   []OperableData
		args args
		want []Data
	}{
		{
			name: "should subtract data",
			args: args{data: IntValue(2)},
			dd: []OperableData{
				IntValue(10),
				IntValue32(10),
				IntValue64(10),
				FloatValue32(10.0),
				FloatValue64(10.0),
				RuneValue(10),
			},
			want: []Data{
				IntValue(8),
				IntValue32(8),
				IntValue64(8),
				FloatValue32(8.0),
				FloatValue64(8.0),
				RuneValue(8),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, s := range tt.dd {
				if got := s.Sub(tt.args.data); !reflect.DeepEqual(got, tt.want[i]) {
					t.Errorf("Sub() for %T got = %v, want %v", s, got, tt.want[i])
				}
			}
		})
	}
}

func TestOperableData_Mod(t *testing.T) {
	type args struct {
		data OperableData
	}
	tests := []struct {
		name string
		dd   []OperableData
		args args
		want []Data
	}{
		{
			name: "should subtract data",
			args: args{data: IntValue(2)},
			dd: []OperableData{
				IntValue(11),
				IntValue32(11),
				IntValue64(11),
				FloatValue32(11.0),
				FloatValue64(11.0),
				RuneValue(11),
			},
			want: []Data{
				IntValue(1),
				IntValue32(1),
				IntValue64(1),
				FloatValue32(1.0),
				FloatValue64(1.0),
				RuneValue(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, s := range tt.dd {
				if got := s.Mod(tt.args.data); !reflect.DeepEqual(got, tt.want[i]) {
					t.Errorf("Mod() for %T got = %v, want %v", s, got, tt.want[i])
				}
			}
		})
	}
}
