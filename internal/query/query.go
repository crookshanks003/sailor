package query

import (
	"reflect"
)

type OrderType string

// shortcut for map[string]interface{}
type M map[string]interface{}

const dbTag = "db"

const (
	ASC  OrderType = "ASC"
	DESC OrderType = "DESC"
)

const (
	whereString          = " WHERE "
	selectString         = "SELECT "
	fromString           = "FROM "
	insertString         = "INSERT INTO "
	andString            = " AND "
	orderByString        = "ORDER BY "
	valuesString         = "VALUES"
	commaSeparatorString = ", "
	spaceRune            = ' '
)

func parseStructWithDBTag(s interface{}) M {
	result := M{}

	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		return result
	}
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		dbTag := field.Tag.Get(dbTag)

		if dbTag != "" && value != reflect.Zero(field.Type).Interface() {
			result[dbTag] = value
		}
	}

	return result
}
