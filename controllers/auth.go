package controllers

import (
    "github.com/BenMz/base_api/app"
    "github.com/BenMz/base_api/models"
    beego "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    jwt "github.com/dgrijalva/jwt-go"
    "time"
    "strconv"
)

func init() {

}
type AuthController struct {
    beego.Controller
}

func (c *AuthController) Post() {
    username := c.GetString("username");
    password := c.GetString("password");
    o := orm.NewOrm();

    var user models.ExternalUser
    //TODO: cipher password
    qs := o.QueryTable("external_user").Filter("username",username).Filter("password",password)
    err := qs.One(&user)

    response := &app.AuthResponse {
        Success: 0,
        Message: "unknown_error",
    }

    if err != nil {
        response.Message = "Invalid username or password";
    } else {
        response.Success = 1;
        response.Message = "Success";
        exp := int64(time.Now().UTC().Unix()+36000);
        strexp := strconv.FormatInt(exp, 10)



        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "token": user.Token,
            "expire_at": strexp,
        });
        
        tokenString, _ := token.SignedString(app.Secret_key);
        var responseUser *app.Auth = new(app.Auth);
        responseUser.Token = tokenString;
        responseUser.Expire_at = exp;
        response.Data  = responseUser;
    }

    c.Data["json"] = response;
    c.ServeJSON();
}
