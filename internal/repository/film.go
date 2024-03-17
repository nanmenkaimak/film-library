package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/nanmenkaimak/film_library/internal/entity"
)

func (r *Repo) CreateFilm(newFilm entity.FilmeWithActors) (uuid.UUID, error) {
	var filmID uuid.UUID

	query := `insert into films (name, description, release_date, rating)
				values ($1, $2, $3, $4) returning id`

	tx := r.main.Db.MustBegin()
	defer tx.Rollback()
	err := tx.QueryRowx(query, newFilm.Name, newFilm.Description, newFilm.ReleaseDate, newFilm.Rating).Scan(&filmID)
	if err != nil {
		return uuid.Nil, err
	}
	query = `insert into films_actors (film_id, actor_id) values ($1, $2)`
	for i := 0; i < len(newFilm.Actors); i++ {
		_, err = tx.MustExec(query, filmID, newFilm.Actors[i].ID).RowsAffected()
		if err != nil {
			return uuid.Nil, err
		}
	}
	if err = tx.Commit(); err != nil {
		return uuid.Nil, err
	}
	return filmID, nil
}

func (r *Repo) UpdateFilm(film entity.UpdateMap) error {
	query := `update films set `

	for k, v := range film.Values {
		str := fmt.Sprintf(" %s = '%v',", k, v)
		query += str
	}
	query = query[:len(query) - 1]

	query += fmt.Sprintf(" where id = '%v'", film.ID)

	err := r.main.Db.QueryRowx(query).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteFilm(filmID uuid.UUID) error {
	query := `delete from films where id = $1`

	err := r.main.Db.QueryRowx(query, filmID).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetFilms(sorting string) ([]entity.FilmeWithActors, error) {
	var films []entity.FilmeWithActors

	if sorting == "" {
		sorting = "rating"
	}

	query := fmt.Sprintf("select * from films order by %s desc", sorting)

	err := r.replica.Db.Select(&films, query)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(films); i++ {
		query = `select a.id, a.name, a.gender, a.birth_day from actors a inner join films_actors fa on fa.actor_id = a.id where fa.film_id = $1`
		err = r.replica.Db.Select(&films[i].Actors, query, films[i].ID)
		if len(films[i].Actors) == 0 {
			films[i].Actors = []entity.Actor{}
		}
		if err != nil {
			return nil, err
		}
	}

	return films, err
}

func (r *Repo) GetFilmsByName(name string) ([]entity.FilmeWithActors, error) {
	var films []entity.FilmeWithActors

	query := fmt.Sprintf(`select * from films where name like '%s%%'`, name)

	err := r.replica.Db.Select(&films, query)
	if err != nil {
		return nil, err
	}
	if len(films) == 0 {
		return []entity.FilmeWithActors{}, nil
	}

	for i := 0; i < len(films); i++ {
		query = `select a.id, a.name, a.gender, a.birth_day from actors a inner join films_actors fa on fa.actor_id = a.id where fa.film_id = $1`
		err = r.replica.Db.Select(&films[i].Actors, query, films[i].ID)
		if len(films[i].Actors) == 0 {
			films[i].Actors = []entity.Actor{}
		}
		if err != nil {
			return nil, err
		}
	}

	return films, err
}

