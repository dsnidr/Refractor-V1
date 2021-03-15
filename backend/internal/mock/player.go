package mock

import "github.com/sniddunc/refractor/refractor"

type mockPlayerRepo struct {
	players map[int64]*refractor.Player
}

func NewMockPlayerRepository(mockPlayers map[int64]*refractor.Player) refractor.PlayerRepository {
	return &mockPlayerRepo{
		players: mockPlayers,
	}
}

func (r *mockPlayerRepo) Create(player *refractor.Player) error {
	newID := int64(len(r.players) + 1)
	r.players[newID] = player

	player.PlayerID = newID

	return nil
}

func (r *mockPlayerRepo) FindByID(id int64) (*refractor.Player, error) {
	foundPlayer := r.players[id]

	if foundPlayer == nil {
		return nil, refractor.ErrNotFound
	}

	return foundPlayer, nil
}

func (r *mockPlayerRepo) FindByPlayFabID(playFabID string) (*refractor.Player, error) {
	for _, player := range r.players {
		if player.PlayFabID == playFabID {
			return player, nil
		}
	}

	return nil, refractor.ErrNotFound
}

func (r *mockPlayerRepo) FindOne(args refractor.FindArgs) (*refractor.Player, error) {
	for _, player := range r.players {
		if args["PlayerID"] != nil && args["PlayerID"].(int64) != player.PlayerID {
			continue
		}

		if args["PlayFabID"] != nil && args["PlayFabID"].(string) != player.PlayFabID {
			continue
		}

		if args["LastSeen"] != nil && args["LastSeen"].(int64) != player.LastSeen {
			continue
		}

		return player, nil
	}

	return nil, refractor.ErrNotFound
}

func (r *mockPlayerRepo) Exists(args refractor.FindArgs) (bool, error) {
	for _, player := range r.players {
		if args["PlayerID"] != nil && args["PlayerID"].(int64) != player.PlayerID {
			continue
		}

		if args["PlayFabID"] != nil && args["PlayFabID"].(string) != player.PlayFabID {
			continue
		}

		if args["LastSeen"] != nil && args["LastSeen"].(int64) != player.LastSeen {
			continue
		}

		return true, nil
	}

	return false, nil
}

func (r *mockPlayerRepo) UpdateName(player *refractor.Player, currentName string) error {
	panic("not implemented")
}

func (r *mockPlayerRepo) Update(id int64, args refractor.UpdateArgs) (*refractor.Player, error) {
	if r.players[id] == nil {
		return nil, refractor.ErrNotFound
	}

	if args["PlayerID"] != nil {
		r.players[id].PlayerID = args["PlayerID"].(int64)
	}

	if args["PlayFabID"] != nil {
		r.players[id].PlayFabID = args["PlayFabID"].(string)
	}

	if args["LastSeen"] != nil {
		r.players[id].LastSeen = args["LastSeen"].(int64)
	}

	return r.players[id], nil
}
