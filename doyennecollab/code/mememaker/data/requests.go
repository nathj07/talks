package data

//
// Request Data
//

// MemePostRequest defines the body of a POST request to use when making a meme
type MemePostRequest struct {
	TemplateID  int       `json:"template_id"`
	Username    string    `json:"username"`
	Password    string    `jsone:"password"`
	Text0       string    `json:"text0"`
	Text1       string    `json:"text1"`
	Font        string    `json:"font,omitempty"`         // api only supports "impact" or "arial"
	MaxFontSize string    `json:"max_fon_size,omitempty"` // default to 50px
	Boxes       []MemeBox `json:"boxes,omitempty"`
}

// MemeBox defines a box for text on a Meme. This is used in a POST request
// when more than 2 text boxes. Using this means text0 and text1 are ignored.
// The caller is responsible for the capitalization of any text in a text box defined this way.
// There are a maximum of 5 text boxes supported by the API
type MemeBox struct {
	Text         string `json:"text"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Color        string `json:"color"`         // the api wants a hex values for this
	OutlineColor string `json:"outline-color"` // the api wants a hex value here too
}
