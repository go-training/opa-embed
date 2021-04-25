opa_test:
	opa test -v *.rego

go_test:
	go test -v .

test: opa_test go_test
