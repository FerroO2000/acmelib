package acmelib

import (
	"fmt"
	"strings"
)

var defaulCANIDBuilder = NewCANIDBuilder("default CAN-ID builder").UseNodeID(0, 4).UseMessageID(4, 7).UseCAN2A()

// CANID is the CAN-ID of a [Message] within a [Bus].
// Every message should have a different CAN-ID.
type CANID uint32

// CANIDBuilderOpKind is the kind of an operation
// perfomed by the [CANIDBuilder].
type CANIDBuilderOpKind int

const (
	// CANIDBuilderOpKindMessagePriority represents an operation
	// that involves the message priority.
	CANIDBuilderOpKindMessagePriority CANIDBuilderOpKind = iota
	// CANIDBuilderOpKindMessageID represents an operation
	// that involves the message id.
	CANIDBuilderOpKindMessageID
	// CANIDBuilderOpKindNodeID represents an operation
	// that involves the node id.
	CANIDBuilderOpKindNodeID
	// CANIDBuilderOpKindBitMask represents a bit masking operation.
	CANIDBuilderOpKindBitMask
)

func (bok CANIDBuilderOpKind) String() string {
	switch bok {
	case CANIDBuilderOpKindMessagePriority:
		return "message-priority"
	case CANIDBuilderOpKindMessageID:
		return "message-id"
	case CANIDBuilderOpKindNodeID:
		return "node-id"
	case CANIDBuilderOpKindBitMask:
		return "bit-mask"
	default:
		return "unknown"
	}
}

// CANIDBuilderOp is an operation performed by the [CANIDBuilder].
type CANIDBuilderOp struct {
	kind CANIDBuilderOpKind
	from int
	len  int
}

func (bo *CANIDBuilderOp) stringify(b *strings.Builder, tabs int) {
	tabStr := getTabString(tabs)

	b.WriteString(fmt.Sprintf("%skind: %s\n", tabStr, bo.kind))
	b.WriteString(fmt.Sprintf("%sfrom: %d; len: %d\n", tabStr, bo.from, bo.len))
}

// Kind returns the kind of the operation.
func (bo *CANIDBuilderOp) Kind() CANIDBuilderOpKind {
	return bo.kind
}

// From returns the index of the first bit on which the operation is performed.
func (bo *CANIDBuilderOp) From() int {
	return bo.from
}

// Len returns the number of bits on which the operation is performed.
func (bo *CANIDBuilderOp) Len() int {
	return bo.len
}

// CANIDBuilder is a builder used to describe how to generate
// the CAN-ID of the messages within a [Bus].
type CANIDBuilder struct {
	*entity
	*withRefs[*Bus]

	operations []*CANIDBuilderOp
}

// NewCANIDBuilder creates a new [CANIDBuilder] with the given name.
func NewCANIDBuilder(name string) *CANIDBuilder {
	return &CANIDBuilder{
		entity:   newEntity(name, EntityKindCANIDBuilder),
		withRefs: newWithRefs[*Bus](),

		operations: []*CANIDBuilderOp{},
	}
}

func (b *CANIDBuilder) stringify(builder *strings.Builder, tabs int) {
	b.entity.stringify(builder, tabs)

	tabStr := getTabString(tabs)

	builder.WriteString(fmt.Sprintf("%soperations:\n", tabStr))
	for _, op := range b.operations {
		op.stringify(builder, tabs+1)
	}
}

func (b *CANIDBuilder) String() string {
	builder := new(strings.Builder)
	b.stringify(builder, 0)
	return builder.String()
}

// Operations returns the operations performed by the [CANIDBuilder].
func (b *CANIDBuilder) Operations() []*CANIDBuilderOp {
	return b.operations
}

// Calculate returns the CAN-ID calculated by applying the operations.
func (b *CANIDBuilder) Calculate(messagePriority MessagePriority, messageID MessageID, nodeID NodeID) CANID {
	canID := uint32(0)

	for _, op := range b.operations {
		if op.kind == CANIDBuilderOpKindBitMask {
			mask := uint32(0xFFFFFFFF) >> uint32(32-op.len)
			canID &= (mask << uint32(op.from))
			continue
		}

		tmpVal := uint32(0)
		switch op.kind {
		case CANIDBuilderOpKindMessagePriority:
			tmpVal = uint32(messagePriority)
		case CANIDBuilderOpKindMessageID:
			tmpVal = uint32(messageID)
		case CANIDBuilderOpKindNodeID:
			tmpVal = uint32(nodeID)
		}

		mask := uint32(0xFFFFFFFF) >> uint32(32-op.len)
		tmpVal &= mask

		tmpVal = tmpVal << uint32(op.from)
		canID |= tmpVal
	}

	return CANID(canID)
}

// UseMessagePriority adds an operation that involves the message priority from the given index.
// The length of the operation is fixed (2 bits).
func (b *CANIDBuilder) UseMessagePriority(from int) *CANIDBuilder {
	b.operations = append(b.operations, &CANIDBuilderOp{
		kind: CANIDBuilderOpKindMessagePriority,
		from: from,
		len:  2,
	})
	return b
}

// UseMessageID adds an operation that involves the message id from the given index and length.
func (b *CANIDBuilder) UseMessageID(from, len int) *CANIDBuilder {
	b.operations = append(b.operations, &CANIDBuilderOp{
		kind: CANIDBuilderOpKindMessageID,
		from: from,
		len:  len,
	})
	return b
}

// UseNodeID adds an operation that involves the node id from the given index and length.
func (b *CANIDBuilder) UseNodeID(from, len int) *CANIDBuilder {
	b.operations = append(b.operations, &CANIDBuilderOp{
		kind: CANIDBuilderOpKindNodeID,
		from: from,
		len:  len,
	})
	return b
}

// UseCAN2A adds a bit mask from 0 with a length of 11,
// which makes the calculated CAN-ID conformed to the CAN 2.0A.
func (b *CANIDBuilder) UseCAN2A() *CANIDBuilder {
	b.operations = append(b.operations, &CANIDBuilderOp{
		kind: CANIDBuilderOpKindBitMask,
		from: 0,
		len:  11,
	})
	return b
}

// UseBitMask adds a bit mask operation from the given index and length.
func (b *CANIDBuilder) UseBitMask(from, len int) {
	b.operations = append(b.operations, &CANIDBuilderOp{
		kind: CANIDBuilderOpKindBitMask,
		from: from,
		len:  len,
	})
}
