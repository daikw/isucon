package caching

import (
	"encoding/json"
	"errors"
)

type QueryResult struct {
	Records    []interface{} `json:"-"`
	RawRecords []byte        `json:"raw"`
}

func (f *QueryResult) UnmarshalJSON(b []byte) error {
	type fleet QueryResult
	err := json.Unmarshal(b, (*fleet)(f))
	if err != nil {
		return err
	}

	for _, raw := range f.RawRecords {
		var v Vehicle
		err = json.Unmarshal(raw, &v)
		if err != nil {
			return err
		}
		var i interface{}
		switch v.Type {
		case "car":
			i = &Car{}
		case "truck":
			i = &Truck{}
		case "bike":
			i = &Bike{}
		default:
			return errors.New("unknown vehicle type")
		}
		err = json.Unmarshal(raw, i)
		if err != nil {
			return err
		}
		f.Vehicles = append(f.Vehicles, i)
	}
	return nil
}
