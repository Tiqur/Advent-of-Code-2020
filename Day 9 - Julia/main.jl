# Disclaimer, I have never used Julia before
input = read("input.txt", String)
data = map(x->parse(Int64, x), split(input, "\r\n"))


function checkIfAvailableSum(data, index, preamble) 
    target = data[index]
    min = (index-preamble > 1 ? index-preamble : 1)
    for i in min:index-1
        for j in i:index-1
            if data[i] + data[j] == target 
                #println(string(data[i]) * "+" * string(data[j]) * "=" * string(target))
                return true
            end
        end
    end

    return false 
end

function part2(data, index)
    target = data[index]
    # for each number
    for i in 1:index-1
        # slice size
        for j in i:index-1
            slice = data[i:j]
            value = 0
            for num in slice
                value += num
            end
            if value == target
                slice = sort(slice)
                return(string(first(slice) + last(slice)))
            end
        end
    end
end


function part1()
    preamble = 25
    for i in preamble+1:size(data)[1]
        if !checkIfAvailableSum(data, i, preamble+1)
            return i
        end
    end
end

println("Part 1: " * string(data[part1()]))
println("Part 2: " * string(part2(data, part1())))


# for num in data
#     println(num)
# end




