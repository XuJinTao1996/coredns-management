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

// Add a dns records
func AddDnsRecord(c *gin.Context) {
	var dnsRecord services.DNSRecord

	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	err := c.ShouldBind(&dnsRecord)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.PARAMETERS_ERR, data)
		return
	}

	resp, err, errCode := dnsRecord.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errCode, data)
		return
	}

	data["list"] = resp
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// Delete a dns records
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
		appG.Response(http.StatusInternalServerError, e.PARAMETERS_ERR, data)
		return
	}

	if record == "" {
		resp, err = services.DeleteZone(zone)
	} else {
		resp, err = services.Delete(zone, record)
	}
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.RECORD_DELETE_FAILED, data)
		return
	}

	data["list"] = resp
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
