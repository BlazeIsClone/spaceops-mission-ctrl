package mission

import (
	"database/sql"
	"fmt"

	"github.com/blazeisclone/spaceops-mission-ctrl/domain"
)

type MySQLMissionRepository struct {
	db *sql.DB
}

func NewMySQLMissionRepository(db *sql.DB) *MySQLMissionRepository {
	return &MySQLMissionRepository{db: db}
}

func (r *MySQLMissionRepository) GetAll() (*[]domain.Mission, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM missions"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var missions []domain.Mission

	for rows.Next() {
		var mission domain.Mission
		var createdAt, updatedAt []uint8

		err := rows.Scan(&mission.ID, &mission.Name, &mission.Description, &createdAt, &updatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning mission: %w", err)
		}

		missions = append(missions, mission)
	}
	if err = rows.Err(); err != nil {
		return &missions, err
	}

	return &missions, nil
}

func (r *MySQLMissionRepository) Create(mission *domain.Mission) error {
	query := "INSERT INTO missions (name, description) VALUES (?, ?)"
	result, err := r.db.Exec(query, mission.Name, mission.Description)
	if err != nil {
		return fmt.Errorf("error creating mission: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert id: %w", err)
	}

	mission.ID = int(id)
	return nil
}

func (r *MySQLMissionRepository) FindByID(id int) (*domain.Mission, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM missions WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var mission domain.Mission
	var createdAt, updatedAt []uint8
	err := row.Scan(&mission.ID, &mission.Name, &mission.Description, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding mission: %w", err)
	}

	return &mission, nil
}

func (r *MySQLMissionRepository) UpdateByID(id int, mission *domain.Mission) error {
	query := "UPDATE missions SET name = ?, description = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?"
	_, err := r.db.Exec(query, mission.Name, mission.Description, id)
	if err != nil {
		return fmt.Errorf("error updating mission: %w", err)
	}
	return nil
}

func (r *MySQLMissionRepository) DeleteByID(id int) error {
	query := "DELETE FROM missions WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting mission: %w", err)
	}
	return nil
}
