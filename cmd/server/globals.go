package main

import (
	internalconsts "storiesservice/internal/consts"
	"storiesservice/internal/db"
	"storiesservice/internal/i18n"
)

var (
	gDatabase = db.NewDatabase(internalconsts.DbConnInfo)
	gI18n     = i18n.NewI18n("ru")
)
