package v1

import (
	"context"
	"github.com/XuJinTao1996/coredns-management/models/etcd"
	"github.com/XuJinTao1996/coredns-management/pkg/app"
	"github.com/XuJinTao1996/coredns-management/pkg/e"
	"github.com/XuJinTao1996/coredns-management/pkg/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// get all dns records
func GetDnsRecords(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	etcdObj := etcd.ETCD{etcd.EtcdCli}
	resp, err := etcdObj.GET("/coredns", context.TODO())

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	data["list"] = resp
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// add a dns records
func AddDnsRecords(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	key := c.PostForm("key")
	value := c.PostForm("value")

	etcdObj := etcd.ETCD{etcd.EtcdCli}
	resp, err := etcdObj.PUT(msg.String2Path(key), value, context.TODO())

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	data["list"] = resp
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
