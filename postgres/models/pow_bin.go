// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// PowBin is an object representing the database table.
type PowBin struct {
	Time         int64       `boil:"time" json:"time" toml:"time" yaml:"time"`
	PoolHashrate null.String `boil:"pool_hashrate" json:"pool_hashrate,omitempty" toml:"pool_hashrate" yaml:"pool_hashrate,omitempty"`
	Workers      null.Int    `boil:"workers" json:"workers,omitempty" toml:"workers" yaml:"workers,omitempty"`
	Bin          string      `boil:"bin" json:"bin" toml:"bin" yaml:"bin"`
	Source       string      `boil:"source" json:"source" toml:"source" yaml:"source"`

	R *powBinR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L powBinL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PowBinColumns = struct {
	Time         string
	PoolHashrate string
	Workers      string
	Bin          string
	Source       string
}{
	Time:         "time",
	PoolHashrate: "pool_hashrate",
	Workers:      "workers",
	Bin:          "bin",
	Source:       "source",
}

// Generated where

var PowBinWhere = struct {
	Time         whereHelperint64
	PoolHashrate whereHelpernull_String
	Workers      whereHelpernull_Int
	Bin          whereHelperstring
	Source       whereHelperstring
}{
	Time:         whereHelperint64{field: "\"pow_bin\".\"time\""},
	PoolHashrate: whereHelpernull_String{field: "\"pow_bin\".\"pool_hashrate\""},
	Workers:      whereHelpernull_Int{field: "\"pow_bin\".\"workers\""},
	Bin:          whereHelperstring{field: "\"pow_bin\".\"bin\""},
	Source:       whereHelperstring{field: "\"pow_bin\".\"source\""},
}

// PowBinRels is where relationship names are stored.
var PowBinRels = struct {
}{}

// powBinR is where relationships are stored.
type powBinR struct {
}

// NewStruct creates a new relationship struct
func (*powBinR) NewStruct() *powBinR {
	return &powBinR{}
}

// powBinL is where Load methods for each relationship are stored.
type powBinL struct{}

var (
	powBinAllColumns            = []string{"time", "pool_hashrate", "workers", "bin", "source"}
	powBinColumnsWithoutDefault = []string{"time", "pool_hashrate", "workers", "bin", "source"}
	powBinColumnsWithDefault    = []string{}
	powBinPrimaryKeyColumns     = []string{"time", "source", "bin"}
)

type (
	// PowBinSlice is an alias for a slice of pointers to PowBin.
	// This should generally be used opposed to []PowBin.
	PowBinSlice []*PowBin

	powBinQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	powBinType                 = reflect.TypeOf(&PowBin{})
	powBinMapping              = queries.MakeStructMapping(powBinType)
	powBinPrimaryKeyMapping, _ = queries.BindMapping(powBinType, powBinMapping, powBinPrimaryKeyColumns)
	powBinInsertCacheMut       sync.RWMutex
	powBinInsertCache          = make(map[string]insertCache)
	powBinUpdateCacheMut       sync.RWMutex
	powBinUpdateCache          = make(map[string]updateCache)
	powBinUpsertCacheMut       sync.RWMutex
	powBinUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single powBin record from the query.
func (q powBinQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PowBin, error) {
	o := &PowBin{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pow_bin")
	}

	return o, nil
}

// All returns all PowBin records from the query.
func (q powBinQuery) All(ctx context.Context, exec boil.ContextExecutor) (PowBinSlice, error) {
	var o []*PowBin

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PowBin slice")
	}

	return o, nil
}

// Count returns the count of all PowBin records in the query.
func (q powBinQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pow_bin rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q powBinQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pow_bin exists")
	}

	return count > 0, nil
}

// PowBins retrieves all the records using an executor.
func PowBins(mods ...qm.QueryMod) powBinQuery {
	mods = append(mods, qm.From("\"pow_bin\""))
	return powBinQuery{NewQuery(mods...)}
}

// FindPowBin retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPowBin(ctx context.Context, exec boil.ContextExecutor, time int64, source string, bin string, selectCols ...string) (*PowBin, error) {
	powBinObj := &PowBin{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pow_bin\" where \"time\"=$1 AND \"source\"=$2 AND \"bin\"=$3", sel,
	)

	q := queries.Raw(query, time, source, bin)

	err := q.Bind(ctx, exec, powBinObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pow_bin")
	}

	return powBinObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PowBin) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pow_bin provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(powBinColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	powBinInsertCacheMut.RLock()
	cache, cached := powBinInsertCache[key]
	powBinInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			powBinAllColumns,
			powBinColumnsWithDefault,
			powBinColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(powBinType, powBinMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(powBinType, powBinMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"pow_bin\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"pow_bin\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into pow_bin")
	}

	if !cached {
		powBinInsertCacheMut.Lock()
		powBinInsertCache[key] = cache
		powBinInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PowBin.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PowBin) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	powBinUpdateCacheMut.RLock()
	cache, cached := powBinUpdateCache[key]
	powBinUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			powBinAllColumns,
			powBinPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update pow_bin, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pow_bin\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, powBinPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(powBinType, powBinMapping, append(wl, powBinPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update pow_bin row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for pow_bin")
	}

	if !cached {
		powBinUpdateCacheMut.Lock()
		powBinUpdateCache[key] = cache
		powBinUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q powBinQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for pow_bin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for pow_bin")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PowBinSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), powBinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"pow_bin\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, powBinPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in powBin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all powBin")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PowBin) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pow_bin provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(powBinColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	powBinUpsertCacheMut.RLock()
	cache, cached := powBinUpsertCache[key]
	powBinUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			powBinAllColumns,
			powBinColumnsWithDefault,
			powBinColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			powBinAllColumns,
			powBinPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert pow_bin, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(powBinPrimaryKeyColumns))
			copy(conflict, powBinPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"pow_bin\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(powBinType, powBinMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(powBinType, powBinMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert pow_bin")
	}

	if !cached {
		powBinUpsertCacheMut.Lock()
		powBinUpsertCache[key] = cache
		powBinUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single PowBin record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PowBin) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PowBin provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), powBinPrimaryKeyMapping)
	sql := "DELETE FROM \"pow_bin\" WHERE \"time\"=$1 AND \"source\"=$2 AND \"bin\"=$3"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from pow_bin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for pow_bin")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q powBinQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no powBinQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pow_bin")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pow_bin")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PowBinSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), powBinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"pow_bin\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, powBinPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from powBin slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pow_bin")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PowBin) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPowBin(ctx, exec, o.Time, o.Source, o.Bin)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PowBinSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PowBinSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), powBinPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"pow_bin\".* FROM \"pow_bin\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, powBinPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PowBinSlice")
	}

	*o = slice

	return nil
}

// PowBinExists checks if the PowBin row exists.
func PowBinExists(ctx context.Context, exec boil.ContextExecutor, time int64, source string, bin string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"pow_bin\" where \"time\"=$1 AND \"source\"=$2 AND \"bin\"=$3 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, time, source, bin)
	}
	row := exec.QueryRowContext(ctx, sql, time, source, bin)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pow_bin exists")
	}

	return exists, nil
}