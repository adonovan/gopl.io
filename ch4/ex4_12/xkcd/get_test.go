/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

package xkcd

import (
	"testing"
)

// TestGet checks an arbitrary comic file, which is known to exist (614).
func TestGet(t *testing.T) {
	id := 614
	comic, err := Get(id)
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if comic.Num != id {
		t.Errorf("Get failed: want num %v, got %v", id, comic.Num)
	}
	if comic.Title != "Woodpecker" {
		t.Errorf("Get failed: want title %v, got %v", "Woodpecker", comic.Title)
	}
}

// TestGetLatest checks the latest comic file, which at the time of writing is
// #2037. Expect the comic.Num to be >= 2037.
func TestGetLatest(t *testing.T) {
	id := 0
	comic, err := Get(id)
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if comic.Num < 2037 {
		t.Errorf("Get failed: want num > 2037, got %v", comic.Num)
	}
}
