package service

import (
	"net/http"

	"github.com/valyala/fasthttp"

	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
)

type AdminService struct {
	AdminRepository *repository.AdminRepository
}

func NewAdminService(repo *repository.AdminRepository) *AdminService {
	return &AdminService{repo}
}

// Cities service methods
func (s *AdminService) GetCities(ctx *fasthttp.RequestCtx) *model.Response {
	cities, err := s.AdminRepository.GetCities(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: cities}
}

func (s *AdminService) CreateCity(ctx *fasthttp.RequestCtx, req *model.CreateCityRequest) *model.Response {
	id, err := s.AdminRepository.CreateCity(ctx, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "City created successfully"}}
}

func (s *AdminService) UpdateCity(ctx *fasthttp.RequestCtx, id int, req *model.UpdateCityRequest) *model.Response {
	err := s.AdminRepository.UpdateCity(ctx, id, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "City updated successfully"}}
}

func (s *AdminService) DeleteCity(ctx *fasthttp.RequestCtx, id int) *model.Response {
	err := s.AdminRepository.DeleteCity(ctx, id)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "City deleted successfully"}}
}

// Brands service methods
func (s *AdminService) GetBrands(ctx *fasthttp.RequestCtx) *model.Response {
	brands, err := s.AdminRepository.GetBrands(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: brands}
}

func (s *AdminService) CreateBrand(ctx *fasthttp.RequestCtx, req *model.CreateBrandRequest) *model.Response {
	id, err := s.AdminRepository.CreateBrand(ctx, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Brand created successfully"}}
}

func (s *AdminService) UpdateBrand(ctx *fasthttp.RequestCtx, id int, req *model.CreateBrandRequest) *model.Response {
	err := s.AdminRepository.UpdateBrand(ctx, id, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Brand updated successfully"}}
}

func (s *AdminService) DeleteBrand(ctx *fasthttp.RequestCtx, id int) *model.Response {
	err := s.AdminRepository.DeleteBrand(ctx, id)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Brand deleted successfully"}}
}

// Models service methods
func (s *AdminService) GetModels(ctx *fasthttp.RequestCtx, brand_id int) *model.Response {
	models, err := s.AdminRepository.GetModels(ctx, brand_id)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: models}
}

func (s *AdminService) CreateModel(ctx *fasthttp.RequestCtx, brand_id int, req *model.CreateModelRequest) *model.Response {
	id, err := s.AdminRepository.CreateModel(ctx, brand_id, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Model created successfully"}}
}

func (s *AdminService) UpdateModel(ctx *fasthttp.RequestCtx, id int, req *model.UpdateModelRequest) *model.Response {
	err := s.AdminRepository.UpdateModel(ctx, id, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Model updated successfully"}}
}

func (s *AdminService) DeleteModel(ctx *fasthttp.RequestCtx, id int) *model.Response {
	err := s.AdminRepository.DeleteModel(ctx, id)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Model deleted successfully"}}
}

// Body Types service methods
func (s *AdminService) GetBodyTypes(ctx *fasthttp.RequestCtx) *model.Response {
	bodyTypes, err := s.AdminRepository.GetBodyTypes(ctx)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: bodyTypes}
}

func (s *AdminService) CreateBodyType(ctx *fasthttp.RequestCtx, req *model.CreateBodyTypeRequest) *model.Response {
	id, err := s.AdminRepository.CreateBodyType(ctx, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Body type created successfully"}}
}

func (s *AdminService) CreateBodyTypeImage(ctx *fasthttp.RequestCtx, id int, paths []string) *model.Response {
	err := s.AdminRepository.CreateBodyTypeImage(ctx, id, paths)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: model.Success{Message: "Body type image created successfully"}}
}

func (s *AdminService) UpdateBodyType(ctx *fasthttp.RequestCtx, id int, req *model.CreateBodyTypeRequest) *model.Response {
	err := s.AdminRepository.UpdateBodyType(ctx, id, req)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Body type updated successfully"}}
}

func (s *AdminService) DeleteBodyType(ctx *fasthttp.RequestCtx, id int) *model.Response {
	err := s.AdminRepository.DeleteBodyType(ctx, id)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return &model.Response{Data: model.Success{Message: "Body type deleted successfully"}}
}

func (s *AdminService) DeleteBodyTypeImage(ctx *fasthttp.RequestCtx, id int) *model.Response {
	err := s.AdminRepository.DeleteBodyTypeImage(ctx, id)

	if err != nil {
		return &model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return &model.Response{Data: model.Success{Message: "Body type image deleted successfully"}}
}

// // Transmissions service methods
// func (s *AdminService) GetTransmissions(ctx *fasthttp.RequestCtx) *model.Response {
// 	transmissions, err := s.AdminRepository.GetTransmissions(ctx)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: transmissions}
// }

// func (s *AdminService) CreateTransmission(ctx *fasthttp.RequestCtx, req *model.CreateTransmissionRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateTransmission(ctx, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Transmission created successfully"}}
// }

// func (s *AdminService) UpdateTransmission(ctx *fasthttp.RequestCtx, id int, req *model.UpdateTransmissionRequest) *model.Response {
// 	err := s.AdminRepository.UpdateTransmission(ctx, id, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: model.Success{Message: "Transmission updated successfully"}}
// }

// func (s *AdminService) DeleteTransmission(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteTransmission(ctx, id)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: model.Success{Message: "Transmission deleted successfully"}}
// }

// // Engines service methods
// func (s *AdminService) GetEngines(ctx *fasthttp.RequestCtx) *model.Response {
// 	engines, err := s.AdminRepository.GetEngines(ctx)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: engines}
// }

// func (s *AdminService) CreateEngine(ctx *fasthttp.RequestCtx, req *model.CreateEngineRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateEngine(ctx, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Engine created successfully"}}
// }

// func (s *AdminService) UpdateEngine(ctx *fasthttp.RequestCtx, id int, req *model.UpdateEngineRequest) *model.Response {
// 	err := s.AdminRepository.UpdateEngine(ctx, id, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: model.Success{Message: "Engine updated successfully"}}
// }

// func (s *AdminService) DeleteEngine(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteEngine(ctx, id)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: model.Success{Message: "Engine deleted successfully"}}
// }

// // Drivetrains service methods
// func (s *AdminService) GetDrivetrains(ctx *fasthttp.RequestCtx) *model.Response {
// 	drivetrains, err := s.AdminRepository.GetDrivetrains(ctx)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: drivetrains}
// }

// func (s *AdminService) CreateDrivetrain(ctx *fasthttp.RequestCtx, req *model.CreateDrivetrainRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateDrivetrain(ctx, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Drivetrain created successfully"}}
// }

// func (s *AdminService) UpdateDrivetrain(ctx *fasthttp.RequestCtx, id int, req *model.UpdateDrivetrainRequest) *model.Response {
// 	err := s.AdminRepository.UpdateDrivetrain(ctx, id, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: model.Success{Message: "Drivetrain updated successfully"}}
// }

// func (s *AdminService) DeleteDrivetrain(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteDrivetrain(ctx, id)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}

// 	return &model.Response{Data: model.Success{Message: "Drivetrain deleted successfully"}}
// }

// // Fuel Types service methods
// func (s *AdminService) GetFuelTypes(ctx *fasthttp.RequestCtx) *model.Response {
// 	fuelTypes, err := s.AdminRepository.GetFuelTypes(ctx)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: fuelTypes}
// }

// func (s *AdminService) CreateFuelType(ctx *fasthttp.RequestCtx, req *model.CreateFuelTypeRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateFuelType(ctx, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Fuel type created successfully"}}
// }

// func (s *AdminService) UpdateFuelType(ctx *fasthttp.RequestCtx, id int, req *model.UpdateFuelTypeRequest) *model.Response {
// 	err := s.AdminRepository.UpdateFuelType(ctx, id, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Fuel type updated successfully"}}
// }

// func (s *AdminService) DeleteFuelType(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteFuelType(ctx, id)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Fuel type deleted successfully"}}
// }

// // Regions service methods
// func (s *AdminService) GetRegions(ctx *fasthttp.RequestCtx, city_id int) *model.Response {
// 	regions, err := s.AdminRepository.GetRegions(ctx, city_id)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: regions}
// }

// func (s *AdminService) CreateRegion(ctx *fasthttp.RequestCtx, city_id int, req *model.CreateRegionRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateRegion(ctx, city_id, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Region created successfully"}}
// }

// func (s *AdminService) UpdateRegion(ctx *fasthttp.RequestCtx, id int, req *model.UpdateRegionRequest) *model.Response {
// 	err := s.AdminRepository.UpdateRegion(ctx, id, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Region updated successfully"}}
// }

// func (s *AdminService) DeleteRegion(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteRegion(ctx, id)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Region deleted successfully"}}
// }

// // Service Types service methods
// func (s *AdminService) GetServiceTypes(ctx *fasthttp.RequestCtx) *model.Response {
// 	serviceTypes, err := s.AdminRepository.GetServiceTypes(ctx)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: serviceTypes}
// }

// func (s *AdminService) CreateServiceType(ctx *fasthttp.RequestCtx, req *model.CreateServiceTypeRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateServiceType(ctx, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Service type created successfully"}}
// }

// func (s *AdminService) UpdateServiceType(ctx *fasthttp.RequestCtx, id int, req *model.UpdateServiceTypeRequest) *model.Response {
// 	err := s.AdminRepository.UpdateServiceType(ctx, id, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Service type updated successfully"}}
// }

// func (s *AdminService) DeleteServiceType(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteServiceType(ctx, id)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Service type deleted successfully"}}
// }

// // Services service methods
// func (s *AdminService) GetServices(ctx *fasthttp.RequestCtx) *model.Response {
// 	services, err := s.AdminRepository.GetServices(ctx)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: services}
// }

// func (s *AdminService) CreateService(ctx *fasthttp.RequestCtx, req *model.CreateServiceRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateService(ctx, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Service created successfully"}}
// }

// func (s *AdminService) UpdateService(ctx *fasthttp.RequestCtx, id int, req *model.UpdateServiceRequest) *model.Response {
// 	err := s.AdminRepository.UpdateService(ctx, id, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Service updated successfully"}}
// }

// func (s *AdminService) DeleteService(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteService(ctx, id)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Service deleted successfully"}}
// }

// // Generations service methods
// func (s *AdminService) GetGenerations(ctx *fasthttp.RequestCtx) *model.Response {
// 	generations, err := s.AdminRepository.GetGenerations(ctx)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: generations}
// }

// func (s *AdminService) CreateGeneration(ctx *fasthttp.RequestCtx, req *model.CreateGenerationRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateGeneration(ctx, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Generation created successfully"}}
// }

// func (s *AdminService) UpdateGeneration(ctx *fasthttp.RequestCtx, id int, req *model.UpdateGenerationRequest) *model.Response {
// 	err := s.AdminRepository.UpdateGeneration(ctx, id, req)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Generation updated successfully"}}
// }

// func (s *AdminService) DeleteGeneration(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteGeneration(ctx, id)

// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Generation deleted successfully"}}
// }

// // Configurations service methods
// func (s *AdminService) GetConfigurations(ctx *fasthttp.RequestCtx) *model.Response {
// 	configurations, err := s.AdminRepository.GetConfigurations(ctx)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: configurations}
// }

// func (s *AdminService) CreateConfiguration(ctx *fasthttp.RequestCtx, req *model.CreateConfigurationRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateConfiguration(ctx, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Configuration created successfully"}}
// }

// func (s *AdminService) UpdateConfiguration(ctx *fasthttp.RequestCtx, id int, req *model.UpdateConfigurationRequest) *model.Response {
// 	err := s.AdminRepository.UpdateConfiguration(ctx, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Configuration updated successfully"}}
// }

// func (s *AdminService) DeleteConfiguration(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteConfiguration(ctx, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Configuration deleted successfully"}}
// }

// // Generation Modifications service methods
// func (s *AdminService) GetGenerationModifications(ctx *fasthttp.RequestCtx, generationId int) *model.Response {
// 	generationModifications, err := s.AdminRepository.GetGenerationModifications(ctx, generationId)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: generationModifications}
// }

// func (s *AdminService) CreateGenerationModification(ctx *fasthttp.RequestCtx, generationId int, req *model.CreateGenerationModificationRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateGenerationModification(ctx, generationId, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Generation modification created successfully"}}
// }

// func (s *AdminService) UpdateGenerationModification(ctx *fasthttp.RequestCtx, generationId int, id int, req *model.UpdateGenerationModificationRequest) *model.Response {
// 	err := s.AdminRepository.UpdateGenerationModification(ctx, generationId, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Generation modification updated successfully"}}
// }

// func (s *AdminService) DeleteGenerationModification(ctx *fasthttp.RequestCtx, generationId int, id int) *model.Response {
// 	err := s.AdminRepository.DeleteGenerationModification(ctx, generationId, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Generation modification deleted successfully"}}
// }

// // Colors service methods
// func (s *AdminService) GetColors(ctx *fasthttp.RequestCtx) *model.Response {
// 	colors, err := s.AdminRepository.GetColors(ctx)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: colors}
// }

// func (s *AdminService) CreateColor(ctx *fasthttp.RequestCtx, req *model.CreateColorRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateColor(ctx, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Color created successfully"}}
// }

// func (s *AdminService) UpdateColor(ctx *fasthttp.RequestCtx, id int, req *model.UpdateColorRequest) *model.Response {
// 	err := s.AdminRepository.UpdateColor(ctx, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Color updated successfully"}}
// }

// func (s *AdminService) DeleteColor(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteColor(ctx, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Color deleted successfully"}}
// }

// // Moto Categories service methods
// func (s *AdminService) GetMotoCategories(ctx *fasthttp.RequestCtx) *model.Response {
// 	motoCategories, err := s.AdminRepository.GetMotoCategories(ctx)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: motoCategories}
// }

// func (s *AdminService) CreateMotoCategory(ctx *fasthttp.RequestCtx, req *model.CreateMotoCategoryRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateMotoCategory(ctx, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Moto category created successfully"}}
// }

// func (s *AdminService) UpdateMotoCategory(ctx *fasthttp.RequestCtx, id int, req *model.UpdateMotoCategoryRequest) *model.Response {
// 	err := s.AdminRepository.UpdateMotoCategory(ctx, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto category updated successfully"}}
// }

// func (s *AdminService) DeleteMotoCategory(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteMotoCategory(ctx, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto category deleted successfully"}}
// }

// // Moto Brands service methods
// func (s *AdminService) GetMotoBrands(ctx *fasthttp.RequestCtx) *model.Response {
// 	motoBrands, err := s.AdminRepository.GetMotoBrands(ctx)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: motoBrands}
// }

// func (s *AdminService) CreateMotoBrand(ctx *fasthttp.RequestCtx, req *model.CreateMotoBrandRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateMotoBrand(ctx, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Moto brand created successfully"}}
// }

// func (s *AdminService) UpdateMotoBrand(ctx *fasthttp.RequestCtx, id int, req *model.UpdateMotoBrandRequest) *model.Response {
// 	err := s.AdminRepository.UpdateMotoBrand(ctx, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto brand updated successfully"}}
// }

// func (s *AdminService) DeleteMotoBrand(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteMotoBrand(ctx, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto brand deleted successfully"}}
// }

// // Moto Models service methods
// func (s *AdminService) GetMotoModels(ctx *fasthttp.RequestCtx) *model.Response {
// 	motoModels, err := s.AdminRepository.GetMotoModels(ctx)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: motoModels}
// }

// func (s *AdminService) CreateMotoModel(ctx *fasthttp.RequestCtx, req *model.CreateMotoModelRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateMotoModel(ctx, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Moto model created successfully"}}
// }

// func (s *AdminService) UpdateMotoModel(ctx *fasthttp.RequestCtx, id int, req *model.UpdateMotoModelRequest) *model.Response {
// 	err := s.AdminRepository.UpdateMotoModel(ctx, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto model updated successfully"}}
// }

// func (s *AdminService) DeleteMotoModel(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteMotoModel(ctx, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto model deleted successfully"}}
// }

// // Moto Parameters service methods
// func (s *AdminService) GetMotoParameters(ctx *fasthttp.RequestCtx) *model.Response {
// 	motoParameters, err := s.AdminRepository.GetMotoParameters(ctx)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: motoParameters}
// }

// func (s *AdminService) CreateMotoParameter(ctx *fasthttp.RequestCtx, req *model.CreateMotoParameterRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateMotoParameter(ctx, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Moto parameter created successfully"}}
// }

// func (s *AdminService) UpdateMotoParameter(ctx *fasthttp.RequestCtx, id int, req *model.UpdateMotoParameterRequest) *model.Response {
// 	err := s.AdminRepository.UpdateMotoParameter(ctx, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto parameter updated successfully"}}
// }

// func (s *AdminService) DeleteMotoParameter(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteMotoParameter(ctx, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto parameter deleted successfully"}}
// }

// // Moto Parameter Values service methods
// func (s *AdminService) GetMotoParameterValues(ctx *fasthttp.RequestCtx, motoParamId int) *model.Response {
// 	motoParameterValues, err := s.AdminRepository.GetMotoParameterValues(ctx, motoParamId)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: motoParameterValues}
// }

// func (s *AdminService) CreateMotoParameterValue(ctx *fasthttp.RequestCtx, motoParamId int, req *model.CreateMotoParameterValueRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateMotoParameterValue(ctx, motoParamId, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Moto parameter value created successfully"}}
// }

// func (s *AdminService) UpdateMotoParameterValue(ctx *fasthttp.RequestCtx, motoParamId int, id int, req *model.UpdateMotoParameterValueRequest) *model.Response {
// 	err := s.AdminRepository.UpdateMotoParameterValue(ctx, motoParamId, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto parameter value updated successfully"}}
// }

// func (s *AdminService) DeleteMotoParameterValue(ctx *fasthttp.RequestCtx, motoParamId int, id int) *model.Response {
// 	err := s.AdminRepository.DeleteMotoParameterValue(ctx, motoParamId, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto parameter value deleted successfully"}}
// }

// // Moto Category Parameters service methods
// func (s *AdminService) GetMotoCategoryParameters(ctx *fasthttp.RequestCtx, categoryId int) *model.Response {
// 	motoCategoryParameters, err := s.AdminRepository.GetMotoCategoryParameters(ctx, categoryId)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: motoCategoryParameters}
// }

// func (s *AdminService) CreateMotoCategoryParameter(ctx *fasthttp.RequestCtx, categoryId int, req *model.CreateMotoCategoryParameterRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateMotoCategoryParameter(ctx, categoryId, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Moto category parameter created successfully"}}
// }

// func (s *AdminService) UpdateMotoCategoryParameter(ctx *fasthttp.RequestCtx, categoryId int, id int, req *model.UpdateMotoCategoryParameterRequest) *model.Response {
// 	err := s.AdminRepository.UpdateMotoCategoryParameter(ctx, categoryId, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto category parameter updated successfully"}}
// }

// func (s *AdminService) DeleteMotoBrand(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteMotoBrand(ctx, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Moto brand deleted successfully"}}
// }

// // Comtrans Categories service methods
// func (s *AdminService) GetComtransCategories(ctx *fasthttp.RequestCtx) *model.Response {
// 	comtransCategories, err := s.AdminRepository.GetComtransCategories(ctx)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: comtransCategories}
// }

// func (s *AdminService) CreateComtransCategory(ctx *fasthttp.RequestCtx, req *model.CreateComtransCategoryRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateComtransCategory(ctx, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Comtrans category created successfully"}}
// }

// func (s *AdminService) UpdateComtransCategory(ctx *fasthttp.RequestCtx, id int, req *model.UpdateComtransCategoryRequest) *model.Response {
// 	err := s.AdminRepository.UpdateComtransCategory(ctx, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans category updated successfully"}}
// }

// func (s *AdminService) DeleteComtransCategory(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteComtransCategory(ctx, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans category deleted successfully"}}
// }

// // Comtrans Brands service methods
// func (s *AdminService) GetComtransBrands(ctx *fasthttp.RequestCtx) *model.Response {
// 	comtransBrands, err := s.AdminRepository.GetComtransBrands(ctx)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: comtransBrands}
// }

// func (s *AdminService) CreateComtransBrand(ctx *fasthttp.RequestCtx, req *model.CreateComtransBrandRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateComtransBrand(ctx, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Comtrans brand created successfully"}}
// }

// func (s *AdminService) UpdateComtransBrand(ctx *fasthttp.RequestCtx, id int, req *model.UpdateComtransBrandRequest) *model.Response {
// 	err := s.AdminRepository.UpdateComtransBrand(ctx, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans brand updated successfully"}}
// }

// func (s *AdminService) DeleteComtransBrand(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteComtransBrand(ctx, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans brand deleted successfully"}}
// }

// // Comtrans Models service methods
// func (s *AdminService) GetComtransModels(ctx *fasthttp.RequestCtx) *model.Response {
// 	comtransModels, err := s.AdminRepository.GetComtransModels(ctx)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: comtransModels}
// }

// func (s *AdminService) CreateComtransModel(ctx *fasthttp.RequestCtx, req *model.CreateComtransModelRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateComtransModel(ctx, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Comtrans model created successfully"}}
// }

// func (s *AdminService) UpdateComtransModel(ctx *fasthttp.RequestCtx, id int, req *model.UpdateComtransModelRequest) *model.Response {
// 	err := s.AdminRepository.UpdateComtransModel(ctx, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans model updated successfully"}}
// }

// func (s *AdminService) DeleteComtransModel(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteComtransModel(ctx, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans model deleted successfully"}}
// }

// // Comtrans Parameters service methods
// func (s *AdminService) GetComtransParameters(ctx *fasthttp.RequestCtx) *model.Response {
// 	comtransParameters, err := s.AdminRepository.GetComtransParameters(ctx)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: comtransParameters}
// }

// func (s *AdminService) CreateComtransParameter(ctx *fasthttp.RequestCtx, req *model.CreateComtransParameterRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateComtransParameter(ctx, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Comtrans parameter created successfully"}}
// }

// func (s *AdminService) UpdateComtransParameter(ctx *fasthttp.RequestCtx, id int, req *model.UpdateComtransParameterRequest) *model.Response {
// 	err := s.AdminRepository.UpdateComtransParameter(ctx, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans parameter updated successfully"}}
// }

// func (s *AdminService) DeleteComtransParameter(ctx *fasthttp.RequestCtx, id int) *model.Response {
// 	err := s.AdminRepository.DeleteComtransParameter(ctx, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans parameter deleted successfully"}}
// }

// // Comtrans Parameter Values service methods
// func (s *AdminService) GetComtransParameterValues(ctx *fasthttp.RequestCtx, parameterId int) *model.Response {
// 	comtransParameterValues, err := s.AdminRepository.GetComtransParameterValues(ctx, parameterId)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: comtransParameterValues}
// }

// func (s *AdminService) CreateComtransParameterValue(ctx *fasthttp.RequestCtx, parameterId int, req *model.CreateComtransParameterValueRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateComtransParameterValue(ctx, parameterId, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Comtrans parameter value created successfully"}}
// }

// func (s *AdminService) UpdateComtransParameterValue(ctx *fasthttp.RequestCtx, parameterId int, id int, req *model.UpdateComtransParameterValueRequest) *model.Response {
// 	err := s.AdminRepository.UpdateComtransParameterValue(ctx, parameterId, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans parameter value updated successfully"}}
// }

// func (s *AdminService) DeleteComtransParameterValue(ctx *fasthttp.RequestCtx, parameterId int, id int) *model.Response {
// 	err := s.AdminRepository.DeleteComtransParameterValue(ctx, parameterId, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans parameter value deleted successfully"}}
// }

// // Comtrans Category Parameters service methods
// func (s *AdminService) GetComtransCategoryParameters(ctx *fasthttp.RequestCtx, categoryId int) *model.Response {
// 	comtransCategoryParameters, err := s.AdminRepository.GetComtransCategoryParameters(ctx, categoryId)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: comtransCategoryParameters}
// }

// func (s *AdminService) CreateComtransCategoryParameter(ctx *fasthttp.RequestCtx, categoryId int, req *model.CreateComtransCategoryParameterRequest) *model.Response {
// 	id, err := s.AdminRepository.CreateComtransCategoryParameter(ctx, categoryId, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.SuccessWithId{Id: id, Message: "Comtrans category parameter created successfully"}}
// }

// func (s *AdminService) UpdateComtransCategoryParameter(ctx *fasthttp.RequestCtx, categoryId int, id int, req *model.UpdateComtransCategoryParameterRequest) *model.Response {
// 	err := s.AdminRepository.UpdateComtransCategoryParameter(ctx, categoryId, id, req)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans category parameter updated successfully"}}
// }

// func (s *AdminService) DeleteComtransCategoryParameter(ctx *fasthttp.RequestCtx, categoryId int, id int) *model.Response {
// 	err := s.AdminRepository.DeleteComtransCategoryParameter(ctx, categoryId, id)
// 	if err != nil {
// 		return &model.Response{Error: err, Status: http.StatusInternalServerError}
// 	}
// 	return &model.Response{Data: model.Success{Message: "Comtrans category parameter deleted successfully"}}
// }
