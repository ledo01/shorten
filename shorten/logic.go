package shorten

import (
	"errors"
	"net/url"
	"time"

	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
)

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type redirectService struct {
	repo RedirectRepository
}

func NewRedirectService(repo RedirectRepository) RedirectService {
	return &redirectService{
		repo,
	}
}

func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.repo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
	if _, err := url.ParseRequestURI(redirect.URL); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}

	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.repo.Store(redirect)
}
