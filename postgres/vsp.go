// Copyright (c) 2018-2019 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/raedahgroup/dcrextdata/app/helpers"
	"github.com/raedahgroup/dcrextdata/cache"
	"github.com/raedahgroup/dcrextdata/datasync"
	"github.com/raedahgroup/dcrextdata/postgres/models"
	"github.com/raedahgroup/dcrextdata/vsp"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/types"
)

var (
	vspTickExistsErr = fmt.Errorf("VSPTick exists")
)

func (pg *PgDb) VspTableName() string {
	return models.TableNames.VSP
}

func (pg *PgDb) VspTickTableName() string {
	return models.TableNames.VSPTick
}

// StoreVSPs attempts to store the vsp responses by calling storeVspResponseG and returning
// a slice of errors
func (pg *PgDb) StoreVSPs(ctx context.Context, data vsp.Response) (int, []error) {
	if ctx.Err() != nil {
		return 0, []error{ctx.Err()}
	}
	errs := make([]error, 0, len(data))
	completed := 0
	for name, tick := range data {
		err := pg.storeVspResponse(ctx, name, tick)
		if err == nil {
			completed++
		} else if err != vspTickExistsErr {
			log.Trace(err)
			errs = append(errs, err)
		}
		if ctx.Err() != nil {
			return 0, append(errs, ctx.Err())
		}
	}
	if completed == 0 {
		log.Info("Unable to store any vsp entry")
	}
	return completed, errs
}

func (pg *PgDb) storeVspResponse(ctx context.Context, name string, resp *vsp.ResposeData) error {
	txr, err := pg.db.Begin()
	if err != nil {
		return err
	}

	pool, err := models.VSPS(models.VSPWhere.Name.EQ(null.StringFrom(name))).One(ctx, pg.db)
	if err == sql.ErrNoRows {
		pool = responseToVSP(name, resp)
		err := pg.tryInsert(ctx, txr, pool)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	vspTick := responseToVSPTick(pool.ID, resp)

	err = vspTick.Insert(ctx, pg.db, boil.Infer())
	if err != nil {
		errR := txr.Rollback()
		if errR != nil {
			return err
		}
		if strings.Contains(err.Error(), "unique constraint") {
			return vspTickExistsErr
		}
		return err
	}

	err = txr.Commit()
	if err != nil {
		return txr.Rollback()
	}
	return nil
}

func responseToVSP(name string, resp *vsp.ResposeData) *models.VSP {
	return &models.VSP{
		Name:                 null.StringFrom(name),
		APIEnabled:           null.BoolFrom(resp.APIEnabled),
		APIVersionsSupported: types.Int64Array(resp.APIVersionsSupported),
		Network:              null.StringFrom(resp.Network),
		URL:                  null.StringFrom(resp.URL),
		Launched:             null.TimeFrom(helpers.UnixTime(resp.Launched)),
	}
}

func responseToVSPTick(poolID int, resp *vsp.ResposeData) *models.VSPTick {
	return &models.VSPTick{
		VSPID:            poolID,
		Immature:         resp.Immature,
		Live:             resp.Live,
		Voted:            resp.Voted,
		Missed:           resp.Missed,
		PoolFees:         resp.PoolFees,
		ProportionLive:   resp.ProportionLive,
		ProportionMissed: resp.ProportionMissed,
		UserCount:        resp.UserCount,
		UsersActive:      resp.UserCountActive,
		Time:             helpers.UnixTime(resp.LastUpdated),
	}
}

func (pg *PgDb) FetchVSPs(ctx context.Context) ([]vsp.VSPDto, error) {
	vspData, err := models.VSPS(qm.OrderBy(models.VSPColumns.URL), qm.OrderBy(models.VSPColumns.Name)).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var result []vsp.VSPDto
	for _, item := range vspData {
		parsedURL, err := url.Parse(item.URL.String)
		if err != nil {
			return nil, err
		}
		result = append(result, vsp.VSPDto{
			Name:                 item.Name.String,
			APIEnabled:           item.APIEnabled.Bool,
			APIVersionsSupported: item.APIVersionsSupported,
			Network:              item.Network.String,
			URL:                  item.URL.String,
			Host:                 parsedURL.Host,
			Launched:             item.Launched.Time,
		})
	}

	return result, nil
}

func (pg *PgDb) AddVspSourceFromSync(ctx context.Context, vspData interface{}) error {
	vspDto := vspData.(vsp.VSPDto)
	count, _ := models.VSPS(models.VSPWhere.Name.EQ(null.StringFrom(vspDto.Name))).Count(ctx, pg.db)
	if count > 0 {
		return nil
	}
	vspModel := models.VSP{
		ID:                   vspDto.ID,
		Name:                 null.StringFrom(vspDto.Name),
		APIEnabled:           null.BoolFrom(vspDto.APIEnabled),
		APIVersionsSupported: vspDto.APIVersionsSupported,
		Network:              null.StringFrom(vspDto.Network),
		URL:                  null.StringFrom(vspDto.URL),
		Launched:             null.TimeFrom(vspDto.Launched),
	}
	err := vspModel.Insert(ctx, pg.db, boil.Infer())
	return err
}

func (pg *PgDb) FetchVspSourcesForSync(ctx context.Context, lastID int64, skip, take int) ([]vsp.VSPDto, int64, error) {
	vspData, err := models.VSPS(
		models.VSPWhere.ID.GT(int(lastID)),
		qm.Offset(skip), qm.Limit(take)).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	var result []vsp.VSPDto
	for _, item := range vspData {
		parsedURL, err := url.Parse(item.URL.String)
		if err != nil {
			return nil, 0, err
		}
		result = append(result, vsp.VSPDto{
			ID:                   item.ID,
			Name:                 item.Name.String,
			APIEnabled:           item.APIEnabled.Bool,
			APIVersionsSupported: item.APIVersionsSupported,
			Network:              item.Network.String,
			URL:                  item.URL.String,
			Host:                 parsedURL.Host,
			Launched:             item.Launched.Time,
		})
	}

	totalCount, err := models.VSPS(models.VSPWhere.ID.GT(int(lastID))).Count(ctx, pg.db)

	return result, totalCount, err
}

// VSPTicks
func (pg *PgDb) FetchVspTicksForSync(ctx context.Context, lastID int64, skip, take int) ([]datasync.VSPTickSyncDto, int64, error) {
	vspIdQuery := models.VSPTickWhere.ID.GT(int(lastID))

	vspTickSlice, err := models.VSPTicks(
		vspIdQuery,
		qm.OrderBy(models.VSPTickColumns.ID),
		qm.Limit(take), qm.Offset(skip)).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	vspTickCount, err := models.VSPTicks(vspIdQuery).Count(ctx, pg.db)

	var vspTicks []datasync.VSPTickSyncDto
	for _, tick := range vspTickSlice {
		vspTicks = append(vspTicks, pg.vspTickModelToSyncDto(tick))
	}

	return vspTicks, vspTickCount, nil
}

func (pg *PgDb) AddVspTicksFromSync(ctx context.Context, tick datasync.VSPTickSyncDto) error {
	if _, err := models.VSPTicks(models.VSPTickWhere.VSPID.EQ(tick.VSPID),
		models.VSPTickWhere.Time.EQ(tick.Time)).One(ctx, pg.db); err == nil {
		return nil // record exists
	}
	tickModel := models.VSPTick{
		ID:               tick.ID,
		VSPID:            tick.VSPID,
		Immature:         tick.Immature,
		Live:             tick.Live,
		Voted:            tick.Voted,
		Missed:           tick.Missed,
		PoolFees:         tick.PoolFees,
		ProportionLive:   tick.ProportionLive,
		ProportionMissed: tick.ProportionMissed,
		UserCount:        tick.UserCount,
		UsersActive:      tick.UsersActive,
		Time:             tick.Time,
	}

	return tickModel.Insert(ctx, pg.db, boil.Infer())
}

func (pg *PgDb) FiltredVSPTicks(ctx context.Context, vspName string, offset, limit int) ([]vsp.VSPTickDto, int64, error) {
	vspInfo, err := models.VSPS(models.VSPWhere.Name.EQ(null.StringFrom(vspName))).One(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	vspIdQuery := models.VSPTickWhere.VSPID.EQ(vspInfo.ID)
	vspTickSlice, err := models.VSPTicks(qm.Load(models.VSPTickRels.VSP), vspIdQuery, qm.Limit(limit), qm.Offset(offset), qm.OrderBy(fmt.Sprintf("%s DESC", models.VSPTickColumns.Time))).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	vspTickCount, err := models.VSPTicks(qm.Load(models.VSPTickRels.VSP), vspIdQuery).Count(ctx, pg.db)

	vspTicks := []vsp.VSPTickDto{}
	for _, tick := range vspTickSlice {
		vspTicks = append(vspTicks, pg.vspTickModelToDto(tick))
	}

	return vspTicks, vspTickCount, nil
}

// VSPTicks
// todo impliment sorting for VSP ticks as it is currently been sorted by time
func (pg *PgDb) AllVSPTicks(ctx context.Context, offset, limit int) ([]vsp.VSPTickDto, int64, error) {
	vspTickSlice, err := models.VSPTicks(qm.Load(models.VSPTickRels.VSP), qm.Limit(limit), qm.Offset(offset), qm.OrderBy(fmt.Sprintf("%s DESC", models.VSPTickColumns.Time))).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	vspTickCount, err := models.VSPTicks().Count(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	vspTicks := []vsp.VSPTickDto{}
	for _, tick := range vspTickSlice {
		vspTicks = append(vspTicks, pg.vspTickModelToDto(tick))
	}

	return vspTicks, vspTickCount, nil
}

func (pg *PgDb) vspTickModelToDto(tick *models.VSPTick) vsp.VSPTickDto {
	return vsp.VSPTickDto{
		ID:               tick.ID,
		VSP:              tick.R.VSP.Name.String,
		Time:             tick.Time.Format(dateTemplate),
		Immature:         tick.Immature,
		Live:             tick.Live,
		Missed:           tick.Missed,
		PoolFees:         tick.PoolFees,
		ProportionLive:   RoundValue(tick.ProportionLive),
		ProportionMissed: RoundValue(tick.ProportionMissed),
		UserCount:        tick.UserCount,
		UsersActive:      tick.UsersActive,
		Voted:            tick.Voted,
	}
}

func (pg *PgDb) vspTickModelToSyncDto(tick *models.VSPTick) datasync.VSPTickSyncDto {
	return datasync.VSPTickSyncDto{
		ID:               tick.ID,
		VSPID:            tick.VSPID,
		Time:             tick.Time,
		Immature:         tick.Immature,
		Live:             tick.Live,
		Missed:           tick.Missed,
		PoolFees:         tick.PoolFees,
		ProportionLive:   tick.ProportionLive,
		ProportionMissed: tick.ProportionMissed,
		UserCount:        tick.UserCount,
		UsersActive:      tick.UsersActive,
		Voted:            tick.Voted,
	}
}

func (pg *PgDb) LastVspTickEntryTime() (time time.Time) {
	rows := pg.db.QueryRow(lastVspTickEntryTime)
	_ = rows.Scan(&time)
	return
}

func (pg *PgDb) VspTickCount(ctx context.Context) (int64, error) {
	return models.VSPTicks().Count(ctx, pg.db)
}

func (pg *PgDb) fetchVSPChartData(ctx context.Context, vspName string, start time.Time, endDate uint64, axisString string) (records models.VSPTickSlice, err error) {
	vspInfo, err := models.VSPS(models.VSPWhere.Name.EQ(null.StringFrom(vspName))).One(ctx, pg.db)
	if err != nil {
		return nil, err
	}
	var queries []qm.QueryMod
	if axisString != "" {
		var col string
		switch strings.ToLower(axisString) {
		case string(cache.ImmatureAxis):
			col = models.VSPTickColumns.Immature
			break

		case string(cache.LiveAxis):
			col = models.VSPTickColumns.Live
			break

		case string(cache.VotedAxis):
			col = models.VSPTickColumns.Voted
			break

		case string(cache.MissedAxis):
			col = models.VSPTickColumns.Missed
			break

		case string(cache.PoolFeesAxis):
			col = models.VSPTickColumns.PoolFees
			break

		case string(cache.ProportionLiveAxis):
			col = models.VSPTickColumns.ProportionLive
			break

		case string(cache.ProportionMissedAxis):
			col = models.VSPTickColumns.ProportionMissed
			break

		case string(cache.UserCountAxis):
			col = models.VSPTickColumns.UserCount
			break

		case string(cache.UsersActiveAxis):
			col = models.VSPTickColumns.UsersActive
			break
		}
		queries = append(queries, qm.Select(models.VSPTickColumns.Time, col))
	}

	queries = append(queries, models.VSPTickWhere.VSPID.EQ(vspInfo.ID), models.VSPTickWhere.Time.GT(start))
	if endDate > 0 {
		queries = append(queries, models.VSPTickWhere.Time.LTE(helpers.UnixTime(int64(endDate))))
	}
	data, err := models.VSPTicks(queries...).All(ctx, pg.db)
	return data, err
}

func (pg *PgDb) allVspTickDates(ctx context.Context, start time.Time, vspSources ...string) ([]time.Time, error) {

	var query = []qm.QueryMod{
		qm.Select(models.VSPTickColumns.Time),
		models.VSPTickWhere.Time.GT(start),
		qm.OrderBy(models.VSPTickColumns.Time),
	}
	var wheres []string
	if len(vspSources) > 0 {
		var args = make([]interface{}, len(vspSources))
		for i, s := range vspSources {
			args[i] = s
			wheres = append(wheres, fmt.Sprintf("%s = $%d", models.VSPColumns.Name, i+1))
		}
		vsps, err := models.VSPS(
			qm.Where(strings.Join(wheres, " OR "), args...),
		).All(ctx, pg.db)
		if err != nil {
			return nil, err
		}

		args = make([]interface{}, len(vsps))
		wheres = make([]string, len(vsps))
		for i, v := range vsps {
			args[i] = v.ID
			wheres[i] = fmt.Sprintf("%s = %d", models.VSPTickColumns.VSPID, v.ID)
		}
		query = append(query, qm.Where(strings.Join(wheres, " OR ")))
	}

	vspDates, err := models.VSPTicks(
		query...,
	).All(ctx, pg.db)

	if err != nil {
		return nil, err
	}

	var dates []time.Time
	var unique = map[time.Time]bool{}

	for _, data := range vspDates {
		if _, found := unique[data.Time]; !found {
			dates = append(dates, data.Time)
			unique[data.Time] = true
		}
	}
	return dates, nil
}

func (pg *PgDb) vspIdByName(ctx context.Context, name string) (id int, err error) {
	vspModel, err := models.VSPS(models.VSPWhere.Name.EQ(null.StringFrom(name))).One(ctx, pg.db)
	if err != nil {
		return 0, err
	}
	return vspModel.ID, nil
}

type vspSet struct {
	time             cache.ChartUints
	immature         map[string]cache.ChartNullUints
	live             map[string]cache.ChartNullUints
	voted            map[string]cache.ChartNullUints
	missed           map[string]cache.ChartNullUints
	poolFees         map[string]cache.ChartNullFloats
	proportionLive   map[string]cache.ChartNullFloats
	proportionMissed map[string]cache.ChartNullFloats
	userCount        map[string]cache.ChartNullUints
	usersActive      map[string]cache.ChartNullUints
}

func (pg *PgDb) fetchEncodeVspChart(ctx context.Context, charts *cache.ChartData, axisString string, binString string, vspSources ...string) ([]byte, error) {
	data, _, err := pg.fetchVspChart(ctx, 0, 0, axisString, vspSources...)
	if err != nil {
		return nil, err
	}
	switch strings.ToLower(axisString) {
	case string(cache.ImmatureAxis):
		var deviations []cache.ChartNullData
		for _, p := range vspSources {
			deviations = append(deviations, data.immature[p])
		}
		return cache.MakeVspChart(charts, data.time, deviations, vspSources)

	case string(cache.LiveAxis):
		var deviations []cache.ChartNullData
		for _, p := range vspSources {
			deviations = append(deviations, data.live[p])
		}
		return cache.MakeVspChart(charts, data.time, deviations, vspSources)

	case string(cache.VotedAxis):
		var deviations []cache.ChartNullData
		for _, p := range vspSources {
			deviations = append(deviations, data.voted[p])
		}
		return cache.MakeVspChart(charts, data.time, deviations, vspSources)

	case string(cache.MissedAxis):
		var deviations []cache.ChartNullData
		for _, p := range vspSources {
			deviations = append(deviations, data.missed[p])
		}
		return cache.MakeVspChart(charts, data.time, deviations, vspSources)

	case string(cache.PoolFeesAxis):
		var deviations []cache.ChartNullData
		for _, p := range vspSources {
			deviations = append(deviations, data.poolFees[p])
		}
		return cache.MakeVspChart(charts, data.time, deviations, vspSources)

	case string(cache.ProportionLiveAxis):
		var deviations []cache.ChartNullData
		for _, p := range vspSources {
			deviations = append(deviations, data.proportionLive[p])
		}
		return cache.MakeVspChart(charts, data.time, deviations, vspSources)

	case string(cache.ProportionMissedAxis):
		var deviations []cache.ChartNullData
		for _, p := range vspSources {
			deviations = append(deviations, data.proportionMissed[p])
		}
		return cache.MakeVspChart(charts, data.time, deviations, vspSources)

	case string(cache.UserCountAxis):
		var deviations []cache.ChartNullData
		for _, p := range vspSources {
			deviations = append(deviations, data.userCount[p])
		}
		return cache.MakeVspChart(charts, data.time, deviations, vspSources)

	case string(cache.UsersActiveAxis):
		var deviations []cache.ChartNullData
		for _, p := range vspSources {
			deviations = append(deviations, data.usersActive[p])
		}
		return cache.MakeVspChart(charts, data.time, deviations, vspSources)
	}
	return nil, cache.UnknownChartErr
}

func (pg *PgDb) fetchCacheVspChart(ctx context.Context, charts *cache.ChartData, page int) (interface{}, func(), bool, error) {
	startDate := charts.VSPTimeTip()
	// Get close to the nearest value after the start date to avoid continue loop for situations where there is a gap
	var receiver time.Time
		statement := fmt.Sprintf("SELECT %s FROM %s WHERE %s > '%s' ORDER BY %s LIMIT 1", models.VSPTickColumns.Time,
		models.TableNames.VSPTick, models.VSPTickColumns.Time, 
		helpers.UnixTime(int64(startDate)).Format("2006-01-02 15:04:05+0700"), models.VSPTickColumns.Time)
	rows := pg.db.QueryRow(statement)
	if err := rows.Scan(&receiver); err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			log.Errorf("Error in getting min vsp date - %s", err.Error())
		}
	}
	if int64(startDate) < receiver.Unix() {
		startDate = uint64(receiver.Unix())
		if startDate > 0 {
			startDate -= 1
		}
	}
	//

	aDay := 86400
	endDate := startDate + uint64(7*aDay)
	data, done, err := pg.fetchVspChart(ctx, startDate, endDate, "")
	return data, func() {}, done, err
}

func (pg *PgDb) fetchVspChart(ctx context.Context, startDate uint64, endDate uint64, axisString string, vspSources ...string) (*vspSet, bool, error) {
	var vspDataSet = vspSet{
		time:             []uint64{},
		immature:         make(map[string]cache.ChartNullUints),
		live:             make(map[string]cache.ChartNullUints),
		voted:            make(map[string]cache.ChartNullUints),
		missed:           make(map[string]cache.ChartNullUints),
		poolFees:         make(map[string]cache.ChartNullFloats),
		proportionLive:   make(map[string]cache.ChartNullFloats),
		proportionMissed: make(map[string]cache.ChartNullFloats),
		userCount:        make(map[string]cache.ChartNullUints),
		usersActive:      make(map[string]cache.ChartNullUints),
	}

	var vsps []string = vspSources
	if len(vsps) == 0 {
		allVspData, err := pg.FetchVSPs(ctx)
		if err != nil {
			return nil, false, err
		}
		for _, vspSource := range allVspData {
			vsps = append(vsps, vspSource.Name)
		}
	}

	dates, err := pg.allVspTickDates(ctx, helpers.UnixTime(int64(startDate)), vspSources...)
	if err != nil && err != sql.ErrNoRows {
		return nil, false, err
	}

	for _, date := range dates {
		vspDataSet.time = append(vspDataSet.time, uint64(date.Unix()))
	}

	var done = true
	for _, vspSource := range vsps {
		points, err := pg.fetchVSPChartData(ctx, vspSource, helpers.UnixTime(int64(startDate)), endDate, axisString)
		if err != nil {
			if err.Error() == sql.ErrNoRows.Error() {
				continue
			}
			return nil, false, fmt.Errorf("error in fetching records for %s: %s", vspSource, err.Error())
		}

		if len(points) > 0 {
			done = false
		}

		var pointsMap = map[time.Time]*models.VSPTick{}
		for _, record := range points {
			pointsMap[record.Time] = record
		}

		var hasFoundOne bool
		for _, date := range dates {
			if record, found := pointsMap[date]; found {
				vspDataSet.immature[vspSource] = append(vspDataSet.immature[vspSource], &null.Uint64{Valid: true, Uint64: uint64(record.Immature)})
				vspDataSet.live[vspSource] = append(vspDataSet.live[vspSource], &null.Uint64{Valid: true, Uint64: uint64(record.Live)})
				vspDataSet.voted[vspSource] = append(vspDataSet.voted[vspSource], &null.Uint64{Valid: true, Uint64: uint64(record.Voted)})
				vspDataSet.missed[vspSource] = append(vspDataSet.missed[vspSource], &null.Uint64{Valid: true, Uint64: uint64(record.Missed)})
				vspDataSet.poolFees[vspSource] = append(vspDataSet.poolFees[vspSource], &null.Float64{Valid: true, Float64: record.PoolFees})
				vspDataSet.proportionLive[vspSource] = append(vspDataSet.proportionLive[vspSource], &null.Float64{Valid: true, Float64: record.ProportionLive})
				vspDataSet.proportionMissed[vspSource] = append(vspDataSet.proportionMissed[vspSource], &null.Float64{Valid: true, Float64: record.ProportionMissed})
				vspDataSet.userCount[vspSource] = append(vspDataSet.userCount[vspSource], &null.Uint64{Valid: true, Uint64: uint64(record.UserCount)})
				vspDataSet.usersActive[vspSource] = append(vspDataSet.usersActive[vspSource], &null.Uint64{Valid: true, Uint64: uint64(record.UsersActive)})
				hasFoundOne = true
			} else {
				if hasFoundOne {
					vspDataSet.immature[vspSource] = append(vspDataSet.immature[vspSource], &null.Uint64{Valid: false})
					vspDataSet.live[vspSource] = append(vspDataSet.live[vspSource], &null.Uint64{Valid: false})
					vspDataSet.voted[vspSource] = append(vspDataSet.voted[vspSource], &null.Uint64{Valid: false})
					vspDataSet.missed[vspSource] = append(vspDataSet.missed[vspSource], &null.Uint64{Valid: false})
					vspDataSet.poolFees[vspSource] = append(vspDataSet.poolFees[vspSource], &null.Float64{Valid: false})
					vspDataSet.proportionLive[vspSource] = append(vspDataSet.proportionLive[vspSource], &null.Float64{Valid: false})
					vspDataSet.proportionMissed[vspSource] = append(vspDataSet.proportionMissed[vspSource], &null.Float64{Valid: false})
					vspDataSet.userCount[vspSource] = append(vspDataSet.userCount[vspSource], &null.Uint64{Valid: false})
					vspDataSet.usersActive[vspSource] = append(vspDataSet.usersActive[vspSource], &null.Uint64{Valid: false})
				} else {
					vspDataSet.immature[vspSource] = append(vspDataSet.immature[vspSource], nil)
					vspDataSet.live[vspSource] = append(vspDataSet.live[vspSource], nil)
					vspDataSet.voted[vspSource] = append(vspDataSet.voted[vspSource], nil)
					vspDataSet.missed[vspSource] = append(vspDataSet.missed[vspSource], nil)
					vspDataSet.poolFees[vspSource] = append(vspDataSet.poolFees[vspSource], nil)
					vspDataSet.proportionLive[vspSource] = append(vspDataSet.proportionLive[vspSource], nil)
					vspDataSet.proportionMissed[vspSource] = append(vspDataSet.proportionMissed[vspSource], nil)
					vspDataSet.userCount[vspSource] = append(vspDataSet.userCount[vspSource], nil)
					vspDataSet.usersActive[vspSource] = append(vspDataSet.usersActive[vspSource], nil)
				}
			}
		}
	}

	return &vspDataSet, done, nil
}

func appendVspChart(charts *cache.ChartData, data interface{}) error {
	vspDataSet := data.(*vspSet)

	if len(vspDataSet.time) == 0 {
		return nil
	}

	if err := charts.AppendChartUintsAxis(cache.VSP+"-"+string(cache.TimeAxis),
		vspDataSet.time); err != nil {
		return err
	}

	keyExists := func (arr []string, key string) bool {
		for _, s := range arr {
			if s == key {
				return true
			}
		}
		return false
	}

	for vspSource, record := range vspDataSet.immature {
		if !keyExists(charts.VSPSources, vspSource) {
			charts.VSPSources = append(charts.VSPSources, vspSource)
		}
		if err := charts.AppendChartNullUintsAxis(cache.VSP+"-"+string(cache.ImmatureAxis)+"-"+vspSource,
			record); err != nil {
			return err
		}
	}

	for vspSource, record := range vspDataSet.live {
		if err := charts.AppendChartNullUintsAxis(cache.VSP+"-"+string(cache.LiveAxis)+"-"+vspSource,
			record); err != nil {
			return err
		}
	}

	for vspSource, record := range vspDataSet.voted {
		if err := charts.AppendChartNullUintsAxis(cache.VSP+"-"+string(cache.VotedAxis)+"-"+vspSource,
			record); err != nil {
			return err
		}
	}

	for vspSource, record := range vspDataSet.missed {
		if err := charts.AppendChartNullUintsAxis(cache.VSP+"-"+string(cache.MissedAxis)+"-"+vspSource,
			record); err != nil {
			return err
		}
	}

	for vspSource, record := range vspDataSet.poolFees {
		if err := charts.AppendChartNullFloatsAxis(cache.VSP+"-"+string(cache.PoolFeesAxis)+"-"+vspSource,
			record); err != nil {
			return err
		}
	}

	for vspSource, record := range vspDataSet.proportionLive {
		if err := charts.AppendChartNullFloatsAxis(cache.VSP+"-"+string(cache.ProportionLiveAxis)+"-"+vspSource,
			record); err != nil {
			return err
		}
	}

	for vspSource, record := range vspDataSet.proportionMissed {
		if err := charts.AppendChartNullFloatsAxis(cache.VSP+"-"+string(cache.ProportionMissedAxis)+"-"+vspSource,
			record); err != nil {
			return err
		}
	}

	for vspSource, record := range vspDataSet.usersActive {
		if err := charts.AppendChartNullUintsAxis(cache.VSP+"-"+string(cache.UsersActiveAxis)+"-"+vspSource,
			record); err != nil {
			return err
		}
	}

	for vspSource, record := range vspDataSet.userCount {
		if err := charts.AppendChartNullUintsAxis(cache.VSP+"-"+string(cache.UserCountAxis)+"-"+vspSource,
			record); err != nil {
			return err
		}
	}

	return nil
}
