package vcard

type ParVal string
type ParName string

// represents a vCard parameter
// https://tools.ietf.org/html/rfc6350#section-5.3
type Par struct {
	Name    ParName
	PosVals []ParVal
	Val     []string
}
