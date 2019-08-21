package statuspage

// statuspage.io data structure

type Res struct {
	Page       Page        `json:"page"`
	Components []Component `json:"components"`
	Status     Status      `json:"status"`
}

type Page struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	TimeZone  string `json:"time_zone"`
	UpdatedAt string `json:"updated_at"`
}

type Component struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	Position           int    `json:"position"`
	Description        string `json:"description"`
	ShowCase           bool   `json:"show_case"`
	GroupId            string `json:"group_id"`
	Group              bool   `json:"group"`
	OnlyShowIfDegraded bool   `json:"only_show_if_degraded"`
}

type Status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}
