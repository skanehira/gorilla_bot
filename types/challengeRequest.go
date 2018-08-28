package types

// ChallengeRequest Verification URL Request
type ChallengeRequest struct {
	Token string `json:"token"`
	RequestType
	Challenge
}

// Challenge Verification URL Response
type Challenge struct {
	Challenge string `json:"challenge"`
}

// NewChallenge New Verification URL Response
func NewChallenge(challenge string) *Challenge {
	return &Challenge{
		Challenge: challenge,
	}
}
