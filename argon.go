package argon

import (
	"errors"
	"fmt"
)

type StateMachine struct {
	currentState State
	entity       StatefulEntity
	config       Config
}

func NewStateMachine(entity StatefulEntity, config Config) (StateMachine, error) {
	var s StateMachine

	if entity == nil {
		return s, errors.New("No stateful entity passed")
	}

	if len(config.States) == 0 {
		return s, errors.New("No valid states passed in config")
	}

	if len(config.Edges) == 0 {
		return s, errors.New("No valid edges passed in config")
	}

	stateExistence := make(map[State]struct{})
	for _, state := range config.States {
		stateExistence[state] = struct{}{}
	}

	if _, startStateExists := stateExistence[config.StartState]; !startStateExists {
		return s, errors.New("Start state not found in list of states")
	}

	actionExistence := make(map[string]struct{})

	for edgeIndex, edge := range config.Edges {
		if _, exists := actionExistence[edge.Action]; exists {
			return s, errors.New(fmt.Sprintf("Duplicate action in edge %v", edgeIndex))
		}
		actionExistence[edge.Action] = struct{}{}

		if _, exists := stateExistence[edge.From]; !exists {
			return s, errors.New(fmt.Sprintf("Invalid start state for edge %v", edgeIndex))
		}
		if _, exists := stateExistence[edge.To]; !exists {
			return s, errors.New(fmt.Sprintf("Invalid end state for edge %v", edgeIndex))
		}
	}

	s.entity = entity
	s.config = config

	return s, nil
}
