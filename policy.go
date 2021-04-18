// +build !go1.16

package main

import (
	"io/ioutil"
	"log"
)

func readPolicy(path string) ([]byte, error) {
	// load policy
	policy, err := ioutil.ReadFile(policyFile)
	if err != nil {
		log.Fatalf("can't read policy file: %v", err)
	}

	return policy, nil
}
