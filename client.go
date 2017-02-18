package incapsula_api_go

import (
	"encoding/json"
	"log"
)

type Client struct {
	response *Response
}

type IpsResponse struct {
	IpRanges   []string
	Ipv6Ranges []string
	Res        int               `json:"res"`
	ResMessage string            `json:"res_message"`
	DebugInfo  map[string]string `json:"debug_info"`
}

func (c *Client) getResponse() *Response {
	if c.response == nil {
		c.response = NewResponse()
	}
	return c.response
}

func (c *Client) getIpsResponse() IpsResponse {
	var ipsResponse = IpsResponse{}

	if err := json.Unmarshal([]byte(c.getResponse().Raw), &ipsResponse); err != nil {
		log.Fatalln("error:", err)
	}
	return ipsResponse
}

func (c *Client) getIpRanges() []string {
	return c.getIpsResponse().IpRanges
}

func (c *Client) getIPv4Ranges() []string {
	return c.getIpRanges()
}

func (c *Client) getIPv6Ranges() []string {
	return c.getIpsResponse().Ipv6Ranges
}

func (c *Client) getRawResponse() string {
	return c.getResponse().Raw
}

func (c *Client) getRes() int {
	return c.getIpsResponse().Res
}

func (c *Client) getResMessage() string {
	return c.getIpsResponse().ResMessage
}

func (c *Client) getDebugInfo() map[string]string {
	return c.getIpsResponse().DebugInfo
}
