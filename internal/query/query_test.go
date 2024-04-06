package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestUser struct {
	Name      string `db:"name"`
	Age       int32  `db:"age"`
	NoDBField string
}

func Test_parseStructWithDBTag(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want M
	}{
		{
			name: "all fields",
			args: args{
				s: TestUser{
					Name:      "Name",
					Age:       200,
					NoDBField: "Field 1",
				},
			},
			want: map[string]interface{}{"name": "Name", "age": int32(200)},
		},
		{
			name: "only name",
			args: args{
				s: TestUser{
					Name: "Name",
				},
			},
			want: map[string]interface{}{"name": "Name"},
		},
		{
			name: "no struct",
			args: args{
				s: []string{"Name"},
			},
			want: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := parseStructWithDBTag(tt.args.s)
			for k, v := range res {
				assert.Equal(t, tt.want[k], v)
			}
		})
	}
}
