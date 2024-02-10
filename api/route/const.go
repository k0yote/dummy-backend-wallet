package route

const (
	APIHeartbeat               = "/heartbeat"
	APIPasswordlessInit        = "/passwordless/init"
	APIPasswordlessAuth        = "/passwordless/authenticate"
	APISessions                = "/sessions"
	APIEmbeddedWalletInit      = "/embeded/wallet/init"
	APIEmbeddedWalletSignature = "/embeded/wallet/signature"
	APIReadonlyWalletGenerate  = "/readonly/wallet/generate"
	APIReadonlyWalletSearch    = "/readonly/wallet/:type/:id"
)
