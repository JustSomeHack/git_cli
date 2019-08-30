package models

// HTTPParams are connection parameters
type HTTPParams struct {
	URL                 string
	Proxy               string
	Timeout             int64
	URLAccessToken      string
	ContentType         string
	AuthorizationBearer string
	AuthorizationKey    string
	AuthorizationToken  string
	BasicAuthUser       string
	BasicAuthPass       string
	Headers             map[string]string
	Queries             map[string]string
}
