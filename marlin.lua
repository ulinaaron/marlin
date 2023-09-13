local t = require("term") -- Lua terminal module

-- Get terminal size
local term_width, term_height = t.getSize() 

-- Content area size
local content_height = term_height - 1

-- Draw initial screen
t.clear()
t.moveTo(1, 1) 

-- Content area 
t.setBackgroundColor(t.lightGray)
for i=1, content_height do
  t.write("\n")
end

-- Status bar  
t.setBackgroundColor(t.black)
t.setForegroundColor(t.white)
t.moveTo(1, term_height)
t.write(filename .. " - Normal")

-- Handle resizing
t.onResize(function()
  term_width, term_height = t.getSize()
  content_height = term_height - 1
  
  t.clear()
  
  t.setBackgroundColor(t.lightGray)
  for i=1, content_height do
    t.write("\n") 
  end

  t.setBackgroundColor(t.black)
  t.setForegroundColor(t.white)
  t.moveTo(1, term_height)
  t.write(filename .. " - Normal") 
end)

-- Main loop
while true do

  -- Draw content
  t.moveTo(1, 1)
  t.write(contents)

  -- Get input
  local input = t.read()

  -- Process commands  

  -- Refresh status bar
  t.moveTo(1, term_height)
  t.write(filename .. " - " .. mode) 

end
