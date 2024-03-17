package utils

const UNKNOWN_DIR_NAME = "unknown"

var NOTES_DIR_NAME = ReadStringConfig("NOTES_DIR_NAME")
var TODO_DIR_NAME = ReadStringConfig("TODO_DIR_NAME")
var MAX_TITLE_LENGTH = ReadIntConfig("MAX_TITLE_LENGTH")
var MAX_CONTENT_LENGTH = ReadIntConfig("MAX_CONTENT_LENGTH")
