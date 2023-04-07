package firezone

type ApiUser struct {
	data User
}

type ApiUsers struct {
	data []User
}

// User -
type User struct {
	ID                 string `json:"id"`
	Email              string `json:"email,omitempty"`
	Role               string `json:"role,omitempty"`
	LastSignedInAt     string `json:"last_signed_in_at,omitempty"`
	LastSignedInMethod string `json:"last_signed_in_method,omitempty"`
	UpdatedAt          string `json:"updated_at,omitempty"`
	InsertedAt         string `json:"inserted_at,omitempty"`
	DisabledAt         string `json:"disabled_at,omitempty"`
}

// ApiRule -
type ApiRule struct {
	data Rule
}

// ApiRules -
type ApiRules struct {
	data []Rule
}

// Rule -
type Rule struct {
	ID          string `json:"id"`
	UserId      string `json:"user_id,omitempty"`
	Action      string `json:"action,omitempty"`
	Destination string `json:"destination,omitempty"`
	PortRange   string `json:"port_range,omitempty"`
	PortType    string `json:"port_type,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	InsertedAt  string `json:"inserted_at,omitempty"`
}

// ApiDevice -
type ApiDevice struct {
	data Device
}

// ApiDevices -
type ApiDevices struct {
	data []Device
}

// Device -
type Device struct {
	ID                            string      `json:"id"`
	AllowedIPs                    []string    `json:"allowed_ips"`
	Description                   string      `json:"description"`
	DNS                           []string    `json:"dns"`
	Endpoint                      string      `json:"endpoint"`
	InsertedAt                    string      `json:"inserted_at"`
	IPv4                          string      `json:"ipv4"`
	IPv6                          string      `json:"ipv6"`
	LatestHandshake               interface{} `json:"latest_handshake"`
	MTU                           int         `json:"mtu"`
	Name                          string      `json:"name"`
	PersistentKeepalive           int         `json:"persistent_keepalive"`
	PresharedKey                  string      `json:"preshared_key"`
	PublicKey                     string      `json:"public_key"`
	RemoteIP                      interface{} `json:"remote_ip"`
	RXBytes                       interface{} `json:"rx_bytes"`
	ServerPublicKey               string      `json:"server_public_key"`
	TXBytes                       interface{} `json:"tx_bytes"`
	UpdatedAt                     string      `json:"updated_at,omitempty"`
	UseDefaultAllowedIPs          bool        `json:"use_default_allowed_ips"`
	UseDefaultDNS                 bool        `json:"use_default_dns"`
	UseDefaultEndpoint            bool        `json:"use_default_endpoint"`
	UseDefaultMTU                 bool        `json:"use_default_mtu"`
	UseDefaultPersistentKeepalive bool        `json:"use_default_persistent_keepalive"`
	UserID                        string      `json:"user_id"`
}

// ApiConfiguration -
type ApiConfiguration struct {
	data Configuration
}

// Configuration -
type Configuration struct {
	ID                                   string   `json:"id"`
	AllowUnprivilegedDeviceConfiguration bool     `json:"allow_unprivileged_device_configuration"`
	AllowUnprivilegedDeviceManagement    bool     `json:"allow_unprivileged_device_management"`
	DefaultClientAllowedIPs              []string `json:"default_client_allowed_ips"`
	DefaultClientDNS                     []string `json:"default_client_dns"`
	DefaultClientEndpoint                string   `json:"default_client_endpoint"`
	DefaultClientMTU                     int      `json:"default_client_mtu"`
	DefaultClientPersistentKeepalive     int      `json:"default_client_persistent_keepalive"`
	DisableVPNOnOIDCError                bool     `json:"disable_vpn_on_oidc_error"`
	InsertedAt                           string   `json:"inserted_at"`
	LocalAuthEnabled                     bool     `json:"local_auth_enabled"`
	Logo                                 struct{} `json:"logo"`
	OpenIDConnectProviders               []struct {
		AutoCreateUsers      bool   `json:"auto_create_users"`
		ClientID             string `json:"client_id"`
		ClientSecret         string `json:"client_secret"`
		DiscoveryDocumentURI string `json:"discovery_document_uri"`
		ID                   string `json:"id"`
		Label                string `json:"label"`
		RedirectURI          string `json:"redirect_uri"`
		ResponseType         string `json:"response_type"`
		Scope                string `json:"scope"`
	} `json:"openid_connect_providers"`
	SAMLIdentityProviders []struct {
		AutoCreateUsers bool   `json:"auto_create_users"`
		BaseURL         string `json:"base_url"`
		ID              string `json:"id"`
		Label           string `json:"label"`
		Metadata        string `json:"metadata"`
	} `json:"saml_identity_providers"`
	UpdatedAt          string `json:"updated_at,omitempty"`
	VpnSessionDuration int    `json:"vpn_session_duration"`
}
