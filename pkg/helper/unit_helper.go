package helper

var m map[string]int64
func init() {
	m = make(map[string]int64, 3)
	m["second"] = 1
	m["minute"] = 60
	m["hour"] = 3600
}

func GetValueByKey(key string) int64 {
	return m[key]
}
