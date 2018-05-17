package argon_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoArgon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoArgon Suite")
}
