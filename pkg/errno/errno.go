package errno

import "fmt"

type Errno struct {
	Code int
	Message string
}

func(err *Errno) Error() string {
	return err.Message
}


type Err struct {
	Code int
	Message string
	Err error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}


func NewErr(errno *Errno, err error) *Err {
	return &Err{Code:errno.Code, Message:errno.Message, Err:err}
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	}
	return InternalServerError.Code, InternalServerError.Message
}

//func (err *Err) Add(message string) {
//	err.Message += " " + message
//}