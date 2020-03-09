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

//
// Response Data
//

// MemeGetResponse holds the data for a GET memes response from api.imgflip.com
type MemeGetResponse struct {
	Success bool     `json:"success"`
	Data    MemeData `json:"data"`
}

// MemeData defines the collection of memes
type MemeData struct {
	Memes []Meme `json:"memes"`
}

// Meme represents the details of an individual meme
type Meme struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	BoxCount int    `json:"box_count"`
}

// MemePostResponse handles the small response information
// returned from the POST call to api.imgflip.com
type MemePostResponse struct {
	Success bool        `json:"success"`
	Data    SuccessData `json:"data"`
}

// SuccessData holds teh links to the meme and the image after a successful call to
// POST api.imgflip.com
type SuccessData struct {
	URL     string `json:"url"`
	PageURL string `json:"page_url"`
}

// MemeError describes the error response received from api.imgflip.com
type MemeError struct {
	Success bool   `json:"success"`
	Data    string `json:"error_message"`
}
