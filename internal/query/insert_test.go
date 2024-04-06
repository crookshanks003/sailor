package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insertQuery_ToSQL(t *testing.T) {
	type fields struct {
		tableName string
		rows      M
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
			name: "default",
			fields: fields{
				tableName: "users",
				rows:      map[string]interface{}{"name": "Name", "age": 22},
				sb:        &sqlBuilder{},
			},
			want:      "INSERT INTO users (name, age) VALUES ($1, $2)",
			want1:     []interface{}{"Name", 22},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := insertQuery{
				tableName: tt.fields.tableName,
				rows:      tt.fields.rows,
				sb:        tt.fields.sb,
			}
			got, got1, err := q.ToSQL()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
