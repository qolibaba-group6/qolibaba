package json

import (
	"database/sql/driver"
	"encoding/json"
)

// JSON is a custom type for handling JSON fields in GORM
type JSON map[string]interface{}

// Scan implements the Scanner interface for JSON
func (j *JSON) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), j)
}

// Value implements the Valuer interface for JSON
func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}
