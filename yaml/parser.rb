require 'yaml'

File.open('sample.yml') do | io |
    YAML.load_stream(io) do | data |
        p data
    end
end
