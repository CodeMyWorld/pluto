// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
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

// DeviceApp is an object representing the database table.
type DeviceApp struct {
	ID        uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	DeviceID  string    `boil:"device_id" json:"device_id" toml:"device_id" yaml:"device_id"`
	AppID     string    `boil:"app_id" json:"app_id" toml:"app_id" yaml:"app_id"`

	R *deviceAppR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L deviceAppL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DeviceAppColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	DeviceID  string
	AppID     string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	DeviceID:  "device_id",
	AppID:     "app_id",
}

// Generated where

var DeviceAppWhere = struct {
	ID        whereHelperuint
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
	DeletedAt whereHelpernull_Time
	DeviceID  whereHelperstring
	AppID     whereHelperstring
}{
	ID:        whereHelperuint{field: "`device_apps`.`id`"},
	CreatedAt: whereHelpernull_Time{field: "`device_apps`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`device_apps`.`updated_at`"},
	DeletedAt: whereHelpernull_Time{field: "`device_apps`.`deleted_at`"},
	DeviceID:  whereHelperstring{field: "`device_apps`.`device_id`"},
	AppID:     whereHelperstring{field: "`device_apps`.`app_id`"},
}

// DeviceAppRels is where relationship names are stored.
var DeviceAppRels = struct {
}{}

// deviceAppR is where relationships are stored.
type deviceAppR struct {
}

// NewStruct creates a new relationship struct
func (*deviceAppR) NewStruct() *deviceAppR {
	return &deviceAppR{}
}

// deviceAppL is where Load methods for each relationship are stored.
type deviceAppL struct{}

var (
	deviceAppAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "device_id", "app_id"}
	deviceAppColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "device_id", "app_id"}
	deviceAppColumnsWithDefault    = []string{"id"}
	deviceAppPrimaryKeyColumns     = []string{"id"}
)

type (
	// DeviceAppSlice is an alias for a slice of pointers to DeviceApp.
	// This should generally be used opposed to []DeviceApp.
	DeviceAppSlice []*DeviceApp
	// DeviceAppHook is the signature for custom DeviceApp hook methods
	DeviceAppHook func(boil.Executor, *DeviceApp) error

	deviceAppQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	deviceAppType                 = reflect.TypeOf(&DeviceApp{})
	deviceAppMapping              = queries.MakeStructMapping(deviceAppType)
	deviceAppPrimaryKeyMapping, _ = queries.BindMapping(deviceAppType, deviceAppMapping, deviceAppPrimaryKeyColumns)
	deviceAppInsertCacheMut       sync.RWMutex
	deviceAppInsertCache          = make(map[string]insertCache)
	deviceAppUpdateCacheMut       sync.RWMutex
	deviceAppUpdateCache          = make(map[string]updateCache)
	deviceAppUpsertCacheMut       sync.RWMutex
	deviceAppUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var deviceAppBeforeInsertHooks []DeviceAppHook
var deviceAppBeforeUpdateHooks []DeviceAppHook
var deviceAppBeforeDeleteHooks []DeviceAppHook
var deviceAppBeforeUpsertHooks []DeviceAppHook

var deviceAppAfterInsertHooks []DeviceAppHook
var deviceAppAfterSelectHooks []DeviceAppHook
var deviceAppAfterUpdateHooks []DeviceAppHook
var deviceAppAfterDeleteHooks []DeviceAppHook
var deviceAppAfterUpsertHooks []DeviceAppHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DeviceApp) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range deviceAppBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DeviceApp) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range deviceAppBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DeviceApp) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range deviceAppBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DeviceApp) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range deviceAppBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DeviceApp) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range deviceAppAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DeviceApp) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range deviceAppAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DeviceApp) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range deviceAppAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DeviceApp) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range deviceAppAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DeviceApp) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range deviceAppAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDeviceAppHook registers your hook function for all future operations.
func AddDeviceAppHook(hookPoint boil.HookPoint, deviceAppHook DeviceAppHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		deviceAppBeforeInsertHooks = append(deviceAppBeforeInsertHooks, deviceAppHook)
	case boil.BeforeUpdateHook:
		deviceAppBeforeUpdateHooks = append(deviceAppBeforeUpdateHooks, deviceAppHook)
	case boil.BeforeDeleteHook:
		deviceAppBeforeDeleteHooks = append(deviceAppBeforeDeleteHooks, deviceAppHook)
	case boil.BeforeUpsertHook:
		deviceAppBeforeUpsertHooks = append(deviceAppBeforeUpsertHooks, deviceAppHook)
	case boil.AfterInsertHook:
		deviceAppAfterInsertHooks = append(deviceAppAfterInsertHooks, deviceAppHook)
	case boil.AfterSelectHook:
		deviceAppAfterSelectHooks = append(deviceAppAfterSelectHooks, deviceAppHook)
	case boil.AfterUpdateHook:
		deviceAppAfterUpdateHooks = append(deviceAppAfterUpdateHooks, deviceAppHook)
	case boil.AfterDeleteHook:
		deviceAppAfterDeleteHooks = append(deviceAppAfterDeleteHooks, deviceAppHook)
	case boil.AfterUpsertHook:
		deviceAppAfterUpsertHooks = append(deviceAppAfterUpsertHooks, deviceAppHook)
	}
}

// One returns a single deviceApp record from the query.
func (q deviceAppQuery) One(exec boil.Executor) (*DeviceApp, error) {
	o := &DeviceApp{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for device_apps")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DeviceApp records from the query.
func (q deviceAppQuery) All(exec boil.Executor) (DeviceAppSlice, error) {
	var o []*DeviceApp

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DeviceApp slice")
	}

	if len(deviceAppAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DeviceApp records in the query.
func (q deviceAppQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count device_apps rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q deviceAppQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if device_apps exists")
	}

	return count > 0, nil
}

// DeviceApps retrieves all the records using an executor.
func DeviceApps(mods ...qm.QueryMod) deviceAppQuery {
	mods = append(mods, qm.From("`device_apps`"))
	return deviceAppQuery{NewQuery(mods...)}
}

// FindDeviceApp retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDeviceApp(exec boil.Executor, iD uint, selectCols ...string) (*DeviceApp, error) {
	deviceAppObj := &DeviceApp{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `device_apps` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, deviceAppObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from device_apps")
	}

	return deviceAppObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DeviceApp) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no device_apps provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(deviceAppColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	deviceAppInsertCacheMut.RLock()
	cache, cached := deviceAppInsertCache[key]
	deviceAppInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			deviceAppAllColumns,
			deviceAppColumnsWithDefault,
			deviceAppColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(deviceAppType, deviceAppMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(deviceAppType, deviceAppMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `device_apps` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `device_apps` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `device_apps` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, deviceAppPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into device_apps")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == deviceAppMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for device_apps")
	}

CacheNoHooks:
	if !cached {
		deviceAppInsertCacheMut.Lock()
		deviceAppInsertCache[key] = cache
		deviceAppInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the DeviceApp.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DeviceApp) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	deviceAppUpdateCacheMut.RLock()
	cache, cached := deviceAppUpdateCache[key]
	deviceAppUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			deviceAppAllColumns,
			deviceAppPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update device_apps, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `device_apps` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, deviceAppPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(deviceAppType, deviceAppMapping, append(wl, deviceAppPrimaryKeyColumns...))
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
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update device_apps row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for device_apps")
	}

	if !cached {
		deviceAppUpdateCacheMut.Lock()
		deviceAppUpdateCache[key] = cache
		deviceAppUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q deviceAppQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for device_apps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for device_apps")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DeviceAppSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deviceAppPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `device_apps` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deviceAppPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in deviceApp slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all deviceApp")
	}
	return rowsAff, nil
}

var mySQLDeviceAppUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DeviceApp) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no device_apps provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(deviceAppColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDeviceAppUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	deviceAppUpsertCacheMut.RLock()
	cache, cached := deviceAppUpsertCache[key]
	deviceAppUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			deviceAppAllColumns,
			deviceAppColumnsWithDefault,
			deviceAppColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			deviceAppAllColumns,
			deviceAppPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert device_apps, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "device_apps", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `device_apps` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(deviceAppType, deviceAppMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(deviceAppType, deviceAppMapping, ret)
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
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for device_apps")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == deviceAppMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(deviceAppType, deviceAppMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for device_apps")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for device_apps")
	}

CacheNoHooks:
	if !cached {
		deviceAppUpsertCacheMut.Lock()
		deviceAppUpsertCache[key] = cache
		deviceAppUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single DeviceApp record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DeviceApp) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DeviceApp provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), deviceAppPrimaryKeyMapping)
	sql := "DELETE FROM `device_apps` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from device_apps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for device_apps")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q deviceAppQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no deviceAppQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from device_apps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for device_apps")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DeviceAppSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(deviceAppBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deviceAppPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `device_apps` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deviceAppPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from deviceApp slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for device_apps")
	}

	if len(deviceAppAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DeviceApp) Reload(exec boil.Executor) error {
	ret, err := FindDeviceApp(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DeviceAppSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DeviceAppSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deviceAppPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `device_apps`.* FROM `device_apps` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deviceAppPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DeviceAppSlice")
	}

	*o = slice

	return nil
}

// DeviceAppExists checks if the DeviceApp row exists.
func DeviceAppExists(exec boil.Executor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `device_apps` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if device_apps exists")
	}

	return exists, nil
}