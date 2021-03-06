package argon

import (
	"errors"
	"fmt"
)

type StateMachine struct {
	entity  StatefulEntity
	config  Config
	edgeMap map[string]Edge
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
	s.edgeMap = make(map[string]Edge)

	for _, edge := range config.Edges {
		s.edgeMap[edge.Action] = edge
	}

	return s, nil
}

func (sm *StateMachine) Do(action string) error {
	var edge Edge
	var edgeExists bool

	if edge, edgeExists = sm.edgeMap[action]; !edgeExists {
		return errors.New("No such action exists")
	}

	if sm.entity.GetState() != edge.From {
		return errors.New("Invalid transition")
	}

	sm.entity.BeforeAction(action)

	sm.entity.SetState(edge.To)

	err := sm.entity.OnAction(action)

	if err != nil {
		sm.entity.SetState(edge.From)
	}

	sm.entity.AfterAction(action, err)

	return err
}
