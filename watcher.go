package main

import "github.com/go-fsnotify/fsnotify"

type watcher struct {
	watcher *fsnotify.Watcher
}

func newWatcher() watcher {
	return watcher{attemptGet(fsnotify.NewWatcher()).(*fsnotify.Watcher)}
}

func watchInput(input string) watcher {
	w := newWatcher()
	attempt(w.watcher.Add(input))
	return w
}

func (w watcher) Close() {
	attempt(w.watcher.Close())
}

func logErrors(w watcher) watcher {
	go func() {
		for {
			attempt(<-w.watcher.Errors)
		}
	}()
	return w
}
