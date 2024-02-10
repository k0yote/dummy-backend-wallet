package issuer

type IssuerNode struct {
	baseURL  string
	username string
	password string
	client   Client
}

type RequestOpts struct {
	BaseURL  string
	Username string
	Password string
}

func Init(c *RequestOpts) IssuerNode {
	issuerNode := IssuerNode{}
	issuerNode.baseURL = c.BaseURL
	issuerNode.username = c.Username
	issuerNode.password = c.Password

	issuerNode.client = Client{
		BaseURL:  issuerNode.baseURL,
		Username: issuerNode.username,
		Password: issuerNode.password,
	}

	return issuerNode
}
