package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetCafeDuration represents the MSG_MHF_GET_CAFE_DURATION
type MsgMhfGetCafeDuration struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetCafeDuration) Opcode() network.PacketID {
	return network.MSG_MHF_GET_CAFE_DURATION
}

// Parse parses the packet from binary
func (m *MsgMhfGetCafeDuration) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetCafeDuration) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}