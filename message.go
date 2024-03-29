package acmelib

import (
	"fmt"
	"strings"
)

type Message struct {
	*entity

	parentNode *Node

	signals     map[EntityID]Signal
	signalNames map[string]EntityID

	signalPayload *signalPayload

	sizeByte int
	sizeBit  int
}

func NewMessage(name, desc string, sizeByte int) *Message {
	return &Message{
		entity: newEntity(name, desc),

		parentNode: nil,

		signals:     make(map[EntityID]Signal),
		signalNames: make(map[string]EntityID),

		signalPayload: newSignalPayload(sizeByte * 8),

		sizeByte: sizeByte,
		sizeBit:  sizeByte * 8,
	}
}

func (m *Message) hasParent() bool {
	return m.parentNode != nil
}

func (m *Message) addSignal(sig Signal) error {
	id := sig.EntityID()

	m.signals[id] = sig
	m.signalNames[sig.Name()] = id

	sig.setParent(m)

	return nil
}

func (m *Message) getSignalByID(sigID EntityID) (Signal, error) {
	if sig, ok := m.signals[sigID]; ok {
		return sig, nil
	}
	return nil, fmt.Errorf("signal not found")
}

// ---------------------------------------------------
// +++ START signalParent interface implementation +++
// ---------------------------------------------------

func (m *Message) errorf(err error) error {
	msgErr := fmt.Errorf(`message "%s": %v`, m.name, err)
	if m.hasParent() {
		return m.parentNode.errorf(msgErr)
	}
	return msgErr
}

func (m *Message) getSignalParentKind() signalParentKind {
	return signalParentKindMessage
}

func (m *Message) verifySignalName(name string) error {
	if _, ok := m.signalNames[name]; ok {
		return fmt.Errorf(`signal name "%s" is duplicated`, name)
	}
	return nil
}

func (m *Message) modifySignalName(sigID EntityID, newName string) {
	sig, err := m.getSignalByID(sigID)
	if err != nil {
		panic(err)
	}

	oldName := sig.Name()
	delete(m.signalNames, oldName)
	m.signalNames[newName] = sigID
}

func (m *Message) modifySignalSize(sigID EntityID, amount int) error {
	if amount == 0 {
		return nil
	}

	sig, err := m.getSignalByID(sigID)
	if err != nil {
		return err
	}

	if amount > 0 {
		return m.signalPayload.modifyStartBitsOnGrow(sig, amount)
	}

	m.signalPayload.modifyStartBitsOnShrink(sig, -amount)

	return nil
}

func (m *Message) toParentMessage() (*Message, error) {
	return m, nil
}

func (m *Message) toParentMultiplexerSignal() (*MultiplexerSignal, error) {
	return nil, fmt.Errorf(`cannot convert to "%s" signal parent is of kind "%s"`,
		signalParentKindMultiplexerSignal, signalParentKindMessage)
}

// -------------------------------------------------
// +++ END signalParent interface implementation +++
// -------------------------------------------------

func (m *Message) String() string {
	var builder strings.Builder

	builder.WriteString("\n+++START MESSAGE+++\n\n")
	builder.WriteString(m.toString())
	builder.WriteString(fmt.Sprintf("size: %d\n", m.sizeByte))

	signalsByPos := m.Signals()
	if len(signalsByPos) == 0 {
		return builder.String()
	}

	builder.WriteString("signals:\n")
	for _, sig := range signalsByPos {
		builder.WriteString(sig.String())
	}

	builder.WriteString("\n+++END MESSAGE+++\n")

	return builder.String()
}

// ----------------------------
// +++ START public getters +++
// ----------------------------

func (m *Message) Size() int {
	return m.sizeByte
}

func (m *Message) Signals() []Signal {
	return m.signalPayload.signals
}

func (m *Message) GetSignalByEntityID(signalEntityID EntityID) (Signal, error) {
	sig, err := m.getSignalByID(signalEntityID)
	if err != nil {
		return nil, m.errorf(fmt.Errorf(`cannot get signal with id "%s" : %v`, signalEntityID, err))
	}
	return sig, nil
}

func (m *Message) GetSignalByName(name string) (Signal, error) {
	id, ok := m.signalNames[name]
	if !ok {
		return nil, fmt.Errorf("signal name not found")
	}

	sig, err := m.getSignalByID(id)
	if err != nil {
		return nil, m.errorf(fmt.Errorf(`cannot get signal with name "%s" : %v`, name, err))
	}

	return sig, nil
}

// --------------------------
// +++ END public getters +++
// --------------------------

func (m *Message) UpdateName(name string) error {
	if m.hasParent() {
		if err := m.parentNode.messages.updateEntityName(m.entityID, m.name, name); err != nil {
			return m.errorf(err)
		}
	}
	return m.entity.UpdateName(name)
}

func (m *Message) AppendSignal(signal Signal) error {
	if err := m.verifySignalName(signal.Name()); err != nil {
		return m.errorf(fmt.Errorf(`cannot append signal "%s" : %v`, signal.Name(), err))
	}

	if err := m.signalPayload.append(signal); err != nil {
		return m.errorf(err)
	}

	return m.addSignal(signal)
}

func (m *Message) InsertSignal(signal Signal, startBit int) error {
	if err := m.verifySignalName(signal.Name()); err != nil {
		return m.errorf(fmt.Errorf(`cannot insert signal "%s" : %v`, signal.Name(), err))
	}

	if err := m.signalPayload.insert(signal, startBit); err != nil {
		return m.errorf(err)
	}

	return m.addSignal(signal)
}

func (m *Message) RemoveSignal(signalEntityID EntityID) error {
	sig, err := m.getSignalByID(signalEntityID)
	if err != nil {
		return m.errorf(fmt.Errorf(`cannot remove signal with entity id "%s" : %v`, signalEntityID, err))
	}

	sig.setParent(nil)

	delete(m.signals, signalEntityID)
	delete(m.signalNames, sig.Name())

	m.signalPayload.remove(signalEntityID)

	return nil
}

func (m *Message) RemoveAllSignals() {
	for tmpSigID, tmpSig := range m.signals {
		tmpSig.setParent(nil)
		delete(m.signals, tmpSigID)
	}

	m.signalPayload.removeAll()
}

func (m *Message) ShiftSignalLeft(signalEntityID EntityID, amount int) int {
	sig, err := m.getSignalByID(signalEntityID)
	if err != nil {
		return 0
	}

	return m.signalPayload.shiftLeft(sig, amount)
}

func (m *Message) ShiftSignalRight(signalEntityID EntityID, amount int) int {
	sig, err := m.getSignalByID(signalEntityID)
	if err != nil {
		return 0
	}

	return m.signalPayload.shiftRight(sig, amount)
}

func (m *Message) CompactSignals() {
	m.signalPayload.compact()
}
