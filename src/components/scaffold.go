package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Scaffold struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *Scaffold) Render() app.UI {
	return app.H1().Style("color", "red").Text("SCAFFOLD")
}
