package namecheap

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	domainsDNSGetHosts   = "namecheap.domains.dns.getHosts"
	domainsDNSSetDefault = "namecheap.domains.dns.setDefault"
	domainsDNSSetHosts   = "namecheap.domains.dns.setHosts"
	domainsDNSGetList    = "namecheap.domains.dns.getList"
	domainsDNSSetCustom  = "namecheap.domains.dns.setCustom"
)

type DomainDNSGetHostsResult struct {
	Domain        string          `xml:"Domain,attr"`
	IsUsingOurDNS bool            `xml:"IsUsingOurDNS,attr"`
	Hosts         []DomainDNSHost `xml:"host"`
}

type DomainDNSGetListResult struct {
	Domain        string   `xml:"Domain,attr"`
	IsUsingOurDNS bool     `xml:"IsUsingOurDNS,attr"`
	Nameservers   []string `xml:"Nameserver"`
}

type DomainDNSHost struct {
	ID      int    `xml:"HostId,attr"`
	Name    string `xml:"Name,attr"`
	Type    string `xml:"Type,attr"`
	Address string `xml:"Address,attr"`
	MXPref  int    `xml:"MXPref,attr"`
	TTL     int    `xml:"TTL,attr"`
}

type DomainDNSSetDefaultResult struct {
	Domain    string `xml:"Domain,attr"`
	IsSuccess bool   `xml:"IsSuccess,attr"`
}

type DomainDNSSetHostsResult struct {
	Domain    string `xml:"Domain,attr"`
	IsSuccess bool   `xml:"IsSuccess,attr"`
}

func (client *Client) DomainsDNSGetHosts(sld, tld string) (*DomainDNSGetHostsResult, error) {
	requestInfo := &ApiRequest{
		command: domainsDNSGetHosts,
		method:  "POST",
		params:  url.Values{},
	}
	requestInfo.params.Set("SLD", sld)
	requestInfo.params.Set("TLD", tld)

	resp, err := client.do(requestInfo)
	if err != nil {
		return nil, err
	}

	return resp.DomainDNSHosts, nil
}

func (client *Client) DomainDNSGetList(sld, tld string) (*DomainDNSGetListResult, error) {
	requestInfo := &ApiRequest{
		command: domainsDNSGetList,
		method:  "POST",
		params:  url.Values{},
	}
	requestInfo.params.Set("SLD", sld)
	requestInfo.params.Set("TLD", tld)

	resp, err := client.do(requestInfo)
	if err != nil {
		return nil, err
	}

	return resp.DomainDNSNameservers, nil
}

func (client *Client) DomainDNSSetDefault(
	sld, tld string, hosts []DomainDNSHost,
) (*DomainDNSSetDefaultResult, error) {
	requestInfo := &ApiRequest{
		command: domainsDNSSetDefault,
		method:  "POST",
		params:  url.Values{},
	}
	requestInfo.params.Set("SLD", sld)
	requestInfo.params.Set("TLD", tld)

	resp, err := client.do(requestInfo)
	if err != nil {
		return nil, err
	}
	return resp.DomainDNSSetDefault, nil
}

func (client *Client) DomainDNSSetHosts(
	sld, tld string, hosts []DomainDNSHost,
) (*DomainDNSSetHostsResult, error) {
	requestInfo := &ApiRequest{
		command: domainsDNSSetHosts,
		method:  "POST",
		params:  url.Values{},
	}
	requestInfo.params.Set("SLD", sld)
	requestInfo.params.Set("TLD", tld)

	for i, h := range hosts {
		requestInfo.params.Set(fmt.Sprintf("HostName%v", i+1), h.Name)
		requestInfo.params.Set(fmt.Sprintf("RecordType%v", i+1), h.Type)
		requestInfo.params.Set(fmt.Sprintf("Address%v", i+1), h.Address)
		if h.Type == "MX" {
			requestInfo.params.Set(fmt.Sprintf("MXPref%v", i+1), strconv.Itoa(h.MXPref))
			requestInfo.params.Set("EmailType", "MX")
		}
		requestInfo.params.Set(fmt.Sprintf("TTL%v", i+1), strconv.Itoa(h.TTL))
	}

	resp, err := client.do(requestInfo)
	if err != nil {
		return nil, err
	}
	return resp.DomainDNSSetHosts, nil
}

type DomainDNSSetCustomResult struct {
	Domain string `xml:"Domain,attr"`
	Update bool   `xml:"Update,attr"`
}

func (client *Client) DomainDNSSetCustom(sld, tld, nameservers string) (*DomainDNSSetCustomResult, error) {
	requestInfo := &ApiRequest{
		command: domainsDNSSetCustom,
		method:  "POST",
		params:  url.Values{},
	}
	requestInfo.params.Set("SLD", sld)
	requestInfo.params.Set("TLD", tld)
	requestInfo.params.Set("Nameservers", nameservers)

	resp, err := client.do(requestInfo)
	if err != nil {
		return nil, err
	}
	return resp.DomainDNSSetCustom, nil
}
