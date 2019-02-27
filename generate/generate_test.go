package generate

import "testing"

func TestString(t *testing.T) {
	type args struct {
		name   string
		list   []string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"basic",
			args{
				"MyEnum",
				[]string{
					"Red",
					"blue",
					"gReen",
				},
				"",
			},
			"enum MyEnum {\n\tRED = 0;\n\tBLUE = 1;\n\tGREEN = 2;\n}",
		},
		{
			"with suffix",
			args{
				"MyEnum",
				[]string{
					"Red",
					"blue",
					"gReen",
				},
				"color",
			},
			"enum MyEnum {\n\tRED_COLOR = 0;\n\tBLUE_COLOR = 1;\n\tGREEN_COLOR = 2;\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.name, tt.args.list, tt.args.suffix); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
