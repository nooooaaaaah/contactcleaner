package contact

import (
	"time"
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
