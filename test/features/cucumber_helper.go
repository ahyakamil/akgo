package features

import (
	"github.com/cucumber/godog"
	"reflect"
)

func mapFields(target interface{}, table *godog.Table) error {
	targetValue := reflect.ValueOf(target).Elem()
	targetType := targetValue.Type()
	fieldMap := make(map[string]int)

	// Create a map of struct field names to their corresponding column index
	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		fieldMap[field.Name] = i
	}

	// Iterate through the table rows and populate the struct fields
	for _, row := range table.Rows {
		fieldName := row.Cells[0].Value
		if fieldIndex, ok := fieldMap[fieldName]; ok {
			fieldValue := targetValue.Field(fieldIndex)
			if fieldValue.CanSet() {
				fieldValue.SetString(row.Cells[1].Value)
			}
		}
	}
	return nil
}
