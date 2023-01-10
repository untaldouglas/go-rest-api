package db

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
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

func (d *Database) PostOpinion(ctx context.Context, opi opinion.Opinion) (opinion.Opinion, error) {
	opi.ID = uuid.NewV4().String()
	postRow := OpinionRow{
		ID:        opi.ID,
		Asunto:    sql.NullString{String: opi.Asunto, Valid: true},
		Contenido: sql.NullString{String: opi.Contenido, Valid: true},
		Autor:     sql.NullString{String: opi.Autor, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO opinions
		(id, asunto, contenido, autor)
		VALUES
		(:id, :asunto, :contenido, :autor)`,
		postRow,
	)
	if err != nil {
		return opinion.Opinion{}, fmt.Errorf("error al insertar opinion: %w", err)
	}
	if err := rows.Close(); err != nil {
		return opinion.Opinion{}, fmt.Errorf("error al cerrar rows: %w", err)
	}

	return opi, nil

}

func (d *Database) DeleteOpinion(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM opinions where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("error al eliminar opinion en la database: %w", err)
	}
	return nil
}

func (d *Database) UpdateOpinion(ctx context.Context, id string, opi opinion.Opinion) (opinion.Opinion, error) {
	opiRow := OpinionRow{
		ID:        id,
		Asunto:    sql.NullString{String: opi.Asunto, Valid: true},
		Contenido: sql.NullString{String: opi.Contenido, Valid: true},
		Autor:     sql.NullString{String: opi.Autor, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATA opinions SET
		asunto = :asunto,
		contenido = :contenido,
		autor = :autor
		WHERE id = :id`,
		opiRow,
	)

	if err != nil {
		return opinion.Opinion{}, fmt.Errorf("error al actualizar opinion: %w", err)
	}
	if err := rows.Close(); err != nil {
		return opinion.Opinion{}, fmt.Errorf("error al cerrar rows: %w", err)
	}

	return convertOpinionRowToOpinion(opiRow), nil
}
