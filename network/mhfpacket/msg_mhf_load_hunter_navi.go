package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfLoadHunterNavi represents the MSG_MHF_LOAD_HUNTER_NAVI
type MsgMhfLoadHunterNavi struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfLoadHunterNavi) Opcode() network.PacketID {
	return network.MSG_MHF_LOAD_HUNTER_NAVI
}

// Parse parses the packet from binary
func (m *MsgMhfLoadHunterNavi) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfLoadHunterNavi) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}