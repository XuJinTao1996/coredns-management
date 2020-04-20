package e

var MsgFlags = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAILED",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
