package vcard

type vCard struct {
	Properties map[PropName]Property
}

func NewVCard() *vCard {
	return &vCard{
		Properties: make(map[PropName]Property),
	}
}

func (v *vCard) AddProperty(prop Property) {
	v.Properties[prop.Name] = prop
}

func (v *vCard) GetProperty(propName PropName) (Property, bool) {
	prop, ok := v.Properties[propName]
	return prop, ok
}
