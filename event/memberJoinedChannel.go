package event

// MemberJoinedChannel Member join channel event
type MemberJoinedChannel struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Channel     string `json:"channel"`
	ChannelType string `json:"channel_type"`
	Team        string `json:"team"`
	Inviter     string `json:"inviter"`
}

// Name MemberJoinedChannel's Name
func (m *MemberJoinedChannel) Name() string {
	return m.User
}

// NewMemberJoinedChannel New struct
func NewMemberJoinedChannel() *MemberJoinedChannel {
	return &MemberJoinedChannel{}
}
