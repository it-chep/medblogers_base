package config

import "medblogers_base/internal/pkg/config"

const (
	// SearchCitiesLimit /* Количество городов которые будут отображаться при поиске ДЕФОЛТ = 5 */
	SearchCitiesLimit = config.ConfigKey("searchCitiesLimit")
	// SearchSpecialitiesLimit /* Количество специальностей которые будут отображаться при поиске ДЕФОЛТ = 5 */
	SearchSpecialitiesLimit = config.ConfigKey("searchSpecialitiesLimit")
	// SearchDoctorsLimit /* Количество врачей которые будут отображаться при поиске ДЕФОЛТ = 30 */
	SearchDoctorsLimit = config.ConfigKey("searchDoctorsLimit")
)
