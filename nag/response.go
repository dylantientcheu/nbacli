package nag

import (
	"encoding/json"
)

// Response represents response from NBA endpoint.
type Response struct {
	Resource string `json:"resource"`
	// Parameters map[string]interface{} `json:"parameters"`
	ResultSet  ResultSet   `json:"resultSet"`
	ResultSets []ResultSet `json:"resultSets"`
}

// ResultSet represents relevant data from Response.
type ResultSet struct {
	Name    string          `json:"name"`
	Headers []string        `json:"headers"`
	RowSet  [][]interface{} `json:"rowSet"`
}

// Map maps ResultSet(s) from Response based on name and headers.
func Map(res Response) map[string]interface{} {
	var resultSets []ResultSet
	if res.ResultSets != nil {
		resultSets = res.ResultSets
	} else {
		resultSets = []ResultSet{res.ResultSet}
	}

	m := make(map[string]interface{})
	for _, rs := range resultSets {
		var rows []map[string]interface{}
		for _, row := range rs.RowSet {
			mm := make(map[string]interface{})
			for i, h := range rs.Headers {
				mm[h] = row[i]
			}
			rows = append(rows, mm)
		}
		m[rs.Name] = rows
	}
	return m
}

// JSON returns a JSON representation of Map
func JSON(res Response) (json.RawMessage, error) {
	return json.Marshal(Map(res))
}
