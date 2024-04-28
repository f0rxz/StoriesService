package db

import (
	"errors"
	"storiesservice/internal/consts"
	"storiesservice/pkg/logger"
)

type Story struct {
	Id       int64  `json:"id"`
	Type     int16  `json:"type"`
	Link     string `json:"link"`
	Text     string `json:"text"`
	TextPosX int    `json:"textPosX"`
	TextPosY int    `json:"testPosY"`
}

type Stories struct {
	List []*Story         `json:"list"`
	ById map[int64]*Story `json:"byId"`
}

func NewStories() *Stories {
	return &Stories{
		List: []*Story{},
		ById: map[int64]*Story{},
	}
}

func (d *Database) GetStories() *Stories {
	stories, ok := d.StoriesCache.Modify(func() interface{} {
		rows, err := d.db.Query(`
			SELECT s.id, s.type, s.link, s.text, s.text_pos_x, s.text_pos_y
			FROM stories s
			ORDER BY s.id DESC
		`)
		if err != nil {
			logger.Eprintln(err)
			return nil
		}

		stories := NewStories()
		for rows.Next() {
			story := &Story{}
			if err := rows.Scan(
				&story.Id,
				&story.Type,
				&story.Link,
				&story.Text,
				&story.TextPosX,
				&story.TextPosY,
			); err != nil {
				logger.Eprintln(err)
				return nil
			}
			stories.List = append(stories.List, story)
			stories.ById[story.Id] = story
		}
		return stories
	}).(*Stories)

	if !ok || stories == nil {
		return NewStories()
	}

	return stories
}

func (d *Database) StoryAdd(storyType int16, link, text string, textPosX, textPosY int) (err error) {
	if _, err = d.db.Exec(`
		INSERT INTO stories (type, link, text, text_pos_x, text_pos_y)
		VALUES ($1, $2, $3, $4, $5)`, storyType, link, text, textPosX, textPosY,
	); err != nil {
		logger.Eprintln(err)
		err = errors.New(consts.I18nDatabaseError)
	}
	return
}
func (d *Database) StoryEdit(storyId int64, storyType int16, link, text string, textPosX, textPosY int) (err error) {
	if _, err = d.db.Exec(`
		UPDATE stories
		SET type = $2, link = $3, text = $4, text_pos_x = $5, text_pos_y = $6
		WHERE id = $1`, storyId, storyType, link, text, textPosX, textPosY,
	); err != nil {
		logger.Eprintln(err)
		err = errors.New(consts.I18nDatabaseError)
	}
	return
}

func (d *Database) StoryRemove(storyId int64) (err error) {
	if _, err = d.db.Exec(`
		DELETE FROM stories
		WHERE id = $1`, storyId,
	); err != nil {
		logger.Eprintln(err)
		err = errors.New(consts.I18nDatabaseError)
	}
	return
}
