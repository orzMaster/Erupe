package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetEarthStatus represents the MSG_MHF_GET_EARTH_STATUS
type MsgMhfGetEarthStatus struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetEarthStatus) Opcode() network.PacketID {
	return network.MSG_MHF_GET_EARTH_STATUS
}

// Parse parses the packet from binary
func (m *MsgMhfGetEarthStatus) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetEarthStatus) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}