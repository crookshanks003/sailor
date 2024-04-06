package query

import (
	"errors"
	"strings"
)

type insertQuery struct {
	tableName string
	rows      M
	sb        *sqlBuilder
}

func (q insertQuery) Rows(rows interface{}) insertQuery {
	row := parseStructWithDBTag(rows)
	
	return insertQuery{
		rows:      row,
		tableName: q.tableName,
		sb:        q.sb,
	}
}

func (q insertQuery) ToSQL() (string, []interface{}, error) {
	if q.tableName == "" {
		return "", []interface{}{}, errors.New("empty tablename")
	}

	q.sb.WriteString(insertString)
	q.sb.WriteString(q.tableName)
	q.sb.WriteSpace()

	q.sb.rowsSQL(q.rows)

	args := []interface{}{}
	for _, v := range q.rows {
		args = append(args, v)
	}

	return q.sb.ToString(), args, nil
}

func InsertInto(tableName string) insertQuery {
	return insertQuery{tableName: tableName, sb: &sqlBuilder{builder: strings.Builder{}}}
}
