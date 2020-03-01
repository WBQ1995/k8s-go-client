package vo

type DeploymentConfig struct {

	Name          string            	`json:"name"`
	Labels        map[string]string 	`json:"labels"`
	Replicas      int32             	`json:"replicas"`
	ContainerName string            	`json:"container_name"`
	ImageName     string            	`json:"image_name"`
	ContainerPort int32             	`json:"port"`
	WorkingDir    string            	`json:"working_dir"`
 }

type ServiceConfig struct {
	Name 		  string 				`json:"name"`
	Labels        map[string]string 	`json:"labels"`
	Port 		  int32					`json:"port"`
}
