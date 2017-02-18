package incapsula_api_go

type IpsResponse struct {
	IpRanges   []string
	Ipv6Ranges []string
	Res        int               `json:"res"`
	ResMessage string            `json:"res_message"`
	DebugInfo  map[string]string `json:"debug_info"`
}
