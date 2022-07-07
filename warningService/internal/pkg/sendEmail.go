package pkg

import "warningService/internal/service"

func DeliverEmail(emailContent string, emailAddr string) error {
	// param prepared
	var notifyMsg = emailContent
	var notifyEmail = emailAddr
	var emailConn = service.EmailConnector{
		// should be written in config
		User:      "1305479162@qq.com",
		PassToken: "hrcvcyndmbashfaf",
		Host:      "smtp.qq.com",
		Port:      587,
	}
	var emailBody = service.EmailBody{
		MailTo: []string{
			notifyEmail,
		},
		Subject: "WARNING",
		Content: notifyMsg,
	}
	// 发送给指定用户
	help := service.SendEmail{}

	err := help.SendMail(emailConn, emailBody)
	if err != nil {
		return err
	}
	return nil
}
