package ecbpoll

import (
	"strconv"
	"time"

	"github.com/diSpector/incrowd.git/internal/models/domain"
	"github.com/diSpector/incrowd.git/internal/models/ecb"
	"github.com/google/uuid"
)

func (s EcbPoll) convertArticle(ecbArticle ecb.Article) domain.Article {
	return domain.Article{
		Id:              s.makeId(),
		ClientId:        s.makeClientId(ecbArticle.AccountId),
		Version:         s.makeVersion(),
		Slug:            s.makeSlug(ecbArticle.TitleUrlSegment),
		PublishDate:     s.makePublishDate(ecbArticle.PublishFrom),
		Pinned:          s.makePinned(),
		SinglePage:      s.makeSinglePage(),
		HeroMedia:       s.makeHeroMedia(&ecbArticle),
		Author:          s.makeAuthor(ecbArticle.Author),
		ReadTimeMinutes: s.makeReadTimeMinutes(ecbArticle.Duration),
		Language:        s.makeLanguage(ecbArticle.Language),
		LinkedIds:       s.makeLinkedIds(ecbArticle.Related),
		Categories:      s.makeCategories(ecbArticle.References),
		DisplayCategory: s.makeDisplayCategory(&ecbArticle),
		Sponsors:        s.makeSponsors(&ecbArticle),
		Source:          s.makeSource(&ecbArticle),
		Tags:            s.makeTags(ecbArticle.Tags),
		ArticleMetadata: s.makeMetadata(ecbArticle.Metadata),
		Content:         s.makeArticleContent(&ecbArticle),
		Auth:            s.makeAuth(),
		Blocked:         s.makeBlocked(),
		LastModified:    s.makeLastModified(ecbArticle.LastModified),
		HasGeoBlocking:  s.makeHasGeoBlocking(),
		Localization:    s.makeLocalization(),
	}
}

func (s EcbPoll) makeId() string {
	return uuid.New().String()
}

func (s EcbPoll) makeClientId(accountId int) string {
	return strconv.Itoa(accountId)
}

func (s EcbPoll) makeVersion() int {
	return VERSION
}

func (s EcbPoll) makeSlug(titleUrlSegment string) string {
	return titleUrlSegment
}

func (s EcbPoll) makePublishDate(timestamp int64) time.Time {
	return time.UnixMilli(timestamp)
}

func (s EcbPoll) makePinned() bool {
	return false
}

func (s EcbPoll) makeSinglePage() bool {
	return false
}

func (s EcbPoll) makeHeroMedia(ecbArticle *ecb.Article) domain.HeroMedia {
	return domain.HeroMedia{
		Title:   ecbArticle.Title,
		Summary: ecbArticle.Summary,
		Content: s.makeContent(ecbArticle),
	}
}

// TODO::
func (s EcbPoll) makeContent(ecbArticle *ecb.Article) domain.Content {
	return domain.Content{}
}

// TODO::
func (s EcbPoll) makeArticleContent(ecbArticle *ecb.Article) []domain.ArticleContent {
	return nil
}

func (s EcbPoll) makeAuthor(author *string) *domain.Author {
	if author == nil {
		return nil
	}

	return &domain.Author{
		Name: *author,
	}
}

func (s EcbPoll) makeReadTimeMinutes(duration *int64) int {
	if duration == nil {
		return 0
	}

	return int(*duration/60 + 1)
}

func (s EcbPoll) makeLanguage(language string) string {
	return language
}

// TODO::
func (s EcbPoll) makeLinkedIds(related []ecb.Related) []domain.LinkedId {
	return nil
}

// TODO::
func (s EcbPoll) makeCategories(references []ecb.Reference) []domain.ArticleCategory {
	return nil
}

// TODO::
func (s EcbPoll) makeDisplayCategory(ecbArticle *ecb.Article) *domain.DisplayCategory {
	return nil
}

// TODO::
func (s EcbPoll) makeSponsors(ecbArticle *ecb.Article) []domain.Sponsor {
	return nil
}

func (s EcbPoll) makeSource(ecbArticle *ecb.Article) *domain.Source {
	return &domain.Source{
		SourceSystem: ECB_SOURCE,
		SourceId:     strconv.Itoa(ecbArticle.Id),
	}
}

func (s EcbPoll) makeTags(tags []ecb.Tag) []string {
	var resTags []string
	for i := range tags {
		resTags = append(resTags, tags[i].Label)
	}
	return resTags
}

func (s EcbPoll) makeMetadata(metadata map[string]interface{}) map[string]interface{} {
	return metadata
}

// TODO::
func (s EcbPoll) makeAuth() domain.Auth {
	return domain.Auth{}
}

func (s EcbPoll) makeBlocked() bool {
	return false
}

func (s EcbPoll) makeLastModified(lastModified int64) domain.LastModified {
	return domain.LastModified{
		Date: time.UnixMilli(lastModified),
	}
}

func (s EcbPoll) makeHasGeoBlocking() bool {
	return false
}

// TODO::
func (s EcbPoll) makeLocalization() domain.Localization {
	return domain.Localization{}
}
