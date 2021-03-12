package ffmpeg_go
const (
 TYPE_STREAMS = "streams"
 TYPE_FORMAT = "format"

)

type Config struct {
	CommandPath map[string]string `json:"command_path"`
	TimeOut int    `json:"time_out"`
}

//获取文件信息参数
func GetInfoParams()[]string  {
	return []string{"-v","quiet","-print_format","json","-show_format"}
}

//获取合并文件信息参数
func GetConcatParams()[]string  {
	return []string{"-v","quiet","-print_format","json","-show_format"}
}