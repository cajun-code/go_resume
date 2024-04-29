package models

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Show        bool   `json:"show"`
}

type Experience struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type Education struct {
	Degree      string `json:"degree"`
	Major       string `json:"major"`
	Institution string `json:"institution"`
	Location    string `json:"location"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type Skill struct {
	Name     string `json:"name"`
	Level    string `json:"level"`
	Years    int    `json:"years"`
	Category string `json:"category"`
	Show     bool   `json:"show"`
}
