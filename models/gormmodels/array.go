package gormmodels

import (
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
