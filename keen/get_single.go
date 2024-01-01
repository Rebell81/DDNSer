package keen

import (
	"encoding/json"
	"errors"
)

type IFace []struct {
	ID            string   `json:"id"`
	Index         int      `json:"index"`
	InterfaceName string   `json:"interface-name"`
	Type          string   `json:"type"`
	Description   string   `json:"description"`
	Traits        []string `json:"traits"`
	Link          string   `json:"link"`
	Connected     string   `json:"connected"`
	State         string   `json:"state"`
	Role          []string `json:"role"`
	Mtu           int      `json:"mtu"`
	TxQueueLength int      `json:"tx-queue-length"`
	Address       string   `json:"address"`
	Mask          string   `json:"mask"`
	Global        bool     `json:"global"`
	Defaultgw     bool     `json:"defaultgw"`
	Priority      int      `json:"priority"`
	SecurityLevel string   `json:"security-level"`
	Usedby        []string `json:"usedby"`
	Ipv6          struct {
		Addresses []struct {
			Address       string `json:"address"`
			PrefixLength  int    `json:"prefix-length"`
			Proto         string `json:"proto"`
			ValidLifetime string `json:"valid-lifetime"`
		} `json:"addresses"`
		Defaultgw bool `json:"defaultgw"`
	} `json:"ipv6"`
	AuthType   string `json:"auth-type"`
	Uptime     int    `json:"uptime"`
	Remote     string `json:"remote"`
	Fail       string `json:"fail"`
	Via        string `json:"via"`
	LastChange string `json:"last-change"`
	SessionID  int    `json:"session-id"`
	AcMac      string `json:"ac-mac"`
	Summary    struct {
		Layer struct {
			Conf string `json:"conf"`
			Link string `json:"link"`
			Ipv4 string `json:"ipv4"`
			Ipv6 string `json:"ipv6"`
			Ctrl string `json:"ctrl"`
		} `json:"layer"`
	} `json:"summary"`
}

func (k *Keenetic) GetSingle(filter string) (IFace, error) {

	var req []map[string]string

	req = append(req, map[string]string{
		"name": filter,
	})

	r, err := k.c.R().SetBody(req).Post("/rci/show/interface/")
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, errors.New(r.String())
	}

	data := IFace{}

	err = json.Unmarshal(r.Body(), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
