package mock

import (
	"github.com/sniddunc/refractor/refractor"
)

type mockInfractionsRepo struct {
	infractions map[int64]*refractor.Infraction
}

func NewMockInfractionRepository(mockInfractions map[int64]*refractor.Infraction) refractor.InfractionRepository {
	return &mockInfractionsRepo{
		infractions: mockInfractions,
	}
}

func (r *mockInfractionsRepo) Create(infraction *refractor.Infraction) (*refractor.Infraction, error) {
	newID := int64(len(r.infractions) + 1)
	r.infractions[newID] = infraction

	infraction.InfractionID = newID

	return infraction, nil
}

func (r *mockInfractionsRepo) FindByID(id int64) (*refractor.Infraction, error) {
	foundServer := r.infractions[id]

	if foundServer == nil {
		return nil, refractor.ErrNotFound
	}

	return foundServer, nil
}

type Infraction struct {
	InfractionID int64  `json:"id"`
	PlayerID     int64  `json:"playerId"`
	UserID       int64  `json:"userId"`
	ServerID     int64  `json:"serverId"`
	Type         string `json:"type"`
	Reason       string `json:"reason"`
	Duration     int    `json:"duration"`
	Timestamp    int64  `json:"timestamp"`
	SystemAction bool   `json:"systemAction"`
	StaffName    string `json:"staffName"` // not a database field
}

func (r *mockInfractionsRepo) FindManyByPlayerID(playerID int64) ([]*refractor.Infraction, error) {
	var foundInfractions []*refractor.Infraction

	for _, infraction := range r.infractions {
		if infraction.PlayerID == playerID {
			foundInfractions = append(foundInfractions, infraction)
		}
	}

	return foundInfractions, nil
}

func (r *mockInfractionsRepo) Exists(args refractor.FindArgs) (bool, error) {
	for _, infraction := range r.infractions {
		if args["InfractionID"] != nil && args["InfractionID"].(int64) != infraction.InfractionID {
			continue
		}

		if args["PlayerID"] != nil && args["PlayerID"].(int64) != infraction.PlayerID {
			continue
		}

		if args["UserID"] != nil && args["UserID"].(int64) != infraction.UserID {
			continue
		}

		if args["ServerID"] != nil && args["ServerID"].(int64) != infraction.ServerID {
			continue
		}

		if args["Type"] != nil && args["Type"].(string) != infraction.Type {
			continue
		}

		if args["Reason"] != nil && args["Reason"].(string) != infraction.Reason {
			continue
		}

		if args["Duration"] != nil && args["Duration"].(int) != infraction.Duration {
			continue
		}

		// If none of the above conditions failed, return true since it's a match
		return true, nil
	}

	// If no matches were found, return false by default
	return false, nil
}

func (r *mockInfractionsRepo) FindOne(args refractor.FindArgs) (*refractor.Infraction, error) {
	for _, infraction := range r.infractions {
		if args["InfractionID"] != nil && args["InfractionID"].(int64) != infraction.InfractionID {
			continue
		}

		if args["PlayerID"] != nil && args["PlayerID"].(int64) != infraction.PlayerID {
			continue
		}

		if args["UserID"] != nil && args["UserID"].(int64) != infraction.UserID {
			continue
		}

		if args["ServerID"] != nil && args["ServerID"].(int64) != infraction.ServerID {
			continue
		}

		if args["Type"] != nil && args["Type"].(string) != infraction.Type {
			continue
		}

		if args["Reason"] != nil && args["Reason"].(string) != infraction.Reason {
			continue
		}

		if args["Duration"] != nil && args["Duration"].(int) != infraction.Duration {
			continue
		}

		// If none of the above conditions failed, return user since it's a match
		return infraction, nil
	}

	// If no matches were found, return ErrNotFound by default
	return nil, refractor.ErrNotFound
}

func (r *mockInfractionsRepo) FindAll() ([]*refractor.Infraction, error) {
	var allServers []*refractor.Infraction

	for _, infraction := range r.infractions {
		allServers = append(allServers, infraction)
	}

	return allServers, nil
}

func (r *mockInfractionsRepo) Update(id int64, args refractor.UpdateArgs) (*refractor.Infraction, error) {
	if r.infractions[id] == nil {
		return nil, refractor.ErrNotFound
	}

	if args["Reason"] != nil {
		r.infractions[id].Reason = args["Reason"].(string)
	}

	if args["Duration"] != nil {
		r.infractions[id].Reason = args["Duration"].(string)
	}

	return r.infractions[id], nil
}

func (r *mockInfractionsRepo) Delete(id int64) error {
	infraction := r.infractions[id]
	if infraction != nil {
		delete(r.infractions, id)
	}

	return nil
}
