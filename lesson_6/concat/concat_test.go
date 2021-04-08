package concat

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestConcat(t *testing.T)  {
//	one := "Hello "
//	two := "World"
//	actual := Concat(one, two)
//	expected := "Hello World"
//	if actual != expected {
//		t.Errorf("%s and %s not equal, but expected", actual, expected)
//	}
//}

//func TestConcat(t *testing.T) {
//	type args struct {
//		one string
//		two string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		{
//			name: "Expecting get equal results",
//			args: args{
//				one: "Hello ",
//				two: "World",
//			},
//			want: "Hello World",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Concat(tt.args.one, tt.args.two); got != tt.want {
//				t.Errorf("Concat() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestConcat(t *testing.T) {
	assert := assert.New(t)
	one := "Hello "
	two := "World"
	actual := Concat(one, two)
	expected := "Hello World"
	assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected", actual, expected))
}
