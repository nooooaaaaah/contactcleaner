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

type ValueType string
type PropName string
type PropValue string

type Property struct {
	Name       PropName
	ValueTypes []ValueType
	PosVals    []PropValue
	Value      PropValue
}

var PROPERTIES = map[PropName]Property{
	BEGIN: {
		Name:       BEGIN,
		ValueTypes: nil,
		PosVals:    nil,
		Value:      "VCARD",
	},
	END: {
		Name:       END,
		ValueTypes: nil,
		PosVals:    nil,
		Value:      "VCARD",
	},
	SOURCE: {
		Name:       SOURCE,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	KIND: {
		Name:       KIND,
		ValueTypes: []ValueType{"text"},
		PosVals:    []PropValue{"individual", "group", "org", "location"},
		Value:      "",
	},
	XML: {
		Name:       XML,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	FN: {
		Name:       FN,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	N: {
		Name:       N,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	NICKNAME: {
		Name:       NICKNAME,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	PHOTO: {
		Name:       PHOTO,
		ValueTypes: []ValueType{"uri", "binary"},
		PosVals:    nil,
		Value:      "",
	},
	BDAY: {
		Name:       BDAY,
		ValueTypes: []ValueType{"date-and-or-time", "text"},
		PosVals:    nil,
		Value:      "",
	},
	ANNIVERSARY: {
		Name:       ANNIVERSARY,
		ValueTypes: []ValueType{"date-and-or-time", "text"},
		PosVals:    nil,
		Value:      "",
	},
	GENDER: {
		Name:       GENDER,
		ValueTypes: []ValueType{"text"},
		PosVals:    []PropValue{"M", "F", "O", "N"},
		Value:      "",
	},
	ADR: {
		Name:       ADR,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	TEL: {
		Name:       TEL,
		ValueTypes: []ValueType{"uri", "text"},
		PosVals:    nil,
		Value:      "",
	},
	EMAIL: {
		Name:       EMAIL,
		ValueTypes: []ValueType{"uri", "text"},
		PosVals:    nil,
		Value:      "",
	},
	IMPP: {
		Name:       IMPP,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	LANG: {
		Name:       LANG,
		ValueTypes: []ValueType{"language-tag"},
		PosVals:    nil, //To many languages to list
		Value:      "",
	},
	TZ: {
		Name:       TZ,
		ValueTypes: []ValueType{"text", "uri", "utc-offset"}, // appreantly utc-offset is not recommended
		PosVals:    nil,
		Value:      "",
	},
	GEO: {
		Name:       GEO,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	TITLE: {
		Name:       TITLE,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	ROLE: {
		Name:       ROLE,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	LOGO: {
		Name:       LOGO,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	ORG: {
		Name:       ORG,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	MEMBER: {
		Name:       MEMBER,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	RELATED: {
		Name:       RELATED,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	CATEGORIES: {
		Name:       CATEGORIES,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	NOTE: {
		Name:       NOTE,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	PRODID: {
		Name:       PRODID,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	REV: {
		Name:       REV,
		ValueTypes: []ValueType{"timestamp"},
		PosVals:    nil,
		Value:      "",
	},
	SOUND: {
		Name:       SOUND,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	UID: {
		Name:       UID,
		ValueTypes: []ValueType{"uri", "text"}, // it shouldnt be text but it can be
		PosVals:    nil,
		Value:      "",
	},
	CLIENTPIDMAP: {
		Name:       CLIENTPIDMAP,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	URL: {
		Name:       URL,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	VERSION: {
		Name:       VERSION,
		ValueTypes: []ValueType{"text"},
		PosVals:    []PropValue{"2.1", "3.0", "4.0"},
		Value:      "",
	},
	KEY: {
		Name:       KEY,
		ValueTypes: []ValueType{"text", "uri"},
		PosVals:    nil,
		Value:      "",
	},
	FBURL: {
		Name:       FBURL,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	CALADRURI: {
		Name:       CALADRURI,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	CALURI: {
		Name:       CALURI,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	BIRTHPLACE: {
		Name:       BIRTHPLACE,
		ValueTypes: []ValueType{"text", "uri"},
		PosVals:    nil,
		Value:      "",
	},
	DEATHPLACE: {
		Name:       DEATHPLACE,
		ValueTypes: []ValueType{"text", "uri"},
		PosVals:    nil,
		Value:      "",
	},
	DEATHDATE: {
		Name:       DEATHDATE,
		ValueTypes: []ValueType{"date-and-or-time", "text"},
		PosVals:    nil,
		Value:      "",
	},
	EXPERTISE: {
		Name:       EXPERTISE,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	HOBBY: {
		Name:       HOBBY,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	INTEREST: {
		Name:       INTEREST,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
	},
	ORG_DIRECTORY: {
		Name:       ORG_DIRECTORY,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	CONTACTURI: {
		Name:       CONTACTURI,
		ValueTypes: []ValueType{"uri"},
		PosVals:    nil,
		Value:      "",
	},
	CREATED: {
		Name:       CREATED,
		ValueTypes: []ValueType{"timestamp"},
		PosVals:    []PropValue{"timestamp"},
		Value:      "",
	},
	GRAMGENDER: {
		Name:       GRAMGENDER,
		ValueTypes: []ValueType{"text"},
		PosVals:    []PropValue{"animate", "common", "feminine", "masculine", "neuter"},
		Value:      "",
	},
	LANGUAGE: {
		Name:       LANGUAGE,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	PRONOUNS: {
		Name:       PRONOUNS,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
	SOCIALPROFILE: {
		Name:       SOCIALPROFILE,
		ValueTypes: []ValueType{"uri", "text"},
		PosVals:    nil,
		Value:      "",
	},
	JSPROP: {
		Name:       JSPROP,
		ValueTypes: []ValueType{"text"},
		PosVals:    nil,
		Value:      "",
	},
}
