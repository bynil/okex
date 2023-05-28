package responses

import "fmt"

type (
	Basic struct {
		Code int    `json:"code,string"`
		Msg  string `json:"msg,omitempty"`
	}
)

func (b Basic) Validate() error {
	if b.Code != 0 {
		return fmt.Errorf("code: %d, msg: %s", b.Code, b.Msg)
	}
	return nil
}
