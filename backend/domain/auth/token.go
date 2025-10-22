package auth

// Token represents an authentication token in the domain
type Token struct {
	Value  string
	UserID string
	Email  string
}

// NewToken creates a new Token
func NewToken(value, userID, email string) *Token {
	return &Token{
		Value:  value,
		UserID: userID,
		Email:  email,
	}
}

