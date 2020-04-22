package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"net/url"
	"smarthead/internal/account"
	"smarthead/internal/geonames"
	"smarthead/internal/platform/auth"
	"smarthead/internal/platform/web"
	"smarthead/internal/platform/web/webcontext"
	"smarthead/internal/platform/web/weberror"
	"smarthead/internal/student"
	"smarthead/internal/user"
	"smarthead/internal/user_account"
	"smarthead/internal/user_auth"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/astaxie/beego"
	"github.com/pborman/uuid"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

type UserController struct {
	beego.Controller

	UserRepo        *user.Repository
	AuthRepo        *user_auth.Repository
	UserAccountRepo *user_account.Repository
	AccountRepo     *account.Repository
	GeoRepo         *geonames.Repository
	MasterDB        *sql.DB
	SecretKey       string
}

func NewUserController(userRepo *user.Repository,
	authRepo        *user_auth.Repository,
	userAccountRepo *user_account.Repository,
	accountRepo     *account.Repository,
	geoRepo         *geonames.Repository,
	masterDB        *sql.DB,
	secretKey       string) *UserController {

	c := &UserController{
		UserRepo: userRepo,
		AuthRepo: authRepo,
		UserAccountRepo: userAccountRepo,
		GeoRepo: geoRepo,
		MasterDB: masterDB,
		SecretKey: secretKey,
	}

	return c
}

// UserLoginRequest extends the AuthenicateRequest with the RememberMe flag.
type UserLoginRequest struct {
	user_auth.AuthenticateRequest
	RememberMe bool
}

func (h *UserController) Login()  {
	r := h.Ctx.Request
	w := h.Ctx.ResponseWriter
	
	ctx := h.Ctx.Request.Context()
	now := time.Now()

	//
	req := new(UserLoginRequest)
	f := func() (bool, error) {

		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				return false, err
			}

			decoder := schema.NewDecoder()
			if err := decoder.Decode(req, r.PostForm); err != nil {
				return false, err
			}

			sessionTTL := time.Hour
			if req.RememberMe {
				sessionTTL = time.Hour * 36
			}

			// Authenticated the user.
			token, err := h.AuthRepo.Authenticate(ctx, user_auth.AuthenticateRequest{
				Email:    req.Email,
				Password: req.Password,
			}, sessionTTL, now)
			if err != nil {
				switch errors.Cause(err) {
				case user.ErrForbidden:
					return false, web.RespondError(ctx, w, weberror.NewError(ctx, err, http.StatusForbidden))
				case user_auth.ErrAuthenticationFailure:
					h.Data["error"] = weberror.NewErrorMessage(ctx, errors.New("Check your username and password and try again"),
					 http.StatusUnauthorized, "Authentication failure.")
					return false, nil
				default:
					if verr, ok := weberror.NewValidationError(ctx, err); ok {
						h.Data["validationErrors"] = verr.(*weberror.Error)
						return false, nil
					} else {
						return false, err 
					}
				}
			}

			// Add the token to the users session.
			err = handleSessionToken(ctx, w, r, token)
			if err != nil {
				return false, err
			}

			redirectUri := "/"
			if qv := r.URL.Query().Get("redirect"); qv != "" {
				redirectUri, err = url.QueryUnescape(qv)
				if err != nil {
					return false, err
				}
			}

			// Redirect the user to the dashboard.
			return true, web.Redirect(ctx, w, r, redirectUri, http.StatusFound)
		}

		return false, nil
	}

	end, err := f()
	if err != nil {
		h.Data["error"] = err.Error()
		return
	} else if end {
		return
	}

	h.Data["form"] = req

	if verr, ok := weberror.NewValidationError(ctx, webcontext.Validator().Struct(UserLoginRequest{})); ok {
		h.Data["validationDefaults"] = verr.(*weberror.Error)
	}


	h.TplName = "user/login.html"
}

// handleSessionToken persists the access token to the session for request authentication.
func handleSessionToken(ctx context.Context, w http.ResponseWriter, r *http.Request, token user_auth.Token) error {
	if token.AccessToken == "" {
		return errors.New("accessToken is required.")
	}

	sess := webcontext.ContextSession(ctx)

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(token.TTL.Seconds()),
		HttpOnly: false,
	}

	sess = webcontext.SessionInit(sess,
		token.AccessToken)
	if err := sess.Save(r, w); err != nil {
		return err
	}

	return nil
}

// updateContextClaims updates the claims in the context.
func updateContextClaims(ctx context.Context, authenticator *auth.Authenticator, claims auth.Claims) (context.Context, error) {
	tkn, err := authenticator.GenerateToken(claims)
	if err != nil {
		return ctx, err
	}

	sess := webcontext.ContextSession(ctx)
	sess = webcontext.SessionUpdateAccessToken(sess, tkn)

	ctx = context.WithValue(ctx, auth.Key, claims)

	return ctx, nil
}

func (c *UserController) Register()  {
	c.TplName = "user/register.html"
}

func (c *UserController) RegisterStudent() {
	flash := beego.NewFlash()
	projectID := beego.AppConfig.String("project_id")
	client, err := firestore.NewClient(c.Ctx.Request.Context(), projectID)
	if err != nil {
		flash.Error("Cannot connect to the database, ", err.Error())
		c.Redirect("/", http.StatusPermanentRedirect)
		return
	}
	defer client.Close()

	req := new(student.CreateRequest)
	_, _, err = client.Collection("students").Add(c.Ctx.Request.Context(), map[string]interface{}{
		"ID": uuid.NewRandom().String(),
		"Name": req.Name,
		"Phone": req.ParentPhone,
		"Email": req.ParentEmail,
		"Class": req.ClassID,
	})
	if err != nil {
		flash.Error("Error in added record to database. Please try again later or contact the admin for help, ", err.Error())
		c.Redirect("/", http.StatusPermanentRedirect)
		return
	}

	flash.Success("Congratulation! Your request has been submitted successfully. You will receive and email from us once your class is scheduled")
	c.Redirect("/", 308)
}