package event

import "gorilla_bot/common"

// MemberJoinedChannel Member join channel event
type MemberJoinedChannel struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Channel     string `json:"channel"`
	ChannelType string `json:"channel_type"`
	Team        string `json:"team"`
	Inviter     string `json:"inviter"`
}

// NewMemberJoinedChannel New struct
func NewMemberJoinedChannel() *MemberJoinedChannel {
	return new(MemberJoinedChannel)
}

// ToMap event struct to map
func (m *MemberJoinedChannel) ToMap() map[string]interface{} {
	return common.StructToMap(m)
}
