package incapsula_api_go

import (
	"reflect"
	"testing"
)

func Test_it_returns_IPv6_ranges_only(t *testing.T) {
	client := Client{}
	expected_ranges := []string{
		"2a02:e980::/29",
	}

	ipv6_ranges := client.getIPv6Ranges()

	if !reflect.DeepEqual(ipv6_ranges, expected_ranges) {
		t.Errorf("Expected '%s' but got '%s'", expected_ranges, ipv6_ranges)
	}
}

func Test_it_returns_IPv4_ranges_only(t *testing.T) {
	client := Client{}
	expected_ranges := []string{
		"199.83.128.0/21",
		"198.143.32.0/19",
		"149.126.72.0/21",
		"103.28.248.0/22",
		"185.11.124.0/22",
		"192.230.64.0/18",
		"45.64.64.0/22",
		"107.154.0.0/16"}

	for _, ip_ranges := range [][]string{client.getIPv4Ranges(), client.getIpRanges()} {
		if !reflect.DeepEqual(ip_ranges, expected_ranges) {
			t.Errorf("Expected '%s' but got '%s'", expected_ranges, ip_ranges)
		}
	}
}

func Test_it_returns_the_raw_json_response(t *testing.T) {
	client := Client{}
	jsonResponse := client.getRawResponse()
	expectedResponse := NewResponse().Raw

	if jsonResponse != expectedResponse {
		t.Errorf("Expected '%s' but got '%s'", expectedResponse, jsonResponse)
	}
}

func Test_it_returns_the_res_code_from_the_response(t *testing.T) {
	client := Client{}
	res := client.getRes()
	expectedRes := 0
	if res != expectedRes {
		t.Errorf("Expected '%v' but got '%v'", expectedRes, res)
	}
}

func Test_it_returns_the_res_message_from_the_response(t *testing.T) {
	client := Client{}
	resMessage := client.getResMessage()
	expectedResMessage := "OK"
	if resMessage != expectedResMessage {
		t.Errorf("Expected '%s' but got '%s'", expectedResMessage, resMessage)
	}
}

func Test_it_returns_the_debug_info_from_the_response(t *testing.T) {
	client := Client{}
	debugInfo := client.getDebugInfo()
	expectedDebugInfo := map[string]string{"id-info": "9088"}
	if !reflect.DeepEqual(debugInfo, expectedDebugInfo) {
		t.Errorf("Expected '%v' but got '%v'", expectedDebugInfo, debugInfo)
	}
}
