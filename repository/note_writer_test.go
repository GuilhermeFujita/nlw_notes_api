package repository_test

import (
	"testing"
	"time"

	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/dto"
	"github.com/GuilhermeFujita/nlw_notes_api/repository"
	"github.com/GuilhermeFujita/nlw_notes_api/repository/testdata"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type NotesWriterSuite struct {
	suite.Suite
	db     *gorm.DB
	writer repository.NotesWriter
	reader repository.NotesReader
}

func (s *NotesWriterSuite) SetupTest() {
	t := s.T()
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	s.Require().NoError(err)
	s.Require().NoError(db.AutoMigrate(&model.Note{}))

	sqlDB, err := db.DB()
	s.Require().NoError(err)
	t.Cleanup(func() {
		sqlDB.Close()
	})

	fixtures, err := testfixtures.New(
		testfixtures.Database(sqlDB),
		testfixtures.Dialect("sqlite"),
		testfixtures.Directory("testdata/fixtures/writer"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)

	s.Require().NoError(err)
	s.Require().NoError(fixtures.Load())

	s.db = db
	s.reader = repository.NewNoteReader(db)
	s.writer = repository.NewNoteWriter(db)
}

func TestNotesWriterSuite(t *testing.T) {
	suite.Run(t, new(NotesWriterSuite))
}

func (s *NotesWriterSuite) Test_ShouldDeleteNoteSuccessfully() {
	noteFound, errFound := s.reader.GetNote(1)
	s.Assert().Equal(testdata.ExpectedNoteFoundToDelete(), noteFound)
	s.Assert().NoError(errFound)

	err := s.writer.DeleteNote(noteFound)
	s.Assert().NoError(err)

	_, err = s.reader.GetNote(1)
	s.Assert().ErrorIs(err, gorm.ErrRecordNotFound)
}

func (s *NotesWriterSuite) Test_ShouldSaveNoteSuccessfully() {
	before := time.Now()

	noteToSave := dto.NoteRequestDTO{
		Content: "Note to save",
	}

	saved, errSave := s.writer.SaveNote(noteToSave)
	s.Assert().NotZero(saved.ID)
	s.Assert().Equal("Note to save", saved.Content)
	s.WithinDuration(before, saved.NoteDate, time.Second)
	s.Assert().NoError(errSave)
}

func (s *NotesWriterSuite) Test_ShouldReturnErrorOnSaveNote() {
	emptyDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	s.Require().NoError(err)

	s.writer = repository.NewNoteWriter(emptyDB)

	noteToSave := dto.NoteRequestDTO{
		Content: "Note to save",
	}
	_, err = s.writer.SaveNote(noteToSave)
	s.Assert().NotNil(err)
	s.Assertions.Error(err, "no such table: notes")
}
