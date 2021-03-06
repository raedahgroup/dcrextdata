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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// VoteReceiveTimeDeviation is an object representing the database table.
type VoteReceiveTimeDeviation struct {
	Bin                   string  `boil:"bin" json:"bin" toml:"bin" yaml:"bin"`
	BlockHeight           int64   `boil:"block_height" json:"block_height" toml:"block_height" yaml:"block_height"`
	BlockTime             int64   `boil:"block_time" json:"block_time" toml:"block_time" yaml:"block_time"`
	ReceiveTimeDifference float64 `boil:"receive_time_difference" json:"receive_time_difference" toml:"receive_time_difference" yaml:"receive_time_difference"`

	R *voteReceiveTimeDeviationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L voteReceiveTimeDeviationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VoteReceiveTimeDeviationColumns = struct {
	Bin                   string
	BlockHeight           string
	BlockTime             string
	ReceiveTimeDifference string
}{
	Bin:                   "bin",
	BlockHeight:           "block_height",
	BlockTime:             "block_time",
	ReceiveTimeDifference: "receive_time_difference",
}

// Generated where

var VoteReceiveTimeDeviationWhere = struct {
	Bin                   whereHelperstring
	BlockHeight           whereHelperint64
	BlockTime             whereHelperint64
	ReceiveTimeDifference whereHelperfloat64
}{
	Bin:                   whereHelperstring{field: "\"vote_receive_time_deviation\".\"bin\""},
	BlockHeight:           whereHelperint64{field: "\"vote_receive_time_deviation\".\"block_height\""},
	BlockTime:             whereHelperint64{field: "\"vote_receive_time_deviation\".\"block_time\""},
	ReceiveTimeDifference: whereHelperfloat64{field: "\"vote_receive_time_deviation\".\"receive_time_difference\""},
}

// VoteReceiveTimeDeviationRels is where relationship names are stored.
var VoteReceiveTimeDeviationRels = struct {
}{}

// voteReceiveTimeDeviationR is where relationships are stored.
type voteReceiveTimeDeviationR struct {
}

// NewStruct creates a new relationship struct
func (*voteReceiveTimeDeviationR) NewStruct() *voteReceiveTimeDeviationR {
	return &voteReceiveTimeDeviationR{}
}

// voteReceiveTimeDeviationL is where Load methods for each relationship are stored.
type voteReceiveTimeDeviationL struct{}

var (
	voteReceiveTimeDeviationAllColumns            = []string{"bin", "block_height", "block_time", "receive_time_difference"}
	voteReceiveTimeDeviationColumnsWithoutDefault = []string{"bin", "block_height", "block_time", "receive_time_difference"}
	voteReceiveTimeDeviationColumnsWithDefault    = []string{}
	voteReceiveTimeDeviationPrimaryKeyColumns     = []string{"block_time", "bin"}
)

type (
	// VoteReceiveTimeDeviationSlice is an alias for a slice of pointers to VoteReceiveTimeDeviation.
	// This should generally be used opposed to []VoteReceiveTimeDeviation.
	VoteReceiveTimeDeviationSlice []*VoteReceiveTimeDeviation

	voteReceiveTimeDeviationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	voteReceiveTimeDeviationType                 = reflect.TypeOf(&VoteReceiveTimeDeviation{})
	voteReceiveTimeDeviationMapping              = queries.MakeStructMapping(voteReceiveTimeDeviationType)
	voteReceiveTimeDeviationPrimaryKeyMapping, _ = queries.BindMapping(voteReceiveTimeDeviationType, voteReceiveTimeDeviationMapping, voteReceiveTimeDeviationPrimaryKeyColumns)
	voteReceiveTimeDeviationInsertCacheMut       sync.RWMutex
	voteReceiveTimeDeviationInsertCache          = make(map[string]insertCache)
	voteReceiveTimeDeviationUpdateCacheMut       sync.RWMutex
	voteReceiveTimeDeviationUpdateCache          = make(map[string]updateCache)
	voteReceiveTimeDeviationUpsertCacheMut       sync.RWMutex
	voteReceiveTimeDeviationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single voteReceiveTimeDeviation record from the query.
func (q voteReceiveTimeDeviationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*VoteReceiveTimeDeviation, error) {
	o := &VoteReceiveTimeDeviation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for vote_receive_time_deviation")
	}

	return o, nil
}

// All returns all VoteReceiveTimeDeviation records from the query.
func (q voteReceiveTimeDeviationQuery) All(ctx context.Context, exec boil.ContextExecutor) (VoteReceiveTimeDeviationSlice, error) {
	var o []*VoteReceiveTimeDeviation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to VoteReceiveTimeDeviation slice")
	}

	return o, nil
}

// Count returns the count of all VoteReceiveTimeDeviation records in the query.
func (q voteReceiveTimeDeviationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count vote_receive_time_deviation rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q voteReceiveTimeDeviationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if vote_receive_time_deviation exists")
	}

	return count > 0, nil
}

// VoteReceiveTimeDeviations retrieves all the records using an executor.
func VoteReceiveTimeDeviations(mods ...qm.QueryMod) voteReceiveTimeDeviationQuery {
	mods = append(mods, qm.From("\"vote_receive_time_deviation\""))
	return voteReceiveTimeDeviationQuery{NewQuery(mods...)}
}

// FindVoteReceiveTimeDeviation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVoteReceiveTimeDeviation(ctx context.Context, exec boil.ContextExecutor, blockTime int64, bin string, selectCols ...string) (*VoteReceiveTimeDeviation, error) {
	voteReceiveTimeDeviationObj := &VoteReceiveTimeDeviation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vote_receive_time_deviation\" where \"block_time\"=$1 AND \"bin\"=$2", sel,
	)

	q := queries.Raw(query, blockTime, bin)

	err := q.Bind(ctx, exec, voteReceiveTimeDeviationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from vote_receive_time_deviation")
	}

	return voteReceiveTimeDeviationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *VoteReceiveTimeDeviation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no vote_receive_time_deviation provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(voteReceiveTimeDeviationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	voteReceiveTimeDeviationInsertCacheMut.RLock()
	cache, cached := voteReceiveTimeDeviationInsertCache[key]
	voteReceiveTimeDeviationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			voteReceiveTimeDeviationAllColumns,
			voteReceiveTimeDeviationColumnsWithDefault,
			voteReceiveTimeDeviationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(voteReceiveTimeDeviationType, voteReceiveTimeDeviationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(voteReceiveTimeDeviationType, voteReceiveTimeDeviationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vote_receive_time_deviation\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vote_receive_time_deviation\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into vote_receive_time_deviation")
	}

	if !cached {
		voteReceiveTimeDeviationInsertCacheMut.Lock()
		voteReceiveTimeDeviationInsertCache[key] = cache
		voteReceiveTimeDeviationInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the VoteReceiveTimeDeviation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *VoteReceiveTimeDeviation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	voteReceiveTimeDeviationUpdateCacheMut.RLock()
	cache, cached := voteReceiveTimeDeviationUpdateCache[key]
	voteReceiveTimeDeviationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			voteReceiveTimeDeviationAllColumns,
			voteReceiveTimeDeviationPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update vote_receive_time_deviation, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vote_receive_time_deviation\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, voteReceiveTimeDeviationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(voteReceiveTimeDeviationType, voteReceiveTimeDeviationMapping, append(wl, voteReceiveTimeDeviationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update vote_receive_time_deviation row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for vote_receive_time_deviation")
	}

	if !cached {
		voteReceiveTimeDeviationUpdateCacheMut.Lock()
		voteReceiveTimeDeviationUpdateCache[key] = cache
		voteReceiveTimeDeviationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q voteReceiveTimeDeviationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for vote_receive_time_deviation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for vote_receive_time_deviation")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VoteReceiveTimeDeviationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), voteReceiveTimeDeviationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vote_receive_time_deviation\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, voteReceiveTimeDeviationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in voteReceiveTimeDeviation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all voteReceiveTimeDeviation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *VoteReceiveTimeDeviation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no vote_receive_time_deviation provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(voteReceiveTimeDeviationColumnsWithDefault, o)

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

	voteReceiveTimeDeviationUpsertCacheMut.RLock()
	cache, cached := voteReceiveTimeDeviationUpsertCache[key]
	voteReceiveTimeDeviationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			voteReceiveTimeDeviationAllColumns,
			voteReceiveTimeDeviationColumnsWithDefault,
			voteReceiveTimeDeviationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			voteReceiveTimeDeviationAllColumns,
			voteReceiveTimeDeviationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert vote_receive_time_deviation, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(voteReceiveTimeDeviationPrimaryKeyColumns))
			copy(conflict, voteReceiveTimeDeviationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"vote_receive_time_deviation\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(voteReceiveTimeDeviationType, voteReceiveTimeDeviationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(voteReceiveTimeDeviationType, voteReceiveTimeDeviationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert vote_receive_time_deviation")
	}

	if !cached {
		voteReceiveTimeDeviationUpsertCacheMut.Lock()
		voteReceiveTimeDeviationUpsertCache[key] = cache
		voteReceiveTimeDeviationUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single VoteReceiveTimeDeviation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *VoteReceiveTimeDeviation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no VoteReceiveTimeDeviation provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), voteReceiveTimeDeviationPrimaryKeyMapping)
	sql := "DELETE FROM \"vote_receive_time_deviation\" WHERE \"block_time\"=$1 AND \"bin\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from vote_receive_time_deviation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for vote_receive_time_deviation")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q voteReceiveTimeDeviationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no voteReceiveTimeDeviationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from vote_receive_time_deviation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for vote_receive_time_deviation")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VoteReceiveTimeDeviationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), voteReceiveTimeDeviationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vote_receive_time_deviation\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, voteReceiveTimeDeviationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from voteReceiveTimeDeviation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for vote_receive_time_deviation")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *VoteReceiveTimeDeviation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVoteReceiveTimeDeviation(ctx, exec, o.BlockTime, o.Bin)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VoteReceiveTimeDeviationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VoteReceiveTimeDeviationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), voteReceiveTimeDeviationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vote_receive_time_deviation\".* FROM \"vote_receive_time_deviation\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, voteReceiveTimeDeviationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VoteReceiveTimeDeviationSlice")
	}

	*o = slice

	return nil
}

// VoteReceiveTimeDeviationExists checks if the VoteReceiveTimeDeviation row exists.
func VoteReceiveTimeDeviationExists(ctx context.Context, exec boil.ContextExecutor, blockTime int64, bin string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vote_receive_time_deviation\" where \"block_time\"=$1 AND \"bin\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, blockTime, bin)
	}
	row := exec.QueryRowContext(ctx, sql, blockTime, bin)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if vote_receive_time_deviation exists")
	}

	return exists, nil
}
