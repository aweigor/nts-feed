package feeds

import "github.com/aweigor/nts-feed-generator/pkg/ntsclient"

type FeedsQueryParams struct {
	format string
	offset int
	limit  int
}

// RSSResponse represents the root of the NTS radio feed.
type RSSResponse struct {
	Channel Channel `json:"channel"`
}

// Channel represents a single NTS radio channel.
type Channel struct {
	Title         string        `json:"title"`
	Link          string        `json:"link"`
	Description   string        `json:"description"`
	Category      string        `json:"category"`
	Image         string        `json:"image"`
	LastBuildDate string        `json:"lastBuildDate"`
	Items         []ChannelItem `json:"item"`
}

// Item represents a single episode of an NTS radio show.
type ChannelItem struct {
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	PubDate     string    `json:"pubDate"`
	Enclosure   Enclosure `json:"enclosure"`
	GUID        string    `json:"guid"`
	ShowID      string    `json:"nts:show_id"`
	EpisodeID   string    `json:"nts:episode_id"`
	AirDate     string    `json:"nts:air_date"`
	Duration    *int      `json:"nts:duration"`
	Tracklist   []Track   `json:"tracklist"`
}

// Enclosure describes the audio enclosure for the episode.
type Enclosure struct {
	URL    string `json:"url"`
	Type   string `json:"type"`
	Length *int   `json:"length"`
}

type Track = ntsclient.Track
