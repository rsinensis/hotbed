package result

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	EXIST    = 501
	NOTEXIST = 502
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	EXIST:          "已存在",
	NOTEXIST:       "不存在",
}
