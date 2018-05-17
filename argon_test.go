package argon_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	argon "github.com/sparkymat/go-argon"
)

var _ = Describe("Argon", func() {
	var createStateMachine = func(e argon.StatefulEntity, c argon.Config) error {
		_, err := argon.NewStateMachine(e, c)
		return err
	}

	Context("With nil entity", func() {
		It("should return error", func() {
			Expect(createStateMachine(nil, validConfig)).ShouldNot(Succeed())
		})
	})

	Context("With valid entity", func() {
		entity := &entityType{}

		Context("With empty config", func() {
			config := argon.Config{}

			It("should return error", func() {
				Expect(createStateMachine(entity, config)).ShouldNot(Succeed())
			})
		})

		Context("With config baving states but no edges", func() {
			config := argon.Config{
				States: []argon.State{Initial, Pending, Final},
				Edges:  []argon.Edge{},
			}

			It("should return error", func() {
				Expect(createStateMachine(entity, config)).ShouldNot(Succeed())
			})
		})

		Context("With valid config", func() {
			config := validConfig

			It("should not return error", func() {
				Expect(createStateMachine(entity, config)).Should(Succeed())
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

var validConfig = argon.Config{
	States: []argon.State{Initial, Pending, Final},
	Edges: []argon.Edge{
		{From: Initial, To: Pending, Action: "start"},
		{From: Pending, To: Final, Action: "finish"},
	},
}
