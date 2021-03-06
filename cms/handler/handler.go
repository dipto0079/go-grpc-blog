package handler

import (
	
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	tpb "go-grpc-blog/gunk/v1/blog"
	
)

const sessionsName = "cms-session"

type Handler struct {
	templates *template.Template
	decoder   *schema.Decoder
	sess      *sessions.CookieStore
	tc       tpb.BlogServiceClient
}

func New(decoder *schema.Decoder, sess *sessions.CookieStore ,tc tpb.BlogServiceClient) *mux.Router {
	h := &Handler{
		decoder: decoder,
		sess:    sess,
		tc:tc,
	}

	h.parseTemplates()
	r := mux.NewRouter()

	r.HandleFunc("/blog/create", h.BlogCreate)
	r.HandleFunc("/blog/store", h.BlogStore)

	// l := r.NewRoute().Subrouter()
	// l.HandleFunc("/login", h.login)
	// l.HandleFunc("/Registration", h.registrationCreate)
	// l.HandleFunc("/User/Store", h.UserStore)
	// l.HandleFunc("/User/login", h.userLogin)
	// l.HandleFunc("/verified", h.userVerified)

	// l.Use(h.loginMiddleware)

	// r.HandleFunc("/User/logout", h.userLogout)
	// r.HandleFunc("/Forgot", h.userForgot)
	// r.HandleFunc("/User/Forgot/check", h.userForgotCheck)
	// r.HandleFunc("/Forgot/verified", h.forgotCheck)
	// r.HandleFunc("/Password", h.password)
	// r.HandleFunc("/Password/Update", h.passwordUpdate)
	//s.Use(h.loginMiddleware)
	//r.HandleFunc("/home/Searching", h.homeSearching)
	//Category
	// s := r.NewRoute().Subrouter()
	// s.HandleFunc("/", h.Home)
	// s.HandleFunc("/Category/List", h.categoryList)
	// s.HandleFunc("/Category/create", h.categoryCreate)
	// s.HandleFunc("/Category/store", h.categoryStore)
	// s.HandleFunc("/Category/{id:[0-9]+}/edit", h.categoryEdit)
	// s.HandleFunc("/Category/{id:[0-9]+}/update", h.categoryUpdate)
	// s.HandleFunc("/Category/{id:[0-9]+}/delete", h.categoryDelete)

	// //Book
	// s.HandleFunc("/Book/List", h.bookList)
	// s.HandleFunc("/Book/Create", h.bookCreate)
	// s.HandleFunc("/Book/store", h.bookStore)
	// s.HandleFunc("/Book/{id:[0-9]+}/active", h.bookActive)
	// s.HandleFunc("/Book/{id:[0-9]+}/deactivate", h.bookDeactivate)
	// s.HandleFunc("/Book/{id:[0-9]+}/edit", h.bookEdit)
	// s.HandleFunc("/Book/{id:[0-9]+}/update", h.bookUpdate)
	// s.HandleFunc("/Book/{id:[0-9]+}/delete", h.bookdelete)
	// //Booking
	// s.HandleFunc("/Booking/{id:[0-9]+}/Create", h.bookingCreate)
	// s.HandleFunc("/Booking/{id:[0-9]+}/Booking", h.bookingStore)
	// s.HandleFunc("/Booking/{id:[0-9]+}/single-list", h.bookiSingleList)
	// s.Use(h.authMiddleware)

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	r.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := h.templates.ExecuteTemplate(rw, "404.html", nil); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	return r
}

func (h *Handler) parseTemplates() {
	h.templates = template.Must(template.ParseFiles(
		// "templates/index.html",
		// "templates/create-category.html",
		// "templates/list-category.html",
		// "templates/edit-category.html",
		// "templates/create-book.html",
		// "templates/list-book.html",
		// "templates/edit-book.html",
		"cms/assets/templates/404.html",
		// "templates/create-booking.html",
		// "templates/registration.html",
		// "templates/login.html",
		// "templates/single-book.html",
		// "templates/single-booking.html",
		// "templates/auth-recoverpw.html",
		// "templates/con_pass.html",
		"cms/assets/templates/create_blog.html",
	))
}

// func (h *Handler) authMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		session, _ := h.sess.Get(r, "library")
// 		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
// 			http.Redirect(rw, r, "/login", http.StatusTemporaryRedirect)
// 			return
// 		}
// 		next.ServeHTTP(rw, r)
// 	})
// // }
// func (h *Handler) loginMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		session, err := h.sess.Get(r, sessionsName)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		authuser := session.Values["authenticated"]
// 		if authuser != nil {
// 			http.Redirect(rw, r, "/", http.StatusTemporaryRedirect)
// 		} else {
// 			next.ServeHTTP(rw, r)
// 		}
// 	})
// }
