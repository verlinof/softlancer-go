package services

import (
	"context"
	"fmt"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
)

type ProjectService struct{}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (p *ProjectService) GetAllProjects(ctx context.Context) ([]models.ProjectDetail, error) {
	var projects []models.ProjectDetail
	query := `
		SELECT 
		projects.id, 
		projects.project_title, 
		projects.project_description, 
		projects.job_type, 
		projects.status, 
		roles.role_name, 
		companies.company_name, 
		companies.company_description, 
		companies.company_logo
		FROM projects JOIN 
		roles on projects.role_id = roles.id JOIN
		companies on projects.company_id = companies.id
		ORDER BY created_at DESC
	`
	tx := database.DB.WithContext(ctx).Raw(query).Scan(&projects)

	return projects, tx.Error
}

func (p *ProjectService) GetOpenProjects(ctx context.Context) ([]models.ProjectDetail, error) {
	var projects []models.ProjectDetail
	query := `
		SELECT 
		projects.id, 
		projects.project_title, 
		projects.project_description, 
		projects.job_type, 
		projects.status, 
		roles.role_name, 
		companies.company_name, 
		companies.company_description, 
		companies.company_logo
		FROM projects JOIN 
		roles on projects.role_id = roles.id JOIN
		companies on projects.company_id = companies.id
		WHERE projects.status = 'open'
		ORDER BY created_at DESC
	`
	tx := database.DB.WithContext(ctx).Raw(query).Scan(&projects)

	return projects, tx.Error
}
func (p *ProjectService) GetProjectByID(ctx context.Context, id string) (*models.ProjectDetail, error) {
	var project models.ProjectDetail
	query := `
		SELECT 
		projects.id, 
		projects.project_title, 
		projects.project_description, 
		projects.job_type, 
		projects.status, 
		roles.role_name, 
		companies.company_name, 
		companies.company_description, 
		companies.company_logo
		FROM projects JOIN 
		roles on projects.role_id = roles.id JOIN
		companies on projects.company_id = companies.id
		WHERE projects.id = ?
	`
	tx := database.DB.WithContext(ctx).Raw(query, id).Scan(&project)

	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("record not found")
	}

	return &project, tx.Error
}

func (p *ProjectService) CreateProject(ctx context.Context, project *models.Project) error {
	tx := database.DB.WithContext(ctx).Create(project)
	return tx.Error
}

func (p *ProjectService) UpdateProject(ctx context.Context, id string, project *requests.ProjectRequest) error {
	tx := database.DB.WithContext(ctx).Table("projects").Where("id = ?", id).Updates(project)
	return tx.Error
}

func (p *ProjectService) DeleteProject(ctx context.Context, id string) error {
	tx := database.DB.WithContext(ctx).Where("id = ?", id).Delete(&models.Project{})
	return tx.Error
}
