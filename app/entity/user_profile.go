package entity

type UserProfile struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"-"`
	IsLogin  bool `json:"-"`
}

type UserProfileBuilder struct {
	profile UserProfile
}

func NewBuilder() *UserProfileBuilder {
	return &UserProfileBuilder{UserProfile{
		IsLogin: false,
	}}
}

func (builder *UserProfileBuilder) WithUserId(userId string) *UserProfileBuilder {
	builder.profile.UserID = userId
	return builder
}

func (builder *UserProfileBuilder) WithEmail(email string) *UserProfileBuilder {
	builder.profile.Email = email
	return builder
}

func (builder *UserProfileBuilder) WithAddress(address string) *UserProfileBuilder {
	builder.profile.Address = address
	return builder
}

func (builder *UserProfileBuilder) WithPassword(password string) *UserProfileBuilder {
	builder.profile.Password = password
	return builder
}

func (builder *UserProfileBuilder) Build() (UserProfile, error) {
	return builder.profile, nil
}
