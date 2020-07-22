package postgre
import (
	"github.com/Raven57/belajargraphql/graph/model"
	"github.com/go-pg/pg/v10"
)

type CommentsRepo struct {
	DB *pg.DB
}

func (u *CommentsRepo) CommentsByUserID(id string) ([]*model.Comment, error) {

	var user []*model.Comment
	err := u.DB.Model(&user).Where("user_id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (u *CommentsRepo) GetAllComments() ([]*model.Comment, error) {
	var users []*model.Comment
	err := u.DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}

	return users, nil
}



