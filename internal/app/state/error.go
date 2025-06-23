package state

import (
	"encoding/json"
	"errors"
)

type JSONSerializableError interface {
	error
	json.Marshaler
	json.Unmarshaler
}

var _ JSONSerializableError = &JSONSerializableErr{}

type JSONSerializableErr struct {
	Err error
}

func (e JSONSerializableErr) Error() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

func (e JSONSerializableErr) MarshalJSON() ([]byte, error) {
	if e.Err == nil {
		return json.Marshal(e.Err)
	}
	return json.Marshal(e.Err.Error())
}

func (e *JSONSerializableErr) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		e.Err = nil
		return nil
	}
	var msg string
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}
	e.Err = errors.New(msg)
	return nil
}
