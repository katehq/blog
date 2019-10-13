package serializer

// Response the basic serializer
type Response struct {
	Status int `json:"status"`
	Msg   string    `json:"msg"`
}

