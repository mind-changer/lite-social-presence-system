package def

type AddFriendRequest struct {
	User string `json:"userId,omitempty"`
}

type UpdateUserStatusRequest struct {
	UserStatus string `json:"userStatus,omitempty"`
}
