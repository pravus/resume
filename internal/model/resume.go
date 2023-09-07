package model

type Resume struct {
	Contact struct {
		FirstName string `json:"first_name" yaml:"first_name"`
		LastName  string `json:"last_name"  yaml:"last_name"`
		Title     string `json:"title"      yaml:"title"`
		Email     string `json:"email"      yaml:"email"`
		Phone     string `json:"phone"      yaml:"phone"`
		Links     []struct {
			Name string `json:"name" yaml:"name"`
			Href string `json:"href" yaml:"href"`
		} `json:"links" yaml:"links"`
		Address struct {
			Line1   string `json:"line1"   yaml:"line1"`
			Line2   string `json:"line2"   yaml:"line2"`
			City    string `json:"city"    yaml:"city"`
			State   string `json:"state"   yaml:"state"`
			ZipCode string `json:"zipcode" yaml:"zipcode"`
		} `json:"address" yaml:"address"`
	} `json:"contact" yaml:"contact"`
	Employments []struct {
		Name       string   `json:"name"       yaml:"name"`
		Title      string   `json:"title"      yaml:"title"`
		Years      string   `json:"years"      yaml:"years"`
		City       string   `json:"city"       yaml:"city"`
		State      string   `json:"state"      yaml:"state"`
		Keywords   []string `json:"keywords"   yaml:"keywords"`
		Experience []string `json:"experience" yaml:"experience"`
	} `json:"employments" yaml:"employments"`
	Skills    []string `json:"skills" yaml:"skills"`
	Education struct {
		Institution string `json:"institution" yaml:"institution"`
		Degree      string `json:"degree"      yaml:"degree"`
		Major       string `json:"major"       yaml:"major"`
		Minor       string `json:"minor"       yaml:"minor"`
		Years       string `json:"years"       yaml:"years"`
		Status      string `json:"status"      yaml:"status"`
	} `json:"education" yaml:"education"`
}
