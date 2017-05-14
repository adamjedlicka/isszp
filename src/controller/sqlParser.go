package controller

import (
	"bytes"
	"fmt"
)

type Queryable interface {
	Query() Query
	ExecQuery(Query)
}

type Query struct {
	table string
	where []string
	data  []interface{}

	groupBy []string
	orderBy []string

	like    []string
	notLike []string
}

func NewQuery(table string) *Query {
	return &Query{
		table: table,
	}
}

func (q *Query) Where(val string, i interface{}) *Query {

	q.where = append(q.where, val)

	q.data = append(q.data, i)

	return q
}

func (q *Query) GroupBy(column ...string) *Query {

	q.groupBy = append(q.groupBy, column...)

	return q
}

func (q *Query) OrderBy(column ...string) *Query {

	q.orderBy = append(q.orderBy, column...)

	return q
}

func (q *Query) Like(val string, i interface{}) *Query {
	q.like = append(q.like, val)

	q.data = append(q.data, i)

	return q
}

func (q *Query) NotLike(val string, i interface{}) *Query {
	q.notLike = append(q.notLike, val)

	q.data = append(q.data, i)

	return q
}

func (q *Query) ToSQL() string {
	var tmp bytes.Buffer

	tmp.WriteString("SELECT * FROM " + q.table)

	if len(q.where) > 0 {

		if len(q.where) > 0 {
			tmp.WriteString("\n\tWHERE ")

			for i := 0; i < len(q.where)-1; i++ {
				tmp.WriteString(q.where[i] + " AND ")

			}

			tmp.WriteString(q.where[len(q.where)-1])
		}

		if len(q.like) > 0 {
			tmp.WriteString("\n\tWHERE ")

			for i := 0; i < len(q.like)-1; i++ {
				tmp.WriteString(q.like[i] + " LIKE ?" + " AND ")
			}

			tmp.WriteString(q.like[len(q.like)-1] + " LIKE ?")
		}

		if len(q.notLike) > 0 {
			tmp.WriteString("\n\tWHERE ")

			for i := 0; i < len(q.notLike)-1; i++ {
				tmp.WriteString(q.notLike[i] + " NOT LIKE ?" + " AND ")
			}

			tmp.WriteString(q.notLike[len(q.notLike)-1] + " NOT LIKE ?")
		}
	}

	if len(q.groupBy) > 0 {
		tmp.WriteString("\n\tGROUP BY ")
		for i := 0; i < len(q.groupBy)-1; i++ {
			tmp.WriteString(q.groupBy[i] + ", ")

		}
		tmp.WriteString(q.groupBy[len(q.groupBy)-1])
	}

	if len(q.orderBy) > 0 {
		tmp.WriteString("\n\tORDER BY ")
		for i := 0; i < len(q.orderBy)-1; i++ {
			tmp.WriteString(q.orderBy[i] + ", ")

		}
		tmp.WriteString(q.orderBy[len(q.orderBy)-1])
	}

	tmp.WriteString(";")
	return tmp.String()
}

func (q *Query) ToData() []interface{} {

	return q.data
}

func (q Query) String() string {
	return fmt.Sprint(q.ToSQL()+"\nVALUES:", q.ToData())
}
