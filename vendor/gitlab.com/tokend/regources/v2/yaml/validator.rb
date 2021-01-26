### DONT CHANGE BELOW ###

require 'yaml'
require 'json-schema'

inner_folder_path = './inner'
resource_folder_path = './resources'

resource_schema_path = './schema/resource.yaml'
inner_schema_path = './schema/inner.yaml'

## Parse schemas

resource_schema = nil
inner_schema = nil

begin
  resource_schema = YAML.safe_load File.read resource_schema_path
rescue Exception => e
  puts "Failed to load '#{resource_schema_path}'. Error: #{e.message}\n"
  exit
end

begin
  inner_schema = YAML.safe_load File.read inner_schema_path
rescue Exception => e
  puts "Failed to load '#{inner_schema_path}'. Error: #{e.message}\n"
  exit
end

## Parse configs

inner_files = Dir[inner_folder_path + '/*.yaml']
resource_files = Dir[resource_folder_path + '/*.yaml']

$has_invalid_configs = false

def parse_configs(config_files, key, schema)
  configs = {}
  config_files.each do |config|
    contens = File.read config
    key_value = nil

    begin
      yaml = YAML.safe_load contens
      key_value = yaml[key]
    rescue Exception => e
      puts "Failed to load '#{config}'. Error: #{e.message}\n\n"
      $has_invalid_configs = true
      next
    end  

    if key_value.nil?
      puts "Resource config '#{config}' has no value for key '#{key}'. Abort!\n\n"
      $has_invalid_configs = true
      next
    end
    
    begin
      JSON::Validator.validate!(schema, yaml)
    rescue Exception => e
      puts "Failed to validate '#{key_value}'. Error: #{e.message}\n\n"
      $has_invalid_configs = true
      next
    end

    configs[key_value.to_s] = yaml
  end
  configs
end

$resource_configs = parse_configs resource_files, 'key', resource_schema
$inner_configs = parse_configs inner_files, 'name', inner_schema

if $has_invalid_configs
  puts "Fix invalid configs!"
  exit
end

puts "Configs are valid!"

### DONT CHANGE ABOVE ###
