package hypnos

// HandleHypnosError ...
func HandleHypnosError(err error) error {
	if IsHypnosError(err) {
		// TODO

		// e := err.(HypnosError)
		// err = errors.WrapStatus(e, hypnosToHTTPStatus(e.Status()))
	}

	return err
}
