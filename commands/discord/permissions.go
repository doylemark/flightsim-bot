package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/doylemark/flightsim-bot/commands/env"
)

// MemberHasPermission returns whether member has specified permission level
func MemberHasPermission(s *discordgo.Session, userID string, permission int64) (bool, error) {
	member, err := s.State.Member(env.GuildID, userID)
	if err != nil {
		if member, err = s.GuildMember(env.GuildID, userID); err != nil {
			return false, err
		}
	}

	for _, roleID := range member.Roles {
		role, err := s.State.Role(env.GuildID, roleID)
		if err != nil {
			return false, err
		}
		if role.Permissions&permission != 0 {
			return true, nil
		}
	}

	return false, nil
}
