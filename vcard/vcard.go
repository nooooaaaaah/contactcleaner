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

// vCard properties
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
	PRONOUNS      = "PRONOUNS"
	SOCIALPROFILE = "SOCIALPROFILE"
	JSPROP        = "JSPROP"
	X             = "X-"
)

// vCard parameters
const (
	LANGUAGE      = "LANGUAGE"
	VALUE         = "VALUE"
	PREF          = "PREF"
	ALTID         = "ALTID"
	PID           = "PID"
	TYPE          = "TYPE"
	MEDIATYPE     = "MEDIATYPE"
	CALSCALE      = "CALSCALE"
	SORT_AS       = "SORT-AS"
	GEO_PARAM     = "GEO"
	TZ_PARAM      = "TZ"
	INDEX         = "INDEX"
	LEVEL         = "LEVEL"
	GROUP         = "GROUP"
	CC            = "CC"
	AUTHOR        = "AUTHOR"
	AUTHORNAME    = "AUTHOR-NAME"
	CREATED_PARAM = "CREATED"
	DERIVED       = "DERIVED"
	LABEL         = "LABEL"
	PHONETIC      = "PHONETIC"
	PROPID        = "PROP-ID"
	SCRIPT        = "SCRIPT"
	SERVICETYPE   = "SERVICE-TYPE"
	USERNAME      = "USERNAME"
	JSPTR         = "JSPTR"
	ANY           = "ANY"
)

var possible_params = map[Prop_Name][]Param_Name{
	BEGIN:         nil,
	END:           nil,
	SOURCE:        {VALUE, PID, PREF, ALTID, MEDIATYPE},
	KIND:          {VALUE},
	XML:           {VALUE, ALTID},
	FN:            {VALUE, TYPE, LANGUAGE, ALTID, PID, PREF, DERIVED},
	N:             {VALUE, SORT_AS, LANGUAGE, ALTID, PHONETIC},
	NICKNAME:      {VALUE, TYPE, LANGUAGE, ALTID, PID, PREF},
	PHOTO:         {VALUE, TYPE, ALTID, TYPE, MEDIATYPE, PREF, PID, PROPID},
	BDAY:          {VALUE, LANGUAGE, ALTID, CALSCALE}, // calender scale can only exist if value is date
	ANNIVERSARY:   {VALUE, ALTID, CALSCALE},           // calender scale can only exist if value is date
	GENDER:        {VALUE},
	ADR:           {VALUE, LABEL, LANGUAGE, GEO_PARAM, TZ_PARAM, ALTID, PID, PREF, TYPE},
	TEL:           {VALUE, TYPE, PID, PREF, ALTID},
	EMAIL:         {VALUE, PID, PREF, TYPE, ALTID},
	IMPP:          {VALUE, PID, PREF, TYPE, MEDIATYPE, ALTID},
	LANG:          {VALUE, PID, PREF, ALTID, TYPE},
	TZ:            {VALUE, PID, PREF, TYPE, MEDIATYPE},
	GEO:           {VALUE, PID, PREF, TYPE, MEDIATYPE, ALTID},
	TITLE:         {VALUE, LANGUAGE, PID, PREF, ALTID, TYPE},
	ROLE:          {VALUE, LANGUAGE, PID, PREF, TYPE, ALTID},
	LOGO:          {VALUE, LANGUAGE, PID, PREF, TYPE, MEDIATYPE, ALTID},
	ORG:           {VALUE, SORT_AS, LANGUAGE, PID, PREF, ALTID, TYPE},
	MEMBER:        {VALUE, PID, PREF, ALTID, MEDIATYPE},
	RELATED:       {VALUE, TYPE, LANGUAGE, PID, PREF, ALTID, MEDIATYPE, LANGUAGE}, // language can only exist if value is text and media type can only exist if value is uri
	CATEGORIES:    {VALUE, PID, PREF, TYPE, ALTID},
	NOTE:          {VALUE, LANGUAGE, PID, PREF, TYPE, ALTID, AUTHOR, AUTHORNAME, CREATED},
	PRODID:        {VALUE},
	REV:           {VALUE},
	SOUND:         {VALUE, LANGUAGE, PID, PREF, TYPE, MEDIATYPE, ALTID},
	UID:           {VALUE},
	CLIENTPIDMAP:  nil,
	URL:           {VALUE, PID, PREF, TYPE, MEDIATYPE, ALTID},
	VERSION:       {VALUE},
	KEY:           {VALUE, ALTID, PID, PREF, TYPE, MEDIATYPE}, // media type can only exist if value is uri
	FBURL:         {VALUE, PID, PREF, TYPE, MEDIATYPE, ALTID}, // busy time
	CALADRURI:     {VALUE, PID, PREF, TYPE, MEDIATYPE, ALTID},
	CALURI:        {VALUE, PID, PREF, TYPE, MEDIATYPE, ALTID},
	BIRTHPLACE:    {VALUE, ALTID, LANGUAGE},
	DEATHPLACE:    {VALUE, ALTID, LANGUAGE},
	DEATHDATE:     {VALUE, ALTID, CALSCALE, LANGUAGE},
	EXPERTISE:     {LEVEL, INDEX, LANGUAGE, PREF, ALTID, TYPE},
	HOBBY:         {LEVEL, INDEX, LANGUAGE, PREF, ALTID, TYPE},
	INTEREST:      {LEVEL, INDEX, LANGUAGE, PREF, ALTID, TYPE},
	ORG_DIRECTORY: {PREF, INDEX, LANGUAGE, PID, ALTID, TYPE},
	CONTACTURI:    {VALUE, PREF},
	CREATED:       {VALUE},
	GRAMGENDER:    {LANGUAGE},
	PRONOUNS:      {LANGUAGE, PREF, TYPE, ALTID},
	SOCIALPROFILE: {VALUE, SERVICETYPE},
	JSPROP:        {JSPTR, VALUE},
}

var valid_values = map[Prop_Name]map[Param_Name][]string{
	SOURCE: {VALUE: {"uri"}},
	KIND:   {VALUE: {"individual", "group", "org", "location"}},
}

type Prop_Name string

type Property struct {
	Name  Prop_Name
	Param []Parameter
	Value string
}

type Param_Name string

type Parameter struct {
	Name  Param_Name
	Value []string
}
