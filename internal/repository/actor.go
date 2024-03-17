package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/nanmenkaimak/film_library/internal/entity"
)

func (r *Repo) CreateActor(newActor entity.Actor) (uuid.UUID, error) {
	query := `insert into actors (name, gender, birth_day) values ($1, $2, $3) returning id`

	var actorID uuid.UUID

	err := r.main.Db.QueryRowx(query, newActor.Name, newActor.Gender, newActor.BirthDay).Scan(&actorID)
	if err != nil {
		return uuid.Nil, err
	}
	return actorID, err
}

func (r *Repo) UpdateActor(actor entity.UpdateMap) error {
	query := `update actors set`

	for k, v := range actor.Values {
		str := fmt.Sprintf(" %s = '%v',", k, v)
		query += str
	}
	query = query[:len(query) - 1]

	query += fmt.Sprintf(" where id = '%v'", actor.ID)

	err := r.main.Db.QueryRowx(query).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteActor(actorID uuid.UUID) error {
	query := `delete from actors where id = $1`

	err := r.main.Db.QueryRowx(query, actorID).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetActors(name string) ([]entity.ActorWithFilms, error) {
	var actorsWithFilms []entity.ActorWithFilms
	queryActor := fmt.Sprintf(`select * from actors where name like '%s%%'`, name)
	err := r.replica.Db.Select(&actorsWithFilms, queryActor)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(actorsWithFilms); i++ {
		query := `select f.id, f.name, f.description, f.release_date, f.rating from films f inner join films_actors fa on fa.film_id = f.id where fa.actor_id = $1`
		err = r.replica.Db.Select(&actorsWithFilms[i].Films, query, actorsWithFilms[i].ID)
		if len(actorsWithFilms[i].Films) == 0 {
			actorsWithFilms[i].Films = []entity.Film{}
		}
		if err != nil {
			return nil, err
		}
	}
	return actorsWithFilms, nil
}