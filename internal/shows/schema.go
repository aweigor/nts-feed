package shows

import "time"

type Resultset struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type Metadata struct {
	Resultset Resultset `json:"resultset"`
}

type Genre struct {
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

type AudioSource struct {
	URL    string `json:"url"`
	Source string `json:"source"`
}

type TracklistItem struct {
	Artist             string  `json:"artist"`
	Title              string  `json:"title"`
	UID                *string `json:"uid"`
	Offset             *int    `json:"offset"`
	Duration           *int    `json:"duration"`
	OffsetEstimate     *int    `json:"offset_estimate"`
	DurationEstimate   *int    `json:"duration_estimate"`
	ACRID              *string `json:"acr_id"`
	DeezerTrackID      *int    `json:"deezer_track_id"`
	ISRCID             *string `json:"isrc_id"`
	MusicbrainzTrackID *string `json:"musicbrainz_track_id"`
}

type Tracklist struct {
	Metadata Metadata        `json:"metadata"`
	Results  []TracklistItem `json:"results"`
}

type Embeds struct {
	Tracklist Tracklist `json:"tracklist"`
}

type EpisodeInfoResponse struct {
	Status          string        `json:"status"`
	Updated         time.Time     `json:"updated"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	DescriptionHTML string        `json:"description_html"`
	ExternalLinks   []interface{} `json:"external_links"`
	Moods           []interface{} `json:"moods"`
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
	Brand           interface{}   `json:"brand"`
	Embeds          Embeds        `json:"embeds"`
}

type EpisodeTracklistResponse struct {
	Results []TracklistItem `json:"results"`
}
