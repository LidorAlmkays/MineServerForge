package configs

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

func loadConfigFromEnvOrFile(cfg interface{}, configFolderPath string) error {
	// Load .env file if it exists
	err := godotenv.Load(configFolderPath)
	if err != nil {
		fmt.Println("No .env file was found. Falling back to environment variables.")
	}

	v := reflect.ValueOf(cfg)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return errors.New("cfg must be a pointer to a struct")
	}

	t := v.Type()
	var missingFields []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Handle embedded structs recursively
		if value.Kind() == reflect.Struct {
			if value.CanAddr() {
				err := loadConfigFromEnvOrFile(value.Addr().Interface(), configFolderPath)
				if err != nil {
					return err
				}
			}
			continue
		}

		// Process the `env` tag
		envTag := field.Tag.Get("env")
		if envTag == "" {
			continue
		}

		// Check for the value in .env file or environment variables
		envValue := os.Getenv(envTag)
		if envValue == "" {
			missingFields = append(missingFields, envTag)
			continue
		}

		// Set the field value based on its type
		err := setFieldValue(value, envValue)
		if err != nil {
			return fmt.Errorf("failed to set value for %s: %w", envTag, err)
		}
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("missing required environment variables: %v", missingFields)
	}

	return nil
}

func setFieldValue(value reflect.Value, envValue string) error {
	if !value.CanSet() {
		return errors.New("field cannot be set")
	}

	switch value.Kind() {
	case reflect.String:
		value.SetString(envValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(envValue, 10, 64)
		if err != nil {
			return err
		}
		value.SetInt(intValue)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(envValue, 10, 64)
		if err != nil {
			return err
		}
		value.SetUint(uintValue)
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(envValue, 64)
		if err != nil {
			return err
		}
		value.SetFloat(floatValue)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(envValue)
		if err != nil {
			return err
		}
		value.SetBool(boolValue)
	case reflect.Struct:
		// Recursively handle nested structs
		return loadConfigFromEnvOrFile(value.Addr().Interface(), "")
	default:
		return fmt.Errorf("unsupported field type: %s", value.Kind())
	}

	return nil
}
