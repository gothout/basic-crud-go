package tokencache

import (
	"basic-crud-go/internal/app/admin/auth/model"
	"sync"
	"time"
)

// cacheEntry holds the user identity and its expiration time
type cacheEntry struct {
	Identity  *model.UserIdentity
	ExpiresAt time.Time
}

// tokenCache is an in-memory cache with mutex protection
type tokenCache struct {
	mu     sync.RWMutex
	tokens map[string]*cacheEntry
	ttl    time.Duration
}

// singleton instance of the cache with default TTL of 10 seconds
var cache = &tokenCache{
	tokens: make(map[string]*cacheEntry),
	ttl:    10 * time.Second,
}

// SaveToken stores the user identity in the cache using email as key
func SaveToken(email string, identity *model.UserIdentity) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.tokens[email] = &cacheEntry{
		Identity:  identity,
		ExpiresAt: time.Now().Add(cache.ttl),
	}
}

// GetToken retrieves the cached user identity if it hasn't expired
// Returns (identity, true) if found and valid; otherwise (nil, false)
func GetToken(email string) (*model.UserIdentity, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	entry, exists := cache.tokens[email]
	if !exists || time.Now().After(entry.ExpiresAt) {
		return nil, false
	}
	return entry.Identity, true
}

// RefreshToken updates the expiration time to 1 hour if the token matches
func RefreshToken(email string, token string) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, exists := cache.tokens[email]
	if !exists || entry.Identity == nil || entry.Identity.TokenUser == nil {
		return false
	}

	if entry.Identity.TokenUser.Token != token {
		return false
	}

	// Update expiration time to 1 hour from now
	entry.ExpiresAt = time.Now().Add(1 * time.Hour)
	return true
}

// Logout removes the token from cache (used for logoff)
func Logout(email string) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	delete(cache.tokens, email)
}

// GetRemainingTTL returns the time remaining before the token expires.
// Returns (duration, true) if found and valid; otherwise (0, false).
func GetRemainingTTL(email string) (time.Duration, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	entry, exists := cache.tokens[email]
	if !exists || time.Now().After(entry.ExpiresAt) {
		return 0, false
	}

	return time.Until(entry.ExpiresAt), true
}
