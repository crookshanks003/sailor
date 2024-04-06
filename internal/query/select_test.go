package query

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	type args struct {
		tableName string
	}
	tests := []struct {
		name string
		args args
		want *selectQuery
	}{
		{
			name: "primary",
			args: args{tableName: "users"},
			want: &selectQuery{
				from:      "users",
				col:       []string{"*"},
				clauses:   nil,
				orderBy:   "",
				orderType: "",
				sb:        &sqlBuilder{builder: strings.Builder{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, From(tt.args.tableName))
		})
	}
}

func Test_selectQuery_ToSQL(t *testing.T) {
	type fields struct {
		from      string
		col       []string
		clauses   map[string]interface{}
		orderBy   string
		orderType OrderType
		sb        *sqlBuilder
	}
	tests := []struct {
		name      string
		fields    fields
		want      string
		want1     []interface{}
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "all",
			fields: fields{
				from:      "users",
				col:       []string{"name", "age", "address"},
				clauses:   map[string]interface{}{"age": 22, "name": "John"},
				orderBy:   "name",
				orderType: "ASC",
				sb:        &sqlBuilder{builder: strings.Builder{}},
			},
			want:      "SELECT name, age, address FROM users WHERE age=$1 AND name=$2 ORDER BY name ASC",
			want1:     []interface{}{22, "John"},
			assertion: assert.NoError,
		},
		{
			name: "no_fields",
			fields: fields{
				from:    "users",
				col:     []string{"*"},
				clauses: map[string]interface{}{},
				sb:      &sqlBuilder{builder: strings.Builder{}},
			},
			want:      "SELECT * FROM users",
			want1:     nil,
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &selectQuery{
				from:      tt.fields.from,
				col:       tt.fields.col,
				clauses:   tt.fields.clauses,
				orderBy:   tt.fields.orderBy,
				orderType: tt.fields.orderType,
				sb:        tt.fields.sb,
			}
			got, got1, err := q.ToSQL()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_selectQuery_Where(t *testing.T) {
	type fields struct {
		from      string
		col       []string
		clauses   map[string]interface{}
		orderBy   string
		orderType OrderType
		sb        *sqlBuilder
	}
	type args struct {
		clause map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *selectQuery
	}{
		{
			name: "primary",
			fields: fields{
				from:      "users",
				col:       []string{"*"},
				clauses:   map[string]interface{}{},
				orderBy:   "",
				orderType: "",
				sb:        &sqlBuilder{builder: strings.Builder{}},
			},
			args: args{
				clause: map[string]interface{}{"age": 22, "name": "John"},
			},
			want: &selectQuery{
				from:      "users",
				col:       []string{"*"},
				clauses:   map[string]interface{}{"age": 22, "name": "John"},
				orderBy:   "",
				orderType: "",
				sb:        &sqlBuilder{builder: strings.Builder{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &selectQuery{
				from:      tt.fields.from,
				col:       tt.fields.col,
				clauses:   tt.fields.clauses,
				orderBy:   tt.fields.orderBy,
				orderType: tt.fields.orderType,
				sb:        &sqlBuilder{builder: strings.Builder{}},
			}
			assert.Equal(t, tt.want, q.Where(tt.args.clause))
		})
	}
}
