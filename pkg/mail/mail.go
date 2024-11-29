package mail

import (
	"github.com/FoyonaCZY/QweeBlog/models"
	"github.com/FoyonaCZY/QweeBlog/pkg/config"
	"github.com/FoyonaCZY/QweeBlog/util"
	"github.com/go-gomail/gomail"
	"github.com/google/uuid"
	"strconv"
)

func generateActivationToken() string {
	return uuid.New().String()
}

func SendActivationEmail(email string) error {
	//检查激活码是否存在
	user, err := models.GetUserByEmail(email)
	if err != nil {
		util.Error("Failed to get user by email: ", err)
	}
	if user.ActivationToken == "nil" {
		//生成激活码
		user.ActivationToken = generateActivationToken()

		//保存激活码
		err = models.DB.Save(&user).Error
		if err != nil {
			util.Error("Failed to save activation token: ", err)
			return err
		}
	}

	//发送激活邮件
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.Configs.Smtp.User)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", config.Configs.Smtp.Nickname+" : 请激活您的账户")
	mailer.SetBody("text/html",
		"请点击此链接激活您的账户: "+
			"<a href='http://"+config.Configs.Site.Domain+"/activate?id="+strconv.Itoa(int(user.ID))+"&token="+user.ActivationToken+"'>激活账户</a>")

	dialer := gomail.NewDialer(config.Configs.Smtp.Host, config.Configs.Smtp.Port, config.Configs.Smtp.User, config.Configs.Smtp.Password)
	dialer.SSL = true
	err = dialer.DialAndSend(mailer)
	if err != nil {
		util.Error("Failed to send activation email: ", err)
		return err
	}
	return nil
}
