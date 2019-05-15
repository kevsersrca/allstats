package lib

type Master struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

// Master isimli struct kullanılarak basit standart json nesleleri üretiminde kullanılır
func JsonData(status bool, message string, data ...interface{}) *Master {
	IsData := data != nil
	d := &Master{}
	d.Status = status
	if status {
		d.Message = message
	} else {
		d.Error = message
	}
	if IsData {
		d.Data = data[0]
	}
	return d
}