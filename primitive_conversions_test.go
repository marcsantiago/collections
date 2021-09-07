package collections

import "testing"

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
			args: args{other: FloatValue64(10)},
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
