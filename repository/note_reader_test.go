package repository_test

import (
	"testing"

	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/repository"
	"github.com/GuilhermeFujita/nlw_notes_api/repository/testdata"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type NotesReaderSuite struct {
	suite.Suite
	db     *gorm.DB
	reader repository.NotesReader
}

func (s *NotesReaderSuite) SetupTest() {
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
		testfixtures.Directory("testdata/fixtures/reader"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)

	s.Require().NoError(err)
	s.Require().NoError(fixtures.Load())

	s.db = db
	s.reader = repository.NewNoteReader(db)
}

func TestNotesReaderSuite(t *testing.T) {
	suite.Run(t, new(NotesReaderSuite))
}

func (s *NotesReaderSuite) TestShouldGetNotesWithoutFilters() {

	notes, err := s.reader.GetNotes("")
	s.Assert().NoError(err)
	s.Assert().Len(notes, 4)
	s.Assert().Equal(testdata.ExpectedNotes(), notes)
}

func (s *NotesReaderSuite) TestShouldGetNotesWhenFilterIsApplied() {
	notes, err := s.reader.GetNotes("content")
	s.Assert().NoError(err)

	s.Assert().Len(notes, 1)
	s.Assert().Equal(testdata.ExpectedFilteredNotes(), notes)
}

func (s *NotesReaderSuite) TestShouldGetASingleNoteSuccessfully() {

	notes, err := s.reader.GetNote(4)
	s.Assert().NoError(err)
	s.Assert().Equal(testdata.ExpectedSingleNote(), notes)
}

func (s *NotesReaderSuite) TestGetNote_Error_NoTable() {
	emptyDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	s.Require().NoError(err)

	s.reader = repository.NewNoteReader(emptyDB)
	_, err = s.reader.GetNote(123)
	s.Assert().NotNil(err)
	s.Assertions.Error(err, "no such table: notes")
}
