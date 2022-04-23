package errmsg

const (
	SUCCES = 200
	ERROR =500
)

var codeMsg = map[int]string{
	SUCCES: 		"OK",
	ERROR:			"FAIL",
}

func GetErrMsg(code int) string{
	return codeMsg[code]
}