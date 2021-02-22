package main

import (
  "net/http"
  "fmt"
  "log"
  "github.com/go-playground/validator"
  "encoding/json"
  "net/url"
)

type Form struct {
  Name string `validate:"required"`
  Email string `validate:"required,email"`
  Text string `validate:"required"`
}

type slackPayload struct {
  Channel string `json:"channel"`
  Username string `json:"username"`
  Message string `json:"text"`
  Icon string `json:"icon_emoji"`
}

var webhookURL string = "https://hooks.slack.com/services/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"

func SendtoSlack(formParams *Form) (err error) {
  
  var msg string
  msg += "[メールアドレス]: "
  msg += string(formParams.Email)
  msg += "\n[送信者]: "
  msg += string(formParams.Name)
  msg += "\n[問い合わせ内容]\n"
  msg += string(formParams.Text)
  fmt.Println(msg)

  payload := slackPayload{
    Channel: "#koeken_dev",
    Username: "お問い合わせ(こえけんHPより)",
    Message: string(msg),
    Icon: ":shipit:",
  }
  p, err := json.Marshal(&payload)
  if err != nil {
    return err
  }

  res, err := http.PostForm(webhookURL, url.Values{"payload": {string(p)}})
  if err != nil {
    return err
  }

  defer res.Body.Close()
  return nil
}

func userValidate(formParams *Form) map[string]string {

  var errorMessages = map[string]string{}

  validate := validator.New()
  errors := validate.Struct(formParams)

  if errors != nil {
    for _, errors := range errors.(validator.ValidationErrors) {

      //var errorMessage map[string]string

      fieldName := errors.Field()

      switch fieldName {
      case "Name":
        errorMessages["Name"] = "名前が正しいか確認してください"
      case "Email":
        errorMessages["Email"] = "メールアドレスが正しいか確認してください"
      case "Text":
        errorMessages["Text"] = "お問い合わせ内容を入力してください"
      }
    }
  }
  return errorMessages
}

func contactFormHandler(w http.ResponseWriter, r *http.Request) {

  //CORS
  w.Header().Set("Access-Control-Allow-Headers", "*")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set( "Access-Control-Allow-Methods", "GET, POST" )

  err := r.ParseForm()

  if err != nil {
    log.Fatalln(err)
  }

  switch(r.Method){
  case "POST":
    name := r.FormValue("name")
    email := r.FormValue("email")
    text := r.FormValue("text")
    
    formParams := &Form{
      Name: name,
      Email: email,
      Text: text,
    }

    errors := userValidate(formParams)
    toJson, err := json.Marshal(errors)
    if err != nil {
      fmt.Println("Json Marshal error:", err)
      return
    }
    fmt.Fprintf(w, string(toJson))
    if string(toJson) != "{}" {
      fmt.Println("Validation Error")
      return
    }
    
    SendtoSlack(formParams)
  
  default:
    fmt.Println("Not Supported!")
  }
}

func main() {
  http.HandleFunc("/api/contactForm", contactFormHandler)
  http.ListenAndServe(":8081", nil)
}