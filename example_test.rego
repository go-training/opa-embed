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
