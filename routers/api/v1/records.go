package v1

import (
	"github.com/XuJinTao1996/coredns-management/pkg/app"
	"github.com/XuJinTao1996/coredns-management/pkg/e"
	"github.com/XuJinTao1996/coredns-management/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// get all dns records
func GetDnsRecords(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})
	zone := c.Query("zone")
	resp, count, err := services.Get(zone)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	data["list"] = resp
	data["count"] = count
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// add a dns records
func AddDnsRecord(c *gin.Context) {
	var dnsRecord services.DNSRecord

	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	err := c.ShouldBindJSON(&dnsRecord)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	resp, err := dnsRecord.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	data["list"] = resp
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// add a dns records
func DeleteDnsRecords(c *gin.Context) {
	var (
		resp interface{}
		err  error
	)
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	zone := c.Query("zone")
	record := c.Query("record")

	if zone == "" && record == "" {
		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	if zone != "" && record != "" {
		resp, err = services.Delete(zone, record)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, data)
			return
		}
	}

	if zone != "" {
		resp, err = services.DeleteZone(zone)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, data)
			return
		}
	}

	data["list"] = resp
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
