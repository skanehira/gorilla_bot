package event

import (
	"reflect"
)

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
	return &MemberJoinedChannel{}
}

// ToMap event struct to map
func (m *MemberJoinedChannel) ToMap() map[string]interface{} {
	event := make(map[string]interface{})
	element := reflect.ValueOf(m).Elem()
	size := element.NumField()

	for i := 0; i < size; i++ {
		key := element.Type().Field(i).Name
		value := element.Field(i).Interface()
		event[key] = value
	}

	return event
}
