package response

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Sd      interface{} `json:"datasd"`
	Smp     interface{} `json:"datasmp"`
	Sma     interface{} `json:"datasma"`
	Smas    interface{} `json:"datasmas"`
	Meta    interface{} `json:"meta,omitempty"`
}
