package config

func ConfigFiles() ([]string, error) {
	//*if ENV['SENSU_CONFIG_FILES']
	//ENV['SENSU_CONFIG_FILES'].split(':')
	//else
	//['/etc/sensu/config.json'] + Dir['/etc/sensu/conf.d/**/*.json']
	return []string{}, nil
}

func LoadConfig(fn string) (map[string]interface{}, error) {
	//JSON.parse(File.open(filename, 'r').read) rescue Hash.new
	m := make(map[string]interface{})
	return m, nil
}

func settings() (map[string]interface{}, error) {
	//@settings ||= config_files.map {|f| load_config(f) }.reduce {|a, b| a.deep_merge(b) }
	m := make(map[string]interface{})
	return m, nil
}
