package e

var MsgFlags = map[int]string{

	SUCCESS: "OK",
	ERROR:   "FAILED",

	PARAMETERS_ERR:           "Request parameters error",
	A_RECORD_ADD_SUCCESS:     "A record add success",
	CNAME_RECORD_ADD_SUCCESS: "CNAME record add success",
	PTR_RECORD_ADD_SUCCESS:   "PTR record add success",
	SRV_RECORD_ADD_SUCCESS:   "SRV record add success",

	A_RECORD_ADD_FAIL:     "A record add FAIL",
	CNAME_RECORD_ADD_FAIL: "CNAME record add FAIL",
	PTR_RECORD_ADD_FAIL:   "PTR record add FAIL",
	SRV_RECORD_ADD_FAIL:   "SRV record add FAIL",

	RECORD_DELETE_SUCCESS: "Record delete success",
	RECORD_DELETE_FAILED:  "Record delete FAIL",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
