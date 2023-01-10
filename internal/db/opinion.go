package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/untaldouglas/go-rest-api/internal/opinion"
)

type OpinionRow struct {
	ID        string
	Asunto    sql.NullString
	Contenido sql.NullString
	Autor     sql.NullString
}

func convertOpinionRowToOpinion(o OpinionRow) opinion.Opinion {
	return opinion.Opinion{
		ID:        o.ID,
		Asunto:    o.Asunto.String,
		Contenido: o.Contenido.String,
		Autor:     o.Autor.String,
	}
}

func (d *Database) GetOpinion(
	ctx context.Context, uuid string,
) (opinion.Opinion, error) {
	var opiRow OpinionRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, asunto, contenido,autor
	FROM opinions
	WHERE id = $1`,
		uuid,
	)
	err := row.Scan(&opiRow.ID, &opiRow.Asunto, &opiRow.Contenido, &opiRow.Autor)
	if err != nil {
		return opinion.Opinion{}, fmt.Errorf("error al leer la opinion por uuid")
	}

	return convertOpinionRowToOpinion(opiRow), nil
}
