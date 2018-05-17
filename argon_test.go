package argon_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	argon "github.com/sparkymat/go-argon"
)

var _ = Describe("Argon", func() {
	var stateMachine argon.StateMachine
	var err error
	var entity argon.StatefulEntity
	var config *argon.Config

	BeforeEach(func() {
		stateMachine, err = argon.NewStateMachine(entity, *config)
	})

	Context("With valid entity", func() {
		entity = &entityType{}

		Context("With empty config", func() {
			config = &argon.Config{}
			It("should return error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Context("With config baving states but no edges", func() {
			config = &argon.Config{
				States: []argon.State{Initial, Pending, Final},
				Edges:  []argon.Edge{},
			}
			It("should return error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

const (
	Initial argon.State = "initial"
	Pending argon.State = "pending"
	Foo     argon.State = "foo"
	Bar     argon.State = "bar"
	Final   argon.State = "final"
)

type entityType struct {
	state argon.State
}

type entityTypeWithOn struct {
	entityType
}

type entityTypeWithAfter struct {
	entityType
}

type entityTypeWithBoth struct {
	entityType
}

func (et *entityType) GetState() argon.State {
	return et.state
}

func (et *entityType) SetState(state argon.State) {
	et.state = state
}
