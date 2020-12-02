package e

var msgFlags = map[int]string{
	SUCCESS:           "ok",
	NOTFOUND:          "paste not found",
	EXPIRED:           "paste expired",
	ReadRequestError:  "read request body error",
	ParseRequestError: "parse request body error",
	CreateError:       "fail to insert data into db",
	NoShortLinkError:  "no short link specified",
	ERROR:             "no such msg",
}

func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}
	return msgFlags[ERROR]
}
