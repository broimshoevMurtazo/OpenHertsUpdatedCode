package emptyfieldcheker

import (
	"fmt"
	"reflect"
)

// EmptyField проверяет, являются ли поля структуры пустыми.
// Поля, указанные в списке исключений, не проверяются.
func EmptyField(obj any, excludedFields ...string) (bool, []string) {
	var empty bool
	var errors []string

	// Получаем значение и тип объекта через reflect
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	// Убедимся, что obj — структура
	if val.Kind() != reflect.Struct {
		return false, []string{"Provided value is not a struct"}
	}

	// Перебираем все поля структуры
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		fieldName := fieldType.Name

		// Проверяем, не является ли текущее поле исключением
		isExcluded := false
		for _, excluded := range excludedFields {
			if fieldName == excluded {
				isExcluded = true
				break
			}
		}

		// Пропускаем проверку, если поле исключено
		if isExcluded {
			continue
		}

		// Проверяем на пустоту значение поля
		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				errors = append(errors, fmt.Sprintf("Field: %v is empty", fieldName))
				empty = true
			}
		case reflect.Int:
			if field.Int() == 0 {
				errors = append(errors, fmt.Sprintf("Field: %v is empty", fieldName))
				empty = true
			}
		case reflect.Float64:
			if field.Float() == 0 {
				errors = append(errors, fmt.Sprintf("Field: %v is empty", fieldName))
				empty = true
			}
		case reflect.Struct:
			// Рекурсивно проверяем вложенную структуру
			isEmpty, nestedErrors := EmptyField(field.Interface(), excludedFields...)
			if isEmpty {
				errors = append(errors, fmt.Sprintf("Nested structure field: %v has empty fields", fieldName))
				empty = true
			}
			errors = append(errors, nestedErrors...)
		// Добавьте другие типы данных, если необходимо
		default:
			// Не обрабатываем другие типы данных
		}
	}

	return empty, errors
}
