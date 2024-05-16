//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetDraftNote(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/projects/1/merge_requests/4329/draft_notes/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_draft_note.json")
	})

	note, _, err := client.DraftNotes.GetDraftNote("1", 4329, 3)
	if err != nil {
		t.Fatal(err)
	}

	want := &DraftNote{
		ID:                37349978,
		AuthorID:          10271899,
		MergeRequestID:    291473309,
		ResolveDiscussion: false,
		DiscussionID:      "",
		Note:              "Some draft note",
		CommitID:          "",
		LineCode:          "",
		Position:          nil,
	}

	if !reflect.DeepEqual(note, want) {
		t.Errorf("DraftNotes.GetDraftNote want %#v, got %#v", note, want)
	}
}

func TestListDraftNotes(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/projects/1/merge_requests/4329/draft_notes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_draft_notes.json")
	})

	notes, _, err := client.DraftNotes.ListDraftNotes("1", 4329, nil)
	if err != nil {
		t.Fatal(err)
	}

	want := []*DraftNote{
		{
			ID:                37349978,
			AuthorID:          10271899,
			MergeRequestID:    291473309,
			ResolveDiscussion: false,
			DiscussionID:      "",
			Note:              "Some draft note",
			CommitID:          "",
			LineCode:          "",
			Position:          nil,
		},
		{
			ID:                37349979,
			AuthorID:          10271899,
			MergeRequestID:    291473309,
			ResolveDiscussion: false,
			DiscussionID:      "",
			Note:              "Some draft note 2",
			CommitID:          "",
			LineCode:          "3dacf79e0d779e2baa1c700cf56510e42f55cf85_10_9",
			Position: &NotePosition{
				BaseSHA:      "64581c4ee41beb44d943d7801f82d9038e25e453",
				StartSHA:     "87bffbff93bf334889780f54ae1922355661f867",
				HeadSHA:      "2c972dbf9094c380f5f00dcd8112d2c69b24c859",
				OldPath:      "src/some-dir/some-file.js",
				NewPath:      "src/some-dir/some-file.js",
				PositionType: "text",
				NewLine:      9,
				LineRange: &LineRange{
					StartRange: &LinePosition{
						LineCode: "3dacf79e0d779e2baa1c700cf56510e42f55cf85_10_9",
						Type:     "new",
						NewLine:  9,
					},
					EndRange: &LinePosition{
						LineCode: "3dacf79e0d779e2baa1c700cf56510e42f55cf85_10_9",
						Type:     "new",
						NewLine:  9,
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(notes, want) {
		t.Errorf("DraftNotes.GetDraftNote want %#v, got %#v", notes, want)
	}
}

func TestCreateDraftNote(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/projects/1/merge_requests/4329/draft_notes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/create_draft_note.json")
	})

	note, _, err := client.DraftNotes.CreateDraftNote("1", 4329, &CreateDraftNoteOptions{
		Note: Ptr("Some new draft note"),
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &DraftNote{
		ID:                37349980,
		AuthorID:          10271899,
		MergeRequestID:    291473309,
		ResolveDiscussion: false,
		DiscussionID:      "",
		Note:              "Some new draft note",
		CommitID:          "",
		LineCode:          "",
		Position:          nil,
	}

	if !reflect.DeepEqual(note, want) {
		t.Errorf("DraftNotes.GetDraftNote want %#v, got %#v", note, want)
	}
}

func TestUpdateDraftNote(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/projects/1/merge_requests/4329/draft_notes/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		mustWriteHTTPResponse(t, w, "testdata/update_draft_note.json")
	})

	note, _, err := client.DraftNotes.UpdateDraftNote("1", 4329, 3, &UpdateDraftNoteOptions{
		Note: Ptr("Some changed draft note"),
	})
	if err != nil {
		t.Fatal(err)
	}

	want := &DraftNote{
		ID:                37349980,
		AuthorID:          10271899,
		MergeRequestID:    291473309,
		ResolveDiscussion: false,
		DiscussionID:      "",
		Note:              "Some changed draft note",
		CommitID:          "",
		LineCode:          "",
		Position:          nil,
	}

	if !reflect.DeepEqual(note, want) {
		t.Errorf("DraftNotes.UpdateDraftNote want %#v, got %#v", note, want)
	}
}

func TestDeleteDraftNote(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/projects/1/merge_requests/4329/draft_notes/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.DraftNotes.DeleteDraftNote("1", 4329, 3)
	if err != nil {
		t.Errorf("DraftNotes.DeleteDraftNote returned error: %v", err)
	}
}

func TestPublishDraftNote(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/projects/1/merge_requests/4329/draft_notes/3/publish", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	_, err := client.DraftNotes.PublishDraftNote("1", 4329, 3)
	if err != nil {
		t.Errorf("DraftNotes.PublishDraftNote returned error: %v", err)
	}
}

func TestPublishAllDraftNotes(t *testing.T) {
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/projects/1/merge_requests/4329/draft_notes/bulk_publish", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	_, err := client.DraftNotes.PublishAllDraftNotes("1", 4329)
	if err != nil {
		t.Errorf("DraftNotes.PublishAllDraftNotes returned error: %v", err)
	}
}
