package handlers

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"bitbucket.org/janpavtel/site/internal/drivers"
	"bitbucket.org/janpavtel/site/internal/forms"
	"bitbucket.org/janpavtel/site/internal/models"
	"bitbucket.org/janpavtel/site/internal/repository"
	"bitbucket.org/janpavtel/site/internal/repository/dbrepo"
	"github.com/go-chi/chi/v5"
)

type TemplateData struct {
	Form *forms.Form
	Data map[string]interface{}
}

type View struct {
	DB repository.DatabaseRepo
}

func NewView(db *drivers.DB) *View {
	return &View{
		DB: dbrepo.NewPostgresRepo(db.SQL),
	}
}

func (view *View) Home(w http.ResponseWriter, r *http.Request) {
	err := renderTemplate(w, "home.tpl", nil)
	if err != nil {
		log.Println("Can't render home page ", err)
	}
}

func (view *View) About(w http.ResponseWriter, r *http.Request) {
	err := renderTemplate(w, "about.tpl", nil)
	if err != nil {
		log.Println("Can't render about page ", err)
	}
}

func (view *View) SignUp(w http.ResponseWriter, r *http.Request) {
	err := renderTemplate(w, "signup.tpl", nil)
	if err != nil {
		log.Println("Can't render sing-up page ", err)
	}
}

func (view *View) ShowUsers(w http.ResponseWriter, r *http.Request) {
	users, err := view.DB.LoadAllUsers()
	if err != nil {
		log.Println("Can't load users", err)
		return
	}

	data := make(map[string]interface{})
	data["users"] = users

	err = renderTemplate(w, "users.tpl", &TemplateData{Data: data})
	if err != nil {
		log.Println("Can't render users page ", err)
		return
	}
}

func (view *View) ShowUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := view.DB.LoadUserById(userID)
	if err != nil {
		log.Println("Can't load user", err)
		return
	}

	data := make(map[string]interface{})
	data["user"] = user

	err = renderTemplate(w, "user.tpl", &TemplateData{Data: data})
	if err != nil {
		log.Println("Can't render users page ", err)
		return
	}
}

func (view *View) NewUserCreation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Can't parse form ", err)
		return
	}

	form := forms.New(r.PostForm)
	form.CheckMinLengh("first_name", 3)
	form.CheckRequiredFields("email")

	if form.IsNotValid() {
		renderTemplate(w, "signup.tpl", &TemplateData{
			Form: form,
		})

		return
	}

	newUser := models.User{
		FirstName: r.Form.Get("first_name"),
		Email:     r.Form.Get("email"),
	}

	id, err := view.DB.AddUser(newUser)
	if err != nil {
		log.Println("Can't add new user", err)
		return
	}

	http.Redirect(w, r, "/users/"+strconv.Itoa(id), http.StatusSeeOther)
}

func renderTemplate(w http.ResponseWriter, templateName string, templateData *TemplateData) error {
	if templateData == nil {
		templateData = &TemplateData{}
	}

	parsedTemplate, errTmp := template.ParseFiles("./www/pages/" + templateName)
	if errTmp != nil {
		return errTmp
	}
	resolvedTemplate, errLay := parsedTemplate.ParseFiles("./www/layouts/base.layout.tpl")
	if errLay != nil {
		return errLay
	}

	buf := new(bytes.Buffer)

	errBufData := resolvedTemplate.Execute(buf, templateData)
	if errBufData != nil {
		return errBufData
	}

	_, errExe := buf.WriteTo(w)
	if errExe != nil {
		return errExe
	}

	return nil
}
