package taglib

import (
	"github.com/navidrome/navidrome/log"
)

type Parser struct{}

type parsedTags = map[string][]string

func (e *Parser) Parse(paths ...string) (map[string]parsedTags, error) {
	fileTags := map[string]parsedTags{}
	for _, path := range paths {
		tags, err := e.extractMetadata(path)
		if err == nil {
			fileTags[path] = tags
		}
	}
	return fileTags, nil
}

func (e *Parser) extractMetadata(filePath string) (parsedTags, error) {
	tags, err := Read(filePath)
	if err != nil {
		log.Warn("Error reading metadata from file. Skipping", "filePath", filePath, err)
	}

	alternativeTags := map[string][]string{
		"title":       {"titlesort"},
		"album":       {"albumsort"},
		"artist":      {"artistsort"},
		"tracknumber": {"trck", "_track"},
	}

	for tagName, alternatives := range alternativeTags {
		for _, altName := range alternatives {
			if altValue, ok := tags[altName]; ok {
				tags[tagName] = append(tags[tagName], altValue...)
			}
		}
	}
	return tags, nil
}
