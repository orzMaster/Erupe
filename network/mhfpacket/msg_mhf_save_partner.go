package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfSavePartner represents the MSG_MHF_SAVE_PARTNER
type MsgMhfSavePartner struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfSavePartner) Opcode() network.PacketID {
	return network.MSG_MHF_SAVE_PARTNER
}

// Parse parses the packet from binary
func (m *MsgMhfSavePartner) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfSavePartner) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}