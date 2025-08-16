package status

import "fmt"

type appsta struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (a appsta) Error() string {
	return fmt.Sprintf("[%d]%s", a.Code, a.Message)
}

func (a appsta) WithMessage(msg string) appsta {
	a.Message = msg
	return a
}
