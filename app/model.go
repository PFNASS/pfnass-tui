package app

// import (
// 	tea "github.com/charmbracelet/bubbletea"
// )

type model struct {
    services  []string           // pfnass services
    choice	int   				// selected service
	// selected map[int]struct{}   // selected service
	chosen	 bool				// is a service currently selected
    cursor   int                // which service the cursor is pointed at
	loaded  bool               // has the model been loaded?
	quitting bool               // is the user quitting?
}