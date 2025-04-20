package usecase_test

import (
	"testing"

	"github.com/GuilhermeFujita/nlw_notes_api/database/model"
	"github.com/GuilhermeFujita/nlw_notes_api/usecase"
	"github.com/GuilhermeFujita/nlw_notes_api/usecase/mocks"
	"github.com/GuilhermeFujita/nlw_notes_api/usecase/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type NoteSuite struct {
	suite.Suite
	writer *mocks.NotesWriter
	finder *mocks.NotesFinder
	uc     usecase.NoteUseCase
}

func (n *NoteSuite) SetupTest() {
	t := n.T()
	t.Helper()

	n.writer = new(mocks.NotesWriter)
	n.finder = new(mocks.NotesFinder)

	n.uc = usecase.NewNoteUseCase(n.writer, n.finder)
}

func TestNoteSuite(t *testing.T) {
	suite.Run(t, new(NoteSuite))
}

func (n *NoteSuite) Test_ShouldCreateNoteSuccessfully() {
	n.writer.On("SaveNote", testdata.NoteToCreate()).
		Return(testdata.ExpectedNoteCreated(), nil).
		Once()

	created, err := n.uc.CreateNote(testdata.NoteToCreate())

	n.Assert().Equal(int64(1), created.ID)
	n.Assert().Equal("Test note", created.Content)
	n.Assert().NoError(err)
}

func (n *NoteSuite) Test_Should_GetNotes_Without_FiltersSuccessfully() {

	n.finder.On("GetNotes", "").Return(testdata.NotesWithoutFiltersResult(), nil).Once()

	notesFound, err := n.uc.GetNotes("")

	n.Assert().NoError(err)
	n.Assert().Equal(testdata.NotesWithoutFiltersResult(), notesFound)
}

func (n *NoteSuite) Test_Should_GetNotes_With_FilterSuccessfully() {
	search := "search"

	n.finder.On("GetNotes", search).Return(testdata.FilteredNotesResult(), nil).Once()

	notesFound, err := n.uc.GetNotes(search)

	n.Assert().NoError(err)
	n.Assert().Equal(testdata.FilteredNotesResult(), notesFound)
}

func (n *NoteSuite) Test_Should_Delete_Note_Successfully() {
	n.finder.
		On("GetNote", 1).
		Return(testdata.ExpectedNoteToDelete(), nil).
		Once()

	n.writer.
		On("DeleteNote", testdata.ExpectedNoteToDelete()).
		Return(nil).
		Once()

	err := n.uc.DeleteNote(1)

	n.Assert().NoError(err)
}

func (n *NoteSuite) Test_Should_Return_Error_When_Create_Note() {
	n.writer.On("SaveNote", testdata.NoteToCreate()).
		Return(model.Note{}, assert.AnError).
		Once()

	created, err := n.uc.CreateNote(testdata.NoteToCreate())

	n.Assert().Equal(model.Note{}, created)
	n.Assert().Error(err)
}

func (n *NoteSuite) Test_Should_Return_Error_When_GetNotes_With_Filters() {
	search := "search"

	n.finder.
		On("GetNotes", search).
		Return([]model.Note{}, assert.AnError).
		Once()

	notesFound, err := n.uc.GetNotes(search)

	n.Assert().Error(err)
	n.Assert().Equal([]model.Note{}, notesFound)
}

func (n *NoteSuite) Test_Should_Return_Error_When_GetNotes_Without_Filters() {
	n.finder.
		On("GetNotes", "").
		Return([]model.Note{}, assert.AnError).
		Once()

	notesFound, err := n.uc.GetNotes("")

	n.Assert().Error(err)
	n.Assert().Equal([]model.Note{}, notesFound)
}

func (n *NoteSuite) Test_Should_Return_Error_When_Search_Note_To_Delete() {
	n.finder.
		On("GetNote", 1).
		Return(model.Note{}, assert.AnError).
		Once()

	err := n.uc.DeleteNote(1)

	n.Assert().Error(err)
	n.Assert().ErrorIs(err, assert.AnError)
	n.writer.AssertNotCalled(n.T(), "DeleteNote", mock.Anything)
}

func (n *NoteSuite) Test_Should_Return_Error_When_Delete_Note() {
	n.finder.
		On("GetNote", 1).
		Return(testdata.ExpectedNoteToDelete(), nil).
		Once()

	n.writer.
		On("DeleteNote", testdata.ExpectedNoteToDelete()).
		Return(assert.AnError).
		Once()

	err := n.uc.DeleteNote(1)

	n.Assert().Error(err)
	n.writer.AssertCalled(n.T(), "DeleteNote", testdata.ExpectedNoteToDelete())
}
