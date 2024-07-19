package types

type RequireEmailError struct{}

func (e *RequireEmailError) Error() string {
	return "Email address required"
}

type CreateNewUserError struct{}

func (e *CreateNewUserError) Error() string {
	return "Unable to sign up. Please try again later"
}

type LoginError struct{}

func (e *LoginError) Error() string {
	return "Unable to sign in. Please try again later"
}
