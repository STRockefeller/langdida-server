package gormmodels

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Define a custom data type for array of strings
type ArrayOfStrings []string

func (aos ArrayOfStrings) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "text"
	case "postgres":
		return "text[]"
	}
	return ""
}

func (aos *ArrayOfStrings) Scan(value interface{}) error {
	if value == nil {
		*aos = nil
		return nil
	}
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, aos)
	case string:
		return json.Unmarshal([]byte(v), aos)
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *ArrayOfStrings", value)
	}
}

func (aos ArrayOfStrings) Value() (driver.Value, error) {
	if aos == nil {
		return nil, nil
	}
	return json.Marshal(aos)
}
