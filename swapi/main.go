package main

import (
	"fmt"
	"strconv"

	"github.com/jonathansudibya/fished"
)

func main() {
	e := fished.New(10)

	e.Rules = []fished.Rule{
		fished.Rule{
			Input:      []string{"luke_id"},
			Output:     "luke",
			Expression: "fetchPeople(luke_id)",
		},
		fished.Rule{
			Input:      []string{"cpo_id"},
			Output:     "cpo",
			Expression: "fetchPeople(cpo_id)",
		},
		fished.Rule{
			Input:      []string{"cpo", "luke"},
			Output:     "isHeavier",
			Expression: "peopleMass(luke) > peopleMass(cpo) ? 'luke' : '3cpo'",
		},
	}
	e.RuleFunctions = map[string]fished.RuleFunction{
		"fetchPeople": func(args ...interface{}) (interface{}, error) {
			if len(args) == 1 {
				index, err := strconv.Atoi(args[0].(string))
				if err != nil {
					return nil, err
				}
				res, err := FetchPeople(index)
				fmt.Println(res)
				return res, err
			}
			return nil, fmt.Errorf("fetchPeople arguments should have: 1 get: %d", len(args))
		},
		"peopleMass": func(args ...interface{}) (interface{}, error) {
			if len(args) == 1 {
				mass := args[0].(*People).Mass
				m, err := strconv.Atoi(mass)
				return (float64)(m), err
			}
			return nil, fmt.Errorf("fetchPeople arguments should have: 1 get: %d", len(args))
		},
	}
	e.Facts["luke_id"] = "1"
	e.Facts["cpo_id"] = "2"
	res, errs := e.Run("isHeavier")
	for _, err := range errs {
		fmt.Println("[ERROR]", err)
	}
	fmt.Println(res)
}
