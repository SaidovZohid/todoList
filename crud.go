package main

import (
	"database/sql"
	"time"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{db: db}
}

type Todo struct {
	id int
	title string
	description string
	assignee string
	status bool
	deadline time.Time
	created_at time.Time
}

type GetAllParam struct {
	limit int
	page int
	thing string
}
 
func (d *DBManager) Create(t *Todo) (*Todo, error) {
	var todo Todo
	query := `
		INSERT INTO todo(
			title,
			description,
			assignee,
			status,
			deadline
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, description, assignee,status, deadline, created_at
	`
	row := d.db.QueryRow(
		query,
		t.title,
		t.description,
		t.assignee,
		t.status,
		t.deadline,
	)
	err := row.Scan(
		&todo.id,
		&todo.title,
		&todo.description,
		&todo.assignee,
		&todo.status,
		&todo.deadline,
		&todo.created_at,
	)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (d *DBManager) Get(id int) (*Todo, error) {
	query := `
		SELECT 
			id,
			title,
			description,
			assignee, 
			status, 
			deadline,
			created_at
		FROM todo WHERE id = $1
	`
	row := d.db.QueryRow(
		query,
		id,
	)
	var todo Todo
	err := row.Scan(
		&todo.id,
		&todo.title,
		&todo.description,
		&todo.assignee,
		&todo.status,
		&todo.deadline,
		&todo.created_at,
	)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (d *DBManager) Update(t *Todo) (*Todo, error){
	query := `
		UPDATE todo SET 
			title = $1,
			description = $2,
			assignee = $3,
			status = $4,
			deadline = $5
		WHERE id = $6
		RETURNING id, title, description, assignee, status, deadline, created_at
	`
	row := d.db.QueryRow(
		query,
		t.title,
		t.description,
		t.assignee,
		t.status,
		t.deadline,
		t.id,
	)
	var todo Todo
	err := row.Scan(
		&todo.id,
		&todo.title,
		&todo.description,
		&todo.assignee,
		&todo.status,
		&todo.deadline,
		&todo.created_at,
	)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (d *DBManager) Delete(id int) error {
	query := `
		DELETE FROM 
			todo 
		WHERE id = $1
	`
	_, err := d.db.Exec(query, id)
	return err
}

func (d *DBManager) GetAll(g *GetAllParam) ([]*Todo, error) {
	offset := (g.page - 1) * g.limit
	query := `
		SELECT 
			id, 
			title,
			description,
			assignee,
			status,
			deadline,
			created_at
		FROM todo ORDER BY id ASC LIMIT $1 OFFSET $2, 
	`
	rows, err := d.db.Query(
		query,
		g.limit,
		offset,
	)
	if err != nil  {
		return nil, err
	}
	var todos []*Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.id,
			&todo.title,
			&todo.description,
			&todo.assignee,
			&todo.status,
			&todo.deadline,
			&todo.created_at,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo) 
	}
	return todos, nil
}