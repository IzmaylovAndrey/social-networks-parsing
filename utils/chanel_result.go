package utils

import "sync"

type APIHandlersResult struct {
	Facebook []string
	VK []string
	Github []string
	Errors []error
	sync.Mutex
}

