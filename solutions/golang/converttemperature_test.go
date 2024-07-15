package golang

import (
	"reflect"
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	t.Parallel()

	type args struct {
		celsius float64
	}

	tests := []struct {
		name string
		args args
		want []float64
	}{
		{name: "smoke 1", args: args{celsius: 36.50}, want: []float64{309.65, 97.7}},
		{name: "smoke 2", args: args{celsius: 122.11}, want: []float64{395.26, 251.798}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := convertTemperature(test.args.celsius); !reflect.DeepEqual(got, test.want) {
				t.Errorf("convertTemperature() = %v, want = %v", got, test.want)
			}
		})
	}
}
