package controller

import (
	"basic-crud-go/internal/app/admin/enterprise/binding"
	"basic-crud-go/internal/app/admin/enterprise/dto"
	"basic-crud-go/internal/app/admin/enterprise/service"
	"basic-crud-go/internal/app/admin/enterprise/util"
	"basic-crud-go/internal/configuration/rest_err"
	"net/http"

	"github.com/gin-gonic/gin"
)

type enterpriseController struct {
	service service.EnterpriseService
}

func NewEnterpriseController(s service.EnterpriseService) EnterpriseController {
	return &enterpriseController{
		service: s,
	}
}

// Ping godoc
// @Summary      Healthcheck do Enterprise
// @Description  Retorna um pong para verificar se o serviço enterprise está ativo
// @Tags         Enterprise
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /enterprise/v1/ping [get]
func (c *enterpriseController) Ping(ctx *gin.Context) {
	result, _ := c.service.Ping(ctx.Request.Context())

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

// CreateEnterpriseHandler godoc
// @Summary 			Create enterprise
// @Description 		Create enterprise by CNPJ and name
// @Tags 				Enterprise
// @Accept       		json
// @Produce      		json
// @Param				request		body 		dto.CreateEnterpriseDTO true "Company data"
// @Success      		200     	{object}  	dto.CreateEnterpriseResponse
// @Failure      		400      	{object}  	rest_err.RestErr
// @Failure     		500      	{object}  	rest_err.RestErr
// @Router       /enterprise/v1/create [post]
func (c *enterpriseController) CreateEnterpriseHandler(ctx *gin.Context) {
	var req dto.CreateEnterpriseDTO
	// Bind JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("missing or invalid required fields", []rest_err.Causes{
			rest_err.NewCause("body", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}
	// Bind DTO
	if err := binding.ValidateCreateEnterpriseDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}
	// Create enterprise
	created, err := c.service.Create(ctx, req.Name, util.RemoveNonDigits(req.Cnpj))
	if err != nil {
		restErr := rest_err.NewInternalServerError("error creating enterprise", []rest_err.Causes{})
		ctx.JSON(restErr.Code, restErr)
		return
	}
	ctx.JSON(http.StatusOK, &dto.CreateEnterpriseResponse{
		Name:      created.Name,
		Cnpj:      created.Cnpj,
		CreatedAt: created.CreateAt,
	})
}

// ReadEnterpriseHandler godoc
// @Summary      Read enterprise
// @Description  Read enterprise by CNPJ
// @Tags         Enterprise
// @Accept       json
// @Produce      json
// @Param        cnpj  path      string  true  "CNPJ of the enterprise"
// @Success      200   {object}  dto.ReadEnterpriseResponse
// @Failure      400   {object}  rest_err.RestErr
// @Failure      500   {object}  rest_err.RestErr
// @Router       /enterprise/v1/read/{cnpj} [get]
func (c *enterpriseController) ReadEnterpriseHandler(ctx *gin.Context) {
	req, err := binding.ValidateReadEnterpriseDTO(ctx)
	if err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// read enterprise
	enterprise, err := c.service.Read(ctx, util.RemoveNonDigits(req.Cnpj))
	if err != nil {
		restErr := rest_err.NewNotFoundError("enterprise not found")
		ctx.JSON(restErr.Code, restErr)
		return
	}
	ctx.JSON(http.StatusOK, &dto.ReadEnterpriseResponse{
		Name:      enterprise.Name,
		Cnpj:      enterprise.Cnpj,
		CreatedAt: enterprise.CreateAt,
		UpdatedAt: enterprise.UpdateAt,
	})
}

// ReadEnterprisesHandler godoc
// @Summary      List enterprises
// @Description  Retrieve a paginated list of enterprises
// @Tags         Enterprise
// @Accept       json
// @Produce      json
// @Param        page   query     int     false   "Page number (min 1)"
// @Param        limit  query     int     false   "Items per page (default: 10)"
// @Success      200    {object}  dto.ReadEnterprisesResponse
// @Failure      400    {object}  rest_err.RestErr
// @Failure      500    {object}  rest_err.RestErr
// @Router       /enterprise/v1/read [get]
func (c *enterpriseController) ReadEnterprisesHandler(ctx *gin.Context) {
	req := binding.ValidateReadEnterprisesDTO(ctx)
	//if err != nil {
	//	restErr := rest_err.NewBadRequestValidationError("invalid request body", []rest_err.Causes{
	//		rest_err.NewCause("validation", err.Error()),
	//	})
	//	ctx.JSON(restErr.Code, restErr)
	//	return
	//}

	// Fetch enterprises
	enterprises, err := c.service.ReadAllEnterprise(ctx, req.Page, req.Limit)
	if err != nil {
		restErr := rest_err.NewInternalServerError("failed to fetch enterprises", []rest_err.Causes{
			rest_err.NewCause("read enteprises", "error"),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}

	// Map to DTO
	var result []dto.ReadEnterpriseResponse
	for _, ent := range enterprises {
		result = append(result, dto.ReadEnterpriseResponse{
			Name:      ent.Name,
			Cnpj:      ent.Cnpj,
			CreatedAt: ent.CreateAt,
			UpdatedAt: ent.UpdateAt,
		})
	}

	ctx.JSON(http.StatusOK, dto.ReadEnterprisesResponse{
		Page:        req.Page,
		Limit:       req.Limit,
		Total:       len(result),
		Enterprises: result,
	})
}

// UpdateEnterpriseHandler godoc
// @Summary 			Update enterprise
// @Description 		Update enterprise by CNPJ
// @Tags 				Enterprise
// @Accept       		json
// @Produce      		json
// @Param				request		body 		dto.UpdateEnterpriseDTO true "Company data"
// @Success      		200     	{object}  	dto.UpdateEnterpriseResponse
// @Failure      		400      	{object}  	rest_err.RestErr
// @Failure     		500      	{object}  	rest_err.RestErr
// @Router       /enterprise/v1/update [put]
func (c *enterpriseController) UpdateEnterpriseHandler(ctx *gin.Context) {
	var req dto.UpdateEnterpriseDTO
	// Bind JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("missing or invalid required fields", []rest_err.Causes{
			rest_err.NewCause("body", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}
	// Bind DTO
	if err := binding.ValidateUpdateEnterpriseDTO(req); err != nil {
		restErr := rest_err.NewBadRequestValidationError("invalid request body", []rest_err.Causes{
			rest_err.NewCause("validation", err.Error()),
		})
		ctx.JSON(restErr.Code, restErr)
		return
	}
	// Update enterprise
	updated, err := c.service.Update(ctx, util.RemoveNonDigits(req.Cnpj), util.RemoveNonDigits(req.NewCnpj), req.NewName)
	if err != nil {
		switch err.Error() {
		case "enterprise not found":
			restErr := rest_err.NewNotFoundError(err.Error())
			ctx.JSON(restErr.Code, restErr)
		case "enterprise_cnpj_key":
			restErr := rest_err.NewBadRequestError("cnpj already exists")
			ctx.JSON(restErr.Code, restErr)
		default:
			restErr := rest_err.NewInternalServerError(err.Error(), []rest_err.Causes{})
			ctx.JSON(restErr.Code, restErr)
		}
		return
	}
	ctx.JSON(http.StatusOK, &dto.UpdateEnterpriseResponse{
		OldCnpj:   &req.Cnpj,
		NewName:   &updated.Name,
		NewCnpj:   updated.Cnpj,
		UpdatedAt: updated.UpdateAt,
	})
}
