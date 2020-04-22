package etcd

import (
	"context"
	"github.com/XuJinTao1996/coredns-management/pkg/msg"
	"github.com/XuJinTao1996/coredns-management/pkg/setting"
	"github.com/coreos/etcd/clientv3"
	"log"
)

var (
	EtcdCli *clientv3.Client
)

func init() {
	EtcdCli = etcdConnection()
}

type ETCD struct {
	Cli *clientv3.Client
}

type EtcdOpt interface {
	Get(key string, ctx *context.Context) (interface{}, error)
	Put(key, value string, ctx *context.Context) (interface{}, error)
	Delete(key string, ctx *context.Context) (interface{}, error)
	DeleteZone(key string, ctx *context.Context) (interface{}, error)
}

func etcdConnection() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{setting.Endpoints},
		DialTimeout: 0,
		//DialKeepAliveTime:    0,
		//DialKeepAliveTimeout: 0,
		//MaxCallSendMsgSize:   0,
		//MaxCallRecvMsgSize:   0,
		//TLS:                  nil,
		//Username:             "",
		//Password:             "",
		//RejectOldCluster:     false,
		//DialOptions:          nil,
		//LogConfig:            nil,
		//Context:              nil,
		//PermitWithoutStream:  false,
	})
	if err != nil {
		log.Fatalf("connect failed, err: %v", err)
	}
	return cli
}

func (e *ETCD) Get(key string, ctx context.Context) ([]interface{}, int, error) {
	var result []interface{}

	resp, err := e.Cli.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		log.Fatalf("get etcd key value error")
		return result, 0, err
	}

	for _, v := range resp.Kvs {
		item := make(map[string]interface{})
		item["key"] = msg.Path2String(msg.String(v.Key))
		item["value"] = msg.String(v.Value)
		item["create_revision"] = v.CreateRevision
		item["mod_revision"] = v.ModRevision
		item["version"] = v.Version
		result = append(result, item)
	}
	return result, int(resp.Count), err
}

func (e *ETCD) Put(key, value string, ctx context.Context) (string, error) {
	_, err := e.Cli.Put(ctx, key, value, clientv3.WithPrevKV())
	if err != nil {
		log.Fatalf("error put key value")
		return "", err
	}
	return "Put Success", err
}

func (e *ETCD) Delete(key string, ctx context.Context) (string, error) {
	_, err := e.Cli.Delete(ctx, key)
	if err != nil {
		log.Fatalf("error delete key value")
		return "", err
	}
	return "Record Deleted", err
}

func (e *ETCD) DeleteZone(key string, ctx context.Context) (string, error) {
	_, err := e.Cli.Delete(ctx, key, clientv3.WithPrefix())
	if err != nil {
		log.Fatalf("error delete key value")
		return "", err
	}
	return "Zone Deleted", err
}
