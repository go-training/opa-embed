opa_test:
	opa test -v policy/*.rego

go_test:
	go test -v .

test: opa_test go_test
