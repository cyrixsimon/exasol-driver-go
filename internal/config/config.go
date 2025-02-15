package config

type Config struct {
	User                      string
	Password                  string
	AccessToken               string
	RefreshToken              string
	Host                      string
	Port                      int
	Params                    map[string]string // Connection parameters
	ApiVersion                int
	ClientName                string
	ClientVersion             string
	Schema                    string
	Autocommit                bool
	FetchSize                 int // Fetch size in kB
	QueryTimeout              int // query timeout in seconds
	Compression               bool
	ResultSetMaxRows          int
	Encryption                bool
	ValidateServerCertificate bool
	CertificateFingerprint    string
}
