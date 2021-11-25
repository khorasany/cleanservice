package echo

func (r *rest) routing() {
	r.echo.POST("api/v1/login", r.handler.login)
	r.echo.POST("api/v1/register", r.handler.register)
}
