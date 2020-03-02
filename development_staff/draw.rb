# this script was used for animated visual debugging in development
# but it is outdated and no longer working
# It's kept because it may be fixed and used later
require 'ruby2d'
require 'json'

# Set the window size
set width: 3000, height: 2000

tick = 0

scale = 100
raw_data = File.read('data/regions2.json')
data = JSON.parse(raw_data)

Circle.new(
  x: (57.605012 - 30) * scale,
  y: (39.746949 - 30) * scale,
  radius: 4 * scale,
  color: 'fuchsia',
  z: 10
)

update do
  coord = data[tick]
  new_coord = data[tick + 1]

  return if new_coord.nil?

  Line.new(
    x1: (coord[1] - 30) * scale, y1: coord[0] * scale - 30 * scale,
    x2: (new_coord[1] - 30) * scale, y2: (new_coord[0] - 30) * scale,
    width: 5,
    color: 'lime'
  )
  tick += 1
end

# Show the window
show
