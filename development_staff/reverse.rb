require 'json'

raw_data = File.read('data/regions2.json')
data = JSON.parse(raw_data)