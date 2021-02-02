package mock

import (
	"github.com/sniddunc/refractor/refractor"
	"golang.org/x/crypto/bcrypt"
)

type MockUser struct {
	UnhashedPassword string `json:"unhashedPassword"`
	*refractor.User
}

func GetMockUsers() map[int64]*MockUser {
	return map[int64]*MockUser{
		1: {
			UnhashedPassword: "password",
			User: &refractor.User{
				UserID:              1,
				Email:               "test@test.com",
				Username:            "tester",
				Password:            HashPassword("password"),
				AccessLevel:         0,
				Activated:           false,
				NeedsPasswordChange: false,
			},
		},
	}
}

type mockUserRepo struct {
	users map[int64]*MockUser
}

func NewMockUserRepository(mockUsers map[int64]*MockUser) refractor.UserRepository {
	return &mockUserRepo{
		users: mockUsers,
	}
}

func (r *mockUserRepo) Create(user *refractor.User) error {
	r.users[user.UserID] = &MockUser{
		UnhashedPassword: user.Password,
		User:             user,
	}

	r.users[user.UserID].Password = HashPassword(user.Password)

	return nil
}

func (r *mockUserRepo) FindByID(id int64) (*refractor.User, error) {
	foundUser := r.users[id]

	if foundUser == nil {
		return nil, refractor.ErrNotFound
	}

	return foundUser.User, nil
}

// Exists loops through all users in the users map and compares each one to the args passed in, returning true if
// a match was found. It currently supports comparing the following args: UserID and Username.
func (r *mockUserRepo) Exists(args refractor.FindArgs) (bool, error) {
	for _, user := range r.users {
		if args["UserID"] != nil && args["UserID"].(int64) != user.UserID {
			continue
		}

		if args["Username"] != nil && args["Username"].(string) != user.Username {
			continue
		}

		// If none of the above conditions failed, return true since it's a match
		return true, nil
	}

	// If no matches were found, return false by default
	return false, nil
}

// FindOne loops through all users in the users map and compares each one to the args passed in, returning the match
// if a match was found. It currently supports comparing the following args: UserID and Username.
func (r *mockUserRepo) FindOne(args refractor.FindArgs) (*refractor.User, error) {
	for _, user := range r.users {
		if args["UserID"] != nil && args["UserID"].(int64) != user.UserID {
			continue
		}

		if args["Username"] != nil && args["Username"].(string) != user.Username {
			continue
		}

		// If none of the above conditions failed, return user since it's a match
		return user.User, nil
	}

	// If no matches were found, return ErrNotFound by default
	return nil, refractor.ErrNotFound
}

func (r *mockUserRepo) FindMany(args refractor.FindArgs) ([]*refractor.User, error) {
	var users []*refractor.User

	for _, user := range r.users {
		if args["UserID"] != nil && args["UserID"].(int64) != user.UserID {
			continue
		}

		if args["Username"] != nil && args["Username"].(string) != user.Username {
			continue
		}

		// If none of the above conditions failed, append since this user is a match
		users = append(users, user.User)
	}

	// If no matches were found, return ErrNotFound
	if len(users) < 1 {
		return nil, refractor.ErrNotFound
	}

	// Otherwise return the matches
	return users, nil
}

// Update updates a user by ID. The following args are currently supported: Username, Email,  Password, AccessLevel
// Activated, and NeedsPasswordChange.
func (r *mockUserRepo) Update(id int64, args refractor.UpdateArgs) (*refractor.User, error) {
	if r.users[id] == nil {
		return nil, refractor.ErrNotFound
	}

	if args["Username"] != nil {
		r.users[id].Username = args["Username"].(string)
	}

	if args["Email"] != nil {
		r.users[id].Email = args["Email"].(string)
	}

	if args["Password"] != nil {
		r.users[id].Password = args["Password"].(string)
	}

	if args["AccessLevel"] != nil {
		r.users[id].AccessLevel = args["AccessLevel"].(int)
	}

	if args["Activated"] != nil {
		r.users[id].Activated = args["Activated"].(bool)
	}

	if args["NeedsPasswordChange"] != nil {
		r.users[id].NeedsPasswordChange = args["NeedsPasswordChange"].(bool)
	}

	return r.users[id].User, nil
}

func (r *mockUserRepo) FindAll() ([]*refractor.User, error) {
	var allUsers []*refractor.User

	for _, user := range r.users {
		allUsers = append(allUsers, user.User)
	}

	return allUsers, nil
}

func (r *mockUserRepo) GetCount() int {
	return len(r.users)
}

func HashPassword(password string) string {
	hashAndSalt, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashAndSalt)
}
