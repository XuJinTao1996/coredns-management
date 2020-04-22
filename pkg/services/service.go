package services

import (
	"context"
	"encoding/json"
	"github.com/XuJinTao1996/coredns-management/models/etcd"
	"github.com/XuJinTao1996/coredns-management/pkg/e"
	"github.com/XuJinTao1996/coredns-management/pkg/msg"
)

type DNSRecord struct {
	Record      string `json:"record "form:"record" binding:"required"`
	Host        string `json:"host" form:"host" binding:"required"`
	Zone        string `json:"zone" form:"zone" binding:"required"`
	DNSType     string `json:"dns_type" form:"dns_type,omitempty"`
	TTL         int    `json:"ttl" form:"ttl,omitempty"`
	Weight      int    `json:"weight" form:"weight,omitempty"`
	Port        int    `json:"port" form:"port,omitempty"`
	Priority    int    `json:"priority" form:"priority,omitempty"`
	TargetStrip int    `json:"target_strip" form:"targetstrip,omitempty"`
	Group       string `json:"group" form:"group,omitempty"`
}

type DNSObj interface {
	Add() (interface{}, error, int)
	Update() (interface{}, error, int)
}

func Get(zone string) (interface{}, int, error) {
	newCLI := etcd.ETCD{etcd.EtcdCli}
	result, count, err := newCLI.Get(msg.String2Path(zone), context.TODO())
	return result, count, err
}

func Delete(zone, record string) (interface{}, error) {
	newCLI := etcd.ETCD{etcd.EtcdCli}
	result, err := newCLI.Delete((msg.String2Path(zone))+"/"+record, context.TODO())
	return result, err
}

func DeleteZone(zone string) (interface{}, error) {
	newCLI := etcd.ETCD{etcd.EtcdCli}
	result, err := newCLI.DeleteZone((msg.String2Path(zone)), context.TODO())
	return result, err
}

// TODO check dns records params
func (dr *DNSRecord) Add() (interface{}, error, int) {
	var (
		result  string
		err     error
		errCode int
	)
	tempMap := make(map[string]interface{})
	newCLI := etcd.ETCD{etcd.EtcdCli}
	if dr.DNSType == "A" {
		tempMap["host"] = dr.Host
		tempMap["ttl"] = dr.TTL
		formStr, err := json.Marshal(tempMap)
		if err != nil {
			errCode = e.A_RECORD_ADD_FAIL
		}
		result, err = newCLI.Put(msg.String2Path(dr.Zone)+"/"+dr.Record, msg.String(formStr), context.TODO())
	}
	if dr.DNSType == "CNAME" {
		tempMap["host"] = dr.Host
		tempMap["ttl"] = dr.TTL
		formStr, err := json.Marshal(tempMap)
		if err != nil {
			errCode = e.CNAME_RECORD_ADD_FAIL
		}
		result, err = newCLI.Put(msg.String2Path(dr.Zone)+"/"+dr.Record, msg.String(formStr), context.TODO())
	}
	if dr.DNSType == "SRV" && dr.Port != 0 {
		tempMap["host"] = dr.Host
		tempMap["ttl"] = dr.TTL
		tempMap["priority"] = dr.Priority
		tempMap["port"] = dr.Port
		formStr, err := json.Marshal(tempMap)
		if err != nil {
			errCode = e.SRV_RECORD_ADD_FAIL
		}
		result, err = newCLI.Put(msg.String2Path(dr.Zone)+"/"+dr.Record, msg.String(formStr), context.TODO())
	}
	if dr.DNSType == "PTR" && dr.TTL == 0 {
		tempMap["host"] = dr.Host
		formStr, err := json.Marshal(tempMap)
		if err != nil {
			errCode = e.PTR_RECORD_ADD_FAIL
		}
		result, err = newCLI.Put(msg.String2Path(dr.Zone)+"/"+dr.Record, msg.String(formStr), context.TODO())
	}
	return result, err, errCode
}
