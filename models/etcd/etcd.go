package etcd

import (
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
