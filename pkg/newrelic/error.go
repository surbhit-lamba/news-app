package newrelic

import "github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"

type stackTracer interface {
	StackTrace() []uintptr
}

type errorClasser interface {
	ErrorClass() string
}

type errorAttributer interface {
	ErrorAttributes() map[string]interface{}
}

type nrpkgerrorsWrappedError interface {
	error
	stackTracer
	errorClasser
}

type fullyEnrichedNewrelicError interface {
	error
	stackTracer
	errorClasser
	errorAttributer
}

type errWithAttributes struct {
	nrpkgerrorsWrappedError
	origErr error
}

var _ fullyEnrichedNewrelicError = errWithAttributes{}

func (e errWithAttributes) ErrorAttributes() map[string]interface{} {
	if err, ok := e.origErr.(errorAttributer); ok {
		return err.ErrorAttributes()
	}
	return map[string]interface{}{}
}

// Wrap is an extension to the nrpkgerrors.Wrap that ensures
// that if the original error implements ErrorAttributer interface
// it is preserved even after wrapping
func Wrap(err error) error {
	wrappedErr := nrpkgerrors.Wrap(err)
	if wrappedErr, ok := wrappedErr.(nrpkgerrorsWrappedError); ok {
		// if the final wrapped error implements all the necessary methods
		// then continue to further wrap it with the attributes
		return errWithAttributes{
			nrpkgerrorsWrappedError: wrappedErr,
			origErr:                 err,
		}
	}
	return err
}
