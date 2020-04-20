package v1

import (
	"context"
	"github.com/XuJinTao1996/coredns-management/models/etcd"
	"github.com/XuJinTao1996/coredns-management/pkg/e"
	"github.com/coreos/etcd/clientv3"
	"github.com/gin-gonic/gin"
)

// get all dns records
func GetDnsRecords(c *gin.Context) {
	data := make(map[string]interface{})
	code := e.SUCCESS

	etcdCli := etcd.EtcdCli
	defer etcdCli.Close()
	resp, err := etcdCli.Get(context.TODO(), "/coredns", clientv3.WithPrevKV())
	if err != nil {
		code = e.ERROR
	} else {
		data["lists"] = resp.Kvs
	}

	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
