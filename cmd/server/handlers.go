package main

import (
	"net/http"
	"storiesservice/internal/consts"
	"storiesservice/internal/sessions"
	"storiesservice/pkg/logger"
	"storiesservice/pkg/utctime"
	"strconv"
	"time"
)

func baseHandler(w http.ResponseWriter, r *http.Request) (lang string, authorized bool, userId int64) {
	r.ParseForm()
	lang = r.Form.Get("lang")
	if lang == "" {
		cookieLang, err := r.Cookie("lang")
		if err == nil {
			lang = cookieLang.Value
		}
	} else {
		http.SetCookie(w, &http.Cookie{Name: "lang", Value: lang, Expires: utctime.Get().Add(365 * 24 * time.Hour)})
	}

	cookieSession, err := r.Cookie("session")
	if err == nil {
		userId, err = sessions.Parse(cookieSession.Value)
		if err == nil {
			authorized = true
		} else {
			authorized = false
			http.SetCookie(w, &http.Cookie{Name: "session", Value: "", Expires: time.Unix(0, 0)})
		}
	}

	return
}

func modifyStory(lang string, w http.ResponseWriter, r *http.Request, storyId ...int64) {
	storyTypeStr := r.Form.Get("type")
	var storyType int16
	switch storyTypeStr {
	case "image":
		storyType = 0
	case "video":
		storyType = 1
	default:
		if err := tmplErrMsg.Execute(w, gI18n.Translate(consts.I18nErrorInvalidStoryType, lang)); err != nil {
			logger.Eprintln(err)
		}
		return
	}
	link := r.Form.Get("link")
	text := r.Form.Get("text")
	textPosXStr := r.Form.Get("text_pos_x")
	textPosYStr := r.Form.Get("text_pos_y")
	textPosX, err := strconv.Atoi(textPosXStr)
	if err != nil {
		if err := tmplErrMsg.Execute(w, gI18n.Translate(consts.I18nErrorInvalidStoryTextPos, lang)); err != nil {
			logger.Eprintln(err)
		}
		return
	}
	textPosY, err := strconv.Atoi(textPosYStr)
	if err != nil {
		if err := tmplErrMsg.Execute(w, gI18n.Translate(consts.I18nErrorInvalidStoryTextPos, lang)); err != nil {
			logger.Eprintln(err)
		}
		return
	}
	if len(storyId) > 0 {
		if err := gDatabase.StoryEdit(storyId[0], storyType, link, text, textPosX, textPosY); err != nil {
			if err := tmplErrMsg.Execute(w, gI18n.Translate(err.Error(), lang)); err != nil {
				logger.Eprintln(err)
			}
			return
		}
	} else {
		if err := gDatabase.StoryAdd(storyType, link, text, textPosX, textPosY); err != nil {
			if err := tmplErrMsg.Execute(w, gI18n.Translate(err.Error(), lang)); err != nil {
				logger.Eprintln(err)
			}
			return
		}
	}

	gDatabase.StoriesCache.ForceUpdate()
	if len(storyId) > 0 {
		if err := tmplMsg.Execute(w, gI18n.Translate(consts.I18nSuccessfulStoryEdit, lang)); err != nil {
			logger.Eprintln(err)
		}
	} else {
		if err := tmplMsg.Execute(w, gI18n.Translate(consts.I18nSuccessfulStoryPublication, lang)); err != nil {
			logger.Eprintln(err)
		}
	}
}

func registerHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lang, authorized, userId := baseHandler(w, r)
		if authorized {
			if err := tmplIndex.Execute(w, indexData{
				IsAdmin:      userId == 1,
				AStoryAdd:    gI18n.Translate(consts.I18nStoryIndexAStoryAdd, lang),
				AStoryEdit:   gI18n.Translate(consts.I18nStoryIndexAStoryEdit, lang),
				AStoryRemove: gI18n.Translate(consts.I18nStoryIndexAStoryRemove, lang),
				Stories:      gDatabase.GetStories().List,
			}); err != nil {
				logger.Eprintln(err)
			}
		} else {
			http.Redirect(w, r, "/sign_in", http.StatusFound)
		}
	})

	http.HandleFunc("/sign_in", func(w http.ResponseWriter, r *http.Request) {
		lang, _, _ := baseHandler(w, r)
		switch r.Method {
		case "POST":
			session, err := gDatabase.SignIn(r.Form.Get("username"), r.Form.Get("password"))
			if err == nil {
				http.SetCookie(w, &http.Cookie{Name: "session", Value: session, Expires: utctime.Get().Add(365 * 24 * time.Hour)})
				if err := tmplMsg.Execute(w, gI18n.Translate(consts.I18nSuccessfulAuthorization, lang)); err != nil {
					logger.Eprintln(err)
				}
			} else {
				if err := tmplErrMsg.Execute(w, gI18n.Translate(err.Error(), lang)); err != nil {
					logger.Eprintln(err)
				}
			}
			break
		default:
			if err := tmplSignIn.Execute(w, signInData{
				Legend:        gI18n.Translate(consts.I18nSignInLegend, lang),
				LabelUsername: gI18n.Translate(consts.I18nSignInLabelUsername, lang),
				LabelPassword: gI18n.Translate(consts.I18nSignInLabelPassword, lang),
				Submit:        gI18n.Translate(consts.I18nSignInSubmit, lang),
			}); err != nil {
				logger.Eprintln(err)
			}
		}
	})

	http.HandleFunc("/story_add", func(w http.ResponseWriter, r *http.Request) {
		lang, authorized, userId := baseHandler(w, r)
		if !authorized || userId != 1 {
			return
		}
		switch r.Method {
		case "POST":
			modifyStory(lang, w, r)
		default:
			if err := tmplStoryModify.Execute(w, storyModifyData{
				EndpointValue: "/story_add",
				TypeValue:     0,
				LinkValue:     "",
				TextValue:     "",
				TextPosXValue: 270 / 2,
				TextPosYValue: 480 / 2,
				Legend:        gI18n.Translate(consts.I18nStoryAddLegend, lang),
				LabelType:     gI18n.Translate(consts.I18nStoryAddLabelType, lang),
				OptionImage:   gI18n.Translate(consts.I18nStoryAddOptionImage, lang),
				OptionVideo:   gI18n.Translate(consts.I18nStoryAddOptionVideo, lang),
				LabelLink:     gI18n.Translate(consts.I18nStoryAddLabelLink, lang),
				LabelText:     gI18n.Translate(consts.I18nStoryLabelTextLink, lang),
				LabelTextPos:  gI18n.Translate(consts.I18nStoryLabelTextPosLink, lang),
				Submit:        gI18n.Translate(consts.I18nStoryAddSubmit, lang),
			}); err != nil {
				logger.Eprintln(err)
			}
		}
	})

	http.HandleFunc("/story_edit", func(w http.ResponseWriter, r *http.Request) {
		lang, authorized, userId := baseHandler(w, r)
		if !authorized || userId != 1 {
			return
		}
		storyIdStr := r.Form.Get("id")
		if storyIdStr == "" {
			cookieEditingStoryId, err := r.Cookie("editing_story_id")
			if err == nil {
				storyIdStr = cookieEditingStoryId.Value
			}
		} else {
			http.SetCookie(w, &http.Cookie{Name: "editing_story_id", Value: storyIdStr, Expires: utctime.Get().Add(365 * 24 * time.Hour)})
		}
		storyId, err := strconv.ParseInt(storyIdStr, 10, 64)
		if err != nil {
			if err := tmplErrMsg.Execute(w, gI18n.Translate(consts.I18nErrorInvalidStoryId, lang)); err != nil {
				logger.Eprintln(err)
			}
			return
		}
		stories := gDatabase.GetStories()
		story, ok := stories.ById[storyId]
		if !ok {
			if err := tmplErrMsg.Execute(w, gI18n.Translate(consts.I18nStoryNotFound, lang)); err != nil {
				logger.Eprintln(err)
			}
			return
		}
		switch r.Method {
		case "POST":
			modifyStory(lang, w, r, story.Id)
		default:
			if err := tmplStoryModify.Execute(w, storyModifyData{
				EndpointValue: "/story_edit",
				TypeValue:     story.Type,
				LinkValue:     story.Link,
				TextValue:     story.Text,
				TextPosXValue: story.TextPosX,
				TextPosYValue: story.TextPosY,
				Legend:        gI18n.Translate(consts.I18nStoryAddLegend, lang),
				LabelType:     gI18n.Translate(consts.I18nStoryAddLabelType, lang),
				OptionImage:   gI18n.Translate(consts.I18nStoryAddOptionImage, lang),
				OptionVideo:   gI18n.Translate(consts.I18nStoryAddOptionVideo, lang),
				LabelLink:     gI18n.Translate(consts.I18nStoryAddLabelLink, lang),
				LabelText:     gI18n.Translate(consts.I18nStoryLabelTextLink, lang),
				LabelTextPos:  gI18n.Translate(consts.I18nStoryLabelTextPosLink, lang),
				Submit:        gI18n.Translate(consts.I18nStoryEditSubmit, lang),
			}); err != nil {
				logger.Eprintln(err)
			}
		}
	})

	http.HandleFunc("/story_remove", func(w http.ResponseWriter, r *http.Request) {
		lang, authorized, userId := baseHandler(w, r)
		if !authorized || userId != 1 {
			return
		}
		storyIdStr := r.Form.Get("id")
		storyId, err := strconv.ParseInt(storyIdStr, 10, 64)
		if err != nil {
			if err := tmplErrMsg.Execute(w, gI18n.Translate(consts.I18nErrorInvalidStoryId, lang)); err != nil {
				logger.Eprintln(err)
			}
			return
		}
		stories := gDatabase.GetStories()
		_, ok := stories.ById[storyId]
		if !ok {
			if err := tmplErrMsg.Execute(w, gI18n.Translate(consts.I18nStoryNotFound, lang)); err != nil {
				logger.Eprintln(err)
			}
			return
		}

		if err := gDatabase.StoryRemove(storyId); err != nil {
			if err := tmplErrMsg.Execute(w, gI18n.Translate(err.Error(), lang)); err != nil {
				logger.Eprintln(err)
			}
			return
		}

		gDatabase.StoriesCache.ForceUpdate()

		if err := tmplMsg.Execute(w, gI18n.Translate(consts.I18nSuccessfulStoryRemove, lang)); err != nil {
			logger.Eprintln(err)
		}
	})
}
