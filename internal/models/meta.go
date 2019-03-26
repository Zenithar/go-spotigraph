package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Metadata represents object metadata
type Metadata map[string]interface{}

// Value exports metadata as json for sql driver
func (p Metadata) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

// Scan deserialize metadata from sql driver
func (p *Metadata) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed.")
	}

	return nil
}
