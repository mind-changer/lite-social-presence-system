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
