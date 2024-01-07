package domain

type ServerRepository interface {
	List() ([]string, error)
	GetRanking(ServerId string) ([]map[string]interface{}, error)
	GetUserRank(ServerId string, UserId string) (map[string]interface{}, error)
	GetTickets(ServerId string) (map[string]interface{}, error)
	EditTickets(ServerId string, tickets map[string]interface{}) error
}
