package dns

import (
	"net"
	"testing"
)

func TestParse(t *testing.T) {
	hasService := false
	lookupSRV = func(service, proto, name string) (cname string, addrs []*net.SRV, err error) {
		if hasService {
			expected := "service"
			if service != expected {
				t.Errorf("expected=%s, actual=%s", expected, service)
			}
		}
		cname = "goccha.org"
		addrs = []*net.SRV{
			{
				Target:   "goccha.org",
				Port:     8080,
				Priority: 1,
				Weight:   1,
			},
		}
		return
	}
	rawurl := "http://srv+test-service.local/p0/p2"
	if addr, err := Parse(rawurl); err != nil {
		t.Errorf("%v", err)
	} else {
		if addr.Scheme != "http" {
			t.Errorf("expected=%s, actual=%s", "http", addr.Scheme)
		}
		expected := "goccha.org:8080"
		if addr.Host != expected {
			t.Errorf("expected=%s, actual=%s", expected, addr.Host)
		}
		expected = "/p0/p2"
		if addr.Path != expected {
			t.Errorf("expected=%s, actual=%s", expected, addr.Path)
		}
	}
	hasService = true
	rawurl = "https://srv+service+test-service.local/p0/p2"
	if addr, err := Parse(rawurl); err != nil {
		t.Errorf("%v", err)
	} else {
		if addr.Scheme != "https" {
			t.Errorf("expected=%s, actual=%s", "https", addr.Scheme)
		}
		expected := "goccha.org:8080"
		if addr.Host != expected {
			t.Errorf("expected=%s, actual=%s", expected, addr.Host)
		}
		expected = "/p0/p2"
		if addr.Path != expected {
			t.Errorf("expected=%s, actual=%s", expected, addr.Path)
		}
	}
	hasService = false
	rawurl = "srv+test-service.local"
	if addr, err := Parse(rawurl); err != nil {
		t.Errorf("%v", err)
	} else {
		if addr.Scheme != "http" {
			t.Errorf("expected=%s, actual=%s", "http", addr.Scheme)
		}
		expected := "goccha.org:8080"
		if addr.Host != expected {
			t.Errorf("expected=%s, actual=%s", expected, addr.Host)
		}
		expected = ""
		if addr.Path != expected {
			t.Errorf("expected=%s, actual=%s", expected, addr.Path)
		}
	}
}
