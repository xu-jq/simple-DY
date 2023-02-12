package key

import "strconv"

func KeyUserFavorite(uid int64) string {
	key := "uid" + strconv.FormatInt(uid, 10)
	return key
}
