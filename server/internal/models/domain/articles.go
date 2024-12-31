package domain

import (
	"time"
)

type ResponseMulti struct {
	Status   string    `json:"status"`
	Data     DataMulti `json:"data,omitempty"`
	Message  string    `json:"message,omitempty"`
	Metadata Metadata  `json:"metadata"`
}

type DataMulti struct {
	Articles   []Article           `json:"articles,omitempty"`
	Categories map[string]Category `json:"categories,omitempty"`
}

type ResponseOne struct {
	Status   string   `json:"status,omitempty"`
	Data     Data     `json:"data,omitempty"`
	Message  string   `json:"message,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}

type Data struct {
	Article    *Article            `json:"article,omitempty"`
	Categories map[string]Category `json:"categories,omitempty"`
}

type ArticleOriginMod struct {
	Id           string       `json:"id" bson:"id"`
	Source       Source       `json:"source" bson:"source"`
	LastModified LastModified `json:"lastModified" bson:"lastModified"`
}

type Article struct {
	Id              string                 `json:"id" bson:"id"`
	ClientId        string                 `json:"clientId" bson:"clientId"`
	Version         int                    `json:"version" bson:"version"`
	Slug            string                 `json:"slug" bson:"slug"`
	PublishDate     time.Time              `json:"publishDate" bson:"publishDate"`
	Pinned          bool                   `json:"pinned" bson:"pinned"`
	SinglePage      bool                   `json:"singlePage" bson:"singlePage"`
	HeroMedia       HeroMedia              `json:"heroMedia" bson:"heroMedia"`
	Author          *Author                `json:"author" bson:"author"`
	ReadTimeMinutes int                    `json:"readTimeMinutes" bson:"readTimeMinutes"`
	Language        string                 `json:"language" bson:"language"`
	LinkedIds       []LinkedId             `json:"linkedIds" bson:"linkedIds"`
	Categories      []ArticleCategory      `json:"categories" bson:"categories"`
	DisplayCategory *DisplayCategory       `json:"displayCategory" bson:"displayCategory"`
	Sponsors        []Sponsor              `json:"sponsors" bson:"sponsors"`
	Source          *Source                `json:"source,omitempty" bson:"source,omitempty"`
	Tags            []string               `json:"tags" bson:"tags"`
	ArticleMetadata map[string]interface{} `json:"metadata" bson:"metadata"`
	Content         []ArticleContent       `json:"content" bson:"content"`
	Auth            Auth                   `json:"auth" bson:"auth"`
	Blocked         bool                   `json:"blocked" bson:"blocked"`
	LastModified    LastModified           `json:"lastModified" bson:"lastModified"`
	HasGeoBlocking  bool                   `json:"hasGeoBlocking" bson:"hasGeoBlocking"`
	Localization    Localization           `json:"localization" bson:"localization"`
}

type Origin struct {
	Source string `json:"source,omitempty" bson:"source,omitempty"`
	Id     string `json:"id,omitempty" bson:"id,omitempty"`
}

type HeroMedia struct {
	Title   string  `json:"title" bson:"title"`
	Summary *string `json:"summary" bson:"summary"`
	Content Content `json:"content" bson:"content"`
}

type Content struct {
	Id          string  `json:"id" bson:"id"`
	ContentType string  `json:"contentType" bson:"contentType"`
	Image       string  `json:"image" bson:"image"`
	AltText     *string `json:"altText" bson:"altText"`
}

type Author struct {
	Name     string `json:"name" bson:"name"`
	ImageURL string `json:"imageUrl" bson:"imageUrl"`
}

type LinkedId struct {
	Text           *string `json:"text" bson:"text"`
	SourceSystem   string  `json:"sourceSystem" bson:"sourceSystem"`
	SourceSystemId string  `json:"sourceSystemId" bson:"sourceSystemId"`
}

type Category struct {
	Id       string  `json:"id" bson:"id"`
	ParentId *string `json:"parentId" bson:"parentId"`
	ClientId string  `json:"clientId" bson:"clientId"`
	Slug     string  `json:"slug" bson:"slug"`
	Text     string  `json:"text" bson:"text"`
	After    *string `json:"after" bson:"after"`
}

type ArticleCategory struct {
	Id   string `json:"id" bson:"id"`
	Text string `json:"text" bson:"text"`
}

type DisplayCategory struct {
	Id   string `json:"id" bson:"id"`
	Text string `json:"text" bson:"text"`
}

type Sponsor struct {
	Text     *string `json:"text" bson:"text"`
	ImageUrl string  `json:"imageUrl" bson:"imageUrl"`
	LinkUrl  string  `json:"linkUrl" bson:"linkUrl"`
}

type Source struct {
	SourceSystem string `json:"sourceSystem,omitempty" bson:"sourceSystem,omitempty"`
	SourceId     string `json:"sourceId,omitempty" bson:"sourceId,omitempty"`
}

type ArticleContent struct {
	Id                string                 `json:"id" bson:"id"`
	ContentType       string                 `json:"contentType" bson:"contentType"`
	Content           *string                `json:"content" bson:"content"`
	CustomContentType *string                `json:"customContentType" bson:"customContentType"`
	CustomContent     map[string]interface{} `json:"customContent" bson:"customContent"`
	IsHtml            *bool                  `json:"isHtml" bson:"isHtml"`
	Children          []ArticleContent       `json:"children" bson:"children"`
	Appearance        *Appearance            `json:"appearance" bson:"appearance"`
	Text              *string                `json:"text" bson:"text"`
	Author            *string                `json:"author" bson:"author"`
	Link              *string                `json:"link" bson:"link"`
	VideoThumbnail    *string                `json:"videoThumbnail" bson:"videoThumbnail"`
	SourceSystemId    *string                `json:"sourceSystemId" bson:"sourceSystemId"`
	Image             *string                `json:"image" bson:"image"`
	Sponsor           *Sponsor               `json:"sponsor" bson:"sponsor"`
}

type Appearance struct {
	Type string `json:"type" bson:"type"`
}

type Auth struct {
	LoginRequired   bool     `json:"loginRequired" bson:"loginRequired"`
	Roles           []string `json:"roles" bson:"roles"`
	Entitlements    []string `json:"entitlements" bson:"entitlements"`
	RestrictionType string   `json:"restrictionType" bson:"restrictionType"`
}

type LastModified struct {
	Date time.Time `json:"date" bson:"date"`
}

type Localization struct {
	Id string `json:"id" bson:"id"`
}

type Metadata struct {
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	PageItems  *int64    `json:"pageItems,omitempty" bson:"pageItems,omitempty"`
	TotalItems *int64    `json:"totalItems,omitempty" bson:"totalItems,omitempty"`
	TotalPages *int64    `json:"totalPages,omitempty" bson:"totalPages,omitempty"`
	PageNumber *int64    `json:"pageNumber,omitempty" bson:"pageNumber,omitempty"`
	PageSize   *int64    `json:"pageSize,omitempty" bson:"pageSize,omitempty"`
	Sort       *string   `json:"sort,omitempty" bson:"sort,omitempty"`
}
