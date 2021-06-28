package refractor

type PlayerInfractionService interface {
	GetPlayerInfractionCount(playerID int64) (int, *ServiceResponse)
}
