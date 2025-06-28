package main

import (
	// "fmt"
	// "os"
	// "strings"
	// "time"
	// "context"
	// "github.com/PFNASS/pfnass-tui"

	// tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

// type model struct {
// 	initError 	error
// 	ctx 		context.Context
// 	cancel 		context.CancelFunc
// 	pages		map[pages.PageName]pages.PageInterface
// 	pageStack 	[]pages.Stack
// }

// func initialModel() model {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	helpPage := pages.NewHelpPage()
// 	pagesMap := map[pages.PageName]pages.PageInterface{
// 		pages.HelpPage: helpPage,
// 	}
// 	pageStack := []pages.Stack{}
// 	m := model{
// 		initError: 	nil,
// 		ctx: 		ctx,
// 		cancel: 	cancel,
// 		pages: 		pagesMap,
// 		pageStack: 	pageStack,
// 	}
// 	return m
// }	