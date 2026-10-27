package model

// Resume is the structured source data used to render the resume.
type Resume struct {
	Contact     Contact      `json:"contact" yaml:"contact"`
	Employments []Employment `json:"employments" yaml:"employments"`
	Skills      []string     `json:"skills" yaml:"skills"`
	Education   Education    `json:"education" yaml:"education"`
}

type Contact struct {
	FirstName string  `json:"first_name" yaml:"first_name"`
	LastName  string  `json:"last_name" yaml:"last_name"`
	Title     string  `json:"title" yaml:"title"`
	Email     string  `json:"email" yaml:"email"`
	Phone     string  `json:"phone" yaml:"phone"`
	Links     []Link  `json:"links" yaml:"links"`
	Address   Address `json:"address" yaml:"address"`
}

type Link struct {
	Name string `json:"name" yaml:"name"`
	Href string `json:"href" yaml:"href"`
}

type Address struct {
	Line1   string `json:"line1" yaml:"line1"`
	Line2   string `json:"line2" yaml:"line2"`
	City    string `json:"city" yaml:"city"`
	State   string `json:"state" yaml:"state"`
	ZipCode string `json:"zipcode" yaml:"zipcode"`
}

type Employment struct {
	Name       string   `json:"name" yaml:"name"`
	Title      string   `json:"title" yaml:"title"`
	Years      string   `json:"years" yaml:"years"`
	City       string   `json:"city" yaml:"city"`
	State      string   `json:"state" yaml:"state"`
	Keywords   []string `json:"keywords" yaml:"keywords"`
	Experience []string `json:"experience" yaml:"experience"`
}

type Education struct {
	Institution string `json:"institution" yaml:"institution"`
	Degree      string `json:"degree" yaml:"degree"`
	Major       string `json:"major" yaml:"major"`
	Minor       string `json:"minor" yaml:"minor"`
	Years       string `json:"years" yaml:"years"`
	Status      string `json:"status" yaml:"status"`
}
