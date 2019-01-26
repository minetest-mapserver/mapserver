package playerdb

import (
	"io/ioutil"
	"time"
)

//https://stackoverflow.com/questions/46746862/list-files-in-a-directory-sorted-by-creation-time

type FilePlayerDB struct{}

func (this *FilePlayerDB) GetActivePlayers() ([]*Player, error) {
	files, err := ioutil.ReadDir("Players")

	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	activeTime := now - 5
	result := make([]*Player, 0)

	for _, file := range files {
		if file.ModTime().Unix() > activeTime {
			p := Player{}

			//TODO: parse

			result = append(result, &p)
		}
	}

	return result, nil
}
