package lib

type Master struct {
	Status  string        `json:"status"`
	İnterval string      `json:"interval"`
	Error   interface{} `json:"errors"`
}


func JsonData(status string, interval string) *Master {
	d := &Master{}
	d.Status = status
	if status == "ok" {
		d.İnterval = interval
	} else {
		d.Error = interval
	}
	return d
}