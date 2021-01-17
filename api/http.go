package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ledo01/shorten/serializer/json"
	"github.com/ledo01/shorten/shorten"
	"github.com/pkg/errors"
)

type RedirectHandler interface {
	Get(c *fiber.Ctx) error
	Post(c *fiber.Ctx) error
}

type handler struct {
	redirectService shorten.RedirectService
}

func NewHandler(service shorten.RedirectService) RedirectHandler {
	return &handler{redirectService: service}
}

func (h *handler) serializer() shorten.RedirectSerializer {
	return &json.Redirect{}
}

func (h *handler) Get(c *fiber.Ctx) error {
	code := c.Params("code")
	redirect, err := h.redirectService.Find(code)
	if err != nil {
		if errors.Cause(err) == shorten.ErrRedirectNotFound {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	c.Redirect(redirect.URL, fiber.StatusMovedPermanently)

	return nil
}

func (h *handler) Post(c *fiber.Ctx) error {
	redirect := &shorten.Redirect{}
	if err := c.BodyParser(redirect); err != nil {
		return err
	}

	err := h.redirectService.Store(redirect)
	if err != nil {
		if errors.Cause(err) == shorten.ErrRedirectInvalid {
			return fiber.ErrBadRequest
		}
	}

	if err := c.Status(fiber.StatusCreated).JSON(redirect); err != nil {
		return err
	}

	return nil
}
