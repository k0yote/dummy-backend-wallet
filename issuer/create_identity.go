package issuer

import (
	"encoding/json"
	"net/http"
	"time"
)

type DidMetadata struct {
	Method     string `json:"method"`
	Blockchain string `json:"blockchain"`
	Network    string `json:"network"`
	Type       string `json:"type"`
}

type CreateEntityArgs struct {
	DidMetadata *DidMetadata `json:"didMetadata,omitempty"`
}

type IdentityState struct {
	ClaimsTreeRoot string    `json:"claimsTreeRoot"`
	CreatedAt      time.Time `json:"createdAt"`
	ModifiedAt     time.Time `json:"modifiedAt"`
	State          string    `json:"state"`
	Status         string    `json:"status"`
}

type CreateEntityResult struct {
	Error      ErrorResult
	StatusCode int
	Address    string        `json:"address"`
	Identifier string        `json:"identifier"`
	State      IdentityState `json:"state"`
}

func (i *IssuerNode) CreateEntity(args *CreateEntityArgs) (*CreateEntityResult, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	resp, err := i.client.Do(CreateEntityEndpoint, b)
	if err != nil {
		return nil, err
	}

	var ret CreateEntityResult

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
		if err := json.Unmarshal(resp.Body, &ret.Error); err != nil {
			return nil, err
		}

		ret.StatusCode = resp.StatusCode

		return &ret, nil
	}

	if err := json.Unmarshal(resp.Body, &ret); err != nil {
		return nil, err
	}

	ret.StatusCode = resp.StatusCode

	return &ret, nil
}
