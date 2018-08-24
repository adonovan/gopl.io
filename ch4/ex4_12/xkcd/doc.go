/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

// The popular web comic xkcd has a JSON interface. For example, a request to
// https://xkcd.com/571/info.0.json produces a detailed description of comic 571
// one of many favorites. Download each URL (once!) and build an offline index.
// Write a tool xkcd that, using this index, prints the URL and transcript of
// each comic that matches a search term provided on the command line.
package xkcd
