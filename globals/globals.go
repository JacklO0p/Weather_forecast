package globals

import "gorm.io/gorm"

var CurrentLocation string

var IsProgramStarted bool = false

var Db *gorm.DB
