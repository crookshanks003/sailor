package query

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sqlBuilder_whereSQL(t *testing.T) {
	type fields struct {
		builder strings.Builder
	}
	type args struct {
		clause M
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "happy",
			fields: fields{
				builder: strings.Builder{},
			},
			args: args{clause: M{"name": "test_user", "age": 22}},
			want: "WHERE name=$1 AND age=$2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := &sqlBuilder{
				builder: tt.fields.builder,
			}
			sb.whereSQL(tt.args.clause)
			value := sb.builder.String()
			assert.Equal(t, tt.want, value)
		})
	}
}

func Test_sqlBuilder_selectSQL(t *testing.T) {
	type fields struct {
		builder strings.Builder
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want string
	}{
		{
			name:   "happy",
			fields: fields{
				builder: strings.Builder{},
			},
			args:   args{
				columns: []string{"name", "age"},
			},
			want:   "SELECT name, age",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := &sqlBuilder{
				builder: tt.fields.builder,
			}
			sb.selectSQL(tt.args.columns)
			value := sb.builder.String()
			assert.Equal(t, tt.want, value)
		})
	}
}
