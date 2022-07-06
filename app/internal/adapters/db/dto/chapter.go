package dto

// CreateChapterInput used by CreateChapter.
type CreateChapterInput struct {
	ChapterName  string
	ChapterNum   string
	RegulationID uint64
}

// CreateChapterResonse returned by CreateChapter.
type CreateChapterOutput struct {
	ChapterID uint64
}
