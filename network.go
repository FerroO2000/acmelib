package acmelib

import (
	"fmt"
	"slices"
	"strings"
)

// Network is the highest level entity in the package.
// Its main purpose is to hold all buses belonging to the same network.
// For example, a car can be seen as a network with multiple buses that
// are serving different areas or ECUs in the vehicle.
type Network struct {
	*entity

	buses    *set[EntityID, *Bus]
	busNames *set[string, EntityID]
}

// NewNetwork returns a new [Network] with the given name.
func NewNetwork(name string) *Network {
	return &Network{
		entity: newEntity(name),

		buses:    newSet[EntityID, *Bus]("bus"),
		busNames: newSet[string, EntityID]("bus name"),
	}
}

func (n *Network) errorf(err error) error {
	return fmt.Errorf(`project "%s" : %w`, n.name, err)
}

func (n *Network) modifyBusName(busEntID EntityID, newName string) {
	bus, err := n.buses.getValue(busEntID)
	if err != nil {
		panic(err)
	}

	oldName := bus.name
	n.busNames.modifyKey(oldName, newName, busEntID)
}

func (n *Network) String() string {
	var builder strings.Builder

	n.entity.stringify(&builder, 0)

	if n.buses.size() == 0 {
		return builder.String()
	}

	builder.WriteString("buses:\n")
	for _, bus := range n.Buses() {
		bus.stringify(&builder, 1)
		builder.WriteRune('\n')
	}

	return builder.String()
}

// UpdateName updates the name of the [Network].
func (n *Network) UpdateName(newName string) {
	n.name = newName
}

// AddBus adds a [Bus] to the [Network].
// It may return an error if the bus name is already taken.
func (n *Network) AddBus(bus *Bus) error {
	if err := n.busNames.verifyKey(bus.name); err != nil {
		return n.errorf(fmt.Errorf(`cannot add bus "%s" : %w`, bus.name, err))
	}

	n.buses.add(bus.entityID, bus)
	n.busNames.add(bus.name, bus.entityID)

	bus.setParentNetwork(n)

	return nil
}

// RemoveBus removes a [Bus] that matches the given entity id from the [Network].
// It may return an error if the bus with the given entity id is not part of the network.
func (n *Network) RemoveBus(busEntityID EntityID) error {
	bus, err := n.buses.getValue(busEntityID)
	if err != nil {
		return n.errorf(fmt.Errorf(`cannot remove bus with entity id "%s" : %w`, busEntityID, err))
	}

	bus.setParentNetwork(nil)

	n.buses.remove(busEntityID)
	n.busNames.remove(bus.name)

	return nil
}

// RemoveAllBuses removes all [Bus]es from the [Network].
func (n *Network) RemoveAllBuses() {
	for _, tmpBus := range n.buses.entries() {
		tmpBus.setParentNetwork(nil)
	}

	n.buses.clear()
	n.busNames.clear()
}

// Buses returns a slice of all [Bus]es in the [Network] sorted by name.
func (n *Network) Buses() []*Bus {
	busSlice := n.buses.getValues()
	slices.SortFunc(busSlice, func(a, b *Bus) int {
		return strings.Compare(a.name, b.name)
	})
	return busSlice
}
