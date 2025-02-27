package hallpers

import (
  "crypto/tls"
  "fmt"
  "io/ioutil"
  "math/rand"
  "time"
//   evt "github.com/masrur-qr/console/Env"
  gomail "gopkg.in/mail.v2"
)
func generateMessageID(domain string) (string, error) {
  b := make([]byte, 16)
  _, err := rand.Read(b)
  if err != nil {
    return "", err
  }
  id := fmt.Sprintf("<%x@%s>", b, domain)
  return id, nil
}
// ? Send email
func SendEmail(email , subject, designPath string) {
  m := gomail.NewMessage()


  m.SetHeader("From", "murtazonroimshoevm4@gmail.com")

  // Set E-Mail receivers
  m.SetHeader("To", email)

  // Set E-Mail subject
  var htmlBody string
  m.SetHeader("Subject", subject)
  ByteTEmp, err := ioutil.ReadFile(designPath)
  if err != nil {
    fmt.Printf("errs: %v\n", err)
  }
  htmlBody = string(ByteTEmp)
  // json.Unmarshal(ByteTEmp,&htmlBody)

  // Set E-Mail body. You can set plain text or html with text/html\
  // text := fmt.Sprintf("Data : %v", Data)
  m.SetBody("text/html", fmt.Sprintf("%v", htmlBody))

  messageID, err := generateMessageID("itisconsole.com")
  if err != nil {
    fmt.Printf("err generateMessageID : %v\n", err)
    return
  }
  m.SetHeader("Message-ID", messageID)

  // Settings for SMTP server
  d := gomail.NewDialer("smtp.gmail.com", 465, "murtazobroimshoevm4@gmail.com","ebnm dgus keae hhut") //evt.Email_Key)
  d.Timeout = 30 * time.Second

  // This is only needed when SSL/TLS certificate is not valid on server.
  // In production this should be set to false.
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

  // Now send E-Mail
  err = d.DialAndSend(m)
  if err != nil {
    fmt.Printf("err DialAndSend : %v\n", err)
  }

  return
}


