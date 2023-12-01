package structs

type roleSummaryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func RoleSummaryResponse(roles []Role) []roleSummaryResponse {
	listRoles := []roleSummaryResponse{}

	for _, role := range roles {
		roleFormatter := roleSummaryResponse{
			ID:   role.ID,
			Name: role.Name,
		}

		listRoles = append(listRoles, roleFormatter)
	}

	return listRoles
}
