package contactcleaner

import (
	// Import necessary packages

	"bufio"
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
	SOCIALPROFILE = "X-SOCIALPROFILE"
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
	Photo            []byte
	Logos            []byte
	Categories       []string // tags
	InstantMessaging []string
	Addresses        []Address
	Emails           []EmailAddr
	SocialProfiles   []SocialMediaProfile
	Telephones       []Telephone
	Items            []Item
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
	Type   string
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

		if p.base64Flag {

			if strings.HasSuffix(p.currentLine, EQUAL) || strings.HasSuffix(p.currentLine, DUBQUAL) {
				p.currentCard.Photo = append(p.currentCard.Photo, p.b64BuffDaddy...)
				p.base64Flag = false
			} else {
				p.b64BuffDaddy = append(p.b64BuffDaddy, p.currentLine...)
			}
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
			// last; first; middle; prefix; suffix
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
			p.b64BuffDaddy = []byte(p.currentLine)

		case strings.HasPrefix(p.currentLine, LOGO):
			p.currentCard.Logos = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, CATEGORIES):
			p.currentCard.Categories = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, IMPP):
			p.currentCard.InstantMessaging = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, ADR):
			p.currentCard.Addresses = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, EMAIL):
			p.currentCard.Emails = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, SOCIALPROFILE):
			p.currentCard.SocialProfiles = strings.Split(p.currentLine, COLON)[1]

		case strings.HasPrefix(p.currentLine, TEL):
			telephone := parseTelephone(p.currentLine)
			p.currentCard.Telephones = append(p.currentCard.Telephones, telephone)

		case strings.HasPrefix(p.currentLine, X):
			p.currentCard.CustomFields = strings.Split(p.currentLine, COLON)[1]
		}
	}
	return nil, nil
}

func (p *Parser) parseName() {
	line := strings.Split(p.currentLine, COLON)
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
