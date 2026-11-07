package ntsclient

import "time"

// Response is the top-level envelope returned by the API.
type ArticlesResponse struct {
	Metadata Metadata      `json:"metadata"`
	Results  []ArticleInfo `json:"results"`
}

// Metadata wraps the pagination info for a response.
type Metadata struct {
	Resultset Resultset `json:"resultset"`
}

// Resultset describes pagination over the full result set.
type Resultset struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// Result is a single item in the results list.
type ArticleInfo struct {
	ArticleType    string          `json:"article_type"`
	Title          string          `json:"title"`
	Artists        []Artist        `json:"artists"`
	Article        Article         `json:"article"`
	AudioSources   []AudioSource   `json:"audio_sources"`
	Description    Description     `json:"description"`
	Image          Image           `json:"image"`
	RelatedEpisode *ArticleEpisode `json:"related_episode,omitempty"`
	LocalDate      string          `json:"local_date"`
	Location       string          `json:"location"`
	Genres         []Genre         `json:"genres"`
	Moods          []Mood          `json:"moods"`
	TrackUID       *string         `json:"track_uid"`
	Brand          *Brand          `json:"brand,omitempty"`
}

// Artist is a performer or contributor on a result.
type Artist struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Article points to the canonical page for a result.
type Article struct {
	Path string `json:"path"`
}

// AudioSource is a playable source for a result.
type AudioSource struct {
	URL    string `json:"url"`
	Source string `json:"source"`
}

// Description is currently an empty object in the API but is modeled
// as a map so future keys do not break decoding.
type Description struct {
	// Fields are currently unknown; keep as a generic map.
	Fields map[string]any `json:"-"`
}

// Image holds the various resized image URLs for a result.
type Image struct {
	Large       string `json:"large"`
	MediumLarge string `json:"medium_large"`
	Medium      string `json:"medium"`
	Small       string `json:"small"`
	Thumb       string `json:"thumb"`
}

// ArticleEpisode is a related episode reference.
type ArticleEpisode struct {
	Path string `json:"path"`
}

// Genre is a musical genre tag.
type Genre struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Mood is a mood tag.
type Mood struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Brand describes the promotional brand block attached to some results.
type Brand struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CTAURL      string `json:"cta_url"`
	CTALabel    string `json:"cta_label"`
}

type EpisodesResponse struct {
	Metadata Metadata      `json:"metadata"`
	Results  []EpisodeInfo `json:"results"`
}

// Same as Episode - need to merge (:todo)
type EpisodeInfo struct {
	Status          string        `json:"status"`
	Updated         string        `json:"updated"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	DescriptionHTML string        `json:"description_html"`
	ExternalLinks   []string      `json:"external_links"`
	Moods           []Taxonomy    `json:"moods"`
	Genres          []Taxonomy    `json:"genres"`
	LocationShort   string        `json:"location_short"`
	LocationLong    string        `json:"location_long"`
	Intensity       string        `json:"intensity"`
	Media           Media         `json:"media"`
	EpisodeAlias    string        `json:"episode_alias"`
	ShowAlias       string        `json:"show_alias"`
	Broadcast       string        `json:"broadcast"`
	Mixcloud        string        `json:"mixcloud"`
	AudioSources    []AudioSource `json:"audio_sources"`
	Brand           struct{}      `json:"brand"`  // {} in data
	Embeds          struct{}      `json:"embeds"` // {} in data
	Links           []Link        `json:"links"`
}

// Taxonomy covers both `moods` and `genres` — same shape:
// { "id": "...", "value": "..." }.
type Taxonomy struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type Media struct {
	BackgroundLarge       string `json:"background_large"`
	BackgroundMediumLarge string `json:"background_medium_large"`
	BackgroundMedium      string `json:"background_medium"`
	BackgroundSmall       string `json:"background_small"`
	BackgroundThumb       string `json:"background_thumb"`
	PictureLarge          string `json:"picture_large"`
	PictureMediumLarge    string `json:"picture_medium_large"`
	PictureMedium         string `json:"picture_medium"`
	PictureSmall          string `json:"picture_small"`
	PictureThumb          string `json:"picture_thumb"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
	Type string `json:"type"`
}

type Episode struct {
	Status          string        `json:"status"`
	Updated         time.Time     `json:"updated"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	DescriptionHTML string        `json:"description_html"`
	ExternalLinks   []string      `json:"external_links"`
	Moods           []Mood        `json:"moods"`
	Genres          []Genre       `json:"genres"`
	LocationShort   string        `json:"location_short"`
	LocationLong    string        `json:"location_long"`
	Intensity       string        `json:"intensity"`
	Media           Media         `json:"media"`
	EpisodeAlias    string        `json:"episode_alias"`
	ShowAlias       string        `json:"show_alias"`
	Broadcast       time.Time     `json:"broadcast"`
	Mixcloud        string        `json:"mixcloud"`
	AudioSources    []AudioSource `json:"audio_sources"`
	Brand           Brand         `json:"brand"`
	Embeds          EpisodeEmbeds `json:"embeds"`
	Links           []Link        `json:"links"`
}

type EpisodeEmbeds struct {
	Tracklist Tracklist `json:"tracklist"`
}

type Tracklist struct {
	Metadata TracklistMetadata `json:"metadata"`
	Results  []Track           `json:"results"`
	Links    []Link            `json:"links"`
}

type TracklistMetadata struct {
	Resultset Resultset `json:"resultset"`
}

type Track struct {
	Artist             string  `json:"artist"`
	Title              string  `json:"title"`
	UID                *string `json:"uid"`
	Offset             *int    `json:"offset"`
	Duration           *int    `json:"duration"`
	OffsetEstimate     *int    `json:"offset_estimate"`
	DurationEstimate   *int    `json:"duration_estimate"`
	ACRID              *string `json:"acr_id"`
	DeezerTrackID      *int64  `json:"deezer_track_id"`
	ISRCID             *string `json:"isrc_id"`
	MusicBrainzTrackID *string `json:"musicbrainz_track_id"`
}

type EpisodeInfoResponse = Episode
