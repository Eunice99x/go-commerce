package api

func (a *API) SetupRoutes() error {
	api := a.app.Group("/api")

	groupUsers := api.Group("/users")

	groupUsers.Name("List All Users").Get("", a.ListUsers)
	groupUsers.Name("Create a User").Post("", a.CreateUser)
	groupUsers.Name("Delete a User").Delete("/:id", a.DeleteUser)
	groupUsers.Name("Update a User").Put("/:id", a.UpdateUser)

	return nil
}
