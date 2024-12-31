package ecb

import (
	"time"
)

type Articles struct {
	PageInfo PageInfo  `json:"pageInfo"`
	Content  []Article `json:"content"`
}

type PageInfo struct {
	Page       int `json:"page"`
	NumPages   int `json:"numPages"`
	PageSize   int `json:"pageSize"`
	NumEntries int `json:"numEntries"`
}

type Article struct {
	Id                int                    `json:"id"`
	AccountId         int                    `json:"accountId"`
	Type              string                 `json:"type"`
	Title             string                 `json:"title"`
	Description       *string                `json:"description"`
	Date              time.Time              `json:"date"`
	Location          *string                `json:"location"`
	Coordinates       []float64              `json:"coordinates"`
	CommentsOn        bool                   `json:"commentsOn"`
	Copyright         *string                `json:"copyright"`
	PublishFrom       int64                  `json:"publishFrom"`
	PublishTo         int64                  `json:"publishTo"`
	Tags              []Tag                  `json:"tags"`
	Platform          string                 `json:"platform"`
	Language          string                 `json:"language"`
	AdditionalInfo    map[string]interface{} `json:"additionalInfo"`
	CanonicalUrl      string                 `json:"canonicalUrl"`
	References        []Reference            `json:"references"`
	Related           []Related              `json:"related"`
	Metadata          map[string]interface{} `json:"metadata"`
	TitleTranslations *string                `json:"titleTranslations"`
	LastModified      int64                  `json:"lastModified"`
	TitleUrlSegment   string                 `json:"titleUrlSegment"`
	Body              *string                `json:"body,omitempty"`
	Author            *string                `json:"author"`
	Subtitle          *string                `json:"subtitle"`
	Variants          []Variant              `json:"variants,omitempty"`
	Summary           *string                `json:"summary,omitempty"`
	HotlinkUrl        *string                `json:"hotlinkUrl"`
	Duration          *int64                 `json:"duration,omitempty"`
	ContentSummary    *string                `json:"contentSummary"`
	LeadMedia         *Article               `json:"leadMedia,omitempty"`
	ImageURL          string                 `json:"imageUrl"`
	OnDemandURL       *string                `json:"onDemandUrl"`
	OriginalDetails   *OriginalDetails       `json:"originalDetails,omitempty"`
}

type Reference struct {
	Label *string `json:"label"`
	Id    int     `json:"id"`
	Type  string  `json:"type"`
	Sid   string  `json:"sid"`
}

type Related struct {
	Label string `json:"label"`
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Sid   string `json:"sid"`
}

type Variant struct {
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	Url    string `json:"url"`
	Tag    Tag    `json:"tag"`
}

type OriginalDetails struct {
	Width       int64   `json:"width"`
	Height      int64   `json:"height"`
	AspectRatio float64 `json:"aspectRatio"`
}

type Tag struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}
