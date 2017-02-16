package incapsula_api_go

import (
	"reflect"
	"testing"
)

func Test_it_returns_IPv4_ranges_only(t *testing.T) {
	client := Client{}
	returned_ipv4_ranges := client.getIPv4Ranges()
	expected_ipv4_ranges := []string{
		"199.83.128.0/21",
		"198.143.32.0/19",
		"149.126.72.0/21",
		"103.28.248.0/22",
		"185.11.124.0/22",
		"192.230.64.0/18",
		"45.64.64.0/22",
		"107.154.0.0/16"}

	if !reflect.DeepEqual(returned_ipv4_ranges, expected_ipv4_ranges) {
		t.Errorf("Expected '%s' but got '%s'", expected_ipv4_ranges, returned_ipv4_ranges)
	}
}

func TestClientReturnsIncapsulaIpRanges(t *testing.T) {
	client := Client{}
	jsonResponse := client.getIpRanges()
	expectedResponse := NewResponse().Raw

	if jsonResponse != expectedResponse {
		t.Errorf("Expected '%s' but got '%s'", expectedResponse, jsonResponse)
	}
}
