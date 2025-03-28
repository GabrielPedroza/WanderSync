// Code generated by ent, DO NOT EDIT.

package ent

import (
	"GabrielPedroza/WanderSync/ent/location"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Location is the model entity for the Location schema.
type Location struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LocationQuery when eager-loading is set.
	Edges        LocationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LocationEdges holds the relations/edges for other nodes in the graph.
type LocationEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedUsers map[string][]*User
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Location) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case location.FieldID:
			values[i] = new(sql.NullInt64)
		case location.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Location fields.
func (l *Location) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case location.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case location.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Location.
// This includes values selected through modifiers, order, etc.
func (l *Location) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Location entity.
func (l *Location) QueryUsers() *UserQuery {
	return NewLocationClient(l.config).QueryUsers(l)
}

// Update returns a builder for updating this Location.
// Note that you need to call Location.Unwrap() before calling this method if this Location
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Location) Update() *LocationUpdateOne {
	return NewLocationClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Location entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Location) Unwrap() *Location {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Location is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Location) String() string {
	var builder strings.Builder
	builder.WriteString("Location(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedUsers returns the Users named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *Location) NamedUsers(name string) ([]*User, error) {
	if l.Edges.namedUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *Location) appendNamedUsers(name string, edges ...*User) {
	if l.Edges.namedUsers == nil {
		l.Edges.namedUsers = make(map[string][]*User)
	}
	if len(edges) == 0 {
		l.Edges.namedUsers[name] = []*User{}
	} else {
		l.Edges.namedUsers[name] = append(l.Edges.namedUsers[name], edges...)
	}
}

// Locations is a parsable slice of Location.
type Locations []*Location
