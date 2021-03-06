// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Class is an object representing the database table.
type Class struct {
	ID          string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string `boil:"name" json:"name" toml:"name" yaml:"name"`
	SchoolOrder int    `boil:"school_order" json:"school_order" toml:"school_order" yaml:"school_order"`

	R *classR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L classL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClassColumns = struct {
	ID          string
	Name        string
	SchoolOrder string
}{
	ID:          "id",
	Name:        "name",
	SchoolOrder: "school_order",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

var ClassWhere = struct {
	ID          whereHelperstring
	Name        whereHelperstring
	SchoolOrder whereHelperint
}{
	ID:          whereHelperstring{field: "\"classes\".\"id\""},
	Name:        whereHelperstring{field: "\"classes\".\"name\""},
	SchoolOrder: whereHelperint{field: "\"classes\".\"school_order\""},
}

// ClassRels is where relationship names are stored.
var ClassRels = struct {
	Deposits      string
	Students      string
	Subscriptions string
}{
	Deposits:      "Deposits",
	Students:      "Students",
	Subscriptions: "Subscriptions",
}

// classR is where relationships are stored.
type classR struct {
	Deposits      DepositSlice
	Students      StudentSlice
	Subscriptions SubscriptionSlice
}

// NewStruct creates a new relationship struct
func (*classR) NewStruct() *classR {
	return &classR{}
}

// classL is where Load methods for each relationship are stored.
type classL struct{}

var (
	classAllColumns            = []string{"id", "name", "school_order"}
	classColumnsWithoutDefault = []string{"id", "name"}
	classColumnsWithDefault    = []string{"school_order"}
	classPrimaryKeyColumns     = []string{"id"}
)

type (
	// ClassSlice is an alias for a slice of pointers to Class.
	// This should generally be used opposed to []Class.
	ClassSlice []*Class

	classQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	classType                 = reflect.TypeOf(&Class{})
	classMapping              = queries.MakeStructMapping(classType)
	classPrimaryKeyMapping, _ = queries.BindMapping(classType, classMapping, classPrimaryKeyColumns)
	classInsertCacheMut       sync.RWMutex
	classInsertCache          = make(map[string]insertCache)
	classUpdateCacheMut       sync.RWMutex
	classUpdateCache          = make(map[string]updateCache)
	classUpsertCacheMut       sync.RWMutex
	classUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single class record from the query.
func (q classQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Class, error) {
	o := &Class{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for classes")
	}

	return o, nil
}

// All returns all Class records from the query.
func (q classQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClassSlice, error) {
	var o []*Class

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Class slice")
	}

	return o, nil
}

// Count returns the count of all Class records in the query.
func (q classQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count classes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q classQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if classes exists")
	}

	return count > 0, nil
}

// Deposits retrieves all the deposit's Deposits with an executor.
func (o *Class) Deposits(mods ...qm.QueryMod) depositQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"deposits\".\"class_id\"=?", o.ID),
	)

	query := Deposits(queryMods...)
	queries.SetFrom(query.Query, "\"deposits\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"deposits\".*"})
	}

	return query
}

// Students retrieves all the student's Students with an executor.
func (o *Class) Students(mods ...qm.QueryMod) studentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"student\".\"class_id\"=?", o.ID),
	)

	query := Students(queryMods...)
	queries.SetFrom(query.Query, "\"student\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"student\".*"})
	}

	return query
}

// Subscriptions retrieves all the subscription's Subscriptions with an executor.
func (o *Class) Subscriptions(mods ...qm.QueryMod) subscriptionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"subscription\".\"class_id\"=?", o.ID),
	)

	query := Subscriptions(queryMods...)
	queries.SetFrom(query.Query, "\"subscription\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"subscription\".*"})
	}

	return query
}

// LoadDeposits allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (classL) LoadDeposits(ctx context.Context, e boil.ContextExecutor, singular bool, maybeClass interface{}, mods queries.Applicator) error {
	var slice []*Class
	var object *Class

	if singular {
		object = maybeClass.(*Class)
	} else {
		slice = *maybeClass.(*[]*Class)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &classR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &classR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`deposits`), qm.WhereIn(`deposits.class_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load deposits")
	}

	var resultSlice []*Deposit
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice deposits")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on deposits")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for deposits")
	}

	if singular {
		object.R.Deposits = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &depositR{}
			}
			foreign.R.Class = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ClassID {
				local.R.Deposits = append(local.R.Deposits, foreign)
				if foreign.R == nil {
					foreign.R = &depositR{}
				}
				foreign.R.Class = local
				break
			}
		}
	}

	return nil
}

// LoadStudents allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (classL) LoadStudents(ctx context.Context, e boil.ContextExecutor, singular bool, maybeClass interface{}, mods queries.Applicator) error {
	var slice []*Class
	var object *Class

	if singular {
		object = maybeClass.(*Class)
	} else {
		slice = *maybeClass.(*[]*Class)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &classR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &classR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`student`), qm.WhereIn(`student.class_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load student")
	}

	var resultSlice []*Student
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice student")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on student")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for student")
	}

	if singular {
		object.R.Students = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &studentR{}
			}
			foreign.R.Class = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.ClassID) {
				local.R.Students = append(local.R.Students, foreign)
				if foreign.R == nil {
					foreign.R = &studentR{}
				}
				foreign.R.Class = local
				break
			}
		}
	}

	return nil
}

// LoadSubscriptions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (classL) LoadSubscriptions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeClass interface{}, mods queries.Applicator) error {
	var slice []*Class
	var object *Class

	if singular {
		object = maybeClass.(*Class)
	} else {
		slice = *maybeClass.(*[]*Class)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &classR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &classR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`subscription`), qm.WhereIn(`subscription.class_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load subscription")
	}

	var resultSlice []*Subscription
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice subscription")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on subscription")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for subscription")
	}

	if singular {
		object.R.Subscriptions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &subscriptionR{}
			}
			foreign.R.Class = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ClassID {
				local.R.Subscriptions = append(local.R.Subscriptions, foreign)
				if foreign.R == nil {
					foreign.R = &subscriptionR{}
				}
				foreign.R.Class = local
				break
			}
		}
	}

	return nil
}

// AddDeposits adds the given related objects to the existing relationships
// of the class, optionally inserting them as new records.
// Appends related to o.R.Deposits.
// Sets related.R.Class appropriately.
func (o *Class) AddDeposits(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Deposit) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ClassID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"deposits\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"class_id"}),
				strmangle.WhereClause("\"", "\"", 2, depositPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ClassID = o.ID
		}
	}

	if o.R == nil {
		o.R = &classR{
			Deposits: related,
		}
	} else {
		o.R.Deposits = append(o.R.Deposits, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &depositR{
				Class: o,
			}
		} else {
			rel.R.Class = o
		}
	}
	return nil
}

// AddStudents adds the given related objects to the existing relationships
// of the class, optionally inserting them as new records.
// Appends related to o.R.Students.
// Sets related.R.Class appropriately.
func (o *Class) AddStudents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Student) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ClassID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"student\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"class_id"}),
				strmangle.WhereClause("\"", "\"", 2, studentPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.ClassID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &classR{
			Students: related,
		}
	} else {
		o.R.Students = append(o.R.Students, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &studentR{
				Class: o,
			}
		} else {
			rel.R.Class = o
		}
	}
	return nil
}

// SetStudents removes all previously related items of the
// class replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Class's Students accordingly.
// Replaces o.R.Students with related.
// Sets related.R.Class's Students accordingly.
func (o *Class) SetStudents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Student) error {
	query := "update \"student\" set \"class_id\" = null where \"class_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Students {
			queries.SetScanner(&rel.ClassID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Class = nil
		}

		o.R.Students = nil
	}
	return o.AddStudents(ctx, exec, insert, related...)
}

// RemoveStudents relationships from objects passed in.
// Removes related items from R.Students (uses pointer comparison, removal does not keep order)
// Sets related.R.Class.
func (o *Class) RemoveStudents(ctx context.Context, exec boil.ContextExecutor, related ...*Student) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ClassID, nil)
		if rel.R != nil {
			rel.R.Class = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("class_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Students {
			if rel != ri {
				continue
			}

			ln := len(o.R.Students)
			if ln > 1 && i < ln-1 {
				o.R.Students[i] = o.R.Students[ln-1]
			}
			o.R.Students = o.R.Students[:ln-1]
			break
		}
	}

	return nil
}

// AddSubscriptions adds the given related objects to the existing relationships
// of the class, optionally inserting them as new records.
// Appends related to o.R.Subscriptions.
// Sets related.R.Class appropriately.
func (o *Class) AddSubscriptions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Subscription) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ClassID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"subscription\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"class_id"}),
				strmangle.WhereClause("\"", "\"", 2, subscriptionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ClassID = o.ID
		}
	}

	if o.R == nil {
		o.R = &classR{
			Subscriptions: related,
		}
	} else {
		o.R.Subscriptions = append(o.R.Subscriptions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &subscriptionR{
				Class: o,
			}
		} else {
			rel.R.Class = o
		}
	}
	return nil
}

// Classes retrieves all the records using an executor.
func Classes(mods ...qm.QueryMod) classQuery {
	mods = append(mods, qm.From("\"classes\""))
	return classQuery{NewQuery(mods...)}
}

// FindClass retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClass(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Class, error) {
	classObj := &Class{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"classes\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, classObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from classes")
	}

	return classObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Class) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no classes provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(classColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	classInsertCacheMut.RLock()
	cache, cached := classInsertCache[key]
	classInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			classAllColumns,
			classColumnsWithDefault,
			classColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(classType, classMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(classType, classMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"classes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"classes\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into classes")
	}

	if !cached {
		classInsertCacheMut.Lock()
		classInsertCache[key] = cache
		classInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Class.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Class) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	classUpdateCacheMut.RLock()
	cache, cached := classUpdateCache[key]
	classUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			classAllColumns,
			classPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update classes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"classes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, classPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(classType, classMapping, append(wl, classPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update classes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for classes")
	}

	if !cached {
		classUpdateCacheMut.Lock()
		classUpdateCache[key] = cache
		classUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q classQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for classes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for classes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClassSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), classPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"classes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, classPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in class slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all class")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Class) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no classes provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(classColumnsWithDefault, o)

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

	classUpsertCacheMut.RLock()
	cache, cached := classUpsertCache[key]
	classUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			classAllColumns,
			classColumnsWithDefault,
			classColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			classAllColumns,
			classPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert classes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(classPrimaryKeyColumns))
			copy(conflict, classPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"classes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(classType, classMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(classType, classMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert classes")
	}

	if !cached {
		classUpsertCacheMut.Lock()
		classUpsertCache[key] = cache
		classUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Class record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Class) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Class provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), classPrimaryKeyMapping)
	sql := "DELETE FROM \"classes\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from classes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for classes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q classQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no classQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from classes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for classes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClassSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), classPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"classes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, classPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from class slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for classes")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Class) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClass(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClassSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClassSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), classPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"classes\".* FROM \"classes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, classPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ClassSlice")
	}

	*o = slice

	return nil
}

// ClassExists checks if the Class row exists.
func ClassExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"classes\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if classes exists")
	}

	return exists, nil
}
