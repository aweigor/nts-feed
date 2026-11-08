# NTS Radio RSS / JSON Feed Generator

A web server for generating RSS / JSON feeds from NTS Radio, configurable via the Linux command line. Supports tracklist loading by URL.

## Table of Contents

- [NTS Radio RSS / JSON Feed Generator](#nts-radio-rss--json-feed-generator)
  - [Table of Contents](#table-of-contents)
  - [Configuration](#configuration)
  - [Feed Structure](#feed-structure)
    - [Endpoints](#endpoints)
    - [Response Structure](#response-structure)
      - [1. RSS](#1-rss)
      - [2. JSON](#2-json)
    - [URL Query Parameters](#url-query-parameters)
  - [API Reference](#api-reference)
    - [1. Episode Info](#1-episode-info)
    - [2. Episode Tracklist](#2-episode-tracklist)
  - [Types](#types)
    - [Tracklist Item](#tracklist-item)
    - [Episode](#episode)
    - [Genre](#genre)
    - [Media](#media)
    - [AudioSource](#audiosource)
  - [Error Codes](#error-codes)

## Running

**Prerequisites:** Go 1.21+

1. Copy the example env file and set your secret:
   ```sh
   cp .env.example .env
   # edit .env and set API_SECRET
   ```

2. Run with `go run` (from the project root):
   ```sh
   go run ./cmd/main.go
   ```

3. Or build and run the binary:
   ```sh
   go build -o nts-feed-generator ./cmd/main.go
   ./nts-feed-generator
   ```

The server starts on port `9000`. Check `GET /heartbeat` to verify it's up.

> **Note:** Both `config.yaml` and `.env` must be present in the directory where you run the binary / `go run` command.

## Configuration

```yaml
nts:
  api_url: "https://www.nts.live/api/v2/"

api:
  public_url: "https://nts-feed/api"

feed:
  output_dir: "./feeds"
  tracklist: true
```

## Feed Structure

### Endpoints

| Section | Method | Path                            |
| ------- | ------ | ------------------------------- |
| Latest  | GET    | `/feeds/latest[?query]`         |
| Show    | GET    | `/feeds/show/{show_id}[?query]` |

### Response Structure

#### 1. RSS

```xml
<channel>
  <title>Channel Title</title>
  <link>https://www.nts.live/{channel_path}</link>
  <description>Channel description</description>
  <category>Channel category</category>
  <image>Channel Image Url</image>
  <lastBuildDate>Last Updated Date</lastBuildDate>
  <item>
      <title>Episode Title</title>
      <link>https://www.nts.live/{episode_path}</link>
      <description>Episode description</description>
      <pubDate>2024-01-15T20:00:00Z</pubDate>
      <enclosure url="{audio_url}" type="audio/mpeg" length="0"/>
      <guid>episode-123</guid>

      <nts:show_id>Show Id</nts:show_id>
      <nts:air_date>2024-01-15</nts:air_date>
      <nts:duration>120</nts:duration>

      <api:tracklist_url>Api Tracklist Url</api:tracklist_url>
  </item>
</channel>
```

#### 2. JSON

```json
{
  "channel": {
    "title": "Channel Title",
    "link": "https://www.nts.live/{channel_path}",
    "description": "Channel description",
    "category": "Channel category",
    "image": "Channel Image Url",
    "lastBuildDate": "Last Updated Date",
    "items": [{
      "title": "Episode Title",
      "link": "https://www.nts.live/{episode_path}",
      "description": "Episode description",
      "pubDate": "2024-01-15T20:00:00Z",
      "enclosure": {
        "url": "{audio_url}",
        "type": "audio/mpeg",
        "length": "0"
      },
      "guid": "episode-123",
      "nts:show_id": "Show Id",
      "nts:air_date": "2024-01-15",
      "nts:duration": "120",
      "tracklist": [
        {
          "artist": "string",
          "title": "string",
          "uid": "string (UUID or null)",
          "offset": "integer or null",
          "duration": "integer or null",
          "offset_estimate": "integer or null",
          "duration_estimate": "integer or null",
          "acr_id": "string or null",
          "deezer_track_id": "integer or null",
          "isrc_id": "string or null",
          "musicbrainz_track_id": "string (UUID or null)"
        }
      ]
    }]
  }
}
```

### URL Query Parameters

| Parameter  | Type     | Default | Description                                 | Example                        |
| ---------- | -------- | ------- | ------------------------------------------- | ------------------------------ |
| `format`   | string   | `rss`   | Response format: `rss`, `atom`, `json`      | `?format=json`                 |
| `limit`    | integer  | `50`    | Max items per page (1–200)                  | `?limit=20`                    |
| `offset`   | integer  | `0`     | Items to skip for pagination                | `?offset=40`                   |
| `page`     | integer  | `1`     | Page number (alternative to `offset`)       | `?page=3`                      |
| `since`    | datetime | —       | Items published after this date (ISO 8601)  | `?since=2026-08-01T00:00:00Z`  |
| `before`   | datetime | —       | Items published before this date (ISO 8601) | `?before=2026-08-01T00:00:00Z` |
| `category` | string   | —       | Filter by category / tag                    | `?category=tech`               | not implemented
| `author`   | string   | —       | Filter by author username / ID              | `?author=johndoe`              | not implemented
| `q`        | string   | —       | Search query in titles and content          | `?q=golang+tutorial`           | not implemented
| `order`    | string   | `desc`  | Sort order: `desc` or `asc`                 | `?order=asc`                   | not implemented
| `full`     | boolean  | `false` | Include full content instead of excerpts    | `?full=true`                   | no effect
| `callback` | string   | —       | JSONP callback function name                | `?callback=myCallback`         | no effect

## API Reference

### 1. Episode Info

`GET /shows/{show_id}/{episode_id}`

Returns an [Episode](#episode) with an additional `embeds` field:

```json
{
  "embeds": {
    "tracklist": {
      "metadata": {
        "resultset": {
          "count": "integer",
          "offset": "integer",
          "limit": "integer"
        }
      },
      "results": ["Tracklist Item"]
    }
  }
}
```

### 2. Episode Tracklist

`GET /shows/{show_id}/{episode_id}/tracklist`

```json
{
  "metadata": {
    "resultset": { "count": "integer", "offset": "integer", "limit": "integer" }
  },
  "results": ["Tracklist Item"]
}
```

## Types

### Tracklist Item

```json
{
  "artist": "string",
  "title": "string",
  "uid": "string (UUID) or null",
  "offset": "integer or null",
  "duration": "integer or null",
  "offset_estimate": "integer or null",
  "duration_estimate": "integer or null",
  "acr_id": "string or null",
  "deezer_track_id": "integer or null",
  "isrc_id": "string or null",
  "musicbrainz_track_id": "string (UUID) or null"
}
```

### Episode

```json
{
  "status": "string",
  "updated": "datetime (ISO 8601)",
  "name": "string",
  "description": "string",
  "description_html": "string (HTML)",
  "external_links": "array",
  "moods": "array",
  "genres": "Genre[]",
  "location_short": "string",
  "location_long": "string",
  "intensity": "string (numeric)",
  "media": "Media",
  "episode_alias": "string",
  "show_alias": "string",
  "broadcast": "datetime (ISO 8601)",
  "mixcloud": "string (URL)",
  "audio_sources": "AudioSource[]",
  "brand": "object"
}
```

### Genre

```json
{
  "id": "string",
  "value": "string"
}
```

### Media

```json
{
  "background_large": "string (URL)",
  "background_medium_large": "string (URL)",
  "background_medium": "string (URL)",
  "background_small": "string (URL)",
  "background_thumb": "string (URL)",
  "picture_large": "string (URL)",
  "picture_medium_large": "string (URL)",
  "picture_medium": "string (URL)",
  "picture_small": "string (URL)",
  "picture_thumb": "string (URL)"
}
```

### AudioSource

```json
{
  "url": "string (URL)",
  "source": "string"
}
```

## Error Codes

| Code | Status                | Description                 |
| ---- | --------------------- | --------------------------- |
| 200  | OK                    | Success                     |
| 304  | Not Modified          | Feed unchanged (ETag match) |
| 400  | Bad Request           | Invalid parameters          |
| 404  | Not Found             | Feed or items not found     |
| 429  | Too Many Requests     | Rate limited                |
| 500  | Internal Server Error | Server error                |
