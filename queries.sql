-- name: GetByModelAndMac :one
SELECT
    display_name, username, password
FROM
    pcr_phone
WHERE
    model = $1 AND
    mac = $2;
