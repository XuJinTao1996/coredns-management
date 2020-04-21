package services

import (
	"context"
	"encoding/json"
	"github.com/XuJinTao1996/coredns-management/models/etcd"
	"github.com/XuJinTao1996/coredns-management/pkg/msg"
	"log"
)

type DNSRecord struct {
	Record      string `json:"record"`
	Host        string `json:"host"`
	Zone        string `json:"zone"`
	DNSType     string `json:"dns_type,omitempty"`
	TTL         int    `json:"ttl,omitempty"`
	Weight      int    `json:"weight,omitempty"`
	Port        int    `json:"port,omitempty"`
	Priority    int    `json:"priority,omitempty"`
	TargetStrip int    `json:"targetstrip,omitempty"`
	Group       string `json:"group,omitempty"`
}

type DNSObj interface {
	Add() (interface{}, error)
	Update() (interface{}, error)
}

func Get(zone string) (interface{}, int, error) {
	newCLI := etcd.ETCD{etcd.EtcdCli}
	result, count, err := newCLI.Get(msg.Path2String(zone), context.TODO())
	return result, count, err
}

func Delete(zone, record string) (interface{}, error) {
	newCLI := etcd.ETCD{etcd.EtcdCli}
	result, err := newCLI.Delete((msg.String2Path(zone))+msg.String2Record(record), context.TODO())
	return result, err
}

func DeleteZone(zone string) (interface{}, error) {
	newCLI := etcd.ETCD{etcd.EtcdCli}
	result, err := newCLI.DeleteZone((msg.String2Path(zone)), context.TODO())
	return result, err
}

// TODO check dns records params
func (dr *DNSRecord) Add() (interface{}, error) {
	var (
		result string
		err    error
	)
	tempMap := make(map[string]interface{})
	newCLI := etcd.ETCD{etcd.EtcdCli}
	if dr.DNSType == "A" {
		tempMap["host"] = dr.Host
		tempMap["ttl"] = dr.TTL
		jsonStr, err := json.Marshal(tempMap)
		if err != nil {
			log.Fatalf("json parse error")
		}
		result, err = newCLI.Put(msg.Path2String(dr.Zone)+msg.String2Record(dr.Record), msg.String(jsonStr), context.TODO())
	}
	if dr.DNSType == "SRV" && dr.Port != 0 {
		tempMap["host"] = dr.Host
		tempMap["ttl"] = dr.TTL
		tempMap["priority"] = dr.Priority
		tempMap["port"] = dr.Port
		jsonStr, err := json.Marshal(tempMap)
		if err != nil {
			log.Fatalf("json parse error")
		}
		result, err = newCLI.Put(msg.String2Path(dr.Zone)+msg.String2Record(dr.Record), msg.String(jsonStr), context.TODO())
	}
	if dr.DNSType == "PTR" && dr.TTL == 0 {
		tempMap["host"] = dr.Host
		jsonStr, err := json.Marshal(tempMap)
		if err != nil {
			log.Fatalf("json parse error")
		}
		result, err = newCLI.Put(msg.Path2String(dr.Zone)+msg.String2Record(dr.Record), msg.String(jsonStr), context.TODO())
	}
	return result, err
}
