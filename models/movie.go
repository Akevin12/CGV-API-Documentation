package models

type Movie struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Age         int    `json:"age"`
	Duration    int    `json:"duration"`
	Review      Review `json:"review"`
	Date        string `json:"Date"`
}

type MovieSchedules struct {
	Schedules []Schedule `json:"data"`
}

type MovieSchedulesResponse struct {
	Response
	MovieSchedules
}

type MovieSchedule struct {
	Schedule Schedule `json:"data"`
}

type MovieScheduleResponse struct {
	Response
	MovieSchedule
}

type MoviesResponse struct {
	Response
	Movies []Movie `json:"data"`
}
type MovieResponse struct {
	Response
	Movie Movie `json:"data"`
}
