package albumsDb

import (
	"errors"
	"fmt"
	. "github.com/danilevy1212/album-service-gin/models"
	"reflect"
	"strconv"
)

type ID string
type AlbumClient struct{}
type AlbumDB interface {
	GetAll() (*[]Album, error)
	Insert(album *AlbumPostBody) (*Album, error)
	GetById(id string) (*Album, error)
	Delete(id string) (*Album, error)
	Patch(id string, albumPatch *AlbumPatchBody) (*Album, error)
}

var albums = map[ID]Album{
	"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
var lastID = uint(len(albums))

func (_ *AlbumClient) GetAll() (*[]Album, error) {
	result := make([]Album, len(albums))
	var idx uint

	for _, v := range albums {
		result[idx] = v
		idx++
	}

	return &result, nil
}

func (_ *AlbumClient) Insert(album *AlbumPostBody) (*Album, error) {
	lastID++
	lastIDStr := strconv.FormatUint(uint64(lastID), 10)
	newAlbum := Album{
		ID:     lastIDStr,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}

	albums[ID(lastIDStr)] = newAlbum

	return &newAlbum, nil
}

func (_ *AlbumClient) GetById(id string) (*Album, error) {
	album, ok := albums[ID(id)]

	if ok {
		return &album, nil
	}

	return nil, errors.New(fmt.Sprintf("Id %s not found", id))
}

func (_ *AlbumClient) Delete(id string) (*Album, error) {
	targetID := ID(id)
	album, ok := albums[targetID]

	if ok {
		delete(albums, targetID)
		return &album, nil
	}

	return nil, errors.New(fmt.Sprintf("Id %s not found", id))
}

func (_ *AlbumClient) Patch(id string, albumPatch *AlbumPatchBody) (*Album, error) {
	targetID := ID(id)
	album, ok := albums[targetID]

	if ok {
		currAlbum := reflect.ValueOf(&album)
		patchAlbum := reflect.ValueOf(albumPatch).Elem()

		for _, field := range reflect.VisibleFields(patchAlbum.Type()) {
			patchAlbumField := patchAlbum.FieldByName(field.Name)

			if patchAlbumField.IsNil() {
				continue
			}

			newFieldValue := patchAlbumField.Elem()
			currAlbum.Elem().FieldByName(field.Name).Set(newFieldValue)
		}
		albums[targetID] = album

		return &album, nil
	}

	return nil, errors.New(fmt.Sprintf("Id %s not found", id))
}

func New() AlbumDB {
	return &AlbumClient{}
}
