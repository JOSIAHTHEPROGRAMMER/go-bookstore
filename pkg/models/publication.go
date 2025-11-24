package models

type Publication struct {
	ID uint `gorm:"primaryKey" json:"id"`

	BookID      uint `json:"book_id"`
	AuthorID    uint `json:"author_id"`
	PublisherID uint `json:"publisher_id"`

	PublishedDate string `json:"published_date"`
	Edition       string `json:"edition"`
	ISBN          string `json:"isbn"`
}
