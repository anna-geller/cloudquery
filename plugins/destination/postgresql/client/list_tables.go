package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

// this returns the following table in sorted manner:
// +----------------+-------------+-------------+------------+---------------+-----------+---------------------+
// | ordinal_position | table_name | column_name | data_type | is_primary_key| not_null  | pk_constraint_name  |
// +----------------+-------------+-------------+------------+---------------+-----------+---------------------+
// |              1 | users       | id          | bigint     | YES           | true 		 | cq_users_pk 	  	   |
// |              2 | users       | name        | text       | NO            | false 	   | 					           |
// |              3 | users       | email       | text       | NO            | true 		 | cq_users_pk         |
// |              1 | posts       | id          | bigint     | YES           | true 		 | cq_posts_pk			   |
// |              2 | posts       | title       | text       | NO            | false 	   | 					           |
const selectTables = `
SELECT
	columns.ordinal_position AS ordinal_position,
	pg_class.relname AS table_name,
	pg_attribute.attname AS column_name,
	CASE
	    -- This is required per the differences in pg_catalog.format_type implementations
	    -- between PostgreSQL & CockroachDB.
	    -- namely, numeric(20,0)[] is returned as numeric[] unless we use the typelem format + []
	    WHEN pg_type.typcategory = 'A' AND pg_type.typelem != 0
		THEN pg_catalog.format_type(pg_type.typelem, pg_attribute.atttypmod) || '[]'
		ELSE pg_catalog.format_type(pg_attribute.atttypid, pg_attribute.atttypmod)
	END AS data_type,
	CASE 
		WHEN primary_keys.conkey IS NOT NULL AND array_position(primary_keys.conkey, pg_attribute.attnum) > 0 THEN true
		ELSE false
	END AS is_primary_key,
	CASE 
		WHEN unique_constraints.conkey IS NOT NULL AND array_position(unique_constraints.conkey, pg_attribute.attnum) > 0 THEN true
		ELSE false
	END AS is_unique,	
	CASE 
		WHEN pg_attribute.attnotnull THEN true
		ELSE false
	END AS not_null,
	COALESCE(primary_keys.conname, '') AS primary_key_constraint_name
FROM
	pg_catalog.pg_attribute
	INNER JOIN
	pg_catalog.pg_type ON pg_type.oid = pg_attribute.atttypid
	INNER JOIN
	pg_catalog.pg_class ON pg_class.oid = pg_attribute.attrelid
	INNER JOIN
	pg_catalog.pg_namespace ON pg_namespace.oid = pg_class.relnamespace
	LEFT JOIN
	pg_catalog.pg_constraint primary_keys ON primary_keys.conrelid = pg_attribute.attrelid
	AND conkey IS NOT NULL AND array_position(conkey, pg_attribute.attnum) > 0
	AND contype = 'p'
	LEFT JOIN
	pg_catalog.pg_constraint unique_constraints ON unique_constraints.conrelid = pg_attribute.attrelid
	AND unique_constraints.conkey IS NOT NULL AND array_position(unique_constraints.conkey, pg_attribute.attnum) > 0
	AND unique_constraints.contype = 'u'
	INNER JOIN
	information_schema.columns ON columns.table_name = pg_class.relname AND columns.column_name = pg_attribute.attname AND columns.table_schema = pg_catalog.pg_namespace.nspname
WHERE
	pg_attribute.attnum > 0
	AND NOT pg_attribute.attisdropped
	AND pg_catalog.pg_namespace.nspname = '%s'
	%s 
ORDER BY
	table_name ASC, ordinal_position ASC;
`

func (c *Client) listTables(ctx context.Context) (schema.Tables, error) {
	c.pgTablesToPKConstraints = map[string]string{}
	var tables schema.Tables
	var whereClause string
	if c.pgType == pgTypeCockroachDB {
		whereClause = " AND information_schema.columns.is_hidden != 'YES'"
	}
	q := fmt.Sprintf(selectTables, c.currentSchemaName, whereClause)
	rows, err := c.conn.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ordinalPosition int
		var tableName, columnName, columnType, pkName string
		var isPrimaryKey, isUnique, notNull bool
		if err := rows.Scan(&ordinalPosition, &tableName, &columnName, &columnType, &isPrimaryKey, &isUnique, &notNull, &pkName); err != nil {
			return nil, err
		}
		if ordinalPosition == 1 {
			tables = append(tables, &schema.Table{
				Name:    tableName,
				Columns: make([]schema.Column, 0),
			})
		}
		table := tables[len(tables)-1]
		// Note: constraints always have a name in PostgreSQL.
		// However, we want not only to store the info about the constraint name,
		// but also the fact that we saw such table.
		switch pkName {
		case "":
			// We still store the fact that we saw the table
			if _, ok := c.pgTablesToPKConstraints[tableName]; !ok {
				// Just store the empty string.
				// This will indicate 2 things:
				// 1. We saw the table with this name
				// 2. If this is still empty on insert, the table in the database doesn't have a PK constraint
				c.pgTablesToPKConstraints[tableName] = ""
			}
		default:
			c.pgTablesToPKConstraints[tableName], table.PkConstraintName = pkName, pkName
		}
		table.Columns = append(table.Columns, schema.Column{
			Name:       columnName,
			PrimaryKey: isPrimaryKey,
			NotNull:    notNull,
			Unique:     isUnique,
			Type:       c.PgToSchemaType(columnType),
		})
	}
	return tables, nil
}
