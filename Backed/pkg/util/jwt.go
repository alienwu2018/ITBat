package util

//var jwtSecret = []byte(conf.JWT_SECRET)
//
//type Claims struct {
//	User User.User
//	jwt.StandardClaims
//}
//
//func GenerateToken(id int) (string, error) {
//	nowTime := time.Now()
//	expireTime := nowTime.Add(3 * time.Hour) //token有效期为3小时
//	user, _ := User.GetInfo(id)
//	claims := Claims{
//		user,
//		jwt.StandardClaims{
//			ExpiresAt: expireTime.Unix(),
//			Issuer:    "gin-init",
//		},
//	}
//
//	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err := tokenClaims.SignedString(jwtSecret)
//
//	return token, err
//}
//
//func ParseToken(token string) (*Claims, error) {
//	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})
//
//	if tokenClaims != nil {
//		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
//			return claims, nil
//		}
//	}
//
//	return nil, err
//}
