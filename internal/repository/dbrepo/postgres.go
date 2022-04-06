package dbrepo

import (
	"context"
	"database/sql"
	"log"
	"time"

	"bitbucket.org/janpavtel/site/internal/models"
)

type postgresDBRepo struct {
	DB *sql.DB
}

func (pg *postgresDBRepo) LoadAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `
	select 
		id, first_name, email
	from 
		users
	`

	var users []models.User

	rows, err := pg.DB.QueryContext(ctx, stmt)
	if err != nil {
		log.Println("Can't query users from postgres", err)
		return users, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.Email,
		)

		if err != nil {
			log.Println("Can't scan result to user model", err)
			return users, err
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Println("Rows are left in error stage", err)
		return users, err
	}

	return users, nil
}

func (pg *postgresDBRepo) LoadUserById(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `
	select 
		id, first_name, email
	from 
		users
	where 
		id = $1
	`

	var user models.User

	row := pg.DB.QueryRowContext(ctx, stmt, id)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.Email,
	)

	if err != nil {
		log.Println("Can't scan result to user model", err)
		return user, err
	}

	return user, nil
}

func (pg *postgresDBRepo) AddUser(user models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `insert into users (first_name,
		email) 
		values ($1, $2)
		RETURNING (id)
		`

	row := pg.DB.QueryRowContext(ctx, stmt,
		user.FirstName,
		user.Email,
	)

	var id int

	err := row.Scan(
		&id,
	)

	if err != nil {
		log.Println("Can't scan result id", err)
		return id, err
	}

	return id, nil
}
