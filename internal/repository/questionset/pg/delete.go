package pg

import "context"

func (r *repo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM question_set WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
