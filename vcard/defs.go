package vcard

const (
	// delimiters
	BEGIN      = "BEGIN"
	END        = "END"
	COLON      = ":"
	SEMICOLON  = ";"
	ENDLINE    = "/n"
	COMMA      = ","
	EQUAL      = "="
	DUBQUAL    = "=="
	DOT        = "."
	FOLDEDLINE = "\r\n"
)

// vCard property Names
const (
	VERSION       = "VERSION"
	SOURCE        = "SOURCE"
	KIND          = "KIND"
	XML           = "XML"
	FN            = "FN"
	N             = "N"
	NICKNAME      = "NICKNAME"
	PHOTO         = "PHOTO"
	BDAY          = "BDAY"
	ANNIVERSARY   = "ANNIVERSARY"
	GENDER        = "GENDER"
	ADR           = "ADR"
	TEL           = "TEL"
	EMAIL         = "EMAIL"
	IMPP          = "IMPP"
	LANG          = "LANG"
	TZ            = "TZ"
	GEO           = "GEO"
	TITLE         = "TITLE"
	ROLE          = "ROLE"
	LOGO          = "LOGO"
	ORG           = "ORG"
	MEMBER        = "MEMBER"
	RELATED       = "RELATED"
	CATEGORIES    = "CATEGORIES"
	NOTE          = "NOTE"
	PRODID        = "PRODID"
	REV           = "REV"
	SOUND         = "SOUND"
	UID           = "UID"
	CLIENTPIDMAP  = "CLIENTPIDMAP"
	URL           = "URL"
	KEY           = "KEY"
	FBURL         = "FBURL"
	CALADRURI     = "CALADRURI"
	CALURI        = "CALURI"
	BIRTHPLACE    = "BIRTHPLACE"
	DEATHPLACE    = "DEATHPLACE"
	DEATHDATE     = "DEATHDATE"
	EXPERTISE     = "EXPERTISE"
	HOBBY         = "HOBBY"
	INTEREST      = "INTEREST"
	ORG_DIRECTORY = "ORG-DIRECTORY"
	CONTACTURI    = "CONTACT-URI"
	CREATED       = "CREATED"
	GRAMGENDER    = "GRAMGENDER"
	LANGUAGE      = "LANGUAGE"
	PRONOUNS      = "PRONOUNS"
	SOCIALPROFILE = "SOCIALPROFILE"
	JSPROP        = "JSPROP"
	X             = "X-"
)

type ParVal string
type ParName string

// represents a vCard parameter
// https://tools.ietf.org/html/rfc6350#section-5.3
type Par struct {
	Name    ParName
	PosVals []ParVal
	Val     []string
}

//	type Par struct {
//		Name    ParamName
//		PosVals []ParamValue
//		Value   []string
//	}
//
// https://tools.ietf.org/html/rfc6350#section-5.3
var (
	LANGUAGEPAR = Par{
		Name:    "LANGUAGE",
		PosVals: nil,
		Val:     nil}
	VALUE = Par{
		Name:    "VALUE",
		PosVals: []ParVal{"uri", "text", "date", "time", "date-time", "date-and-or-time", "timestamp", "boolean", "integer", "float", "utc-offset", "language-tag"},
		Val:     nil}
	PREF = Par{
		Name:    "PREF",
		PosVals: nil,
		Val:     nil}
	ALTID = Par{
		Name:    "ALTID",
		PosVals: nil,
		Val:     nil}
	PID = Par{
		Name:    "PID",
		PosVals: nil,
		Val:     nil}
	TYPE = Par{
		Name:    "TYPE",
		PosVals: []ParVal{"work", "home", "text", "voice", "fax", "cell", "video", "pager", "textphone", "textphone", "main", "other"},
		Val:     nil}
	MEDIATYPE = Par{
		Name:    "MEDIATYPE",
		PosVals: nil,
		Val:     nil}
	CALSCALE = Par{
		Name:    "CALSCALE",
		PosVals: []ParVal{"gregorian"},
		Val:     nil}
	SORT_AS = Par{
		Name:    "SORT-AS",
		PosVals: nil,
		Val:     nil}
	GEO_PARAM = Par{
		Name:    "GEO",
		PosVals: nil,
		Val:     nil}
	TZ_PARAM = Par{
		Name:    "TZ",
		PosVals: nil,
		Val:     nil}
	INDEX = Par{
		Name:    "INDEX",
		PosVals: nil,
		Val:     nil}
	LEVEL = Par{
		Name:    "LEVEL",
		PosVals: nil,
		Val:     nil}
	GROUP = Par{
		Name:    "GROUP",
		PosVals: nil,
		Val:     nil}
	CC = Par{
		Name:    "CC",
		PosVals: nil,
		Val:     nil}
	AUTHOR = Par{
		Name:    "AUTHOR",
		PosVals: nil,
		Val:     nil}
	AUTHORNAME = Par{
		Name:    "AUTHOR-NAME",
		PosVals: nil,
		Val:     nil}
	CREATED_PARAM = Par{
		Name:    "CREATED",
		PosVals: nil,
		Val:     nil}
	DERIVED = Par{
		Name:    "DERIVED",
		PosVals: nil,
		Val:     nil}
	LABEL = Par{
		Name:    "LABEL",
		PosVals: nil,
		Val:     nil}
	PHONETIC = Par{
		Name:    "PHONETIC",
		PosVals: nil,
		Val:     nil}
	PROPID = Par{
		Name:    "PROP-ID",
		PosVals: nil,
		Val:     nil}
	SCRIPT = Par{
		Name:    "SCRIPT",
		PosVals: nil,
		Val:     nil}
	SERVICETYPE = Par{
		Name:    "SERVICE-TYPE",
		PosVals: nil,
		Val:     nil}
	USERNAME = Par{
		Name:    "USERNAME",
		PosVals: nil,
		Val:     nil}
	JSPTR = Par{
		Name:    "JSPTR",
		PosVals: nil,
		Val:     nil}
)

type ValueType string
type PropName string
type PropValue string

type Property struct {
	Name       PropName
	PosPars    []Par
	ValueTypes []ValueType
	PosVals    []PropValue
	Value      PropValue
}

var PROPERTIES = map[PropName]Property{
	BEGIN: {
		Name:       BEGIN,
		PosPars:    nil,
		ValueTypes: nil,
		PosVals:    nil,
		Value:      "VCARD",
	},
	END: {
		Name:       END,
		PosPars:    nil,
		ValueTypes: nil,
		PosVals:    nil,
		Value:      "VCARD",
	},
	SOURCE: {
		Name:       SOURCE,
		PosPars:    nil,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	KIND: {
		Name:       KIND,
		PosPars:    nil,
		ValueTypes: []ValueType{"text"},
		PosVals:    []PropValue{"individual", "group", "org", "location"},
		Value:      "",
	},
	XML: {
		Name:       XML,
		PosPars:    []Par{ALTID},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	FN: {
		Name:       FN,
		PosPars:    []Par{TYPE, LANGUAGEPAR, ALTID, PID, PREF, DERIVED},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	N: {
		Name:       N,
		PosPars:    []Par{SORT_AS, LANGUAGEPAR, ALTID, PHONETIC},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	NICKNAME: {
		Name:       NICKNAME,
		PosPars:    []Par{TYPE, LANGUAGEPAR, ALTID, PID, PREF},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	PHOTO: {
		Name:       PHOTO,
		PosPars:    []Par{TYPE, ALTID, MEDIATYPE, PREF, PID},
		ValueTypes: []ValueType{"uri", "binary"},
		PosVals:    nil,
		Value:      "",
	},
	BDAY: {
		Name:       BDAY,
		PosPars:    []Par{VALUE, ALTID, LANGUAGEPAR, CALSCALE}, // only possible if VALUE is date and true
		ValueTypes: []ValueType{"date-and-or-time", "text"},
		PosVals:    nil,
		Value:      "",
	},
	ANNIVERSARY: {
		Name:       ANNIVERSARY,
		PosPars:    []Par{VALUE, ALTID, LANGUAGEPAR, CALSCALE}, // only possible if VALUE is date and true
		ValueTypes: []ValueType{"date-and-or-time", "text"},
		PosVals:    nil,
		Value:      "",
	},
	GENDER: {
		Name:       GENDER,
		PosPars:    nil,
		ValueTypes: []ValueType{"text"},
		PosVals:    []PropValue{"M", "F", "O", "N"},
		Value:      "",
	},
	ADR: {
		Name:       ADR,
		PosPars:    []Par{LABEL, LANGUAGEPAR, GEO_PARAM, TZ_PARAM, ALTID, PID, PREF, TYPE},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	TEL: {
		Name:       TEL,
		PosPars:    []Par{TYPE, VALUE, PREF, PID, ALTID, MEDIATYPE},
		ValueTypes: []ValueType{"uri", "text"},
		PosVals:    nil,
		Value:      "",
	},
	EMAIL: {
		Name:       EMAIL,
		PosPars:    []Par{TYPE, PREF, PID, ALTID},
		ValueTypes: []ValueType{"uri", "text"},
		PosVals:    nil,
		Value:      "",
	},
	IMPP: {
		Name:       IMPP,
		PosPars:    []Par{TYPE, PREF, PID, ALTID, MEDIATYPE},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	LANG: {
		Name:       LANG,
		PosPars:    []Par{PREF, PID, ALTID, TYPE},
		ValueTypes: []ValueType{"language-tag"},
		PosVals:    nil, //To many languages to list
		Value:      "",
	},
	TZ: {
		Name:       TZ,
		PosPars:    []Par{VALUE, PREF, PID, ALTID, TYPE, MEDIATYPE},
		ValueTypes: []ValueType{"text", "uri", "utc-offset"}, // appreantly utc-offset is not recommended
		PosVals:    nil,
		Value:      "",
	},
	GEO: {
		Name:       GEO,
		PosPars:    []Par{PREF, PID, ALTID, TYPE, MEDIATYPE},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	TITLE: {
		Name:       TITLE,
		PosPars:    []Par{LANGUAGEPAR, ALTID, PID, PREF, TYPE},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	ROLE: {
		Name:       ROLE,
		PosPars:    []Par{LANGUAGEPAR, ALTID, PID, PREF, TYPE},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	LOGO: {
		Name:       LOGO,
		PosPars:    []Par{LANGUAGEPAR, ALTID, MEDIATYPE, PREF, PID, TYPE},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	ORG: {
		Name:       ORG,
		PosPars:    []Par{SORT_AS, LANGUAGEPAR, PID, PREF, ALTID, TYPE},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	MEMBER: {
		Name:       MEMBER,
		PosPars:    []Par{MEDIATYPE, ALTID, PID, PREF},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	RELATED: {
		Name:       RELATED,
		PosPars:    []Par{TYPE, VALUE, PREF, PID, ALTID, MEDIATYPE},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	CATEGORIES: {
		Name:       CATEGORIES,
		PosPars:    []Par{ALTID, PID, PREF, TYPE},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	NOTE: {
		Name:       NOTE,
		PosPars:    []Par{LANGUAGEPAR, ALTID, PID, PREF, TYPE},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	PRODID: {
		Name:       PRODID,
		PosPars:    nil,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	REV: {
		Name:       REV,
		PosPars:    nil,
		ValueTypes: []ValueType{"timestamp"},
		PosVals:    nil,
		Value:      "",
	},
	SOUND: {
		Name:       SOUND,
		PosPars:    []Par{LANGUAGEPAR, ALTID, MEDIATYPE, PREF, PID, TYPE},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	UID: {
		Name:       UID,
		PosPars:    []Par{VALUE},
		ValueTypes: []ValueType{"uri", "text"}, // it shouldnt be text but it can be
		PosVals:    nil,
		Value:      "",
	},
	CLIENTPIDMAP: {
		Name:       CLIENTPIDMAP,
		PosPars:    nil,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	URL: {
		Name:       URL,
		PosPars:    []Par{MEDIATYPE, ALTID, PID, PREF, TYPE},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	VERSION: {
		Name:       VERSION,
		PosPars:    nil,
		ValueTypes: []ValueType{"text"},
		PosVals:    []PropValue{"2.1", "3.0", "4.0"},
		Value:      "",
	},
	KEY: {
		Name:       KEY,
		PosPars:    []Par{TYPE, MEDIATYPE, ALTID, PID, PREF},
		ValueTypes: []ValueType{"text", "uri"},
		PosVals:    nil,
		Value:      "",
	},
	FBURL: {
		Name:       FBURL,
		PosPars:    []Par{ALTID, MEDIATYPE, PREF, PID, TYPE},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	CALADRURI: {
		Name:       CALADRURI,
		PosPars:    []Par{ALTID, MEDIATYPE, PREF, PID, TYPE},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	CALURI: {
		Name:       CALURI,
		PosPars:    []Par{ALTID, MEDIATYPE, PREF, PID, TYPE},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	BIRTHPLACE: {
		Name:       BIRTHPLACE,
		PosPars:    []Par{VALUE, LANGUAGEPAR, ALTID},
		ValueTypes: []ValueType{"text", "uri"},
		PosVals:    nil,
		Value:      "",
	},
	DEATHPLACE: {
		Name:       DEATHPLACE,
		PosPars:    []Par{VALUE, LANGUAGEPAR, ALTID},
		ValueTypes: []ValueType{"text", "uri"},
		PosVals:    nil,
		Value:      "",
	},
	DEATHDATE: {
		Name:       DEATHDATE,
		PosPars:    []Par{VALUE, CALSCALE, LANGUAGEPAR, ALTID},
		ValueTypes: []ValueType{"date-and-or-time", "text"},
		PosVals:    nil,
		Value:      "",
	},
	EXPERTISE: {
		Name:       EXPERTISE,
		PosPars:    []Par{LEVEL, INDEX, LANGUAGEPAR, PREF, ALTID, TYPE},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	HOBBY: {
		Name:       HOBBY,
		PosPars:    []Par{LEVEL, INDEX, LANGUAGEPAR, PREF, ALTID, TYPE},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	INTEREST: {
		Name:       INTEREST,
		PosPars:    []Par{LEVEL, INDEX, LANGUAGEPAR, PREF, ALTID, TYPE},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
	},
	ORG_DIRECTORY: {
		Name:       ORG_DIRECTORY,
		PosPars:    []Par{PREF, INDEX, LANGUAGEPAR, PID, PREF, ALTID, TYPE},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	CONTACTURI: {
		Name:       CONTACTURI,
		PosPars:    []Par{PREF},
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	CREATED: {
		Name:       CREATED,
		PosPars:    []Par{VALUE},
		ValueTypes: []ValueType{"timestamp"},
		PosVals:    []PropValue{"timestamp"},
		Value:      "",
	},
	GRAMGENDER: {
		Name:       GRAMGENDER,
		PosPars:    []Par{LANGUAGEPAR},
		ValueTypes: []ValueType{"text"},
		PosVals:    []PropValue{"animate", "common", "feminine", "masculine", "neuter"},
		Value:      "",
	},
	LANGUAGE: {
		Name:       LANGUAGE,
		PosPars:    nil,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	PRONOUNS: {
		Name:       PRONOUNS,
		PosPars:    []Par{LANGUAGEPAR, ALTID, PREF, TYPE},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	SOCIALPROFILE: {
		Name:       SOCIALPROFILE,
		PosPars:    []Par{SERVICETYPE},
		ValueTypes: []ValueType{"uri", "text"},
		PosVals:    nil,
		Value:      "",
	},
	JSPROP: {
		Name:       JSPROP,
		PosPars:    []Par{JSPTR},
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
}
