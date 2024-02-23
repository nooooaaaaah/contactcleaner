package vcard

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	LANGUAGE_PARAM     ParamName = "LANGUAGE"
	VALUE_PARAM        ParamName = "VALUE"
	PREF_PARAM         ParamName = "PREF"
	ALTID_PARAM        ParamName = "ALTID"
	PID_PARAM          ParamName = "PID"
	TYPE_PARAM         ParamName = "TYPE"
	MEDIATYPE_PARAM    ParamName = "MEDIATYPE"
	CALSCALE_PARAM     ParamName = "CALSCALE"
	SORTAS_PARAM       ParamName = "SORT-AS"
	GEO_PARAM          ParamName = "GEO"
	TZ_PARAM           ParamName = "TZ"
	INDEX_PARAM        ParamName = "INDEX"
	LEVEL_PARAM        ParamName = "LEVEL"
	GROUP_PARAM        ParamName = "GROUP"
	CC_PARAM           ParamName = "CC"
	AUTHOR_PARAM       ParamName = "AUTHOR"
	AUTHOR_NAME_PARAM  ParamName = "AUTHOR-NAME"
	CREATED_PARAM      ParamName = "CREATED"
	DERIVED_PARAM      ParamName = "DERIVED"
	LABEL_PARAM        ParamName = "LABEL"
	PHONETIC_PARAM     ParamName = "PHONETIC"
	PROP_ID_PARAM      ParamName = "PROP-ID"
	SCRIPT_PARAM       ParamName = "SCRIPT"
	SERVICE_TYPE_PARAM ParamName = "SERVICE-TYPE"
	USERNAME_PARAM     ParamName = "USERNAME"
	JSPTR_PARAM        ParamName = "JSPTR"
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

type TzParam struct {
	BaseParam
}

// Creates a new TzParam with the provided tz value.
// If the tz value is invalid, an error is returned.
// https://tools.ietf.org/html/rfc6350#section-5.11
func NewTzParam(tz string) (*TzParam, error) {
	tp := &TzParam{
		BaseParam: BaseParam{
			Name: TZ_PARAM,
			Val:  []string{tz},
		},
	}
	if err := tp.validate(); err != nil {
		return nil, err
	}
	return tp, nil
}

func (tp *TzParam) validate() error {
	if len(tp.Val) == 0 {
		return ErrNilTz.Error(string(TZ_PARAM))
	}

	validateTzParam := func(paramValue ParamVal) bool {
		var paramRegex = regexp.MustCompile(`^[^"]+$`)
		var uriRegex = regexp.MustCompile(`^"([^"]+)"$`)

		return paramRegex.MatchString(string(paramValue)) || uriRegex.MatchString(string(paramValue))
	}
	if !validateTzParam(ParamVal(tp.Val[0])) {
		return ErrInvalidTz.Error(tp.Val[0])
	}
	return nil
}

type IndexParam struct {
	BaseParam
}

// Creates a new IndexParam with the provided index value.
// If the index value is invalid, an error is returned.
// https://tools.ietf.org/html/rfc6715#section-3.1
func NewIndexParam(index string) (*IndexParam, error) {
	ip := &IndexParam{
		BaseParam: BaseParam{
			Name: INDEX_PARAM,
			Val:  []string{index},
		},
	}
	if err := ip.validate(); err != nil {
		return nil, err
	}
	return ip, nil
}

func (ip *IndexParam) validate() error {
	param := ip.Val[0]

	if len(param) == 0 {
		return ErrNilIndex.Error(string(INDEX_PARAM))
	}

	var indexRegex = regexp.MustCompile(`^[1-9][0-9]*$`)
	if !indexRegex.MatchString(param) {
		return ErrInvalidIndex.Error(param)
	}

	index, err := strconv.Atoi(param)
	if err != nil {
		return ErrInvalidIndex.Error(param)
	}
	if index < 1 {
		return ErrInvalidIndex.Error(param)
	}

	return nil
}

type LevelParam struct {
	BaseParam
}

// Creates a new LevelParam with the provided level value.
// If the level value is invalid, an error is returned.
// https://tools.ietf.org/html/rfc6715#section-3.2
func NewLevelParam(level string) (*LevelParam, error) {
	lp := &LevelParam{
		BaseParam: BaseParam{
			Name: LEVEL_PARAM,
			Val:  []string{level},
		},
	}
	if err := lp.validate(); err != nil {
		return nil, err
	}
	return lp, nil
}

func (lp *LevelParam) validate() error {
	param := lp.Val[0]

	if len(param) == 0 {
		return ErrNilLevel.Error(string(LEVEL_PARAM))
	}

	allowedLevels := map[string]bool{
		"beginner": true,
		"average":  true,
		"expert":   true,
		"high":     true,
		"medium":   true,
		"low":      true,
	}

	if !allowedLevels[strings.ToLower(param)] {
		return ErrInvalidLevel.Error(param)
	}

	return nil
}

type GroupParam struct {
	BaseParam
}

// Creates a new GroupParam with the provided group value.
// If the group value is invalid, an error is returned.
// Only use with jCard objects.
// https://tools.ietf.org/html/rfc7095#section-8.1
func NewGroupParam(group string) (*GroupParam, error) {
	gp := &GroupParam{
		BaseParam: BaseParam{
			Name: GROUP_PARAM,
			Val:  []string{group},
		},
	}
	if err := gp.validate(); err != nil {
		return nil, err
	}
	return gp, nil
}

func (gp *GroupParam) validate() error {
	param := gp.Val[0]
	if len(param) == 0 {
		return ErrNilGroup.Error(string(GROUP_PARAM))
	}

	for _, char := range gp.Val[0] {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '-' {
			return ErrInvalidGroup.Error(string(char))
		}
	}
	return nil
}

type CcParam struct {
	BaseParam
}

// Creates a new CcParam with the provided cc value.
// If the cc value is invalid, an error is returned.
// https://tools.ietf.org/html/rfc8605#section-3.1
func NewCcParam(cc string) (*CcParam, error) {
	cp := &CcParam{
		BaseParam: BaseParam{
			Name: CC_PARAM,
			Val:  []string{cc},
		},
	}
	if err := cp.validate(); err != nil {
		return nil, err
	}
	return cp, nil
}

func (cp *CcParam) validate() error {
	param := cp.Val[0]
	if len(param) == 0 {
		return ErrNilCc.Error(string(CC_PARAM))
	}

	validateCcParam := func(paramValue ParamVal) bool {
		var ccRegex = regexp.MustCompile(`^[A-Z]{2}$`)
		return ccRegex.MatchString(string(paramValue))
	}
	if !validateCcParam(ParamVal(param)) {
		return ErrInvalidCc.Error(param)
	}
	return nil
}

type AuthorParam struct {
	BaseParam
}

// Creates a new AuthorParam with the provided author uri value.
// If the author value is invalid, an error is returned.
// https://www.ietf.org/archive/id/draft-ietf-calext-vcard-jscontact-extensions-11.html#name-author
func NewAuthorParam(author string) (*AuthorParam, error) {
	ap := &AuthorParam{
		BaseParam: BaseParam{
			Name: AUTHOR_PARAM,
			Val:  []string{author},
		},
	}
	if err := ap.validate(); err != nil {
		return nil, err
	}
	return ap, nil
}

func (ap *AuthorParam) validate() error {
	param := ap.Val[0]
	if len(param) == 0 {
		return ErrNilAuthor.Error(string(AUTHOR_PARAM))
	}

	validateAuthorParam := func(paramValue ParamVal) bool {
		var uriRegex = regexp.MustCompile(`^".+"$`)
		return uriRegex.MatchString(string(paramValue))
	}
	if !validateAuthorParam(ParamVal(param)) {
		return ErrInvalidAuthor.Error(param)
	}

	return nil
}

type AuthorNameParam struct {
	BaseParam
}

// Creates a new AuthorNameParam with the provided author name value.
// If the author name value is invalid, an error is returned.
// https://www.ietf.org/archive/id/draft-ietf-calext-vcard-jscontact-extensions-11.html#name-author-name
func NewAuthorNameParam(authorName string) (*AuthorNameParam, error) {
	anp := &AuthorNameParam{
		BaseParam: BaseParam{
			Name: AUTHOR_NAME_PARAM,
			Val:  []string{authorName},
		},
	}
	if err := anp.validate(); err != nil {
		return nil, err
	}
	return anp, nil
}

func (anp *AuthorNameParam) validate() error {
	param := anp.Val[0]
	if len(param) == 0 {
		return ErrNilAuthorName.Error(string(AUTHOR_NAME_PARAM))
	}

	validateAuthorNameParam := func(paramValue ParamVal) bool {
		var anpRegex = regexp.MustCompile(`^"[^"]+"$`)
		return anpRegex.MatchString(string(paramValue))
	}
	if !validateAuthorNameParam(ParamVal(param)) {
		return ErrInvalidAuthorName.Error(param)
	}

	return nil
}

type CreatedParam struct {
	BaseParam
}

// Creates a new CreatedParam with the provided created timestamp value.
// If the created value is invalid, an error is returned.
// https://www.ietf.org/archive/id/draft-ietf-calext-vcard-jscontact-extensions-11.html#name-created
func NewCreatedParam(created string) (*CreatedParam, error) {
	cp := &CreatedParam{
		BaseParam: BaseParam{
			Name: CREATED_PARAM,
			Val:  []string{created},
		},
	}
	if err := cp.validate(); err != nil {
		return nil, err
	}
	return cp, nil
}

func (cp *CreatedParam) validate() error {
	param := cp.Val[0]
	if len(param) == 0 {
		return ErrNilCreated.Error(string(CREATED_PARAM))
	}

	validateCreatedParam := func(paramValue ParamVal) bool {
		var createdRegex = regexp.MustCompile(`^\d{8}T\d{6}Z$`)
		return createdRegex.MatchString(string(paramValue))
	}
	if !validateCreatedParam(ParamVal(param)) {
		return ErrInvalidCreated.Error(param)
	}

	return nil
}

type DerivedParam struct {
	BaseParam
}

// Creates a new DerivedParam with the provided derived timestamp value.
// If the derived value is invalid, an error is returned.
// https://www.ietf.org/archive/id/draft-ietf-calext-vcard-jscontact-extensions-11.html#name-derived
func NewDerivedParam(derived string) (*DerivedParam, error) {
	dp := &DerivedParam{
		BaseParam: BaseParam{
			Name: DERIVED_PARAM,
			Val:  []string{derived},
		},
	}
	if err := dp.validate(); err != nil {
		return nil, err
	}
	return dp, nil
}

func (dp *DerivedParam) validate() error {
	param := dp.Val[0]
	if len(param) == 0 {
		return ErrNilDerived.Error(string(DERIVED_PARAM))
	}

	validateDerivedParam := func(paramValue string) bool {
		derived := strings.ToLower(paramValue)
		return derived == "true" || derived == "false"
	}
	if !validateDerivedParam(param) {
		return ErrInvalidDerived.Error(param)
	}

	return nil
}

type LabelParam struct {
	BaseParam
}

// Creates a new LabelParam with the provided label value.
// If the label value is invalid, an error is returned.
// https://www.ietf.org/archive/id/draft-ietf-calext-vcard-jscontact-extensions-11.html#name-label
func NewLabelParam(label string) (*LabelParam, error) {
	lp := &LabelParam{
		BaseParam: BaseParam{
			Name: LABEL_PARAM,
			Val:  []string{label},
		},
	}
	if err := lp.validate(); err != nil {
		return nil, err
	}
	return lp, nil
}

func (lp *LabelParam) validate() error {
	param := lp.Val[0]
	if len(param) == 0 {
		return ErrNilLabel.Error(string(LABEL_PARAM))
	}

	return nil
}

type PhoneticParam struct {
	BaseParam
}

// Creates a new PhoneticParam with the provided phonetic value.
// If the phonetic value is invalid, an error is returned.
// https://www.ietf.org/archive/id/draft-ietf-calext-vcard-jscontact-extensions-11.html#name-phonetic
func NewPhoneticParam(phonetic string) (*PhoneticParam, error) {
	pp := &PhoneticParam{
		BaseParam: BaseParam{
			Name: PHONETIC_PARAM,
			Val:  []string{phonetic},
		},
	}
	if err := pp.validate(); err != nil {
		return nil, err
	}
	return pp, nil
}

func (pp *PhoneticParam) validate() error {
	param := pp.Val[0]
	if len(param) == 0 {
		return ErrNilPhonetic.Error(string(PHONETIC_PARAM))
	}
	allowedSystems := map[string]bool{"ipa": true, "piny": true, "jyut": true, "script": true}
	if !allowedSystems[param] {
		return ErrInvalidPhonetic.Error(param)
	}

	return nil
}

type PropIDParam struct {
	BaseParam
}

// Creates a new propParam with the provided prop value.
// If the prop value is invalid, an error is returned.
// https://www.ietf.org/archive/id/draft-ietf-calext-vcard-jscontact-extensions-11.html#name-prop
func NewPropIDParam(prop string) (*PropIDParam, error) {
	pp := &PropIDParam{
		BaseParam: BaseParam{
			Name: PROP_ID_PARAM,
			Val:  []string{prop},
		},
	}
	if err := pp.validate(); err != nil {
		return nil, err
	}
	return pp, nil
}

func (pp *PropIDParam) validate() error {
	param := pp.Val[0]
	if len(param) == 0 {
		return ErrNilPropID.Error(string(PROP_ID_PARAM))
	}

	validatePropIDParam := func(paramValue ParamVal) bool {
		var propIDRegex = regexp.MustCompile(`^[A-Za-z0-9\\-_]{1,255}$`)
		return propIDRegex.MatchString(string(paramValue))
	}
	if !validatePropIDParam(ParamVal(param)) {
		return ErrInvalidPropID.Error(param)
	}

	return nil
}

type ScriptParam struct {
	BaseParam
}

// Creates a new ScriptParam with the provided script value.
// If the script value is invalid, an error is returned.
// https://www.ietf.org/archive/id/draft-ietf-calext-vcard-jscontact-extensions-11.html#name-script
func NewScriptParam(script string) (*ScriptParam, error) {
	sp := &ScriptParam{
		BaseParam: BaseParam{
			Name: SCRIPT_PARAM,
			Val:  []string{script},
		},
	}
	if err := sp.validate(); err != nil {
		return nil, err
	}
	return sp, nil
}

func (sp *ScriptParam) validate() error {
	param := sp.Val[0]
	if len(param) == 0 {
		return ErrNilScript.Error(string(SCRIPT_PARAM))
	}

	pattern := regexp.MustCompile(`^[A-Za-z]{4}$`)
	if !pattern.MatchString(param) {
		return ErrInvalidScript.Error(param)
	}

	return nil
}

type ServiceTypeParam struct {
	BaseParam
}

// Creates a new ServiceTypeParam with the provided service type value.
// If the service type value is invalid, an error is returned.
// https://www.ietf.org/archive/id/draft-ietf-calext-vcard-jscontact-extensions-11.html#name-service-type
func NewServiceTypeParam(serviceType string) (*ServiceTypeParam, error) {
	stp := &ServiceTypeParam{
		BaseParam: BaseParam{
			Name: SERVICE_TYPE_PARAM,
			Val:  []string{serviceType},
		},
	}
	if err := stp.validate(); err != nil {
		return nil, err
	}
	return stp, nil
}

func (stp *ServiceTypeParam) validate() error {
	param := stp.Val[0]
	if len(param) == 0 {
		return ErrNilServiceType.Error(string(SERVICE_TYPE_PARAM))
	}
	pattern := regexp.MustCompile(`^.*$`)
	if !pattern.MatchString(param) {
		return ErrInvalidServiceType.Error(param)
	}

	return nil
}

type UsernameParam struct {
	BaseParam
}

// Creates a new UsernameParam with the provided username value.
// If the username value is invalid, an error is returned.
// https://www.ietf.org/archive/id/draft-ietf-calext-vcard-jscontact-extensions-11.html#name-username
func NewUsernameParam(username string) (*UsernameParam, error) {
	up := &UsernameParam{
		BaseParam: BaseParam{
			Name: USERNAME_PARAM,
			Val:  []string{username},
		},
	}
	if err := up.validate(); err != nil {
		return nil, err
	}
	return up, nil
}

func (up *UsernameParam) validate() error {
	param := up.Val[0]
	if len(param) == 0 {
		return ErrNilUsername.Error(string(USERNAME_PARAM))
	}

	pattern := regexp.MustCompile(`^.*$`)
	if !pattern.MatchString(param) {
		return ErrInvalidUsername.Error(param)
	}

	return nil
}

type JsptrParam struct {
	BaseParam
}

// Creates a new JsptrParam with the provided jsptr value.
// If the jsptr value is invalid, an error is returned.
// https://www.ietf.org/archive/id/draft-ietf-calext-vcard-jscontact-extensions-11.html#name-jsptr
func NewJsptrParam(jsptr string) (*JsptrParam, error) {
	jp := &JsptrParam{
		BaseParam: BaseParam{
			Name: JSPTR_PARAM,
			Val:  []string{jsptr},
		},
	}
	if err := jp.validate(); err != nil {
		return nil, err
	}
	return jp, nil
}

func (jp *JsptrParam) validate() error {
	param := jp.Val[0]
	if len(param) == 0 {
		return ErrNilJsptr.Error(string(JSPTR_PARAM))
	}

	pattern := regexp.MustCompile(`^"\/(?:[^\/"]|\\")*\/"$`)
	if !pattern.MatchString(param) {
		return ErrInvalidJsptr.Error(param)
	}

	return nil
}
