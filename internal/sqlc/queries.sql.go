// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getByModelAndMac = `-- name: GetByModelAndMac :one
SELECT
    display_name, username, password
FROM
    pcr_phone
WHERE
    model = $1 AND
    mac = $2
`

type GetByModelAndMacParams struct {
	Model pgtype.Text
	Mac   pgtype.Text
}

type GetByModelAndMacRow struct {
	DisplayName pgtype.Text
	Username    pgtype.Text
	Password    pgtype.Text
}

func (q *Queries) GetByModelAndMac(ctx context.Context, arg GetByModelAndMacParams) (GetByModelAndMacRow, error) {
	row := q.db.QueryRow(ctx, getByModelAndMac, arg.Model, arg.Mac)
	var i GetByModelAndMacRow
	err := row.Scan(&i.DisplayName, &i.Username, &i.Password)
	return i, err
}
