package errors

// HandleError ...
func HandleError(err error, doers ...Doer) error {

	for _, do := range doers {
		err = do(err)
	}

	return err
}
