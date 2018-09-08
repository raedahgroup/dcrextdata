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

func testPowData(t *testing.T) {
	t.Parallel()

	query := PowData()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPowDataDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
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

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPowDataQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PowData().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPowDataSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PowDatumSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPowDataExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PowDatumExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if PowDatum exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PowDatumExists to return true, but got false.")
	}
}

func testPowDataFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	powDatumFound, err := FindPowDatum(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if powDatumFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPowDataBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PowData().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPowDataOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PowData().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPowDataAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	powDatumOne := &PowDatum{}
	powDatumTwo := &PowDatum{}
	if err = randomize.Struct(seed, powDatumOne, powDatumDBTypes, false, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, powDatumTwo, powDatumDBTypes, false, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = powDatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = powDatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PowData().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPowDataCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	powDatumOne := &PowDatum{}
	powDatumTwo := &PowDatum{}
	if err = randomize.Struct(seed, powDatumOne, powDatumDBTypes, false, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, powDatumTwo, powDatumDBTypes, false, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = powDatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = powDatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func powDatumBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PowDatum) error {
	*o = PowDatum{}
	return nil
}

func powDatumAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PowDatum) error {
	*o = PowDatum{}
	return nil
}

func powDatumAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PowDatum) error {
	*o = PowDatum{}
	return nil
}

func powDatumBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PowDatum) error {
	*o = PowDatum{}
	return nil
}

func powDatumAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PowDatum) error {
	*o = PowDatum{}
	return nil
}

func powDatumBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PowDatum) error {
	*o = PowDatum{}
	return nil
}

func powDatumAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PowDatum) error {
	*o = PowDatum{}
	return nil
}

func powDatumBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PowDatum) error {
	*o = PowDatum{}
	return nil
}

func powDatumAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PowDatum) error {
	*o = PowDatum{}
	return nil
}

func testPowDataHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PowDatum{}
	o := &PowDatum{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, powDatumDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PowDatum object: %s", err)
	}

	AddPowDatumHook(boil.BeforeInsertHook, powDatumBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	powDatumBeforeInsertHooks = []PowDatumHook{}

	AddPowDatumHook(boil.AfterInsertHook, powDatumAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	powDatumAfterInsertHooks = []PowDatumHook{}

	AddPowDatumHook(boil.AfterSelectHook, powDatumAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	powDatumAfterSelectHooks = []PowDatumHook{}

	AddPowDatumHook(boil.BeforeUpdateHook, powDatumBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	powDatumBeforeUpdateHooks = []PowDatumHook{}

	AddPowDatumHook(boil.AfterUpdateHook, powDatumAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	powDatumAfterUpdateHooks = []PowDatumHook{}

	AddPowDatumHook(boil.BeforeDeleteHook, powDatumBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	powDatumBeforeDeleteHooks = []PowDatumHook{}

	AddPowDatumHook(boil.AfterDeleteHook, powDatumAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	powDatumAfterDeleteHooks = []PowDatumHook{}

	AddPowDatumHook(boil.BeforeUpsertHook, powDatumBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	powDatumBeforeUpsertHooks = []PowDatumHook{}

	AddPowDatumHook(boil.AfterUpsertHook, powDatumAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	powDatumAfterUpsertHooks = []PowDatumHook{}
}

func testPowDataInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPowDataInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(powDatumColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPowDataReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
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

func testPowDataReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PowDatumSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPowDataSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PowData().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	powDatumDBTypes = map[string]string{`Actual124h`: `double precision`, `BTC24H`: `double precision`, `Blocks24h`: `double precision`, `Blocksfound`: `double precision`, `Blocksper`: `double precision`, `Btcprice`: `character varying`, `Coinprice`: `character varying`, `CreatedOn`: `double precision`, `CreatedTime`: `date`, `Currentheight`: `double precision`, `Currentnetworkblock`: `double precision`, `Dev`: `double precision`, `Efficiency`: `double precision`, `Est`: `double precision`, `Estimatecurrent`: `double precision`, `Estimatelast24h`: `double precision`, `Estshare`: `double precision`, `Esttime`: `timestamp without time zone`, `Fees`: `double precision`, `Hashrate`: `double precision`, `Hashratelast24h`: `double precision`, `Height`: `double precision`, `ID`: `integer`, `Lastblock`: `double precision`, `Lastupdate`: `double precision`, `Luck`: `double precision`, `Mbtcmhfactor`: `double precision`, `Name`: `character varying`, `Nethashrate`: `double precision`, `Networkdiff`: `double precision`, `Networkdifficulty`: `double precision`, `Nextnetworkblock`: `double precision`, `Port`: `double precision`, `Pos`: `double precision`, `Pow`: `double precision`, `Powid`: `double precision`, `Ppshare`: `double precision`, `Progress`: `double precision`, `Rentalcurrent`: `double precision`, `Success`: `character varying`, `Timesincelast`: `double precision`, `Total`: `double precision`, `Totalkickback`: `double precision`, `Totalminers`: `double precision`, `Workers`: `double precision`}
	_               = bytes.MinRead
)

func testPowDataUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(powDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(powDatumColumns) == len(powDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPowDataSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(powDatumColumns) == len(powDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PowDatum{}
	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, powDatumDBTypes, true, powDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(powDatumColumns, powDatumPrimaryKeyColumns) {
		fields = powDatumColumns
	} else {
		fields = strmangle.SetComplement(
			powDatumColumns,
			powDatumPrimaryKeyColumns,
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

	slice := PowDatumSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPowDataUpsert(t *testing.T) {
	t.Parallel()

	if len(powDatumColumns) == len(powDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PowDatum{}
	if err = randomize.Struct(seed, &o, powDatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PowDatum: %s", err)
	}

	count, err := PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, powDatumDBTypes, false, powDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PowDatum struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PowDatum: %s", err)
	}

	count, err = PowData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
