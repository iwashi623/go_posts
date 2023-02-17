package store

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/iwashi623/go_posts/clock"
	"github.com/iwashi623/go_posts/entity"
	"github.com/iwashi623/go_posts/testutil"
	"github.com/jmoiron/sqlx"
)

func TestRepository_ListPosts(t *testing.T) {
	ctx := context.Background()

	// entity.Postsを取得する際、他のテストケースと混ざるとTestが失敗する
	// そのため､トランザクションを利用してテストケースごとにデータを分離する
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	t.Cleanup(func() {
		_ = tx.Rollback()
	})
	if err != nil {
		t.Fatal(err)
	}
	expected := preparePosts(ctx, t, tx)

	sut := &Repository{}
	actual, err := sut.ListPosts(ctx, tx)
	if err != nil {
		t.Fatalf("failed to ListPosts: %s", err)
	}
	if d := cmp.Diff(expected, actual); d != "" {
		t.Errorf("ListPosts() mismatch (-want +got):\n%s", d)
	}
}

func preparePosts(ctx context.Context, t *testing.T, tx *sqlx.Tx) entity.Posts {
	t.Helper()
	// 一度きれいにする
	if _, err := tx.ExecContext(ctx, "DELETE FROM post"); err != nil {
		t.Fatalf("failed to DELETE FROM post: %s", err)
	}
	c := clock.FixedClocker{}

	posts := entity.Posts{
		{
			Title: "title1", Body: "body1", Status: entity.PostStatusDraft,
			CreatedAt: c.Now(), UpdatedAt: c.Now(),
		},
		{
			Title: "title2", Body: "body2", Status: entity.PostStatusPublished,
			CreatedAt: c.Now(), UpdatedAt: c.Now(),
		},
		{
			Title: "title3", Body: "body3", Status: entity.PostStatusPrivate,
			CreatedAt: c.Now(), UpdatedAt: c.Now(),
		},
	}

	result, err := tx.ExecContext(ctx,
		`INSERT INTO post (title, body, status, created_at, updated_at)
		 VALUES
		    (?, ?, ?, ?, ?),
			(?, ?, ?, ?, ?),
			(?, ?, ?, ?, ?);`,
		posts[0].Title, posts[0].Body, posts[0].Status, posts[0].CreatedAt, posts[0].UpdatedAt,
		posts[1].Title, posts[1].Body, posts[1].Status, posts[1].CreatedAt, posts[1].UpdatedAt,
		posts[2].Title, posts[2].Body, posts[2].Status, posts[2].CreatedAt, posts[2].UpdatedAt,
	)
	if err != nil {
		t.Fatalf("failed to INSERT INTO post: %s", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	posts[0].ID = entity.PostID(id)
	posts[1].ID = entity.PostID(id + 1)
	posts[2].ID = entity.PostID(id + 2)
	return posts
}

func TestRepository_AddPost(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	c := clock.FixedClocker{}
	var wantID int = 20
	okPost := &entity.Post{
		Title: "OK title", Body: "OK 本文", Status: entity.PostStatusDraft,
		CreatedAt: c.Now(), UpdatedAt: c.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = db.Close()
	})
	mock.ExpectExec(
		// エスケープが必要
		`INSERT INTO post \(title, body, status, created_at, updated_at\) VALUES \(\?, \?, \?, \?, \?\)`,
	).WithArgs(okPost.Title, okPost.Body, okPost.Status, okPost.CreatedAt, okPost.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(int64(wantID), 1))

	xdb := sqlx.NewDb(db, "mysql")
	r := &Repository{Clocker: c}
	if err := r.AddPost(ctx, xdb, okPost); err != nil {
		t.Fatalf("failed to AddPost: %s", err)
	}
}
