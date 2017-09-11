package errors

import (
	"fmt"

	baseErrors = "errors"

	pkgErrors "github.com/pkg/errors"
)

type ErrorStatus int

const (
	delim = ";"

	StatusBadReq ErrorStatus = iota
	StatusInternal
	StatusUnreachable
)

// New wraps standard library errors
func New(msg string) error {
	return pkgErrors.New(msg)
}

// Wrap wraps pkg error
func Wrap(err error, message string) error {
	return pkgErrors.Wrap(err, message)
}

// IError ...
type IError interface {
	error
	Status() ErrorStatus
	isIError() bool
}

type errorImp struct {
	error
	status ErrorStatus
}

func (e *errorImp) Error() string {
	return e.Error()
}

func (e *errorImp) Status() ErrorStatus {
	return e.status
}

func (e *errorImp) isIError() bool {
	return true
}

// IsIError ...
func IsIError(err error) bool {
	e, ok := err.(IError)
	return ok && e.isIError()
}

// NewStatus ...
func NewStatus(status ErrorStatus, err error, messages ...string) error {
	for _, message := range messages {
		err = pkgErrors.Wrap(err, message+delim)
	}

	return &errorImp{
		error:  err,
		status: status,
	}
}

// NewStatusTest ...
func NewStatusTest(status ErrorStatus, err error, message string) error {
	fmt.Println("NewStatusTest is wrapping")
	err = pkgErrors.Wrap(err, message)
	fmt.Println("NewStatusTest has wrapped")

	return &errorImp{
		error:  err,
		status: status,
	}
}

// BadReq constructor
func BadReq(err error, messages ...string) error {
	return NewStatus(StatusBadReq, err, messages...)
}

// Internal constructor
func Internal(err error, messages ...string) error {
	return NewStatus(StatusInternal, err, messages...)
}

// Unreachable constructor
func Unreachable(err error, messages ...string) error {
	return NewStatus(StatusUnreachable, err, messages...)
}

// Test ...
func Test() {
	fmt.Println("Running Test")
	err := pkgErrors.New("Original Error;")
	err = pkgErrors.Wrap(err, "Wrapping Error;")
	err = pkgErrors.Wrap(err, "Double Wrapping Error;")
	fmt.Println("Error: ", err)
}

// Test2 ...
func Test2() {
	fmt.Println("Running Test2")
	err := New("Original Error")
	err = Wrap(err, "Wrapping Error")
	err = Wrap(err, "Double Wrapping Error")
	fmt.Println("Error: ", err)
}

// Test3 ...
func Test3() {
	fmt.Println("Running Test3")

	wrapMessages := []string{"Wrapping Error", "Double Wrapping Error", "Triple Wrapping Error"}

	err := New("Original Error")

	for _, msg := range wrapMessages {
		err = Wrap(err, msg)
	}

	fmt.Println("Error: ", err)
}

// Test4 ...
func Test4() {
	fmt.Println("Running Test4")

	err := New("Original Error")

	err = NewStatusTest(StatusBadReq, err, "Wrap Error")

	fmt.Println("Error: ", err)
}

// Test5 ...
func Test5() {
	fmt.Println("Running Test5")

	err := pkgErrors.New("Hello")

	err = &errorImp{
		error:  err,
		status: StatusBadReq,
	}

	fmt.Println(err)
}

// // Test6 ...
// func Test6() {
// 	fmt.Println("Running Test6")

// 	err := baseErrors.New("Hello")

// 	err = &errorImp{
// 		error:  err,
// 		status: StatusBadReq,
// 	}

// 	fmt.Println(err)
// }
