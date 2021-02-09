package mock

import (
	"github.com/sniddunc/refractor/refractor"
)

func GetMockServers() map[int64]*refractor.Server {
	mockUsers := map[int64]*refractor.Server{
		1: {
			ServerID:     1,
			Name:         "Test Server #1",
			Address:      "127.0.0.1",
			RCONPort:     "1337",
			RCONPassword: "rconpassword",
		},
		2: {
			ServerID:     2,
			Name:         "Test Server #2",
			Address:      "192.168.0.1",
			RCONPort:     "8002",
			RCONPassword: "passwordrcon2",
		},
	}

	return mockUsers
}

type mockServerRepo struct {
	servers map[int64]*refractor.Server
}

func NewMockServerRepository(mockServers map[int64]*refractor.Server) refractor.ServerRepository {
	return &mockServerRepo{
		servers: mockServers,
	}
}

func (r *mockServerRepo) Create(server *refractor.Server) error {
	newID := int64(len(r.servers) + 1)
	r.servers[newID] = server

	server.ServerID = newID

	return nil
}

func (r *mockServerRepo) FindByID(id int64) (*refractor.Server, error) {
	foundServer := r.servers[id]

	if foundServer == nil {
		return nil, refractor.ErrNotFound
	}

	return foundServer, nil
}

func (r *mockServerRepo) Exists(args refractor.FindArgs) (bool, error) {
	for _, server := range r.servers {
		if args["ServerID"] != nil && args["ServerID"].(int64) != server.ServerID {
			continue
		}

		if args["Name"] != nil && args["Name"].(string) != server.Name {
			continue
		}

		// If none of the above conditions failed, return true since it's a match
		return true, nil
	}

	// If no matches were found, return false by default
	return false, nil
}

func (r *mockServerRepo) FindOne(args refractor.FindArgs) (*refractor.Server, error) {
	for _, server := range r.servers {
		if args["ServerID"] != nil && args["ServerID"].(int64) != server.ServerID {
			continue
		}

		if args["Name"] != nil && args["Name"].(string) != server.Name {
			continue
		}

		if args["Address"] != nil && args["Address"].(string) != server.Address {
			continue
		}

		if args["RCONPort"] != nil && args["RCONPort"].(string) != server.RCONPort {
			continue
		}

		if args["RCONPassword"] != nil && args["RCONPassword"].(string) != server.RCONPassword {
			continue
		}

		// If none of the above conditions failed, return user since it's a match
		return server, nil
	}

	// If no matches were found, return ErrNotFound by default
	return nil, refractor.ErrNotFound
}

func (r *mockServerRepo) FindAll() ([]*refractor.Server, error) {
	var allServers []*refractor.Server

	for _, server := range r.servers {
		allServers = append(allServers, server)
	}

	return allServers, nil
}

func (r *mockServerRepo) Update(id int64, args refractor.UpdateArgs) (*refractor.Server, error) {
	if r.servers[id] == nil {
		return nil, refractor.ErrNotFound
	}

	if args["Name"] != nil {
		r.servers[id].Name = args["Name"].(string)
	}

	if args["Address"] != nil {
		r.servers[id].Address = args["Address"].(string)
	}

	if args["RCONPort"] != nil {
		r.servers[id].RCONPort = args["RCONPort"].(string)
	}

	if args["RCONPassword"] != nil {
		r.servers[id].RCONPassword = args["RCONPassword"].(string)
	}

	return r.servers[id], nil
}

func (r *mockServerRepo) Delete(id int64) error {
	server := r.servers[id]
	if server != nil {
		delete(r.servers, id)
	}

	return nil
}
