package set

import "context"

func (qss *serv) Delete(ctx context.Context, id int64) error {
	err := qss.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
