package email

import (
	"colleague/taoyuan-shop-api/kit/httpreq"
)

const sendEmailUrl string = "https://gateway.p2shop.com.cn/alert-service/mail"

func SendEmail(from string, to []string, subject string, message string) error {
	var body struct {
		From    string   `json:"from"`
		To      []string `json:"to"`
		Subject string   `json:"subject"`
		Message string   `json:"message"`
	}
	body.From = from
	body.To = to
	body.Subject = subject
	body.Message = message

	if err := httpreq.POST("", sendEmailUrl, body, nil); err != nil {
		return err
	}

	return nil
}

func Emailtemplate(headline string, emailType string, emailUrl string) string {
	message := `
	<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Pangpang</title>
		</head>
		<body>
		<table border="0" cellpadding="0" cellspacing="0" align="center" style="width:100%; background-color:#f8f8f9;">
			<tbody>
			<tr>
				<td align="center">
					<div style="max-width:595px; margin:0 auto">
						<table border="0" cellpadding="0" cellspacing="0" align="center" style="max-width:595px; width:100%; background-color:#fff; font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; text-align:left">
							<tbody>
							<tr>
								<td height="30"></td>
							</tr>
							<tr>
								<td style="padding-right: 24px; padding-left: 24px">
									<table border="0" cellpadding="0" cellspacing="0" style="font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; text-align:left">
										<tbody>
										<tr>
											<td style="font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; font-size:20px; font-weight: 700">Pangpang</td>
											<td style="font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; font-size:14px; font-weight: 400">` + headline + `</td>
										</tr>
										</tbody>
									</table>
								</td>
							</tr>
							<tr>
								<td height="20"></td>
							</tr>
							<tr>
								<td style="padding-right: 24px; padding-left: 24px">
									<table border="0" cellpadding="0" cellspacing="0" style="width: 100%; font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; text-align:left">
										<tbody>
										<tr>
											<td style="font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif">
												<strong>请打开一下链接绑定微信.</strong>
											</td>
										</tr>
										<tr>
										<td height="20"></td>
								     	</tr>
										<tr>
											<td style="font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; font-size:20px;">
												<strong style="color: #20c1dc"><a href=` + emailUrl + `>` + emailUrl + `</a></strong>
											</td>
										</tr>
										<tr>
											<td height="70"></td>
										</tr>
										<tr>
											<td style="font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; font-size: 14px; color: #666666">
												感谢你使用Pangpang.
											</td>
										</tr>
										<tr>
											<td style="font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; font-size: 14px; color: #666666">
											    为了提供更便利的服务，我们会尽最大的努力.
											</td>
										</tr>
										<tr>
											<td height="10"></td>
										</tr>
										</tbody>
									</table>
								</td>
							</tr>
							<tr>
								<td height="20"></td>
							</tr>
							</tbody>
						</table>
					</div>
					<div style="max-width:595px; margin:0 auto">
						<table border="0" cellpadding="0" cellspacing="0" align="center" style="max-width:595px; width:100%; background-color:#f9f9f9; font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; text-align:left">
							<tbody>
							<tr>
								<td style="padding-top: 24px; padding-right: 24px; padding-bottom: 2px; padding-left: 24px">
									<table border="0" cellpadding="0" cellspacing="0" style="font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; text-align:left">
										<tbody>
										<tr>
											<td style="font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; font-size: 12px; color: #8899aa">
										    	电子邮件是发电专用.<br />
												关于pangpang服务相关的问题，请在pangpang<strong style="text-decoration: underline">顾客中心</strong>确认
											</td>
										</tr>
										<tr>
											<td style="padding-top: 20px; font-family:'맑은고딕',Malgun Gothic,'돋움',Dotum,Helvetica,'Apple SD Gothic Neo',Sans-serif; font-size: 12px; color: #8899aa">
												Copyright ⓒ Pangpang Corp. All Rights Reserved.
											</td>
										</tr>
										</tbody>
									</table>
								</td>
							</tr>
							</tbody>
						</table>
					</div>
				</td>
			</tr>
			<tr>
				<td height="22" style="background:#f9f9f9;"></td>
			</tr>
			</tbody>
		</table>
		</body>
		</html>
	`
	return message
}
