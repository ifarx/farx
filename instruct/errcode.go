/**
 * @create  2019/12/28 17:40
 * @since 1.0.0.0
 * @author  ant
 */
package instruct

import "strconv"

type Errcode interface {
	Code() int32
	Message() string
	Ok() bool
	String() string
	Error() string
}

var ERRCODE_SUCC = NewErrcode(0, "succ")
var ERRCODE_FAIL = NewErrcode(-1, "fail")

type errCode struct {
	_c int32  /*code*/
	_m string /*message*/
}

func (e *errCode) Code() int32 {
	return e._c
}

func (e *errCode) Message() string {
	return e._m
}

func (e *errCode) Ok() bool {
	return e._c == 0
}

func (e *errCode) String() string {
	return "[code=" + strconv.FormatUint(uint64(e.Code()), 10) + ",message=" + e.Message() + "]"
}

func (e *errCode) Error() string {
	return e.String()
}

func NewErrcode(c int32, m string) Errcode {
	return &errCode{c, m}
}
