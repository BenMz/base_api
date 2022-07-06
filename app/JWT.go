package app

import(
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/beego/beego/v2/server/web/context"
    "strconv"
    "time"
    // "fmt"
)

func hasExpiredJWT(exp string) bool{
	expire_at, _ := strconv.Atoi(exp);
	time_now := int(time.Now().UTC().Unix());
        
  return time_now>expire_at
}

func ParseJWT(ctx *context.Context) (bool, string){
	var expire_at string;
	jwt_string := ctx.Input.Header("Authorization");
	var key, _ = jwt.Parse(jwt_string, func(token *jwt.Token) (interface{}, error) {
		expire_at = token.Claims.(jwt.MapClaims)["expire_at"].(string)
		
	    return Secret_key, nil
	})
    if key != nil {
		if key.Valid {
			if(!hasExpiredJWT(expire_at)){
				return true, "Success"
			} else {
				return false, "Token has expired"
			}
		}
	}

	return false,"Invalid Token"
	
}