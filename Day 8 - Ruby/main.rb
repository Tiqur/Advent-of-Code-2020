# Disclaimer, I have never used Ruby before
require 'set'
file = File.read("input.txt")
lines = file.split("\n")


parsedInstructions = []

lines.each{ |line|
    ins = line.split(" ")
    parsedInstructions.push([ins[0], ins[1]])
}


def checkForRepeat(array)
    set = Set.new
    min = -1
    for i in (array.length-1).downto(0)
        if set.include?(array[i])
            min = i
        else
            set.add(array[i])
        end
    end
    return min != -1
end


def part1(instructions)
    executedPositions = []
    accumulator = 0;
    pointer = 0
    while (pointer < instructions.length) do
        ins = instructions[pointer][0]
        value = Integer(instructions[pointer][1])

        checkForRepeat(executedPositions)
        # if contains duplicate pointer
        if (executedPositions.index(pointer) != nil) 
            break
        else
            executedPositions.push(pointer)
        end

        if ins == "acc"
            accumulator += value
        elsif ins == "jmp"
            pointer += value-1
        end

        pointer += 1
    end
    return accumulator
end


def part2(instructions)
    pointer = 0
    corrupeted_pointer = 0

    while (pointer < instructions.length) do
        ins = instructions[pointer][0]

        if ins == "jmp" || ins == "nop"
            tempPointer = 0
            executedPositions = []
            repeats = false

            # modify instructions
            tempInstructions = instructions.map(&:clone)
            tempInstructions[pointer][0] = (ins == "jmp" ? "nop" : "jmp")

            # calculate modified instructions
            while (tempPointer < tempInstructions.length) do
                tempIns = tempInstructions[tempPointer][0]
                tempValue = Integer(tempInstructions[tempPointer][1])
                executedPositions.push(tempPointer) 
                
                if tempIns == "jmp"
                    tempPointer += tempValue-1 if tempIns == "jmp"
                end

                # if program still repeats, then targeted pointer is not "corrupt"
                if (checkForRepeat(executedPositions)) 
                    repeats = true
                    break
                end
                
                tempPointer += 1
            end
            
            # if program terminates ( doesn't repeat ) then we found the corrupted pointer
            if (!repeats)
                corrupeted_pointer = pointer
            end
        end

        pointer = pointer + 1
    end

    # fix "corrupted" instruction
    instructions[corrupeted_pointer][0] = (instructions[corrupeted_pointer][0] == "jmp" ? "nop" : "jmp")
    return part1(instructions)
end


puts "Part 1: " + String(part1(parsedInstructions))
puts "Part 2: " + String(part2(parsedInstructions))