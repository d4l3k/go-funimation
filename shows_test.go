package funimation

import (
	"reflect"
	"testing"
)

func TestShowsURL(t *testing.T) {
	cases := []struct {
		limit, offset int
		want          string
	}{
		{
			1000, 0,
			"http://www.funimation.com/feeds/ps/shows?limit=1000&offset=0",
		},
	}

	for i, c := range cases {
		out := showsURL(c.limit, c.offset)
		if c.want != out {
			t.Errorf("%d. showsURL(%v, %v) = %v; not %v", i, c.limit, c.offset, out, c.want)
		}
	}
}

func TestVideosURL(t *testing.T) {
	cases := []struct {
		limit, offset int
		want          string
	}{
		{
			1000, 0,
			"http://www.funimation.com/feeds/ps/videos?limit=1000&offset=0",
		},
	}

	for i, c := range cases {
		out := videosURL(c.limit, c.offset)
		if c.want != out {
			t.Errorf("%d. showsURL(%v, %v) = %v; not %v", i, c.limit, c.offset, out, c.want)
		}
	}
}

func TestShows(t *testing.T) {
	limit := 5
	shows, err := Shows(limit, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(shows) != limit {
		t.Errorf("expected len(shows) = %d, not %v", limit, shows)
	}
	for i, s := range shows {
		if reflect.DeepEqual(s, Show{}) {
			t.Errorf("expected shows[%d] to be populated", i)
		}
	}
}

func TestVideos(t *testing.T) {
	limit := 5
	videos, err := Videos(limit, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(videos) != limit {
		t.Errorf("expected len(videos) = %d, not %v", limit, videos)
	}
	for i, s := range videos {
		if reflect.DeepEqual(s, Video{}) {
			t.Errorf("expected videos[%d] to be populated", i)
		}
	}
}
