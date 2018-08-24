/*
 * Copyright © 2018 Alex G Rice
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 */

/* ex3.4 Following the approach of the lissajous example in section 1.7,
 * construct a web server that computes surfaces and writes svg data to the
 * client. The server must set the content-type header like this:

w.Header().Set("content-type", "image/svg+xml")

This step was not required in the lissajous example because the server uses
standard heuristics to recognize common formats like PNG from the first 512
bytes of the response, and generates the proper header. Allow the client to
specify values like height, width and color as http request parameters.

*/
package ex3_4
