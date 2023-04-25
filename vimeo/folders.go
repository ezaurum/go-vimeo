package vimeo

import (
	"fmt"
	"time"
)

// FoldersService handles communication with the folders related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/folders
type FoldersService service

type dataListFolder struct {
	Data []*Folder `json:"data"`
	pagination
}

// Folder represents a category.
type Folder struct {
	CreatedTime           time.Time `json:"created_time"`
	URI                   string    `json:"uri,omitempty"`
	Link                  string    `json:"link,omitempty"`
	Name                  string    `json:"name,omitempty"`
	TopLevel              bool      `json:"top_level"`
	Pictures              *Pictures `json:"pictures,omitempty"`
	LastVideoFeaturedTime string    `json:"last_video_featured_time,omitempty"`
	Parent                *Folder   `json:"parent,omitempty"`
	SubFolders            []*Folder `json:"subfolders,omitempty"`
	ResourceKey           string    `json:"resource_key,omitempty"`
	MetaData              *MetaData `json:"metadata,omitempty"`
}

type Interaction struct {
	URI  string `json:"URI,omitempty"`
	Name string `json:"name,omitempty"`
	Link string `json:"link,omitempty"`
}

type Interactions struct {
	WatchLater *Interaction `json:"watchlater,omitempty"`
	Like       *Interaction `json:"like,omitempty"`
}

type MetaData struct {
	Interactions *Interactions `json:"interactions,omitempty"`
	ParentFolder *Folder       `json:"parent_folder,omitempty"`
}

func listFolder(c *Client, url string, opt ...CallOption) ([]*Folder, *Response, error) {
	u, err := addOptions(url, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	folders := &dataListFolder{}

	resp, err := c.Do(req, folders)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(folders)

	return folders.Data, resp, err
}

func getFolder(c *Client, url string, opt ...CallOption) (*Folder, *Response, error) {
	u, err := addOptions(url, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	category := &Folder{}

	resp, err := c.Do(req, category)
	if err != nil {
		return nil, resp, err
	}

	return category, resp, err
}

// List method gets all existing folders.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/folders#get_folders
func (s *FoldersService) List(opt ...CallOption) ([]*Folder, *Response, error) {
	folders, resp, err := listFolder(s.client, "folders", opt...)

	return folders, resp, err
}

// Get method gets a single category.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/folders#get_category
func (s *FoldersService) Get(cat string, opt ...CallOption) (*Folder, *Response, error) {
	u := fmt.Sprintf("folders/%s", cat)
	category, resp, err := getFolder(s.client, u, opt...)

	return category, resp, err
}

// ListVideo method gets all the videos that belong to a folder.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/folders#get_folder_items
func (s *FoldersService) ListVideo(cat string, opt ...CallOption) ([]*Video, *Response, error) {
	u := fmt.Sprintf("folders/%s/videos", cat)
	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}
