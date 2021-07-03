package device

var errors []error

func pushError(err error) {
	errors = append(errors, err)
}

func PopError() error {
	err := errors[0]
	errors[0] = nil
	errors = errors[1:]

	return err
}
