package query

import (
	"fmt"
	"strings"
)

type sqlBuilder struct {
	builder strings.Builder
}

func (sb *sqlBuilder) whereSQL(clause M) {
	sb.builder.WriteString(whereString)
	sb.builder.WriteString(sb.clauseSQL(clause))
}

// return string of format `param1=$1 AND param2=$2` format
func (sb *sqlBuilder) clauseSQL(clause M) string {
	clauses := []string{}
	i := 1
	for k := range clause {
		clauses = append(clauses, fmt.Sprintf("%s=$%d", k, i))
		i++
	}
	return strings.Join(clauses, andString)
}

func (sb *sqlBuilder) rowsSQL(values M) {
	i := 1
	sb.builder.WriteRune('(')

	keys := []string{}
	vals := []string{}
	for k := range values {
		keys = append(keys, k)
		vals = append(vals, fmt.Sprintf("$%d", i))
		i++
	}
	sb.WriteString(strings.Join(keys, commaSeparatorString))
	sb.builder.WriteRune(')')

	sb.WriteSpace()

	sb.WriteString(valuesString)
	sb.WriteSpace()

	sb.builder.WriteRune('(')
	sb.builder.WriteString(strings.Join(vals, commaSeparatorString))
	sb.builder.WriteRune(')')
}

func (sb *sqlBuilder) selectSQL(columns []string) {
	sb.builder.WriteString(selectString)
	sb.builder.WriteString(strings.Join(columns, commaSeparatorString))
}

func (sb *sqlBuilder) orderBySQL(orderBy string, orderType string) {
	sb.builder.WriteRune(spaceRune)
	sb.WriteString(orderByString)
	sb.WriteString(orderBy)
	sb.builder.WriteRune(spaceRune)
	sb.WriteString(string(orderType))
}

func (sb *sqlBuilder) WriteString(val string) {
	sb.builder.WriteString(val)
}

func (sb *sqlBuilder) WriteSpace() {
	sb.builder.WriteRune(' ')
}

func (sb *sqlBuilder) ToString() string {
	return sb.builder.String()
}
