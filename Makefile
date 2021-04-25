opa_test:
	opa test -v *.rego

rbac_test:
	opa test -v ./rbac/*.rego

go_test:
	go test -v .

test: rbac_test opa_test go_test
