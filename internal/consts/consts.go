package consts

import "time"

const (
	UserSessionTime = 72 * time.Hour
)

var (
	UserSessionPepper = [32]uint8{211, 67, 10, 222, 87, 9, 71, 240, 162, 157, 146, 177, 212, 146, 61, 200, 55, 96, 24, 70, 62, 223, 41, 103, 207, 230, 231, 135, 77, 35, 248, 152}
)

const (
	I18nGetUsersFailed = "get_users_failed"
	I18nUserNotFound   = "user_not_found"
	I18nWrongPassword  = "wrong_password"

	I18nSessionDecodeFailed = "session_decode_failed"
	I18nSessionVerifyFailed = "session_verify_failed"
	I18nSessionExpired      = "session_expired"

	I18nSuccessfulAuthorization = "successful_authorization"

	I18nSignInLegend        = "sign_in_legend"
	I18nSignInLabelUsername = "sign_in_label_username"
	I18nSignInLabelPassword = "sign_in_label_password"
	I18nSignInSubmit        = "sign_in_submit"

	I18nStoryAddLegend      = "story_add_legend"
	I18nStoryAddLabelType   = "story_add_label_type"
	I18nStoryAddOptionImage = "story_add_option_image"
	I18nStoryAddOptionVideo = "story_add_option_video"
	I18nStoryAddLabelLink   = "story_add_label_link"

	I18nStoryLabelTextLink    = "story_add_label_text"
	I18nStoryLabelTextPosLink = "story_add_label_text_pos"

	I18nStoryAddSubmit  = "story_add_submit"
	I18nStoryEditSubmit = "story_edit_submit"

	I18nStoryIndexAStoryAdd    = "index_a_story_add"
	I18nStoryIndexAStoryEdit   = "index_a_story_edit"
	I18nStoryIndexAStoryRemove = "index_a_story_remove"

	I18nStoryNotFound = "index_story_not_found"

	I18nSuccessfulStoryPublication = "successful_story_publication"
	I18nSuccessfulStoryEdit        = "successful_story_edit"
	I18nSuccessfulStoryRemove      = "successful_story_remove"

	I18nErrorInvalidStoryId      = "error_invalid_story_id"
	I18nErrorInvalidStoryType    = "error_invalid_story_type"
	I18nErrorInvalidStoryTextPos = "error_invalid_story_text_pos"
	I18nDatabaseError            = "database_error"
)
