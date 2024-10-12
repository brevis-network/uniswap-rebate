package webapi

import (
	"database/sql/driver"

	"google.golang.org/protobuf/encoding/protojson"
)

// impl db required interface for objects
func (c *NewProofReq) Value() (driver.Value, error) {
	return protojson.Marshal(c)
}

// tried double pointer but not work
func (u *NewProofReq) Scan(value interface{}) error {
	if u == nil {
		u = new(NewProofReq)
	}
	return protojson.Unmarshal(value.([]byte), u)
}
