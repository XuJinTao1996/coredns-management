package routers

import (
	"github.com/XuJinTao1996/coredns-management/pkg/setting"
	v1 "github.com/XuJinTao1996/coredns-management/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/records", v1.GetDnsRecords)
	}
	return r
}
