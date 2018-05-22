package argon_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	argon "github.com/sparkymat/go-argon"
)

var _ = Describe("argon.NewStateMachine", func() {
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

		Context("With config having states but no edges", func() {
			config := argon.Config{
				States: []argon.State{Initial, Pending, Final},
				Edges:  []argon.Edge{},
			}

			It("should return error", func() {
				Expect(createStateMachine(entity, config)).ShouldNot(Succeed())
			})
		})

		Context("With config having valid states and edges but no start state", func() {
			config := argon.Config{
				States: []argon.State{Initial, Pending, Final},
				Edges: []argon.Edge{
					{From: Pending, To: Final, Action: "DoThings"},
				},
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

type entityTypeWithCallbacks struct {
	entityType
}

func (et *entityType) GetState() argon.State {
	return et.state
}

func (et *entityType) SetState(state argon.State) {
	et.state = state
}

func (et *entityType) BeforeAction(action string) {
}

func (et *entityType) AfterAction(action string, err error) {
}

func (et *entityType) OnAction(action string) error {
	return nil
}

var validConfig = argon.Config{
	States:     []argon.State{Initial, Pending, Final},
	StartState: Initial,
	Edges: []argon.Edge{
		{From: Initial, To: Pending, Action: "Start"},
		{From: Pending, To: Final, Action: "Finish"},
	},
}
