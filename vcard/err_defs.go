package vcard

import "fmt"

type err struct {
	message string
}

func (e *err) Error(val string) error {
	return fmt.Errorf(e.message, val)
}

var (
	ErrLangTag            = &err{"Invalid language tag"}
	ErrParam              = &err{"Invalid parameter"}
	ErrNilAltId           = &err{"Missing required parameter"}
	ErrParamSyntax        = &err{"Parameter syntax error"}
	ErrValParamValue      = &err{"Invalid value-type for Value tag"}
	ErrPrefParam          = &err{"Invalid preference value"}
	ErrNilPref            = &err{"Missing required preference value"}
	ErrInvalidPid         = &err{"Invalid property id"}
	ErrNilType            = &err{"Missing required type"}
	ErrInvalidType        = &err{"Invalid type"}
	ErrNilMediaType       = &err{"Missing required media type value"}
	ErrInvalidMediaType   = &err{"Invalid media type value"}
	ErrNilCalscale        = &err{"Missing required calscale value"}
	ErrInvalidCalscale    = &err{"Invalid calscale value"}
	ErrNilSortAs          = &err{"Missing required sort-as value"}
	ErrInvalidSortAs      = &err{"Invalid sort-as value"}
	ErrNilGeo             = &err{"Missing required geo value"}
	ErrInvalidGeo         = &err{"Invalid geo value"}
	ErrNilTz              = &err{"Missing required timezone value"}
	ErrInvalidTz          = &err{"Invalid timezone value"}
	ErrNilIndex           = &err{"Missing required index value"}
	ErrInvalidIndex       = &err{"Invalid index value"}
	ErrNilLevel           = &err{"Missing required level value"}
	ErrInvalidLevel       = &err{"Invalid level value"}
	ErrInvalidGroup       = &err{"Invalid group value"}
	ErrNilGroup           = &err{"Missing required group value"}
	ErrNilCc              = &err{"Missing required country code"}
	ErrInvalidCc          = &err{"Invalid country code"}
	ErrNilAuthor          = &err{"Missing required author value"}
	ErrInvalidAuthor      = &err{"Invalid author value"}
	ErrNilAuthorName      = &err{"Missing required author name"}
	ErrInvalidAuthorName  = &err{"Invalid author name"}
	ErrNilCreated         = &err{"Missing required created value"}
	ErrInvalidCreated     = &err{"Invalid created value"}
	ErrNilDerived         = &err{"Missing required derived value"}
	ErrInvalidDerived     = &err{"Invalid derived value"}
	ErrNilLabel           = &err{"Missing required label value"}
	ErrInvalidLabel       = &err{"Invalid label value"}
	ErrNilPhonetic        = &err{"Missing required phonetic value"}
	ErrInvalidPhonetic    = &err{"Invalid phonetic value"}
	ErrNilPropID          = &err{"Missing required property id"}
	ErrInvalidPropID      = &err{"Invalid property id"}
	ErrNilScript          = &err{"Missing required script value"}
	ErrInvalidScript      = &err{"Invalid script value"}
	ErrNilServiceType     = &err{"Missing required service type value"}
	ErrNilUsername        = &err{"Missing required username value"}
	ErrInvalidUsername    = &err{"Invalid username value"}
	ErrInvalidServiceType = &err{"Invalid service type value"}
	ErrNilJsptr           = &err{"Missing required json-pointer value"}
	ErrInvalidJsptr       = &err{"Invalid json-pointer value"}
)
