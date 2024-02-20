package parsing

import (
	"ContactCleaner/contact"
	"ContactCleaner/vcard"
	"bufio"
	"errors"
	"strings"
	"time"
)

type Parser struct {
	error        error
	currentCard  *contact.ContactCard
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

func (p *Parser) Parse() (*contact.ContactCard, error) {
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
		case strings.HasPrefix(p.currentLine, vcard.BEGIN):
			p.currentCard = &contact.ContactCard{}

		case strings.HasPrefix(p.currentLine, vcard.END):
			return p.currentCard, nil

		case strings.HasPrefix(p.currentLine, vcard.VERSION):
			p.currentCard.Version = strings.Split(p.currentLine, vcard.COLON)[1]

		case strings.HasPrefix(p.currentLine, vcard.PRODID):
			p.currentCard.ProdID = strings.Split(p.currentLine, vcard.COLON)[1]

		case strings.HasPrefix(p.currentLine, vcard.N):
			p.parseName()

		case strings.HasPrefix(p.currentLine, vcard.FN):
			p.currentCard.FullName = strings.Split(p.currentLine, vcard.COLON)[1]

		case strings.HasPrefix(p.currentLine, vcard.BDAY):
			p.currentCard.Birthday, err = StringtoDateParser(strings.Split(p.currentLine, vcard.COLON)[1])
			if err != nil {
				return nil, err
			}

		case strings.HasPrefix(p.currentLine, vcard.UID):
			p.currentCard.UID = strings.Split(p.currentLine, vcard.COLON)[1]

		case strings.HasPrefix(p.currentLine, vcard.NICKNAME):
			p.currentCard.Nickname = strings.Split(p.currentLine, vcard.COLON)[1]

		case strings.HasPrefix(p.currentLine, vcard.ORG):
			p.currentCard.Organization = removeSemiColon(strings.Split(p.currentLine, vcard.COLON)[1])

		case strings.HasPrefix(p.currentLine, vcard.URL):
			p.currentCard.URL = strings.Split(p.currentLine, vcard.COLON)[1]

		case strings.HasPrefix(p.currentLine, vcard.NOTE):
			p.currentCard.Notes = strings.Split(p.currentLine, vcard.COLON)[1]

		case strings.HasPrefix(p.currentLine, vcard.TITLE):
			p.currentCard.Titles = strings.Split(p.currentLine, vcard.COLON)[1]

		case strings.HasPrefix(p.currentLine, vcard.PHOTO):
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
	line := strings.SplitN(currentLine, vcard.COLON, 2)
	if len(line) < 2 {
		return nil, "", errors.New("Invalid line: " + currentLine)
	}
	// Split the parameters
	params := strings.Split(line[0], vcard.SEMICOLON)
	value := line[1]
	return params, value, nil
}

func (p *Parser) parseBase64() {
	if strings.HasSuffix(p.currentLine, vcard.EQUAL) || strings.HasSuffix(p.currentLine, vcard.DUBQUAL) {
		p.currentCard.Photo = contact.EncodedImage(string(p.b64BuffDaddy))
		p.base64Flag = false
	} else {
		p.b64BuffDaddy = append(p.b64BuffDaddy, p.currentLine...)
		p.NextLine()
		p.parseBase64()
	}
}

func (p *Parser) parseName() {
	line := strings.Split(p.currentLine, vcard.COLON)
	p.currentCard.FullName = strings.TrimSpace(strings.ReplaceAll(line[1], vcard.SEMICOLON, " "))
	names := strings.Split(line[1], vcard.SEMICOLON)
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
	return strings.TrimRight(s, vcard.SEMICOLON)
}

// Parse the birthday string into a time.Time. (YYYY-MM-D)
func StringtoDateParser(date string) (*time.Time, error) {
	day, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	return &day, err
}
