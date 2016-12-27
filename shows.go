package funimation

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Show represents a single Funimation show.
type Show struct {
	AssetID                    string      `json:"asset_id"`
	PubDate                    string      `json:"pubDate"`
	SeriesName                 string      `json:"series_name"`
	Link                       string      `json:"link"`
	SeriesDescription          string      `json:"series_description"`
	SeasonCount                string      `json:"season_count"`
	EpisodeCount               int         `json:"episode_count"`
	Genres                     string      `json:"genres"`
	Simulcast                  string      `json:"simulcast"`
	Popularity                 string      `json:"popularity"`
	OfficialMarketingWebsite   string      `json:"official_marketing_website"`
	LatestVideoFree            interface{} `json:"latest_video_free"`
	LatestVideoFreeReleaseDate interface{} `json:"latest_video_free_release_date"`
	LatestVideoSubscription    struct {
		VideoID     string `json:"video_id"`
		ReleaseDate string `json:"release_date"`
		Title       string `json:"title"`
	} `json:"latest_video_subscription"`
	LatestVideoSubscriptionReleaseDate string `json:"latest_video_subscription_release_date"`
	ShowRating                         string `json:"show_rating"`
	ActiveHD1080                       string `json:"active_hd_1080"`
	HasClosedCaptions                  string `json:"has_closed_captions"`
	ThumbnailSmall                     string `json:"thumbnail_small"`
	ThumbnailMedium                    string `json:"thumbnail_medium"`
	ThumbnailLarge                     string `json:"thumbnail_large"`
	PosterArt                          string `json:"poster_art"`
	PosterArtLarge                     string `json:"poster_art_large"`
	ContactLink                        string `json:"contactLink"`
	DisplayOrder                       int    `json:"display_order"`
	ElementPosition                    int    `json:"element_position"`
	RatingSystem                       string `json:"rating_system"`
	Quality                            string `json:"quality"`
	Languages                          string `json:"languages,omitempty"`
}

// Video represents a single Funimation video.
type Video struct {
	AssetID               string   `json:"asset_id"`
	FunimationID          string   `json:"funimation_id"`
	PubDate               string   `json:"pubDate"`
	Rating                string   `json:"rating"`
	Quality               string   `json:"quality"`
	Language              string   `json:"language"`
	Duration              int      `json:"duration"`
	Simulcast             string   `json:"simulcast"`
	ClosedCaptioning      string   `json:"closed_captioning"`
	URL                   string   `json:"url"`
	Promo                 string   `json:"promo"`
	ShowName              string   `json:"show_name"`
	Popularity            string   `json:"popularity"`
	Title                 string   `json:"title"`
	Description           string   `json:"description"`
	Sequence              string   `json:"sequence"`
	Number                int      `json:"number"`
	VideoType             string   `json:"video_type"`
	ShowID                string   `json:"show_id"`
	Thumbnail             string   `json:"thumbnail"`
	SeasonID              string   `json:"season_id"`
	SeasonNumber          string   `json:"season_number"`
	Genre                 string   `json:"genre"`
	ReleaseDate           string   `json:"releaseDate"`
	ThumbnailURL          string   `json:"thumbnail_url"`
	ThumbnailSmall        string   `json:"thumbnail_small"`
	ThumbnailMedium       string   `json:"thumbnail_medium"`
	ThumbnailLarge        string   `json:"thumbnail_large"`
	VideoURL              string   `json:"video_url"`
	ClosedCaptionLocation string   `json:"closed_caption_location"`
	Aips                  []string `json:"aips"`
	DubSub                string   `json:"dub_sub"`
	Featured              string   `json:"featured"`
	Highdef               string   `json:"highdef"`
	HasSubtitles          string   `json:"has_subtitles"`
	ElementPosition       int      `json:"element_position"`
	TvOrMove              string   `json:"tv_or_move"`
	RatingSystem          string   `json:"rating_system"`
	DisplayOrder          int      `json:"display_order"`
	ExtendedTitle         string   `json:"extended_title"`
}

type videosResponse struct {
	Videos []Video `json:"videos"`
}

// Funimation API URL endpoints
var (
	BaseURL    = "http://www.funimation.com/"
	ShowsPath  = "feeds/ps/shows"
	VideosPath = "feeds/ps/videos"
	LoginPath  = "feeds/ps/login.json?v=2"
)

func showsURL(limit, offset int) string {
	return fmt.Sprintf("%s%s?limit=%d&offset=%d", BaseURL, ShowsPath, limit, offset)
}

// Shows returns all the shows on Funimation.
func Shows(limit, offset int) ([]Show, error) {
	var shows []Show
	resp, err := http.Get(showsURL(limit, offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&shows); err != nil {
		return nil, err
	}
	return shows, nil
}

func videosURL(limit, offset int) string {
	return fmt.Sprintf("%s%s?limit=%d&offset=%d", BaseURL, VideosPath, limit, offset)
}

// Videos return all the videos on Funimation.
func Videos(limit, offset int) ([]Video, error) {
	var videos videosResponse
	resp, err := http.Get(videosURL(limit, offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&videos); err != nil {
		return nil, err
	}
	return videos.Videos, nil
}
