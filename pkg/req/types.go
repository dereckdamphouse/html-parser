package req

type Property struct {
	Name      string
	Selector  string
	Attribute string
}

type Data struct {
	HTML       string
	Properties []Property
}
