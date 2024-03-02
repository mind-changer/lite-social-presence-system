package def

type AddFriendRequest struct {
	UserId string `json:"userId,omitempty"`
}

type UpdateUserStatusRequest struct {
	UserStatus string `json:"userStatus,omitempty"`
}

type SendPartyInvitationRequest struct {
	UserId string `json:"userId,omitempty"`
}
