// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRedditInfos(t *testing.T) {
	t.Parallel()

	query := RedditInfos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRedditInfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RedditInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRedditInfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RedditInfos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RedditInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRedditInfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RedditInfoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RedditInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRedditInfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RedditInfoExists(ctx, tx, o.Date)
	if err != nil {
		t.Errorf("Unable to check if RedditInfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RedditInfoExists to return true, but got false.")
	}
}

func testRedditInfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	redditInfoFound, err := FindRedditInfo(ctx, tx, o.Date)
	if err != nil {
		t.Error(err)
	}

	if redditInfoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRedditInfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RedditInfos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRedditInfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RedditInfos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRedditInfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	redditInfoOne := &RedditInfo{}
	redditInfoTwo := &RedditInfo{}
	if err = randomize.Struct(seed, redditInfoOne, redditInfoDBTypes, false, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, redditInfoTwo, redditInfoDBTypes, false, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = redditInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = redditInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RedditInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRedditInfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	redditInfoOne := &RedditInfo{}
	redditInfoTwo := &RedditInfo{}
	if err = randomize.Struct(seed, redditInfoOne, redditInfoDBTypes, false, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, redditInfoTwo, redditInfoDBTypes, false, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = redditInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = redditInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RedditInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testRedditInfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RedditInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRedditInfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(redditInfoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RedditInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRedditInfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRedditInfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RedditInfoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRedditInfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RedditInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	redditInfoDBTypes = map[string]string{`Date`: `timestamp without time zone`, `Subscribers`: `integer`, `AccountsActive`: `integer`}
	_                 = bytes.MinRead
)

func testRedditInfosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(redditInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(redditInfoAllColumns) == len(redditInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RedditInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRedditInfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(redditInfoAllColumns) == len(redditInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RedditInfo{}
	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RedditInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, redditInfoDBTypes, true, redditInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(redditInfoAllColumns, redditInfoPrimaryKeyColumns) {
		fields = redditInfoAllColumns
	} else {
		fields = strmangle.SetComplement(
			redditInfoAllColumns,
			redditInfoPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RedditInfoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRedditInfosUpsert(t *testing.T) {
	t.Parallel()

	if len(redditInfoAllColumns) == len(redditInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RedditInfo{}
	if err = randomize.Struct(seed, &o, redditInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RedditInfo: %s", err)
	}

	count, err := RedditInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, redditInfoDBTypes, false, redditInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RedditInfo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RedditInfo: %s", err)
	}

	count, err = RedditInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}