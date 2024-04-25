package models

type Movie struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	Age         int    `json:"age"`
	Duration    int    `json:"duration"`
	Review      Review `json:"review"`
	Date        string `json:"Date"`
}
type NowShowingMovies struct {
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title"`
	Location string `json:"location"`
	Age      int    `json:"age"`
}
type UpComingMovies struct {
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title"`
	Likes    int    `json:"likes"`
	Location string `json:"location"`
	Age      int    `json:"age"`
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
type NowShowingMoviesResponse struct {
	Response
	NowShowingMovies []NowShowingMovies `json:"data"`
}
type NowShowingMovieResponse struct {
	Response
	NowShowingMovie NowShowingMovies `json:"data"`
}
type UpComingMoviesResponse struct {
	Response
	UpComingMovies []UpComingMovies `json:"data"`
}
type UpComingMovieResponse struct {
	Response
	UpComingMovie UpComingMovies `json:"data"`
}
