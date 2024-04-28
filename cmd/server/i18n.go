package main

import "storiesservice/internal/consts"

func init() {
	// ru
	gI18n.AddTranslation("ru", consts.I18nGetUsersFailed, "не удалось получить список пользователей")
	gI18n.AddTranslation("ru", consts.I18nUserNotFound, "пользователь не найден")
	gI18n.AddTranslation("ru", consts.I18nWrongPassword, "неверный пароль")
	gI18n.AddTranslation("ru", consts.I18nSessionDecodeFailed, "не удалось декодировать сессию")
	gI18n.AddTranslation("ru", consts.I18nSessionVerifyFailed, "не удалось проверить сессию")
	gI18n.AddTranslation("ru", consts.I18nSessionExpired, "сессия истекла")
	// en
	gI18n.AddTranslation("en", consts.I18nGetUsersFailed, "failed to get users list")
	gI18n.AddTranslation("en", consts.I18nUserNotFound, "user not found")
	gI18n.AddTranslation("en", consts.I18nWrongPassword, "wrong password")
	gI18n.AddTranslation("en", consts.I18nSessionDecodeFailed, "failed to decode session")
	gI18n.AddTranslation("en", consts.I18nSessionVerifyFailed, "failed to verify session")
	gI18n.AddTranslation("en", consts.I18nSessionExpired, "session expired")

	// ru
	gI18n.AddTranslation("ru", consts.I18nSuccessfulAuthorization, "успешная авторизация")
	// en
	gI18n.AddTranslation("en", consts.I18nSuccessfulAuthorization, "successful authorization")

	// ru
	gI18n.AddTranslation("ru", consts.I18nSignInLegend, "Авторизация")
	gI18n.AddTranslation("ru", consts.I18nSignInLabelUsername, "Имя пользователя")
	gI18n.AddTranslation("ru", consts.I18nSignInLabelPassword, "Пароль")
	gI18n.AddTranslation("ru", consts.I18nSignInSubmit, "Войти")
	// en
	gI18n.AddTranslation("en", consts.I18nSignInLegend, "Sign In")
	gI18n.AddTranslation("en", consts.I18nSignInLabelUsername, "Username")
	gI18n.AddTranslation("en", consts.I18nSignInLabelPassword, "Password")
	gI18n.AddTranslation("en", consts.I18nSignInSubmit, "Sign In")

	// ru
	gI18n.AddTranslation("ru", consts.I18nStoryAddLegend, "Новая сторис")
	gI18n.AddTranslation("ru", consts.I18nStoryAddLabelType, "Тип сторис")
	gI18n.AddTranslation("ru", consts.I18nStoryAddOptionImage, "Картинка")
	gI18n.AddTranslation("ru", consts.I18nStoryAddOptionVideo, "Видео")
	gI18n.AddTranslation("ru", consts.I18nStoryAddLabelLink, "Ссылка")
	gI18n.AddTranslation("ru", consts.I18nStoryLabelTextLink, "Текст")
	gI18n.AddTranslation("ru", consts.I18nStoryLabelTextPosLink, "Позиция текста")
	gI18n.AddTranslation("ru", consts.I18nStoryAddSubmit, "Опубликовать сторис")
	gI18n.AddTranslation("ru", consts.I18nStoryEditSubmit, "Редактировать сторис")
	// en
	gI18n.AddTranslation("en", consts.I18nStoryAddLegend, "New story")
	gI18n.AddTranslation("en", consts.I18nStoryAddLabelType, "Story type")
	gI18n.AddTranslation("en", consts.I18nStoryAddOptionImage, "Image")
	gI18n.AddTranslation("en", consts.I18nStoryAddOptionVideo, "Video")
	gI18n.AddTranslation("en", consts.I18nStoryAddLabelLink, "Link")
	gI18n.AddTranslation("en", consts.I18nStoryLabelTextLink, "Text")
	gI18n.AddTranslation("en", consts.I18nStoryLabelTextPosLink, "Text position")
	gI18n.AddTranslation("en", consts.I18nStoryAddSubmit, "Post story")
	gI18n.AddTranslation("en", consts.I18nStoryEditSubmit, "Edit story")

	// ru
	gI18n.AddTranslation("ru", consts.I18nStoryIndexAStoryAdd, "Новая сторис")
	gI18n.AddTranslation("ru", consts.I18nStoryIndexAStoryEdit, "Редактировать")
	gI18n.AddTranslation("ru", consts.I18nStoryIndexAStoryRemove, "Удалить")
	// en
	gI18n.AddTranslation("en", consts.I18nStoryIndexAStoryAdd, "New story")
	gI18n.AddTranslation("en", consts.I18nStoryIndexAStoryEdit, "Edit")
	gI18n.AddTranslation("en", consts.I18nStoryIndexAStoryRemove, "Remove")

	// ru
	gI18n.AddTranslation("ru", consts.I18nStoryNotFound, "сторис не найдена")
	// en
	gI18n.AddTranslation("en", consts.I18nStoryNotFound, "story not found")

	// ru
	gI18n.AddTranslation("ru", consts.I18nSuccessfulStoryPublication, "успешная публикация сторис")
	gI18n.AddTranslation("ru", consts.I18nSuccessfulStoryEdit, "успешное редактирование сторис")
	gI18n.AddTranslation("ru", consts.I18nSuccessfulStoryRemove, "успешное удаление сторис")
	// en
	gI18n.AddTranslation("en", consts.I18nSuccessfulStoryPublication, "successful story publication")
	gI18n.AddTranslation("en", consts.I18nSuccessfulStoryEdit, "successful story edit")
	gI18n.AddTranslation("en", consts.I18nSuccessfulStoryRemove, "successful story remove")

	// ru
	gI18n.AddTranslation("ru", consts.I18nErrorInvalidStoryId, "некорректный id сторис")
	gI18n.AddTranslation("ru", consts.I18nErrorInvalidStoryType, "некорректный тип сторис")
	gI18n.AddTranslation("ru", consts.I18nErrorInvalidStoryTextPos, "некорректная позиция текста")
	gI18n.AddTranslation("ru", consts.I18nDatabaseError, "ошибка базы данных")
	// en
	gI18n.AddTranslation("en", consts.I18nErrorInvalidStoryId, "invalid story id")
	gI18n.AddTranslation("en", consts.I18nErrorInvalidStoryType, "invalid story type")
	gI18n.AddTranslation("en", consts.I18nErrorInvalidStoryTextPos, "invalid story text position")
	gI18n.AddTranslation("en", consts.I18nDatabaseError, "database error")
}
