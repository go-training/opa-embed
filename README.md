# opa-demo

[![Run Tests](https://github.com/go-training/opa-demo/actions/workflows/go.yml/badge.svg)](https://github.com/go-training/opa-demo/actions/workflows/go.yml)
[![Open Policy Agent Testing](https://github.com/go-training/opa-demo/actions/workflows/opa.yml/badge.svg)](https://github.com/go-training/opa-demo/actions/workflows/opa.yml)

How to use Role-based access control (RBAC) with the Open Policy Agent. See the [reference guide](https://www.openpolicyagent.org/docs/latest/comparison-to-other-systems/#role-based-access-control-rbac).

* [Chinese Blog](https://blog.wu-boy.com/2021/04/setup-rbac-role-based-access-control-using-open-policy-agent/)
* [Chinese Video](https://www.youtube.com/watch?v=AkMVh5XRcuI)

## Create RBAC policy

[embedmd]:# (example.rego)
```rego
package rbac.authz

# user-role assignments
group_roles := {
	"design_group_kpi_editor": ["kpi_editor_design", "viewer_limit_ds"],
	"system_group_kpi_editor": ["kpi_editor_system", "viewer_limit_ds"],
	"manufacture_group_kpi_editor": ["kpi_editor_manufacture", "viewer"],
	"project_leader": ["viewer_limit_ds", "viewer_limit_m"],
}

# role-permissions assignments
role_permissions := {
	"admin": [
		{"action": "view_all", "object": "design"},
		{"action": "edit", "object": "design"},
		{"action": "view_all", "object": "system"},
		{"action": "edit", "object": "system"},
		{"action": "view_all", "object": "manufacture"},
		{"action": "edit", "object": "manufacture"},
	],
	"quality_head_design": [
		{"action": "view_all", "object": "design"},
		{"action": "edit", "object": "design"},
		{"action": "view_all", "object": "system"},
		{"action": "view_all", "object": "manufacture"},
	],
	"quality_head_system": [
		{"action": "view_all", "object": "design"},
		{"action": "view_all", "object": "system"},
		{"action": "edit", "object": "system"},
		{"action": "view_all", "object": "manufacture"},
	],
	"quality_head_manufacture": [
		{"action": "view_all", "object": "design"},
		{"action": "view_all", "object": "system"},
		{"action": "view_all", "object": "manufacture"},
		{"action": "edit", "object": "manufacture"},
	],
	"kpi_editor_design": [
		{"action": "view_all", "object": "design"},
		{"action": "edit", "object": "design"},
	],
	"kpi_editor_system": [
		{"action": "view_all", "object": "system"},
		{"action": "edit", "object": "system"},
	],
	"kpi_editor_manufacture": [
		{"action": "view_all", "object": "manufacture"},
		{"action": "edit", "object": "manufacture"},
	],
	"viewer": [
		{"action": "view_all", "object": "design"},
		{"action": "view_all", "object": "system"},
		{"action": "view_all", "object": "manufacture"},
	],
	"viewer_limit_ds": [
		{"action": "view_all", "object": "design"},
		{"action": "view_all", "object": "system"},
	],
	"viewer_limit_m": [{"action": "view_l3_project", "object": "manufacture"}],
}

# logic that implements RBAC.
default allow = false

allow {
	# lookup the list of roles for the user
	roles := group_roles[input.user[_]]

	# for each role in that list
	r := roles[_]

	# lookup the permissions list for role r
	permissions := role_permissions[r]

	# for each permission
	p := permissions[_]

	# check if the permission granted to r matches the user's request
	p == {"action": input.action, "object": input.object}
}
```

## Write Testing

Please download [OPA Binary](https://www.openpolicyagent.org/docs/latest/#running-opa) first.

[embedmd]:# (example_test.rego)
```rego
package rbac.authz

test_design_group_kpi_editor {
	allow with input as {"user": ["design_group_kpi_editor"], "action": "view_all", "object": "design"}
	allow with input as {"user": ["design_group_kpi_editor"], "action": "edit", "object": "design"}
	allow with input as {"user": ["design_group_kpi_editor"], "action": "view_all", "object": "system"}
	not allow with input as {"user": ["design_group_kpi_editor"], "action": "edit", "object": "system"}
	not allow with input as {"user": ["design_group_kpi_editor"], "action": "view_all", "object": "manufacture"}
	not allow with input as {"user": ["design_group_kpi_editor"], "action": "edit", "object": "manufacture"}
}

test_system_group_kpi_editor {
	allow with input as {"user": ["system_group_kpi_editor"], "action": "view_all", "object": "design"}
	not allow with input as {"user": ["system_group_kpi_editor"], "action": "edit", "object": "design"}
	allow with input as {"user": ["system_group_kpi_editor"], "action": "view_all", "object": "system"}
	allow with input as {"user": ["system_group_kpi_editor"], "action": "edit", "object": "system"}
	not allow with input as {"user": ["system_group_kpi_editor"], "action": "view_all", "object": "manufacture"}
	not allow with input as {"user": ["system_group_kpi_editor"], "action": "edit", "object": "manufacture"}
}

test_manufacture_group_kpi_editor {
	allow with input as {"user": ["manufacture_group_kpi_editor"], "action": "view_all", "object": "design"}
	not allow with input as {"user": ["manufacture_group_kpi_editor"], "action": "edit", "object": "design"}
	allow with input as {"user": ["manufacture_group_kpi_editor"], "action": "view_all", "object": "system"}
	not allow with input as {"user": ["manufacture_group_kpi_editor"], "action": "edit", "object": "system"}
	allow with input as {"user": ["manufacture_group_kpi_editor"], "action": "view_all", "object": "manufacture"}
	allow with input as {"user": ["manufacture_group_kpi_editor"], "action": "edit", "object": "manufacture"}
}

test_project_leader {
	allow with input as {"user": ["project_leader"], "action": "view_all", "object": "design"}
	not allow with input as {"user": ["project_leader"], "action": "edit", "object": "design"}
	allow with input as {"user": ["project_leader"], "action": "view_all", "object": "system"}
	not allow with input as {"user": ["project_leader"], "action": "edit", "object": "system"}
	not allow with input as {"user": ["project_leader"], "action": "view_all", "object": "manufacture"}
	not allow with input as {"user": ["project_leader"], "action": "edit", "object": "manufacture"}
	allow with input as {"user": ["project_leader"], "action": "view_l3_project", "object": "manufacture"}
}

test_design_group_kpi_editor_and_system_group_kpi_editor {
	allow with input as {"user": ["design_group_kpi_editor", "system_group_kpi_editor"], "action": "edit", "object": "design"}
	allow with input as {"user": ["design_group_kpi_editor", "system_group_kpi_editor"], "action": "edit", "object": "system"}
}
```

run test command:

```bash
$ opa test -v *.rego
data.rbac.authz.test_design_group_kpi_editor: PASS (8.604833ms)
data.rbac.authz.test_system_group_kpi_editor: PASS (7.260166ms)
data.rbac.authz.test_manufacture_group_kpi_editor: PASS (2.217125ms)
data.rbac.authz.test_project_leader: PASS (1.823833ms)
data.rbac.authz.test_design_group_kpi_editor_and_system_group_kpi_editor: PASS (1.150791ms)
--------------------------------------------------------------------------------
PASS: 5/5
```
