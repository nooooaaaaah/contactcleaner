package vcard

import "fmt"

type err struct {
	message string
}

func (e *err) Error(val string) error {
	return fmt.Errorf(e.message, val)
}

var (
	ErrLangTag          = &err{"Invalid language tag"}
	ErrParam            = &err{"Invalid parameter"}
	ErrNilAltId         = &err{"Missing required parameter"}
	ErrParamSyntax      = &err{"Parameter syntax error"}
	ErrValParamValue    = &err{"Invalid value-type for Value tag"}
	ErrPrefParam        = &err{"Invalid preference value"}
	ErrNilPref          = &err{"Missing required preference value"}
	ErrInvalidPid       = &err{"Invalid property id"}
	ErrNilType          = &err{"Missing required type"}
	ErrInvalidType      = &err{"Invalid type"}
	ErrNilMediaType     = &err{"Missing required media type value"}
	ErrInvalidMediaType = &err{"Invalid media type value"}
	ErrNilCalscale      = &err{"Missing required calscale value"}
	ErrInvalidCalscale  = &err{"Invalid calscale value"}
	ErrNilSortAs        = &err{"Missing required sort-as value"}
	ErrInvalidSortAs    = &err{"Invalid sort-as value"}
	ErrNilGeo           = &err{"Missing required geo value"}
	ErrInvalidGeo       = &err{"Invalid geo value"}
)
