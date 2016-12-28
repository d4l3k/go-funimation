package funimation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	funimation "golang.ssttevee.com/funimation/lib"
)

// Client is a client to the Funimation API.
type Client struct {
	Client http.Client
	User   User
}

// NewClient returns a new Client.
func NewClient() *Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	c := &Client{
		Client: http.Client{
			Jar: jar,
		},
	}
	return c
}

// Cookies returns all cookies for the client.
func (c Client) Cookies() []*http.Cookie {
	return c.Client.Jar.Cookies(c.baseURL())
}

// SetCookies sets the cookies on the client.
func (c Client) SetCookies(cookies []*http.Cookie) {
	c.Client.Jar.SetCookies(c.baseURL(), cookies)
}

var _ json.Marshaler = &Client{}

// MarshalJSON implements json.Marshaler.
func (c Client) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Cookies())
}

var _ json.Unmarshaler = &Client{}

// UnmarshalJSON implements json.Unmarshaler.
func (c Client) UnmarshalJSON(data []byte) error {
	var cookies []*http.Cookie
	if err := json.Unmarshal(data, &cookies); err != nil {
		return err
	}
	c.Client.Jar.SetCookies(c.baseURL(), cookies)
	return nil
}

func (c Client) baseURL() *url.URL {
	url, err := url.Parse(BaseURL)
	if err != nil {
		log.Fatal(err)
	}
	return url
}

func loginURL() string {
	return BaseURL + LoginPath
}

type loginRequest struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	PlayStationID string `json:"playstation_id"`
}

// User is the user info of the currently logged in user.
type User struct {
	UserID           string `json:"user_id"`
	UserType         string `json:"user_type"`
	UT               string `json:"ut"`
	UserRole         string `json:"user_role"`
	SubscriberStatus string `json:"subscriber_status"`
	UserBirthday     string `json:"user_birthday"`
	UserAge          int    `json:"user_age"`
	Country          string `json:"country"`

	// Response fields.
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Login logs into a users Funimation account with the specified username and
// password.
func (c *Client) Login(username, password string) (User, error) {
	req := loginRequest{
		Username: username,
		Password: password,
	}
	body, err := json.Marshal(req)
	if err != nil {
		return User{}, err
	}
	resp, err := c.Client.Post(loginURL(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()
	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return User{}, err
	}
	if len(user.Message) > 0 && !user.Success {
		return User{}, errors.New(user.Message)
	}
	c.User = user

	form := url.Values{}
	form.Add("logged", "1")
	form.Add("email_field", username)
	form.Add("password_field", password)
	form.Add("currenturl", "http://shop.funimation.com/")
	form.Add("loginpopup", "yes")

	resp2, err := c.Client.Post("http://www.funimation.com/login", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		return User{}, err
	}
	defer resp2.Body.Close()
	return user, nil
}

// QueueResponse is all the shows queued.
type QueueResponse struct {
	ExecutionTime string `json:"execution_time"`
	UserID        string `json:"user_id"`
	User          struct {
		Userid          string      `json:"userid"`
		Username        string      `json:"username"`
		Email           string      `json:"email"`
		Membergroupids  string      `json:"membergroupids"`
		Usergroupid     string      `json:"usergroupid"`
		IsUsernameSet   string      `json:"is_username_set"`
		AvatarURL       string      `json:"avatar_url"`
		AvatarSize      interface{} `json:"avatar_size"`
		Subscriber      bool        `json:"subscriber"`
		OnceASubscriber bool        `json:"once_a_subscriber"`
	} `json:"user"`
	Msg   string `json:"msg"`
	Queue []struct {
		QueueID                string      `json:"queue_id"`
		Order                  string      `json:"order"`
		OnlyShow               string      `json:"only_show"`
		Recaps                 string      `json:"recaps"`
		Promotionals           string      `json:"promotionals"`
		ShowURL                string      `json:"show_url"`
		Published              string      `json:"published"`
		ShowSquare             string      `json:"show_square"`
		UserRank               interface{} `json:"user_rank"`
		ShowID                 string      `json:"show_id"`
		FunimationWebsite      string      `json:"funimation_website"`
		Copyright              string      `json:"copyright"`
		Title                  string      `json:"title"`
		OriginalTitle          string      `json:"original_title"`
		VodSummary255          string      `json:"vod_summary_255"`
		VodSummary400          string      `json:"vod_summary_400"`
		FullSummary            string      `json:"full_summary"`
		ActiveDub              string      `json:"active_dub"`
		ActiveSub              string      `json:"active_sub"`
		ActiveHd               string      `json:"active_hd"`
		ActiveVideos           string      `json:"active_videos"`
		ActiveEpisodes         string      `json:"active_episodes"`
		ActiveFree             string      `json:"active_free"`
		ActiveSvod             string      `json:"active_svod"`
		ActiveSvodEpisodes     string      `json:"active_svod_episodes"`
		ActiveSvodExclusive    string      `json:"active_svod_exclusive"`
		ActiveClips            string      `json:"active_clips"`
		OriginalProductionYear string      `json:"original_production_year"`
		WeHaveEpisode          string      `json:"we_have_episode"`
		WeHaveMovie            string      `json:"we_have_movie"`
		WeHaveOva              string      `json:"we_have_ova"`
		WeHaveSpecial          string      `json:"we_have_special"`
		Simulcast              string      `json:"simulcast"`
		Pageviews              string      `json:"pageviews"`
		OriginalLanguage       string      `json:"original_language"`
		LanguageAbbreviation   string      `json:"language_abbreviation"`
		Genres                 string      `json:"genres"`
		GenresID               string      `json:"genres_id"`
		TvRatings              string      `json:"tv_ratings"`
		TvRatingsID            string      `json:"tv_ratings_id"`
		Rank                   string      `json:"rank"`
		ShowThumbnail          string      `json:"show_thumbnail"`
		FeaturedTrailerSw      string      `json:"featured_trailer_sw"`
		FeaturedTrailerID      string      `json:"featured_trailer_id"`
		FeaturedProductSw      string      `json:"featured_product_sw"`
		FeaturedProductID      string      `json:"featured_product_id"`
		ForumID                string      `json:"forum_id"`
		FeaturedTrailerURL     string      `json:"featured_trailer_url"`
		RecapsCount            int         `json:"recaps_count"`
		VideosCount            int         `json:"videos_count"`
	} `json:"queue"`
	ShowsCount  int    `json:"shows_count"`
	VideosCount int    `json:"videos_count"`
	Duration    int    `json:"duration"`
	CurrentPage int    `json:"current_page"`
	Limit       string `json:"limit"`
	NextVideo   map[string]struct {
		ShowID         string      `json:"show_id"`
		ActiveVideos   string      `json:"active_videos"`
		ShowURL        string      `json:"show_url"`
		VideoID        string      `json:"video_id"`
		Watched        string      `json:"watched"`
		Checkpoint     string      `json:"checkpoint"`
		RecapID        string      `json:"recap_id"`
		Title          string      `json:"title"`
		Thumbnail      string      `json:"thumbnail"`
		VideoType      string      `json:"video_type"`
		Number         string      `json:"number"`
		VideoSequence  string      `json:"video_sequence"`
		VideoURL       string      `json:"video_url"`
		VideoSimulcast string      `json:"video_simulcast"`
		Language       interface{} `json:"language"`
		VideoCategory  string      `json:"video_category"`
		VideoNumber    string      `json:"video_number"`
		VideoTitle     string      `json:"video_title"`
		VideoDuration  string      `json:"video_duration"`
		Exclusive      string      `json:"exclusive"`
		WidgetTitle    string      `json:"widget_title"`
	} `json:"next_video"`
	Offset int    `json:"offset"`
	Items  string `json:"items"`
}

func queueURL(userID string, limit, offset int) string {
	return fmt.Sprintf("%sprofile/queue_search_ajax?userid=%s&offset=%d&limit=%d", BaseURL, userID, offset, limit)
}

// Queue returns all shows that have been queued up.
func (c *Client) Queue(limit, offset int) (QueueResponse, error) {
	var qr QueueResponse
	resp, err := c.Client.Get(queueURL(c.User.UserID, limit, offset))
	if err != nil {
		return qr, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&qr); err != nil {
		return qr, errors.Wrap(err, "decode queue response, probably login error")
	}
	return qr, nil
}

// DownloadClient returns a client that can be used to download an episode.
func (c *Client) DownloadClient() *funimation.Client {
	return funimation.New(c.Client.Jar.(*cookiejar.Jar))
}
