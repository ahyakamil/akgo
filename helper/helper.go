package helper

import (
	"akgo/constant/error_message"
	"errors"
	"github.com/jackc/pgconn"
	uuid2 "github.com/nu7hatch/gouuid"
	"net/http"
	"reflect"
	"strconv"
)

func MapQueryParamsToStruct(r *http.Request, queryParams interface{}) error {
	// Get the query parameters from the request
	queryValues := r.URL.Query()

	// Use reflection to map query parameters to struct fields
	structValue := reflect.ValueOf(queryParams).Elem()
	structType := structValue.Type()

	for i := 0; i < structValue.NumField(); i++ {
		field := structType.Field(i)
		fieldName := field.Tag.Get("json")
		paramValue := queryValues.Get(fieldName)

		// Handle different data types as needed
		switch field.Type.Kind() {
		case reflect.String:
			structValue.Field(i).SetString(paramValue)
		case reflect.Int:
			intValue, err := strconv.Atoi(paramValue)
			if err != nil {
				return err
			}
			structValue.Field(i).SetInt(int64(intValue))
		}
	}

	return nil
}

func GetUUID() string {
	uuid, _ := uuid2.NewV4()
	return uuid.String()
}

func ValidateUpdate(result pgconn.CommandTag) error {
	var err error
	if result.String() == "UPDATE 0" {
		err = errors.New(error_message.ERROR_DATA_NOT_FOUND)
	}
	return err
}

func ValidateDelete(result pgconn.CommandTag) error {
	var err error
	if result.String() == "DELETE 0" {
		err = errors.New(error_message.ERROR_DATA_NOT_FOUND)
	}
	return err
}
