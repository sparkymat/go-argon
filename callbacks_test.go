package argon_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	argon "github.com/sparkymat/go-argon"
)

var _ = Describe("Callbacks", func() {
	var entity = &entityTypeWithCallbacks{entityType{State: Initial}}
	var stateMachine, _ = argon.NewStateMachine(entity, validConfig)

	It("should invoke On callback", func() {
		Expect(stateMachine.Do("Start")).Should(Succeed())
	})
})
