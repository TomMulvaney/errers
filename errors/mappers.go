package errors

// StatusMapper ...
type StatusMapper interface {
	Do(status int) int
}

type statusMapper struct {
	do statusMapperFunc
}

// NewStatusMapper ...
func NewStatusMapper(do statusMapperFunc) StatusMapper {
	return &statusMapper{
		do: do,
	}
}

func (sm *statusMapper) Do(status int) int {
	return sm.do(status)
}

type statusMapperFunc func(status int) int

// ToHTTPStatus ...
func ToHTTPStatus(status int) int {
	return StatusUnknown
}

// FromHTTPStatus ...
func FromHTTPStatus(status int) int {
	return StatusUnknown
}

// ToErrerAPIStatus is for converting errer statuses for internal use to errer statuses for API clients
// For example, convert Unreachable to Internal because the client doesn't care that the server couldn't reach the other server
// Should this be done in the handlers?
func ToErrerAPIStatus(status int) int {
	return StatusUnknown
}
