package incapsula_api_go

import "testing"

func TestClientReturnsIncapsulaIpRanges(t *testing.T) {
	client := Client{};
	jsonResponse := client.getIpRanges();
	expectedResponse := `{"ipRanges":["199.83.128.0/21","198.143.32.0/19","149.126.72.0/21","103.28.248.0/22","185.11.124.0/22","192.230.64.0/18","45.64.64.0/22","107.154.0.0/16"],"ipv6Ranges":["2a02:e980::/29"],"res":0,"res_message":"OK","debug_info":{"id-info":"9088"}}`

	if jsonResponse != expectedResponse {
		t.Errorf("Expected '%s' but got '%s'", jsonResponse, expectedResponse);
	}
}