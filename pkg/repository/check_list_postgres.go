package repository

import (
	"Brachium/pkg/entity"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type CheckListPostgres struct {
	db *sql.DB
}

func NewCheckListPostgres(db *sql.DB) *CheckListPostgres {
	return &CheckListPostgres{db: db}
}

func (r *CheckListPostgres) Create(userId int, list entity.CheckList) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (overall_percent_complete, start_date, end_date, user_id) "+
		"VALUES ($1, $2, $3, $4) RETURNING list_id", checkListsTable)
	row := r.db.QueryRow(query, list.OverallPercentComplete, list.StartDate, list.EndDate, userId)
	if err := row.Scan(&list.ListId); err != nil {
		return 0, err
	}
	return list.ListId, nil
}

func (r *CheckListPostgres) GetAll(userId int) ([]entity.CheckList, error) {
	var lists []entity.CheckList
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", checkListsTable)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		l := entity.CheckList{}
		err := rows.Scan(&l.ListId, &l.OverallPercentComplete, &l.StartDate, &l.EndDate, &l.UserId)
		if err != nil {
			return nil, err
		}
		lists = append(lists, l)
	}

	return lists, nil
}

func (r *CheckListPostgres) GetById(userId, listId int) (entity.CheckList, error) {
	var list = entity.CheckList{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND list_id = $2", checkListsTable)
	row := r.db.QueryRow(query, userId, listId)
	if err := row.Scan(&list.ListId, &list.OverallPercentComplete, &list.StartDate, &list.EndDate, &list.UserId); err != nil {
		return entity.CheckList{}, err
	}
	return list, nil
}

func (r *CheckListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND list_id = $2", checkListsTable)
	_, err := r.db.Exec(query, userId, listId)
	return err
}

func (r *CheckListPostgres) Update(userId, listId int, input entity.UpdateCheckListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.OverallPercentComplete != nil {
		setValues = append(setValues, fmt.Sprintf("overall_percent_complete=$%d", argId))
		args = append(args, *input.OverallPercentComplete)
		argId++
	}
	if input.StartDate != nil {
		setValues = append(setValues, fmt.Sprintf("start_date=$%d", argId))
		args = append(args, *input.StartDate)
		argId++
	}
	if input.EndDate != nil {
		setValues = append(setValues, fmt.Sprintf("end_date=$%d", argId))
		args = append(args, *input.EndDate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id=$%d AND list_id=$%d",
		checkListsTable, setQuery, argId, argId+1)
	args = append(args, userId, listId)

	result, err := r.db.Exec(query, args...)
	if num, _ := result.RowsAffected(); num == 0 {
		return errors.New("no entries affected")
	}
	return err
}
