package contactcleaner

import (
	// Import necessary packages

	"bufio"
	"errors"
	"strings"
	"time"
)

const (
	// delimiters
	Begin     = "BEGIN"
	END       = "END"
	COLON     = ":"
	SEMICOLON = ";"
	ENDLINE   = "/n"
	COMMA     = ","
	EQUAL     = "="
	DUBQUAL   = "=="
	// vCard fields
	// standard fields
	VERSION       = "VERSION"
	PRODID        = "PRODID"
	N             = "N"
	FN            = "FN"
	TITLE         = "TITLE"
	PHOTO         = "PHOTO"
	LOGO          = "LOGO"
	REV           = "REV"
	UID           = "UID"
	NICKNAME      = "NICKNAME"
	CATEGORIES    = "CATEGORIES"
	IMPP          = "IMPP"
	ADR           = "ADR"
	X             = "X-"
	ITEM          = "item"
	EMAIL         = "EMAIL"
	SOCIALPROFILE = "SOCIALPROFILE"
	BDAY          = "BDAY"
	TEL           = "TEL"
	ORG           = "ORG"
	TYPE          = "TYPE"
	URL           = "URL"
	NOTE          = "NOTE"
	GEO           = "GEO"
	ROLE          = "ROLE"
	TZ            = "TZ"
	MAILER        = "MAILER"
	LABEL         = "LABEL"
	AGENT         = "AGENT"
)

type ContactCard struct {
	Revision         time.Time
	CustomFields     map[string]string
	Birthday         *time.Time
	Version          string
	ProdID           string
	FullName         string
	FirstName        string
	LastName         string
	MiddleName       string
	Prefix           string
	Suffix           string
	UID              string
	Nickname         string
	Organization     string
	URL              string
	Notes            string
	Titles           string
	Photo            Image
	Logos            Image
	Categories       []string // tags
	InstantMessaging []string
	Addresses        []Address
	Emails           []EmailAddr
	SocialProfiles   []SocialMediaProfile
	Telephones       []Telephone
	Items            []Item
	ExtendedFields   []XField
}

type Image interface {
	isEncodedImage() bool
	data() string
}

type EncodedImage string

func (e EncodedImage) isEncodedImage() bool {
	return true
}

func (e EncodedImage) data() string {
	return string(e)
}

type ImageURL string

func (i ImageURL) isEncodedImage() bool {
	return false
}

func (i ImageURL) data() string {
	return string(i)
}

type XField struct {
	Type string
	Data string
}

type Item struct {
	ItemNumber int
	ItemName   string
	ItemValue  string
}

type Address struct {
	Type    string // Type of address (home, work, etc.)
	Street  string
	City    string
	State   string
	Zip     string
	Country string
	Label   string // Custom label (e.g., "Vacation Home")
}

type EmailAddr struct {
	Type    string
	Address string
}

type Telephone struct {
	Type   []string
	Number string
}

type SocialMediaProfile struct {
	Type string
	URL  string
}

type Parser struct {
	error        error
	currentCard  *ContactCard
	scanner      *bufio.Scanner
	currentLine  string
	base64Flag   bool
	b64BuffDaddy []byte
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) NextLine() {
	p.currentLine = p.scanner.Text()
}

func (p *Parser) Parse() (*ContactCard, error) {
	// Parse the vCard string
	var err error

	for p.scanner.Scan() {
		if p.error != nil {
			return nil, p.error
		}
		// If we are in the middle of a base64 encoded block, keep reading until we find the end
		if p.base64Flag {
			p.parseBase64()
			continue
		}
		p.NextLine()

		switch {
		case strings.HasPrefix(p.currentLine, Begin):
			p.currentCard = &ContactCard{}

		case strings.HasPrefix(p.currentLine, END):
			return p.currentCard, nil

		case strings.HasPrefix(p.currentLine, VERSION):
			p.currentCard.Version = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, PRODID):
			p.currentCard.ProdID = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, N):
			p.parseName()

		case strings.HasPrefix(p.currentLine, FN):
			p.currentCard.FullName = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, BDAY):
			p.currentCard.Birthday, err = StringtoDateParser(strings.Split(p.currentLine, COLON)[1])
			if err != nil {
				return nil, err
			}

		case strings.HasPrefix(p.currentLine, UID):
			p.currentCard.UID = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, NICKNAME):
			p.currentCard.Nickname = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, ORG):
			p.currentCard.Organization = removeSemiColon(strings.Split(p.currentLine, COLON)[1])

		case strings.HasPrefix(p.currentLine, URL):
			p.currentCard.URL = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, NOTE):
			p.currentCard.Notes = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, TITLE):
			p.currentCard.Titles = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, PHOTO):
			p.base64Flag = true
			// check if url or base64
			p.b64BuffDaddy = []byte(p.currentLine)

		}
	}
	return nil, nil
}

/*
Takes into account all the parameters
eg. TEL;TYPE=WORK,VOICE:(111) 555-1212
Splits the field into the correct type
*/
func parseLine(currentLine string) ([]string, string, error) {
	line := strings.SplitN(currentLine, COLON, 2)
	if len(line) < 2 {
		return nil, "", errors.New("Invalid line: " + currentLine)
	}
	// Split the parameters
	params := strings.Split(line[0], SEMICOLON)
	value := line[1]
	return params, value, nil
}

func (p *Parser) parseBase64() {
	if strings.HasSuffix(p.currentLine, EQUAL) || strings.HasSuffix(p.currentLine, DUBQUAL) {
		p.currentCard.Photo = EncodedImage(string(p.b64BuffDaddy))
		p.base64Flag = false
	} else {
		p.b64BuffDaddy = append(p.b64BuffDaddy, p.currentLine...)
		p.NextLine()
		p.parseBase64()
	}
}

func (p *Parser) parseName() {
	line := strings.Split(p.currentLine, COLON)
	p.currentCard.FullName = strings.TrimSpace(strings.ReplaceAll(line[1], SEMICOLON, " "))
	names := strings.Split(line[1], SEMICOLON)
	for i, name := range names {
		switch i {
		case 0:
			if name != "" {
				p.currentCard.LastName = name
			}
		case 1:
			if name != "" {
				p.currentCard.FirstName = name
			}
		case 2:
			if name != "" {
				p.currentCard.MiddleName = name
			}
		case 3:
			if name != "" {
				p.currentCard.Prefix = name
			}
		case 4:
			if name != "" {
				p.currentCard.Suffix = name
			}
		}
	}

}

// removes trailing semicolon
// e.g. taco; -> taco
func removeSemiColon(s string) string {
	return strings.TrimRight(s, SEMICOLON)
}

// Parse the birthday string into a time.Time. (YYYY-MM-D)
func StringtoDateParser(date string) (*time.Time, error) {
	day, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	return &day, err
}
