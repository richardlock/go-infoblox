package infoblox

import "fmt"

func (c *Client) RecordA() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "record:a",
	}
}

type RecordAObject struct {
	Object
	Ipv4Addr HostIpv4Addr `json:"ipv4addr,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Client) RecordAObject(ref string) *RecordAObject {
	a := RecordAOBject{}
	a.Object = Object {
		Ref: ref,
		r: c.RecordA(),`a
	}
	return &a
}

func (c *Client) GetRecordA(ref string) (*RecordAObject, error) {
	resp, err := c.RecordAObject(ref).get(nil)
	if err != nil {
		return nil, fmt.Errorf("Could not get created host record: %s", err)
	}
	var out RecordAObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) FindRecordA(name string) ([]RecordAObject, error) {
	field := "name"
	conditions := []Condition{Condition{Field:&field, Value:name}}
	resp, err := c.RecordA().find(conditions, nil)
	if err != nil {
		return nil, err
	}

	var out []RecordAObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
} 
