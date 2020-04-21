package msg

import (
	"github.com/XuJinTao1996/coredns-management/pkg/setting"
	"strings"
	"unsafe"
)

func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String2Path(p string) string {
	strlist := strings.Split(p, ".")
	for i, j := 0, len(strlist)-1; i < j; i, j = i+1, j-1 {
		strlist[i], strlist[j] = strlist[j], strlist[i]
	}
	return setting.RootKey + "/" + strings.Join(strlist, "/")
}

//func String2Record(p string) string {
//	strlist := strings.Split(p, ".")
//	for i, j := 0, len(strlist)-1; i < j; i, j = i+1, j-1 {
//		strlist[i], strlist[j] = strlist[j], strlist[i]
//	}
//	return "/" + strings.Join(strlist, "/")
//}

func Path2String(p string) string {
	strlist := strings.Split(p, "/")[2:]
	for i, j := 0, len(strlist)-1; i < j; i, j = i+1, j-1 {
		strlist[i], strlist[j] = strlist[j], strlist[i]
	}
	return strings.Join(strlist, ".")
}
