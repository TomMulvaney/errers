package errors

// StatusMapper ...
type StatusMapper interface {
	Do(status int) int
}

type statusMapper struct {
	Do statusMapperFunc
}

// NewStatusMapper ...
func NewStatusMapper(do statusMapperFunc) *statusMapper {
	return &statusMapper{
		Do: do,
	}
}

type statusMapperFunc func(status int) int

// ErrerToHTTPStatus ...
func ErrerToHTTPStatus(status int) int {
	return StatusUnknown
}

// HTTPToErrerStatus ...
func HTTPToErrerStatus(status int) int {
	return StatusUnknown
}

// ErrerToGRPCStatus ...
func ErrerToGRPCStatus(status int) int { // Calling function needs to convert to uint32
	return StatusUnknown
}

// GRPCToErrerStatus ...
func GRPCToErrerStatus(status int) int {
	return StatusUnknown
}

// ErrerToErrerAPIStatus is for converting errer statuses for internal use to errer statuses for API clients
// For example, convert Unreachable to Internal because the client doesn't care that the server couldn't reach the other server
// Should this be done in the handlers?
func ErrerToErrerAPIStatus(status int) int {
	return StatusUnknown
}
