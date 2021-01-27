package multiselector

import (
	"encoding/json"
	"testing"

	"github.com/globalsign/mgo/bson"
)

type selectorTestItem struct {
	workflow string
	Input    interface{}
	Output   interface{}
}

func TestSelector(t *testing.T) {

	tests := []selectorTestItem{
		selectorTestItem{
			workflow: "bson",
			Input:    "created <= %now%,something == 'okey'",
			Output: bson.M{
				"created": bson.M{
					"$lte": "%now%",
				},
				"something": "okey",
			},
		},
		selectorTestItem{
			workflow: "bson",
			Input: []string{
				"test == 'another'",
			},
			Output: bson.M{
				"test": "another",
			},
		},

		selectorTestItem{
			workflow: "bson",
			Input: []string{
				"test != another",
			},
			Output: bson.M{
				"test": bson.M{
					"$ne": "another",
				},
			},
		},
	}

	for _, test := range tests {
		s, _ := NewSelector(test.Input)
		switch test.workflow {
		case "bson":
			res, err := s.ToBson()
			if err != nil {
				t.Fatalf("cant parse inputs: %s", err.Error())
			}
			expected, err := json.Marshal(test.Output)
			if err != nil {
				t.Fatalf("cant marshal expected: %s", err.Error())
			}
			actual, err := json.Marshal(res)
			if err != nil {
				t.Fatalf("cant marshal actual: %s", err.Error())
			}
			if string(expected) != string(actual) {
				t.Fatalf(
					"dont work\n expected: %#v\n actual: %#v",
					test.Output,
					res,
				)
			}
		}
	}

}
