package tmdb

const (
	Host    = "www.themoviedb.org"
	ApiHost = "api.themoviedb.org"

	MovieURL = "https://www.themoviedb.org/movie"
	PicURL   = "https://image.tmdb.org/t/p"
)

type SearchMovieRespItem struct {
	Adult            bool    `json:"adult,omitempty"`
	BackdropPath     string  `json:"backdrop_path,omitempty"`
	GenreIds         []int   `json:"genre_ids,omitempty"`
	ID               int     `json:"id,omitempty"`
	OriginalLanguage string  `json:"original_language,omitempty"`
	OriginalTitle    string  `json:"original_title,omitempty"`
	Overview         string  `json:"overview,omitempty"`
	Popularity       float64 `json:"popularity,omitempty"`
	PosterPath       string  `json:"poster_path,omitempty"`
	ReleaseDate      string  `json:"release_date,omitempty"`
	Title            string  `json:"title,omitempty"`
	Video            bool    `json:"video,omitempty"`
	VoteAverage      float64 `json:"vote_average,omitempty"`
	VoteCount        int     `json:"vote_count,omitempty"`
}

type SearchMovieResp struct {
	Page         int                   `json:"page,omitempty"`
	Results      []SearchMovieRespItem `json:"results,omitempty"`
	TotalPages   int                   `json:"total_pages,omitempty"`
	TotalResults int                   `json:"total_results,omitempty"`
}

type SearchMovieOption struct {
	Adult    bool
	Language string
	Page     int
	Region   string
	Year     int
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Company struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type MovieDetailResp struct {
	Adult               bool      `json:"adult"`
	BackdropPath        string    `json:"backdrop_path"`
	Budget              int       `json:"budget"`
	Genres              []Genre   `json:"genres"`
	Homepage            string    `json:"homepage"`
	ID                  int       `json:"id"`
	ImdbID              string    `json:"imdb_id"`
	OriginalLanguage    string    `json:"original_language"`
	OriginalTitle       string    `json:"original_title"`
	Overview            string    `json:"overview"`
	Popularity          float64   `json:"popularity"`
	PosterPath          string    `json:"poster_path"`
	ProductionCompanies []Company `json:"production_companies"`
	ReleaseDate         string    `json:"release_date"`
	Revenue             int       `json:"revenue"`
	Runtime             int       `json:"runtime"`
	Status              string    `json:"status"`
	Tagline             string    `json:"tagline"`
	Title               string    `json:"title"`
	Video               bool      `json:"video"`
	VoteAverage         float64   `json:"vote_average"`
	VoteCount           int       `json:"vote_count"`
}

type MovieDetailOption struct {
	Language string
}
