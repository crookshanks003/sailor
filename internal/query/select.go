package query

import (
	"errors"
	"strings"
)

type selectQuery struct {
	from      string
	col       []string
	clauses   map[string]interface{}
	orderBy   string
	orderType OrderType
	sb        *sqlBuilder
}

func (q *selectQuery) Select(columns ...string) *selectQuery {
	return &selectQuery{
		from:      q.from,
		col:       columns,
		clauses:   q.clauses,
		orderBy:   q.orderBy,
		orderType: q.orderType,
		sb:        q.sb,
	}
}

func (q *selectQuery) Where(clause M) *selectQuery {
	return &selectQuery{
		from:      q.from,
		orderBy:   q.orderBy,
		orderType: q.orderType,
		col:       q.col,
		clauses:   clause,
		sb:        q.sb,
	}
}

func (q *selectQuery) OrderBy(column string, orderType OrderType) *selectQuery {
	return &selectQuery{
		from:      q.from,
		col:       q.col,
		clauses:   q.clauses,
		orderBy:   column,
		orderType: orderType,
		sb:        q.sb,
	}
}

func (q *selectQuery) ToSQL() (string, []interface{}, error) {
	if q.from == "" {
		return "", nil, errors.New("empty tablename")
	}

	q.sb.selectSQL(q.col)

	q.sb.builder.WriteRune(spaceRune)
	q.sb.WriteString(fromString)
	q.sb.WriteString(q.from)

	var args []interface{} = nil
	if len(q.clauses) != 0 {
		q.sb.whereSQL(q.clauses)

		args = make([]interface{}, 0, len(q.clauses))
		for _, v := range q.clauses {
			args = append(args, v)
		}
	}

	if q.orderBy != "" {
		q.sb.orderBySQL(q.orderBy, string(q.orderType))
	}

	return q.sb.ToString(), args, nil
}

func From(tableName string) *selectQuery {
	return &selectQuery{
		from: tableName,
		col:  []string{"*"},
		sb:   &sqlBuilder{builder: strings.Builder{}},
	}
}
