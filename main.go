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
	End       = "END"
	Colon     = ":"
	SemiColon = ";"
	EndLine   = "/n"
	COMMA     = ","
	// vCard fields
	// standard fields
	Version       = "VERSION"
	Prodid        = "PRODID"
	N             = "N"
	FN            = "FN"
	Title         = "TITLE"
	Photo         = "PHOTO"
	Logo          = "LOGO"
	Rev           = "REV"
	UID           = "UID"
	Nickname      = "NICKNAME"
	Categories    = "CATEGORIES"
	Impp          = "IMPP"
	Adr           = "ADR"
	X             = "X-"
	Email         = "EMAIL"
	SocialProfile = "X-SOCIALPROFILE"
	BDay          = "BDAY"
	Tel           = "TEL"
	Org           = "ORG"
	Type          = "TYPE"
	Url           = "URL"
	Note          = "NOTE"
	Geo           = "GEO"
	Role          = "ROLE"
	Tz            = "TZ"
	Mailer        = "MAILER"
	Label         = "LABEL"
	Agent         = "AGENT"
)

type ContactCard struct {
	Revision         time.Time
	CustomFields     map[string]string
	Birthday         *time.Time
	Version          string
	ProdID           string
	FullName         string
	FistName         string
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

// Address represents a structured address.
type Address struct {
	Type    string // Type of address (home, work, etc.)
	Street  string
	City    string
	State   string
	Zip     string
	Country string
	Label   string // Custom label (e.g., "Vacation Home")
}

// Email represents an email address with type.
type EmailAddr struct {
	Type    string
	Address string
}

// Telephone represents a telephone number with type.
type Telephone struct {
	Type   string
	Number string
}

// SocialProfile represents a social media profile.
type SocialMediaProfile struct {
	Type string
	URL  string
}

type Parser struct {
	error       error
	currentCard *ContactCard
	scanner     *bufio.Scanner
	currentLine string
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
		p.NextLine()
		switch {
		case strings.HasPrefix(p.currentLine, Begin):
			p.currentCard = &ContactCard{}
		case strings.HasPrefix(p.currentLine, End):
			return p.currentCard, nil
		case strings.HasPrefix(p.currentLine, Version):
			p.currentCard.Version = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Prodid):
			p.currentCard.ProdID = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, N):
			// last; first; middle; prefix; suffix
			p.parseName()
		case strings.HasPrefix(p.currentLine, FN):
			p.currentCard.FullName = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, BDay):
			p.currentCard.Birthday, err = StringtoDateParser(strings.Split(p.currentLine, Colon)[1])
			if err != nil {
				return nil, err
			}
		case strings.HasPrefix(p.currentLine, UID):
			p.currentCard.UID = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Nickname):
			p.currentCard.Nickname = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Org):
			p.currentCard.Organization = removeSemiColon(strings.Split(p.currentLine, Colon)[1])
		case strings.HasPrefix(p.currentLine, Url):
			p.currentCard.URL = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Note):
			p.currentCard.Notes = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Title):
			p.currentCard.Titles = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Photo):
			p.currentCard.Photos = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Logo):
			p.currentCard.Logos = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Categories):
			p.currentCard.Categories = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Impp):
			p.currentCard.InstantMessaging = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Adr):
			p.currentCard.Addresses = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Email):
			p.currentCard.Emails = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, SocialProfile):
			p.currentCard.SocialProfiles = strings.Split(p.currentLine, Colon)[1]
		case strings.HasPrefix(p.currentLine, Tel):
			p.currentCard.Telephones = map[string]string{strings.Split(p.currentLine, Colon)[1]}
		case strings.HasPrefix(p.currentLine, X):
			p.currentCard.CustomFields = strings.Split(p.currentLine, Colon)[1]
		}
	}
	return nil, nil
}

// removes trailing semicolon
// e.g. taco; -> taco
func removeSemiColon(s string) string {
	return strings.TrimRight(s, SemiColon)
}

func (p *Parser) parseName() {
	line := strings.Split(p.currentLine, Colon)
	names := strings.Split(line[1], SemiColon)
	for i, name := range names {
		switch i {
		case 0:
			if name != "" {
				p.currentCard.LastName = name
			}
		case 1:
			if name != "" {
				p.currentCard.FistName = name
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

func StringtoDateParser(date string) (*time.Time, error) {
	// Parse the birthday string
	// The format of the birthday string is YYYY-MM-DD
	day, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	return &day, err
}
