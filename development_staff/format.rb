require 'json'
require 'byebug'

raw_data = File.read('data/raw.json')
data = JSON.parse(raw_data)

loops = data['geometries'].first['coordinates']

name = ARGV[0]
fias_id = ARGV[1]

loops.map! { |l| l.reverse! }

result = [{ name: name, fias_id: fias_id, loops: loops }]

result_file = File.open("data/region_#{name}.json", 'wb')

result_file.write(result.to_json)
