package lib

import (
	"os"

	"github.com/go-git/go-git/v5"
)

// IsRepoClean will check if a git repo directory is clean
func IsRepoClean(dir string) (bool, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false, err
	}
	r, err := git.PlainOpen(dir)
	if err != nil {
		return false, err
	}
	w, err := r.Worktree()
	if err != nil {
		return false, err
	}
	stats, err := w.Status()
	if err != nil {
		return false, err
	}
	// for k, fileStatus := range stats {
	// 	fmt.Println(k)
	// 	stagingStatusCode := *&fileStatus.Staging
	// 	fmt.Println(stagingStatusCode)
	// }
	return stats.IsClean(), nil
}
