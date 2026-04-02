package utils

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ParsePGIDs(s string) (pgtype.UUID, error) {

	parsedId, err := uuid.Parse(s)

	if err != nil {
		return pgtype.UUID{}, err
	}

	dbId := pgtype.UUID{
		Bytes: parsedId,
		Valid: true,
	}

	return dbId, nil
}

// func ParseUUID(s pgtype.UUID) (uuid.UUID, error) {

// }
