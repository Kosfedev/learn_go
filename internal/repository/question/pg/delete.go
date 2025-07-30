package pg

import "context"

func (r *repo) Delete(_ context.Context, id int) error {
	query := `DELETE FROM question WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
