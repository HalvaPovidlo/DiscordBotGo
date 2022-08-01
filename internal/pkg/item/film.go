package item

import (
	"time"
)

type SortCode int8

const (
	RandomSort    SortCode = 0
	TitleSort     SortCode = 1
	TitleBackSort SortCode = 2
	ScoreSort     SortCode = 3
	ScoreBackSort SortCode = 4
)

type Film struct {
	ID                       string             `firestore:"-" json:"id"`
	Title                    string             `firestore:"title,omitempty" json:"title"`
	TitleOriginal            string             `firestore:"title_original,omitempty" json:"title_original,omitempty"`
	Poster                   string             `firestore:"cover,omitempty" json:"cover,omitempty"`
	Cover                    string             `firestore:"poster,omitempty" json:"poster,omitempty"`
	Director                 string             `firestore:"director,omitempty" json:"director,omitempty"`
	Description              string             `firestore:"description,omitempty" json:"description,omitempty"`
	Duration                 string             `firestore:"duration,omitempty" json:"duration,omitempty"`
	Score                    int                `firestore:"score" json:"score"`
	UserScore                *int               `firestore:"user_score" json:"user_score,omitempty"`
	Average                  float64            `firestore:"average,omitempty" json:"average,omitempty"`
	Scores                   map[string]int     `firestore:"scores" json:"scores,omitempty"`
	Comments                 map[string]Comment `firestore:"-" json:"comments,omitempty"`
	WithComments             bool               `firestore:"-" json:"-"`
	URL                      string             `firestore:"kinopoisk,omitempty" json:"kinopoisk,omitempty"`
	RatingKinopoisk          float64            `firestore:"rating_kinopoisk,omitempty" json:"rating_kinopoisk,omitempty"`
	RatingKinopoiskVoteCount int                `firestore:"rating_kinopoisk_vote_count,omitempty" json:"rating_kinopoisk_vote_count,omitempty"`
	RatingImdb               float64            `firestore:"rating_imdb,omitempty" json:"rating_imdb,omitempty"`
	RatingImdbVoteCount      int                `firestore:"rating_imdb_vote_count,omitempty" json:"rating_imdb_vote_count,omitempty"`
	Year                     int                `firestore:"year,omitempty" json:"year,omitempty"`
	FilmLength               int                `firestore:"film_length,omitempty" json:"film_length,omitempty"`
	Serial                   bool               `firestore:"serial" json:"serial"`
	ShortFilm                bool               `firestore:"short_film" json:"short_film"`
	Genres                   []string           `firestore:"genres,omitempty" json:"genres,omitempty"`
}

type Comment struct {
	UserID    string    `firestore:"user_id" json:"user_id"`
	Text      string    `firestore:"text" json:"text"`
	CreatedAt time.Time `firestore:"created_at" json:"created_at"`
}

type FireScore struct {
	Score int `firestore:"score"`
}