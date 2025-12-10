package main

import (
	"context"
	"errors"
	"ftpclient/ftpclient"
	"os"
	"path"
	"time"
)

// App struct
type App struct {
	ctx    context.Context
	client *ftpclient.FtpClient
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		client: ftpclient.NewClient(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Connect(server, username, password string, anon bool) error {
	if server == "" || (username == "" && anon == false) {
		return errors.New("please enter valid connection credentials")
	}

	if anon == false {
		a.client.Username = username
		a.client.Password = password
	} else {
		a.client.Username = "anonymous"
		a.client.Password = ""
	}

	a.client.Server = server

	err := a.client.Connect()
	if err != nil {
		return err
	}

	time.Sleep(1 * time.Second) // Unironically will help with any timing issues. Fix this later

	return nil
}

func (a *App) Close() {
	a.client.Close()
}

func (a *App) ListDirectory() ([]*ftpclient.FtpEntry, error) {
	list, err := a.client.GetDirectoryList()
	if err != nil {
		return nil, errors.New("unable to fetch directory list")
	}

	parsed, err := ftpclient.ParseDir(list)
	if err != nil {
		return nil, errors.New("unable to parse directory listing")
	}

	return parsed, nil
}

func (a *App) ChangeDirectory(path string) error {
	return a.client.ChangeDirectory(path)
}

func (a *App) DownloadFile(name string) error {
	file, err := a.client.GetFile(name)
	if err != nil {
		return errors.New("unable to download file")
	}

	if _, err := os.Stat("./files"); errors.Is(err, os.ErrNotExist) {
		_ = os.Mkdir("./files", 0666)
	}

	err = os.WriteFile(path.Join("./files/", name), file, 0666)
	if err != nil {
		return errors.New("unable to save downloaded file")
	}

	return nil
}
