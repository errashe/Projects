package main

type SteamApiMatches struct {
	Result struct {
		Matches []struct {
			DireTeamID  float64 `json:"dire_team_id"`
			LobbyType   float64 `json:"lobby_type"`
			MatchID     float64 `json:"match_id"`
			MatchSeqNum float64 `json:"match_seq_num"`
			Players     []struct {
				AccountID  float64 `json:"account_id"`
				HeroID     float64 `json:"hero_id"`
				PlayerSlot float64 `json:"player_slot"`
			} `json:"players"`
			RadiantTeamID float64 `json:"radiant_team_id"`
			StartTime     float64 `json:"start_time"`
		} `json:"matches"`
		NumResults       float64 `json:"num_results"`
		ResultsRemaining float64 `json:"results_remaining"`
		Status           float64 `json:"status"`
		TotalResults     float64 `json:"total_results"`
	} `json:"result"`
}
