package pack_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pack Suite")
}
