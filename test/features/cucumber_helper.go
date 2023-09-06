package features

import (
	"fmt"
	"github.com/cucumber/godog"
	"reflect"
)

func mapFields(target interface{}, table *godog.Table) error {
	targetValue := reflect.ValueOf(target).Elem()

	// Create a map of struct field names to their corresponding column names
	fieldMap := make(map[int]string)

	// Iterate through the table header to create the mapping
	for i, headerCell := range table.Rows[0].Cells {
		// Map the table column name to the struct field name
		fieldName := headerCell.Value
		if fieldName != "" {
			fieldMap[i] = fieldName
		}
	}

	// Iterate through the table rows and populate the struct fields
	for _, row := range table.Rows[1:] {
		for i, cell := range row.Cells {
			// Get the corresponding struct field name based on the table column name
			structFieldName := fieldMap[i]
			// Look for the struct field and set its value
			fieldValuePtr := targetValue.FieldByName(structFieldName).Addr().Interface()
			if err := setValue(cell.Value, fieldValuePtr); err != nil {
				return err
			}
		}
	}
	return nil
}

func setValue(value string, target interface{}) error {
	switch v := target.(type) {
	case *string:
		*v = value
	default:
		return fmt.Errorf("unsupported field type")
	}
	return nil
}
