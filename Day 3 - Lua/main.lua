-- Disclaimer, I have never used Lua before

-- For easier indexing
getmetatable('').__index = function(str,i)
    return type(i) == 'number' and string.sub(str,i,i) or string[i]
end

-- parse input
function helperFunc()
    local parseddata = {}
    local inputFile = assert(io.open("input.txt", "r"))
    local data = inputFile:read("*all")
    inputFile:close()
    local line = {}
    
    -- convert data into 2d array
    for i = 1, string.len(data) do
        if data[i] == '\n' or i == #data then
            parseddata[#parseddata+1] = line  -- append table
            line = {}
        else
            line[#line+1] = data[i]  -- append character to table
        end
    end

    return parseddata
end


function part1(data, rise, run)
    local x = 1
    local y = 1
    local trees = 0

    while y <= #data do

        -- calculate overflow 
        local newX = x - (#data[1] * math.floor(x / #data[1]))

        -- increment trees
        if data[y][(newX ~= 0 and newX or #data[1])] == "#" then
            trees = trees + 1
        end
  
        x = x + run
        y = y + rise

    end

    return trees
end


function part2(data, set)
    local product = 1
    for i = 1,  #set do
        product = product * part1(data, set[i][1], set[i][2])
    end
    return product
end


data = helperFunc()
print("Part 1: " .. part1(data, 1, 3))
print("Part 2: " .. part2(data, {{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}))