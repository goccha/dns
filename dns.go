package dns

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

func Parse(rawurl string, protocol ...string) (*url.URL, error) {
	scheme := "http"
	uri := rawurl
	if index := strings.Index(uri, "://"); index > 0 {
		scheme = uri[:index]
		uri = uri[index+3:]
	}
	hostName := uri
	if strings.HasPrefix(hostName, "srv+") {
		path := ""
		if index := strings.Index(uri, "/"); index > 0 {
			hostName = uri[:index]
			path = uri[index:]
		}
		service := ""
		hostName = hostName[4:]
		if index := strings.Index(hostName, "+"); index > 0 {
			service = hostName[:index]
			hostName = hostName[index:]
		}
		proto := ""
		if protocol != nil && len(protocol) > 0 {
			proto = protocol[0]
		}
		if _, addrs, err := lookupSRV(service, proto, hostName); err != nil {
			return nil, err
		} else {
			hostName = fmt.Sprintf("%s:%d", strings.TrimRight(addrs[0].Target, "."), addrs[0].Port)
			return url.Parse(fmt.Sprintf("%s://%s%s", scheme, hostName, path))
		}
	}
	return url.Parse(rawurl)
}

var lookupSRV = func(service, proto, name string) (cname string, addrs []*net.SRV, err error) {
	return net.LookupSRV(service, proto, name)
}
