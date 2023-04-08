package client

import "fmt"

type ApiUser struct {
	Data User `json:"data"`
}

type ApiUsers struct {
	Data []User `json:"data"`
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

type CreateUser struct {
	User struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	} `json:"user"`
}

// PrintUser -	Prints user in a human readable format
func (u *User) PrintUser() string {
	return fmt.Sprintf("User: %s	Role: %s	Email: %s", u.ID, u.Role, u.Email)
}

// PrintUsers -	Prints users in a human readable format
func PrintUsers(u *[]User) string {
	var users string
	for _, user := range *u {
		users += user.PrintUser() + "\n"
	}
	return users
}

// ApiRule -
type ApiRule struct {
	Data Rule `json:"data"`
}

// ApiRules -
type ApiRules struct {
	Data []Rule `json:"data"`
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

type CreateRule struct {
	Rule struct {
		Action      string `json:"action"`
		Destination string `json:"destination"`
		PortRange   string `json:"port_range"`
		PortType    string `json:"port_type"`
		UserId      string `json:"user_id,omitempty"`
	} `json:"rule"`
}

// PrintRule -	Prints rule in a human readable format
func (r *Rule) PrintRule() string {
	return fmt.Sprintf("Rule: %s	%s	%s	%s", r.Action, r.Destination, r.PortRange, r.PortType)
}

// PrintRules -	Prints rules in a human readable format
func PrintRules(r *[]Rule) string {
	var rules string
	for _, rule := range *r {
		rules += rule.PrintRule() + "\n"
	}
	return rules
}

// ApiDevice -
type ApiDevice struct {
	Data Device `json:"data"`
}

// ApiDevices -
type ApiDevices struct {
	Data []Device `json:"data"`
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
	UserId                        string      `json:"user_id"`
}

type CreateDevice struct {
	Device struct {
		AllowedIPs             []string `json:"allowed_ips,omitempty"`
		Description            string   `json:"description"`
		DNS                    []string `json:"dns,omitempty"`
		Endpoint               string   `json:"endpoint,omitempty"`
		IPv4                   string   `json:"ipv4,omitempty"`
		IPv6                   string   `json:"ipv6,omitempty"`
		MTU                    int      `json:"mtu,omitempty"`
		Name                   string   `json:"name"`
		PersistentKeepalive    int      `json:"persistent_keepalive,omitempty"`
		PresharedKey           string   `json:"preshared_key,omitempty"`
		PublicKey              string   `json:"public_key"`
		UseDefaultAllowedIPs   bool     `json:"use_default_allowed_ips,omitempty"`
		UseDefaultDNS          bool     `json:"use_default_dns,omitempty"`
		UseDefaultEndpoint     bool     `json:"use_default_endpoint,omitempty"`
		UseDefaultMTU          bool     `json:"use_default_mtu,omitempty"`
		UseDefaultPersistentKA bool     `json:"use_default_persistent_keepalive,omitempty"`
		UserId                 string   `json:"user_id"`
	} `json:"device"`
}

// PrintDevice -	Prints device in a human readable format
func (d *Device) PrintDevice() string {
	return fmt.Sprintf("Device: %s	IPv4: %s	IPv6: %s	Endpoint: %s	Allowed IPs: %s", d.Name, d.IPv4, d.IPv6, d.Endpoint, d.AllowedIPs)
}

// PrintDevices -	Prints devices in a human readable format
func PrintDevices(d *[]Device) string {
	var devices string
	for _, device := range *d {
		devices += device.PrintDevice() + "\n"
	}
	return devices
}

// ApiConfiguration -
type ApiConfiguration struct {
	Data Configuration `json:"data"`
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

// PrintConfiguration -	Prints configuration in a human readable format

func (c *Configuration) PrintConfiguration() string {
	return fmt.Sprintf("Configuration: %s	VPN Session Duration: %d", c.ID, c.VpnSessionDuration)
}
