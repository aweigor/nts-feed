package feeds

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

// Track represents one track within an episode's tracklist.
type Track struct {
	Artist             string  `json:"artist"`
	Title              string  `json:"title"`
	UID                *string `json:"uid"`                  // *string: UUID or null
	Offset             *int    `json:"offset"`               // *int: integer or null
	Duration           *int    `json:"duration"`             // *int: integer or null
	OffsetEstimate     *int    `json:"offset_estimate"`      // *int: integer or null
	DurationEstimate   *int    `json:"duration_estimate"`    // *int: integer or null
	ACRID              *string `json:"acr_id"`               // *string: string or null
	DeezerTrackID      *int    `json:"deezer_track_id"`      // *int: integer or null
	ISRCID             *string `json:"isrc_id"`              // *string: string or null
	MusicBrainzTrackID *string `json:"musicbrainz_track_id"` // *string: UUID or null
}
