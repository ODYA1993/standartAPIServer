package api

func (a *API) configureRouterField() {
	a.router.Get("/users/api", a.GetUsersByAPI)
	a.router.Post("/user", a.PostUser)
	a.router.Get("/users", a.GetAllUsers)
	a.router.Get("/user/{id}", a.GetUserByID)
	a.router.Delete("/user/{id}", a.DeleteUserByID)
	a.router.Delete("/users", a.DeleteUsers)
	a.router.Put("/user/{id}", a.UpdateUserByID)
}
