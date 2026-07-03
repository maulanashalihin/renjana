package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/maulanashalihin/laju-go/app/models"
	"github.com/maulanashalihin/laju-go/app/queries"
	"github.com/maulanashalihin/laju-go/app/services"
	"github.com/maulanashalihin/laju-go/app/session"
)

// UserAdminHandler handles admin-only user management pages.
type UserAdminHandler struct {
	store          *session.Store
	inertiaService *services.InertiaService
	userAdminSvc   *services.UserAdminService
	querier        *queries.Querier
}

func NewUserAdminHandler(
	store *session.Store,
	inertiaService *services.InertiaService,
	userAdminSvc *services.UserAdminService,
	querier *queries.Querier,
) *UserAdminHandler {
	return &UserAdminHandler{
		store:          store,
		inertiaService: inertiaService,
		userAdminSvc:   userAdminSvc,
		querier:        querier,
	}
}

func (h *UserAdminHandler) authUser(c *fiber.Ctx) (int64, *models.User, error) {
	userID := c.Locals("user_id")
	if userID == nil {
		return 0, nil, fiber.ErrUnauthorized
	}
	id := userID.(int64)
	u, err := h.querier.GetUserByID(c.Context(), id)
	if err != nil {
		return 0, nil, err
	}
	return id, u, nil
}

// Index — list admin users only.
func (h *UserAdminHandler) Index(c *fiber.Ctx) error {
	_, user, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "20"))
	search := c.Query("search", "")

	result, err := h.userAdminSvc.ListUsers(c.Context(), services.UserFilter{
		Role:   models.RoleAdmin,
		Search: search,
	}, page, perPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load users: " + err.Error(),
		})
	}

	adminCount, _ := h.userAdminSvc.CountByRole(c.Context(), models.RoleAdmin)

	return h.inertiaService.Render(c, "app/Users", fiber.Map{
		"user":           user,
		"users":          result,
		"current_search": search,
		"admin_count":    adminCount,
		"all_roles":      []models.UserRole{models.RoleAdmin},
	})
}

// Create — render create form.
func (h *UserAdminHandler) Create(c *fiber.Ctx) error {
	return c.Redirect("/admin/users?action=create")
}

// Store — handle POST /admin/users.
func (h *UserAdminHandler) Store(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	role := models.UserRole(c.FormValue("role"))
	districtID, _ := strconv.ParseInt(c.FormValue("district_id", "0"), 10, 64)
	volunteerID, _ := strconv.ParseInt(c.FormValue("volunteer_id", "0"), 10, 64)

	if name == "" || email == "" || password == "" {
		h.store.Flash(c, "error", "Name, email, and password are required")
		return c.Redirect("/admin/users?action=create")
	}

	_, err = h.userAdminSvc.CreateUser(c.Context(), name, email, password, role, districtID, volunteerID)
	if err != nil {
		h.store.Flash(c, "error", "Failed to create user: "+err.Error())
		return c.Redirect("/admin/users?action=create")
	}

	h.store.Flash(c, "success", "User created successfully")
	return c.Redirect("/admin/users?success=created")
}

// Edit — render edit form.
func (h *UserAdminHandler) Edit(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	return c.Redirect(fmt.Sprintf("/admin/users?action=edit&id=%d", id))
}

// UpdateRole — handle PUT /admin/users/:id/role.
func (h *UserAdminHandler) UpdateRole(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	role := models.UserRole(c.FormValue("role"))
	districtID, _ := strconv.ParseInt(c.FormValue("district_id", "0"), 10, 64)
	volunteerID, _ := strconv.ParseInt(c.FormValue("volunteer_id", "0"), 10, 64)

	if err := h.userAdminSvc.UpdateUserRole(c.Context(), id, role, districtID, volunteerID); err != nil {
		h.store.Flash(c, "error", "Failed to update role: "+err.Error())
		return c.Redirect(fmt.Sprintf("/admin/users?action=edit&id=%d", id))
	}

	h.store.Flash(c, "success", "User role updated")
	return c.Redirect("/admin/users?success=updated")
}

// ToggleActive — toggle is_active status.
func (h *UserAdminHandler) ToggleActive(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	isActive := c.FormValue("active") == "true"

	if err := h.userAdminSvc.SetActive(c.Context(), id, isActive); err != nil {
		h.store.Flash(c, "error", "Failed to toggle active: "+err.Error())
		return c.Redirect("/admin/users?error=" + err.Error())
	}

	h.store.Flash(c, "success", "User status updated")
	return c.Redirect("/admin/users?success=status-updated")
}

// Destroy — delete user.
func (h *UserAdminHandler) Destroy(c *fiber.Ctx) error {
	_, _, err := h.authUser(c)
	if err != nil {
		return c.Redirect("/login")
	}
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	if err := h.userAdminSvc.DeleteUser(c.Context(), id); err != nil {
		h.store.Flash(c, "error", "Failed to delete user: "+err.Error())
		return c.Redirect("/admin/users?error=" + err.Error())
	}

	h.store.Flash(c, "success", "User deleted")
	return c.Redirect("/admin/users?success=deleted")
}
