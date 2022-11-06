package nag_test

import (
	"encoding/json"
	"testing"

	"nba-cli/nag"
)

func TestGetJSON(t *testing.T) {
	tests := []struct {
		name string
		res  json.RawMessage
		want json.RawMessage
	}{
		{
			name: "result set",
			res: []byte(`
{
	"resultSets": [
		{
			"name": "name",
			"headers": ["H1", "H2"],
			"rowSet": [
				["R1.1", "R1.2"],
				["R2.1", "R2.2"]
			]
		}
	]
}
			`),
			want: []byte(`
{
	"name": [
		{
			"H1": "R1.1",
			"H2": "R1.2"
		},
		{
			"H1": "R2.1",
			"H2": "R2.2"
		}
	]
}
			`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO this is just a concatenation function
			var m map[string]interface{}
			json.Unmarshal(tt.want, &m)
			want, _ := json.Marshal(m)

			var res nag.Response
			json.Unmarshal(tt.res, &res)
			got, _ := nag.GetJSON(res)
			if string(got) != string(want) {
				t.Errorf("want:\n%s\ngot:\n%s\n", string(want), string(got))
			}
		})
	}
}
