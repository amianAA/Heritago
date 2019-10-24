package transformations

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
	gql "heritago/backend/gql/models"
	dbm "heritago/backend/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *dbm.User) (o *gql.User, err error) {
	o = &gql.User{
		ID:        i.ID.String(),
		Email:     i.Email,
		UserID:    i.UserID,
		FirstName: i.FirstName,
		LastName:  i.LastName,
		Country:   i.Country,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
	return o, err
}

// GQLInputUserToDBUser transforms [user] gql input to db model
func GQLInputUserToDBUser(i *gql.UserInput, update bool, ids ...string) (o *dbm.User, err error) {
	userID := strings.ToLower((*i.FirstName)[:1] + *i.LastName)
	hashedPassword := dbm.HashAndSaltPwd(i.Password)

	o = &dbm.User{
		UserID:    &userID,
		Email:     *i.Email,
		Password:  &hashedPassword,
		FirstName: i.FirstName,
		LastName:  i.LastName,
		Country:   i.Country,
	}
	if i.Email == nil && !update {
		return nil, errors.New("Field [email] is required")
	}
	if i.Email != nil {
		o.Email = *i.Email
	}
	if len(ids) > 0 {
		updID, err := uuid.FromString(ids[0])
		if err != nil {
			return nil, err
		}
		o.ID = updID
	}
	return o, err
}
