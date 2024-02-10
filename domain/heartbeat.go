package domain

type HeartbeatRequest struct{}

type HeartbeatResponse struct {
	Redis bool `json:"redis"`
	Mongo bool `json:"mongo"`
}

type HeartbeatUsecase interface {
	Ping() *HeartbeatResponse
}
