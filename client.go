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
	Ip6Ranges  []string
	Res        int
	ResMessage string
	DebugInfo  struct {
		IdInfo string
	}
}

func (c *Client) getResponse() *Response {
	if c.response == nil {
		c.response = NewResponse()
	}
	return c.response
}

func (c *Client) getIpRanges() string {
	return c.getResponse().Raw
}

func (c *Client) getIPv4Ranges() []string {

	var ipsResponse = IpsResponse{}

	if err := json.Unmarshal([]byte(c.getResponse().Raw), &ipsResponse); err != nil {
		log.Fatalln("error:", err)
	}

	return ipsResponse.IpRanges
}
