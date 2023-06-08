package main

import (
	"fmt"

	"github.com/gofrs/uuid"
	"go.riyazali.net/sqlite"
)

type v7String struct{}

func (*v7String) Args() int           { return 0 }
func (*v7String) Deterministic() bool { return false }
func (*v7String) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	u, err := uuid.NewV7()
	if err != nil {
		ctx.ResultError(fmt.Errorf("error generating uuidv7: %w", err))
		return
	}

	ctx.ResultText(u.String())
}

type v7Byte struct{}

func (*v7Byte) Args() int           { return 0 }
func (*v7Byte) Deterministic() bool { return false }
func (*v7Byte) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	u, err := uuid.NewV7()
	if err != nil {
		ctx.ResultError(fmt.Errorf("error generating uuidv7: %w", err))
		return
	}

	ctx.ResultBlob(u.Bytes())
}

func init() {
	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("uuid_generate_v7", &v7String{}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("uuid_generate_v7_bytes", &v7Byte{}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		return sqlite.SQLITE_OK, nil
	})
}

func main() {}
