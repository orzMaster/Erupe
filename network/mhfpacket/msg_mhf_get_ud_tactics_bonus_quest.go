package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetUdTacticsBonusQuest represents the MSG_MHF_GET_UD_TACTICS_BONUS_QUEST
type MsgMhfGetUdTacticsBonusQuest struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetUdTacticsBonusQuest) Opcode() network.PacketID {
	return network.MSG_MHF_GET_UD_TACTICS_BONUS_QUEST
}

// Parse parses the packet from binary
func (m *MsgMhfGetUdTacticsBonusQuest) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetUdTacticsBonusQuest) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}