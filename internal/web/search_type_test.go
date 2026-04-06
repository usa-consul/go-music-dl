package web

import (
	"reflect"
	"strings"
	"testing"

	"github.com/guohuiyuan/go-music-dl/core"
)

func TestDefaultSourcesForSearchType(t *testing.T) {
	wantAlbum := core.GetAlbumSourceNames()
	if got := defaultSourcesForSearchType("album"); !reflect.DeepEqual(got, wantAlbum) {
		t.Fatalf("defaultSourcesForSearchType(album) = %v, want %v", got, wantAlbum)
	}

	if got := defaultSourcesForSearchType("playlist"); len(got) == 0 {
		t.Fatal("defaultSourcesForSearchType(playlist) returned empty sources")
	}

	if got := defaultSourcesForSearchType("song"); len(got) == 0 {
		t.Fatal("defaultSourcesForSearchType(song) returned empty sources")
	}
}

func TestSearchPlaceholderForType(t *testing.T) {
	tests := []struct {
		searchType string
		want       string
	}{
		{searchType: "song", want: "歌曲"},
		{searchType: "playlist", want: "歌单"},
		{searchType: "album", want: "专辑"},
	}

	for _, tt := range tests {
		if got := searchPlaceholderForType(tt.searchType); !strings.Contains(got, tt.want) {
			t.Fatalf("searchPlaceholderForType(%q) = %q, want contains %q", tt.searchType, got, tt.want)
		}
	}
}

func TestCollectionLabelsForSearchType(t *testing.T) {
	if got := collectionLabelForSearchType("album"); got != "专辑" {
		t.Fatalf("collectionLabelForSearchType(album) = %q, want 专辑", got)
	}
	if got := collectionCreatorLabelForSearchType("album"); got != "歌手" {
		t.Fatalf("collectionCreatorLabelForSearchType(album) = %q, want 歌手", got)
	}
	if got := collectionLabelForSearchType("playlist"); got != "歌单" {
		t.Fatalf("collectionLabelForSearchType(playlist) = %q, want 歌单", got)
	}
}
