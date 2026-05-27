package psiphon

func getAndroidPackageName() (string, error) { _ = "STUB: not implemented"; return "", nil }

// /proc/self/cmdline is NULL-terminated; read up to the first NULL.

// On success the NULL will be included; trim regardless.

// On Android, each app has a private data directory /data/user/<user-id>/<package>/ (something like /data/user/0/com.example.app/)
// Older devices, before multi user, used /data/data/<packageName>/. For backwards-compatibility, the OS still maintains symlinks so /data/data/... resolves to /data/user/<user-id>/ for the current user.
// This uses the directory /data/data/<packageName>/cache/
// It validates the directory exists (or creates it).
func androidPrivateCacheDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Prefer legacy symlink location; it resolves to /data/user/<id>/<pkg>.

// Ensure it exists (normally already does).

func getUserCacheDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

// For every other system os.UserCacheDir works okay
