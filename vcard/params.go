package vcard

import (
	"regexp"
	"strings"
)

const (
	LANGUAGE_PARAM  ParamName = "LANGUAGE"
	VALUE_PARAM     ParamName = "VALUE"
	PREF_PARAM      ParamName = "PREF"
	ALTID_PARAM     ParamName = "ALTID"
	PID_PARAM       ParamName = "PID"
	TYPE_PARAM      ParamName = "TYPE"
	MEDIATYPE_PARAM ParamName = "MEDIATYPE"
	CALSCALE_PARAM  ParamName = "CALSCALE"
	SORTAS_PARAM    ParamName = "SORT-AS"
	GEO_PARAM       ParamName = "GEO"
	TZ_PARAM        ParamName = "TZ"
)
const (
	WORK         ParamVal = "work"
	HOME         ParamVal = "home"
	VOICE        ParamVal = "voice"
	FAX          ParamVal = "fax"
	CELL         ParamVal = "cell"
	VIDEO        ParamVal = "video"
	PAGER        ParamVal = "pager"
	TEXT         ParamVal = "text"
	TEXTPHONE    ParamVal = "textphone"
	GREGORIAN    ParamVal = "gregorian"
	CONTACT      ParamVal = "contact"
	ACQUAINTANCE ParamVal = "acquaintance"
	FRIEND       ParamVal = "friend"
	MET          ParamVal = "met"
	COWORKER     ParamVal = "co-worker"
	COLLEAGUE    ParamVal = "colleague"
	CORESIDENT   ParamVal = "co-resident"
	NEIGHBOR     ParamVal = "neighbor"
	CHILD        ParamVal = "child"
	PARENT       ParamVal = "parent"
	SIBLING      ParamVal = "sibling"
	SPOUSE       ParamVal = "spouse"
	KIN          ParamVal = "kin"
	MUSE         ParamVal = "muse"
	CRUSH        ParamVal = "crush"
	DATE         ParamVal = "date"
	SWEETHEART   ParamVal = "sweetheart"
	ME           ParamVal = "me"
	AGENT        ParamVal = "agent"
	EMERGENCY    ParamVal = "emergency"
)

var TYPEGROUP = map[ParamVal]map[PropName]bool{
	WORK: {
		FN:         true,
		NICKNAME:   true,
		PHOTO:      true,
		ADR:        true,
		TEL:        true,
		EMAIL:      true,
		IMPP:       true,
		LANG:       true,
		TZ:         true,
		GEO:        true,
		TITLE:      true,
		ROLE:       true,
		LOGO:       true,
		ORG:        true,
		RELATED:    true,
		CATEGORIES: true,
		NOTE:       true,
		SOUND:      true,
		URL:        true,
		KEY:        true,
		FBURL:      true,
		CALADRURI:  true,
		CALURI:     true,
	},
	HOME: {
		FN:         true,
		NICKNAME:   true,
		PHOTO:      true,
		ADR:        true,
		TEL:        true,
		EMAIL:      true,
		IMPP:       true,
		LANG:       true,
		TZ:         true,
		GEO:        true,
		TITLE:      true,
		ROLE:       true,
		LOGO:       true,
		ORG:        true,
		RELATED:    true,
		CATEGORIES: true,
		NOTE:       true,
		SOUND:      true,
		URL:        true,
		KEY:        true,
		FBURL:      true,
		CALADRURI:  true,
		CALURI:     true,
	},
	VOICE: {
		TEL: true,
	},
	FAX: {
		TEL: true,
	},
	CELL: {
		TEL: true,
	},
	VIDEO: {
		TEL: true,
	},
	PAGER: {
		TEL: true,
	},
	TEXT: {
		TEL: true,
	},
	TEXTPHONE: {
		TEL: true,
	},
	GREGORIAN: {
		BDAY:        true,
		ANNIVERSARY: true,
	},
	CONTACT: {
		RELATED: true,
	},
	ACQUAINTANCE: {
		RELATED: true,
	},
	FRIEND: {
		RELATED: true,
	},
	MET: {
		RELATED: true,
	},
	COWORKER: {
		RELATED: true,
	},
	COLLEAGUE: {
		RELATED: true,
	},
	CORESIDENT: {
		RELATED: true,
	},
	NEIGHBOR: {
		RELATED: true,
	},
	CHILD: {
		RELATED: true,
	},
	PARENT: {
		RELATED: true,
	},
	SIBLING: {
		RELATED: true,
	},
	SPOUSE: {
		RELATED: true,
	},
	KIN: {
		RELATED: true,
	},
	MUSE: {
		RELATED: true,
	},
	CRUSH: {
		RELATED: true,
	},
	DATE: {
		RELATED: true,
	},
	SWEETHEART: {
		RELATED: true,
	},
	ME: {
		RELATED: true,
	},
	AGENT: {
		RELATED: true,
	},
	EMERGENCY: {
		RELATED: true,
	},
}

type Param interface {
	GetName() ParamName
	GetVal() []string
	validate() error
}

type ParamName string
type ParamVal string

// represents a vCard parameter
// https://tools.ietf.org/html/rfc6350#section-5.3
type BaseParam struct {
	Name ParamName
	Val  []string
}

func (bp *BaseParam) GetName() ParamName {
	return bp.Name
}

// returns a slice of strings representing the value of the parameter
func (bp *BaseParam) GetVal() []string {
	return bp.Val
}

type LanguageParam struct {
	BaseParam
}

// Creates a new LanguageParam based on the value of the
// found with the provided language tag.
// If the language tag is invalid, an error is returned.
// The language tag must conform to the RFC 5646 standard.
// https://tools.ietf.org/html/rfc5646
// https://tools.ietf.org/html/rfc6350#section-5.1
func NewLanguageParam(languageTag string) (*LanguageParam, error) {
	lp := &LanguageParam{
		BaseParam: BaseParam{
			Name: LANGUAGE_PARAM,
			Val:  []string{languageTag},
		},
	}
	if err := lp.validate(); err != nil {
		return nil, err
	}
	return lp, nil
}

func (lp *LanguageParam) GetLanguageTag() string {
	if len(lp.Val) > 0 {
		return lp.Val[0]
	}
	return ""
}

// Validates if the provided language tag conforms to the RFC 5646 standard.
func (lp *LanguageParam) validate() error {
	ltag := lp.GetLanguageTag()
	languageTagRegex := regexp.MustCompile(`^[a-zA-Z]{2,3}(-[a-zA-Z]{3})?$`)
	if languageTagRegex.MatchString(ltag) {
		return nil
	}
	return ErrLangTag.Error(ltag)
}

type ValueParam struct {
	BaseParam
}

// Creates a new ValueParam based on the value of the
// found with the provided value.
// If the value is invalid, an error is returned.
// The value must be either "text", "uri", "date-time", "date", "time" or "date-and-or-time".
// https://tools.ietf.org/html/rfc6350#section-5.2
func NewValueParam(valueType string) (*ValueParam, error) {
	vp := &ValueParam{
		BaseParam: BaseParam{
			Name: VALUE_PARAM,
			Val:  []string{valueType},
		},
	}
	if err := vp.validate(); err != nil {
		return nil, err
	}
	return vp, nil
}

// Validates if the provided value type is a valid value
// based on https://tools.ietf.org/html/rfc6350#section-5.2
func (vp *ValueParam) validate() error {
	valueType := strings.ToLower(vp.GetVpValue())
	switch valueType {
	case "text", "uri", "date-time", "date", "time",
		"date-and-or-time", "timestamp", "boolean", "integer",
		"float", "utc-offset", "language-tag", "unknown":
		return nil
	default:
		return ErrValParamValue.Error(valueType)
	}
}

func (vp *ValueParam) GetVpValue() string {
	if len(vp.Val) > 0 {
		return vp.Val[0]
	}
	return ""
}

type PrefParam struct {
	BaseParam
}

// Creates a new PrefParam based on the value of the
// found with the provided preference value.
// If the preference value is invalid, an error is returned.
// The preference value must be a number between 1 and 100.
// https://tools.ietf.org/html/rfc6350#section-5.3
func NewPrefParam(prefValue string) (*PrefParam, error) {
	pp := &PrefParam{
		BaseParam: BaseParam{
			Name: PREF_PARAM,
			Val:  []string{prefValue},
		},
	}
	if err := pp.validate(); err != nil {
		return nil, err
	}
	return pp, nil
}

// Returns the preference value of the parameter
func (pp *PrefParam) GetPrefValue() string {
	if len(pp.Val) > 0 {
		return pp.Val[0]
	}
	return ""
}

// Validates if the provided preference value is a valid value
// based on https://tools.ietf.org/html/rfc6350#section-5.3
func (pp *PrefParam) validate() error {
	prefValue := pp.GetPrefValue()
	prefValueRegex := regexp.MustCompile(`^[1-9][0-9]?$|^100$`)
	if prefValueRegex.MatchString(prefValue) {
		return nil
	}
	return ErrPrefParam.Error(prefValue)
}

type AltIdParam struct {
	BaseParam
}

// Creates a new AltIdParam based on the value of the
// found with the provided altId value.
// If the altId value is invalid, an error is returned.
// When used you must ensure there is a corresponding
// property with the same value.
// https://tools.ietf.org/html/rfc6350#section-5.4
func NewAltIdParam(altId string) (*AltIdParam, error) {
	ap := &AltIdParam{
		BaseParam: BaseParam{
			Name: ALTID_PARAM,
			Val:  []string{altId},
		},
	}
	if err := ap.validate(); err != nil {
		return nil, err
	}
	return ap, nil
}

func (ap *AltIdParam) validate() error {
	if len(ap.Val) == 0 {
		return ErrNilAltId.Error(string(ALTID_PARAM))
	}
	return nil
}

// Returns the altId value of the parameter
func (ap *AltIdParam) GetAltId() string {
	if len(ap.Val) > 0 {
		return ap.Val[0]
	}
	return ""
}

type PIDParam struct {
	BaseParam
}

// Creates a new PIDParam with the provided pid value.
// If the pid value is invalid, an error is returned.
// When used you must ensure there are multiple instances
// of the same property with the same value.
// https://tools.ietf.org/html/rfc6350#section-5.5
func NewPIDParam(pid string) (*PIDParam, error) {
	pip := &PIDParam{
		BaseParam: BaseParam{
			Name: PID_PARAM,
			Val:  []string{pid},
		},
	}
	if err := pip.validate(); err != nil {
		return nil, err
	}
	return pip, nil
}

func (pip *PIDParam) validate() error {
	pid := pip.GetPid()
	pidRegex := regexp.MustCompile(`^\d+$`)
	if pidRegex.MatchString(pid) {
		return nil
	}
	return ErrInvalidPid.Error(pid)
}

// Returns the pid value of the parameter
func (pip *PIDParam) GetPid() string {
	if len(pip.Val) > 0 {
		return pip.Val[0]
	}
	return ""
}

type TypeParam struct {
	BaseParam
}

// Creates a new TypeParam with the provided type value.
// If the type value is invalid, an error is returned.
// https://tools.ietf.org/html/rfc6350#section-5.6
func NewTypeParam(typeValue string, propName PropName) (*TypeParam, error) {
	tp := &TypeParam{
		BaseParam: BaseParam{
			Name: TYPE_PARAM,
			Val:  []string{typeValue},
		},
	}
	if err := tp.validate(propName); err != nil {
		return nil, err
	}
	return tp, nil
}

func (tp *TypeParam) validate(propName PropName) error {
	if len(tp.Val) == 0 {
		return ErrNilType.Error(string(TYPE_PARAM))
	}

	// Function to validate the TYPE parameter against the allowed values for a property
	validateTypeParam := func(property PropName, paramValue ParamVal) bool {
		if props, ok := TYPEGROUP[paramValue]; ok {
			return props[property]
		}
		return false
	}
	if !validateTypeParam(propName, ParamVal(tp.Val[0])) {
		return ErrInvalidType.Error(tp.Val[0])
	}
	return nil
}

type MediaTypeParam struct {
	BaseParam
}

// Creates a new MediatypeParam with the provided mediatype value.
// If the mediatype value is invalid, an error is returned.
// https://tools.ietf.org/html/rfc6350#section-5.7
func NewMediatypeParam(mediatype string) (*MediaTypeParam, error) {
	mtp := &MediaTypeParam{
		BaseParam: BaseParam{
			Name: MEDIATYPE_PARAM,
			Val:  []string{mediatype},
		},
	}
	if err := mtp.validate(); err != nil {
		return nil, err
	}
	return mtp, nil
}

func (mtp *MediaTypeParam) validate() error {
	if len(mtp.Val) == 0 {
		return ErrNilMediaType.Error(string(MEDIATYPE_PARAM))
	}
	var mediaTypeRegex = regexp.MustCompile(`^[a-zA-Z0-9!#$&-^_]+/[a-zA-Z0-9!#$&-^_]+(;[a-zA-Z0-9!#$&-^_]+=[a-zA-Z0-9!#$&-^_]+)*$`)
	if !mediaTypeRegex.MatchString(mtp.Val[0]) {
		return ErrInvalidMediaType.Error(mtp.Val[0])
	}
	return nil
}

type CalscaleParam struct {
	BaseParam
}

// Creates a new CalscaleParam with the provided calscale value.
// If the calscale value is invalid, an error is returned.
// https://tools.ietf.org/html/rfc6350#section-5.8
func NewCalscaleParam(calscale string) (*CalscaleParam, error) {
	cp := &CalscaleParam{
		BaseParam: BaseParam{
			Name: CALSCALE_PARAM,
			Val:  []string{calscale},
		},
	}
	if err := cp.validate(); err != nil {
		return nil, err
	}
	return cp, nil
}

func (cp *CalscaleParam) validate() error {
	if len(cp.Val) == 0 {
		return ErrNilCalscale.Error(string(CALSCALE_PARAM))
	}

	var calscaleRegex = regexp.MustCompile(`^(gregorian|[a-zA-Z0-9-]+)$`)
	if !calscaleRegex.MatchString(cp.Val[0]) {
		return ErrInvalidCalscale.Error(cp.Val[0])
	}

	return nil
}

type SortAsParam struct {
	BaseParam
}

// Creates a new SortAsParam with the provided sortAs value.
// If the sortAs value is invalid, an error is returned.
// https://tools.ietf.org/html/rfc6350#section-5.9
func NewSortAsParam(sortAs string, propCompCount int) (*SortAsParam, error) {
	sap := &SortAsParam{
		BaseParam: BaseParam{
			Name: SORTAS_PARAM,
			Val:  []string{sortAs},
		},
	}
	if err := sap.validate(propCompCount); err != nil {
		return nil, err
	}
	return sap, nil
}

func (sap *SortAsParam) validate(propCompCount int) error {
	sapVal := strings.Split(sap.Val[0], ",")
	if len(sapVal) == 0 {
		return ErrNilSortAs.Error(string(SORTAS_PARAM))
	}
	if propCompCount > len(sapVal) {
		return ErrInvalidSortAs.Error(sapVal[0])
	}
	for _, value := range sapVal {
		// If any element is empty, return false
		if strings.TrimSpace(value) == "" {
			return ErrInvalidSortAs.Error(value)
		}
	}
	return nil
}

type GeoParam struct {
	BaseParam
}

// Creates a new GeoParam with the provided geo value.
// If the geo value is invalid, an error is returned.
// https://tools.ietf.org/html/rfc6350#section-5.10
func NewGeoParam(geo string) (*GeoParam, error) {
	gp := &GeoParam{
		BaseParam: BaseParam{
			Name: GEO_PARAM,
			Val:  []string{geo},
		},
	}
	if err := gp.validate(); err != nil {
		return nil, err
	}
	return gp, nil
}

func (gp *GeoParam) validate() error {
	if len(gp.Val) == 0 {
		return ErrNilGeo.Error(string(GEO_PARAM))
	}

	validateGeoParam := func(paramValue ParamVal) bool {
		var geoRegex = regexp.MustCompile(`^"([^"]+)"$`)
		return geoRegex.MatchString(string(paramValue))
	}
	if !validateGeoParam(ParamVal(gp.Val[0])) {
		return ErrInvalidGeo.Error(gp.Val[0])
	}
	return nil
}
