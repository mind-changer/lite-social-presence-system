package def

type ViewFriendsResponse struct {
	Friends []string `json:"friends,omitempty"`
}

type AddFriendResponse struct {
	Status string `json:"status,omitempty"`
}

type RemoveFriendResponse struct {
	Status string `json:"status,omitempty"`
}

type UpdateUserStatusResponse struct {
	Status string `json:"status,omitempty"`
}

type AcceptFriendRequestResponse struct {
	Status string `json:"status,omitempty"`
}

type RejectRequestResponse struct {
	Status string `json:"status,omitempty"`
}

type CreatePartyResponse struct {
	PartyId string `json:"partyId,omitempty"`
}

type SendPartyInvitationResponse struct {
	Status string `json:"status,omitempty"`
}

type AcceptPartyInvitationResponse struct {
	Status string `json:"status,omitempty"`
}

type RejectPartyInvitationResponse struct {
	Status string `json:"status,omitempty"`
}

type KickPartyMemberResponse struct {
	Status string `json:"status,omitempty"`
}
type LeavePartyResponse struct {
	Status string `json:"status,omitempty"`
}
