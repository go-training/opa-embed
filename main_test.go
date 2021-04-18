package main

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/open-policy-agent/opa/rego"
)

var query rego.PreparedEvalQuery

func setup() {
	var err error
	policy, err := readPolicy(policyFile)
	if err != nil {
		log.Fatal(err)
	}

	query, err = rego.New(
		rego.Query(defaultQuery),
		rego.Module(policyFile, string(policy)),
	).PrepareForEval(context.TODO())

	if err != nil {
		log.Fatalf("initial rego error: %v", err)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func Test_result(t *testing.T) {
	ctx := context.TODO()
	type args struct {
		ctx   context.Context
		query rego.PreparedEvalQuery
		input map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test_design_group_kpi_editor_edit_design",
			args: args{
				ctx:   ctx,
				query: query,
				input: map[string]interface{}{
					"user":   []string{"design_group_kpi_editor"},
					"action": "edit",
					"object": "design",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test_design_group_kpi_editor_edit_system",
			args: args{
				ctx:   ctx,
				query: query,
				input: map[string]interface{}{
					"user":   []string{"design_group_kpi_editor"},
					"action": "edit",
					"object": "system",
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "test_design_group_kpi_editor_and_system_group_kpi_editor_for_edit_design",
			args: args{
				ctx:   ctx,
				query: query,
				input: map[string]interface{}{
					"user":   []string{"design_group_kpi_editor", "system_group_kpi_editor"},
					"action": "edit",
					"object": "design",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test_design_group_kpi_editor_and_system_group_kpi_editor_for_edit_system",
			args: args{
				ctx:   ctx,
				query: query,
				input: map[string]interface{}{
					"user":   []string{"design_group_kpi_editor", "system_group_kpi_editor"},
					"action": "edit",
					"object": "system",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := result(tt.args.ctx, tt.args.query, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("result() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("result() = %v, want %v", got, tt.want)
			}
		})
	}
}
