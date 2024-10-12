package handler

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"pcr/internal/phone"
	"pcr/internal/sqlc"
	"text/template"
)

type Cursor interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	sqlc.DBTX
}

type Pcr struct {
	Cursor Cursor
	Root   string
}

func (p *Pcr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Join(p.Root, r.URL.Path)
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		slog.Error("Failed to load template", "reason", err, "filename", filename)
		if errors.Is(err, os.ErrNotExist) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	ua := r.Header.Get("User-Agent")
	identity := phone.Identify(ua)

	if (phone.Identity{}) == identity {
		err = tmpl.Execute(w, nil)
	} else {
		err = tmpl.Execute(w, p.getPhoneValues(r.Context(), identity))
	}

	if err != nil {
		slog.Error("Failed to execute template", "reason", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (p *Pcr) getPhoneValues(ctx context.Context, id phone.Identity) phone.Values {
	queries := sqlc.New(p.Cursor)
	res, err := queries.GetByModelAndMac(ctx, sqlc.GetByModelAndMacParams{
		Model: pgtype.Text{
			String: id.Model,
			Valid:  true,
		},
		Mac: pgtype.Text{
			String: id.MAC,
			Valid:  true,
		},
	})

	if err != nil {
		slog.Error("Failed to retrieve phone values", "reason", err, "model", id.Model, "mac", id.MAC)
		return phone.Values{}
	}

	return phone.Values{
		DisplayName: res.DisplayName.String,
		Username:    res.Username.String,
		Password:    res.Password.String,
	}
}
