// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// CommStat is an object representing the database table.
type CommStat struct {
	Date                 time.Time `boil:"date" json:"date" toml:"date" yaml:"date"`
	RedditSubscribers    int       `boil:"reddit_subscribers" json:"reddit_subscribers" toml:"reddit_subscribers" yaml:"reddit_subscribers"`
	RedditAccountsActive int       `boil:"reddit_accounts_active" json:"reddit_accounts_active" toml:"reddit_accounts_active" yaml:"reddit_accounts_active"`

	R *commStatR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L commStatL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CommStatColumns = struct {
	Date                 string
	RedditSubscribers    string
	RedditAccountsActive string
}{
	Date:                 "date",
	RedditSubscribers:    "reddit_subscribers",
	RedditAccountsActive: "reddit_accounts_active",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var CommStatWhere = struct {
	Date                 whereHelpertime_Time
	RedditSubscribers    whereHelperint
	RedditAccountsActive whereHelperint
}{
	Date:                 whereHelpertime_Time{field: "\"comm_stat\".\"date\""},
	RedditSubscribers:    whereHelperint{field: "\"comm_stat\".\"reddit_subscribers\""},
	RedditAccountsActive: whereHelperint{field: "\"comm_stat\".\"reddit_accounts_active\""},
}

// CommStatRels is where relationship names are stored.
var CommStatRels = struct {
}{}

// commStatR is where relationships are stored.
type commStatR struct {
}

// NewStruct creates a new relationship struct
func (*commStatR) NewStruct() *commStatR {
	return &commStatR{}
}

// commStatL is where Load methods for each relationship are stored.
type commStatL struct{}

var (
	commStatAllColumns            = []string{"date", "reddit_subscribers", "reddit_accounts_active"}
	commStatColumnsWithoutDefault = []string{"date", "reddit_subscribers", "reddit_accounts_active"}
	commStatColumnsWithDefault    = []string{}
	commStatPrimaryKeyColumns     = []string{"date"}
)

type (
	// CommStatSlice is an alias for a slice of pointers to CommStat.
	// This should generally be used opposed to []CommStat.
	CommStatSlice []*CommStat

	commStatQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	commStatType                 = reflect.TypeOf(&CommStat{})
	commStatMapping              = queries.MakeStructMapping(commStatType)
	commStatPrimaryKeyMapping, _ = queries.BindMapping(commStatType, commStatMapping, commStatPrimaryKeyColumns)
	commStatInsertCacheMut       sync.RWMutex
	commStatInsertCache          = make(map[string]insertCache)
	commStatUpdateCacheMut       sync.RWMutex
	commStatUpdateCache          = make(map[string]updateCache)
	commStatUpsertCacheMut       sync.RWMutex
	commStatUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single commStat record from the query.
func (q commStatQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CommStat, error) {
	o := &CommStat{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for comm_stat")
	}

	return o, nil
}

// All returns all CommStat records from the query.
func (q commStatQuery) All(ctx context.Context, exec boil.ContextExecutor) (CommStatSlice, error) {
	var o []*CommStat

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CommStat slice")
	}

	return o, nil
}

// Count returns the count of all CommStat records in the query.
func (q commStatQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count comm_stat rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q commStatQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if comm_stat exists")
	}

	return count > 0, nil
}

// CommStats retrieves all the records using an executor.
func CommStats(mods ...qm.QueryMod) commStatQuery {
	mods = append(mods, qm.From("\"comm_stat\""))
	return commStatQuery{NewQuery(mods...)}
}

// FindCommStat retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCommStat(ctx context.Context, exec boil.ContextExecutor, date time.Time, selectCols ...string) (*CommStat, error) {
	commStatObj := &CommStat{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"comm_stat\" where \"date\"=$1", sel,
	)

	q := queries.Raw(query, date)

	err := q.Bind(ctx, exec, commStatObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from comm_stat")
	}

	return commStatObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CommStat) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no comm_stat provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(commStatColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	commStatInsertCacheMut.RLock()
	cache, cached := commStatInsertCache[key]
	commStatInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			commStatAllColumns,
			commStatColumnsWithDefault,
			commStatColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(commStatType, commStatMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(commStatType, commStatMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"comm_stat\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"comm_stat\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into comm_stat")
	}

	if !cached {
		commStatInsertCacheMut.Lock()
		commStatInsertCache[key] = cache
		commStatInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CommStat.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CommStat) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	commStatUpdateCacheMut.RLock()
	cache, cached := commStatUpdateCache[key]
	commStatUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			commStatAllColumns,
			commStatPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update comm_stat, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"comm_stat\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, commStatPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(commStatType, commStatMapping, append(wl, commStatPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update comm_stat row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for comm_stat")
	}

	if !cached {
		commStatUpdateCacheMut.Lock()
		commStatUpdateCache[key] = cache
		commStatUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q commStatQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for comm_stat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for comm_stat")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CommStatSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commStatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"comm_stat\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, commStatPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in commStat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all commStat")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CommStat) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no comm_stat provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(commStatColumnsWithDefault, o)

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

	commStatUpsertCacheMut.RLock()
	cache, cached := commStatUpsertCache[key]
	commStatUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			commStatAllColumns,
			commStatColumnsWithDefault,
			commStatColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			commStatAllColumns,
			commStatPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert comm_stat, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(commStatPrimaryKeyColumns))
			copy(conflict, commStatPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"comm_stat\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(commStatType, commStatMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(commStatType, commStatMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
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
		return errors.Wrap(err, "models: unable to upsert comm_stat")
	}

	if !cached {
		commStatUpsertCacheMut.Lock()
		commStatUpsertCache[key] = cache
		commStatUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CommStat record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CommStat) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CommStat provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), commStatPrimaryKeyMapping)
	sql := "DELETE FROM \"comm_stat\" WHERE \"date\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from comm_stat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for comm_stat")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q commStatQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no commStatQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from comm_stat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for comm_stat")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CommStatSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commStatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"comm_stat\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commStatPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from commStat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for comm_stat")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CommStat) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCommStat(ctx, exec, o.Date)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommStatSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CommStatSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commStatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"comm_stat\".* FROM \"comm_stat\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commStatPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CommStatSlice")
	}

	*o = slice

	return nil
}

// CommStatExists checks if the CommStat row exists.
func CommStatExists(ctx context.Context, exec boil.ContextExecutor, date time.Time) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"comm_stat\" where \"date\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, date)
	}

	row := exec.QueryRowContext(ctx, sql, date)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if comm_stat exists")
	}

	return exists, nil
}
