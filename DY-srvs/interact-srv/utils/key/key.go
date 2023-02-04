package key

import "fmt"

func KeyUserFavorite(uid interface{}) string {
	key := fmt.Sprintf("uid:%s", uid)
	return key
}
