package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysReserve20E represents the MSG_SYS_reserve20E
type MsgSysReserve20E struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysReserve20E) Opcode() network.PacketID {
	return network.MSG_SYS_reserve20E
}

// Parse parses the packet from binary
func (m *MsgSysReserve20E) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysReserve20E) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}