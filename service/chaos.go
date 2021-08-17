package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const ()

// StartChaosTest install instance with custom mycnf file
func StartChaosTest(serverAddr, cmd string) (response []byte, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/chaosblade?cmd=%s", serverAddr, cmd))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if response, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	return
}
