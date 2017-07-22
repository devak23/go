package syncutils

import "sync"

var Wg sync.WaitGroup

/**
	the purpose of this file/code is to only provide a global variable Wg = WaitGroup and package it
	so it can be included in all the other code which need it.
 */
