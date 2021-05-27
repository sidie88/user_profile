package request

type UserProfileRequest struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
	RetypePassword string `json:"retype_password"`
}
