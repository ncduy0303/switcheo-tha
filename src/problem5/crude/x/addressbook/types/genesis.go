package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ContactList: []Contact{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in contact
	contactIdMap := make(map[uint64]bool)
	contactCount := gs.GetContactCount()
	for _, elem := range gs.ContactList {
		if _, ok := contactIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for contact")
		}
		if elem.Id >= contactCount {
			return fmt.Errorf("contact id should be lower or equal than the last id")
		}
		contactIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
