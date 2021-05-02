// +build !go1.16

package policy

import (
	"io/ioutil"
	"log"
)

func ReadPolicy(path string) ([]byte, error) {
	// load policy
	policy, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("can't read policy file: %v", err)
	}

	return policy, nil
}
