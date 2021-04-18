package main

import (
	"context"
	_ "embed"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

//go:embed example.rego
var policy []byte
var policyFile = "example.rego"
var defaultQuery = "x = data.rbac.authz.allow"

type input struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Object string `json:"object"`
}

func main() {
	s := input{
		User:   "design_group_kpi_editor",
		Action: "view_all",
		Object: "design",
	}

	input := map[string]interface{}{
		"user":   []string{s.User},
		"action": s.Action,
		"object": s.Object,
	}

	ctx := context.TODO()
	query, err := rego.New(
		rego.Query(defaultQuery),
		rego.Module(policyFile, string(policy)),
	).PrepareForEval(ctx)

	if err != nil {
		log.Fatalf("initial rego error: %v", err)
	}

	ok, _ := result(ctx, query, input)
	log.Println("", ok)
}

func result(ctx context.Context, query rego.PreparedEvalQuery, input map[string]interface{}) (bool, error) {
	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatalf("evaluation error: %v", err)
	} else if len(results) == 0 {
		log.Fatal("undefined result", err)
		// Handle undefined result.
	} else if result, ok := results[0].Bindings["x"].(bool); !ok {
		log.Fatalf("unexpected result type: %v", result)
	}

	return results[0].Bindings["x"].(bool), nil
}
