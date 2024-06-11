package main

import (
	"os"
	"testing"
)

func TestCorrectRequest(t *testing.T) {
	t.Log("Start Correct Request Test")
	os.Setenv("RS_IMAGE_NAME", "marketplace.us1.greenlake-hpe.com/ezmeral/5.6.4/k8s.gcr.io/git-sync/git-sync:v3.6.1")
	os.Setenv("RS_TEST", "Yes")
	os.Setenv("RS_DEBUG", "Yes")
	os.Setenv("RS_DURATION_SEC", "10")
	main()
}
