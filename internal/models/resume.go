package models

type Resume struct {
	Name       string       `json:"name"`
	Address    string       `json:"address"`
	Phone      string       `json:"phone"`
	Email      string       `json:"email"`
	Experience []Experience `json:"experience"`
	Education  []Education  `json:"education"`
	Skills     []Skill      `json:"skills"`
	Projects   []Project    `json:"projects"`
}

func NewResume() *Resume {
	return &Resume{}
}
func FromJsonFile(path string) (*Resume, error) {
	return nil, nil
}

func (r *Resume) SaveToFile(path string) error {
	return nil
}
