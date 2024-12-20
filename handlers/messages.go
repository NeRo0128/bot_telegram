package handlers

//todo: Estructurar bien los mensajes
//	pasar a ingles

const (
	helpMsg = `
		*Bienvenido al Bot de AniList!*
		Un bot para obtener y mostrar información sobre tu anime y manga favorito.

		*Puedes realizar las siguientes acciones:*
			- Buscar un anime por nombre.
			- Obtener una lista de los animes que se están emitiendo actualmente.
			- Ver esta ayuda nuevamente.
	`

	unknownCommandMsg = `
		*Command not recognized *
			help :/help
	`
	starMsg string = `
		Hello, I'm @%s.\nI am a sample bot to demonstrate how file sending works.
		Try the /help command!
		_A bot to fetch and display information about your favorite anime and manga\!_
	`
)
