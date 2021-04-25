opa_test:
	opa test -v *.rego

server_test:
	opa test -v ./opa/*.rego

go_test:
	go test -v .

test: server_test opa_test go_test
