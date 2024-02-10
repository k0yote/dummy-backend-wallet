package types

type KeyInfo struct {
	PrivateKey string
	Address    string
	AccountID  string
}

type InitializeRequest struct {
	Email string `json:"email" binding:"required"`
}

type InitializeResponse struct {
	Success bool `json:"success"`
}

type InitializeWalletRequest struct {
	Nonce string `json:"nonce" binding:"required"`
}

type InitializeWalletResponse struct{}

type PasscodeResult struct {
	Code   string
	Secret string
}

type AuthenticateRequest struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

type User struct {
	User                UserDetail `json:"user"`
	IsNewUser           bool       `json:"isNewUser"`
	Token               string     `json:"token"`
	RefreshToken        string     `json:"refreshToken"`
	SessionUpdateAction string     `json:"sessionUpdateAction"`
}

type UserDetail struct {
	ID             string        `json:"id"`
	CreatedAt      int64         `json:"createdAt"`
	LinkedAccounts []interface{} `json:"linkedAccounts"`
}

type LinkedEmailAccount struct {
	Type       string `json:"type"`
	Address    string `json:"address"`
	VerifiedAt int64  `json:"verifiedAt"`
}
type LinkedWalletAccount struct {
	Type       string `json:"type"`
	Address    string `json:"address"`
	VerifiedAt int64  `json:"verifiedAt"`
}

type JWTToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SearchType string

const (
	Address   SearchType = "address"
	AccountID SearchType = "accountId"
)

func (s SearchType) IsValid() bool {
	switch s {
	case Address, AccountID:
		return true
	}
	return false
}

func (s SearchType) String() string {
	return string(s)
}

type GenerateReadonlyWalletRequest struct {
	AccountID string `json:"accountId" binding:"required"`
}

type ReadonlyWalletResponse struct {
	Identifier string `json:"identifier,omitempty"`
	AccountID  string `json:"accountId,omitempty"`
	Address    string `json:"address,omitempty"`
	CipherText string `json:"cipherText,omitempty"`
	IV         string `json:"iv,omitempty"`
	EncKey     string `json:"encKey,omitempty"`
	Nonce      string `json:"nonce,omitempty"`
}

type EncryptedInfo struct {
	EncKey     string
	IV         string
	CipherText string
}

type ShareSecret struct {
	EncryptedRecoveryShare   string `json:"encryptedRecoveryShare"`
	EncryptedRecoveryShareIV string `json:"encryptedRecoveryShareIv"`
}

type EmbeddedWalletInitializeRequest struct {
	Identifier string `json:"identifier" binding:"required"`
}

type EmbeddedWalletInitializeResponse struct {
	Identifier string `json:"identifier" binding:"required"`
}
