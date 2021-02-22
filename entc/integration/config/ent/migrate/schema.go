// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "Users" table.
	UsersColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true, Size: 128},
		{Name: "label", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "Users" table.
	UsersTable = &schema.Table{
		Name:        "Users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
	UsersTable.Annotation = &entsql.Annotation{
		Table:   "Users",
		Charset: "utf8mb4",
	}
	UsersTable.Annotation.Incremental = new(bool)
	*UsersTable.Annotation.Incremental = false
}
