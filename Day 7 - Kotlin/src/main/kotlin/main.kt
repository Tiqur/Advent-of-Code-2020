// Disclaimer, I have never used Kotlin before ( I have used Java )
import java.io.File


class Bag(bag_id: String, bag_contents: MutableSet<Pair<String, Int>>) {
    val id = bag_id
    val contents = bag_contents
}


fun main() {
    val uniqueBags = parseInput()
    println("Part 1: ${part1(uniqueBags)}")
    println("Part 2: ${part2(uniqueBags)}")
}

fun parseInput(): Set<Bag> {
    val inputFile = File("input.txt").inputStream().bufferedReader().use { it.readText() }

    val uniqueBags = mutableSetOf<Bag>()
    for (bag in inputFile.split("\n")) {
        val bagContents = mutableSetOf<Pair<String, Int>>()
        val bagId = bag.substring(0, bag.indexOf(" bags"))
        val unparsedContents = bag.substring(bag.indexOf(" contain")+9, bag.length).split(", ")

        // parse bag contents
        for (i in unparsedContents) {
            val amount = i[0]
            if (amount.isDigit()) { // if can contain bags
                val contents = Pair(i.substring(i.indexOf(" ")+1, i.indexOf("bag")-1), amount.toString().toInt())
                bagContents.add(contents)
            }
        }
        uniqueBags.add(Bag(bagId, bagContents))
    }
    return uniqueBags
}

// Didn't see a Set.find function, so this is a basic implementation of that
fun get(type: String, unique_bags: Set<Bag>): Bag? {
    for (bag in unique_bags)
        if (bag.id == type) return bag
    return null
}


// recursively checks each bag to see if it contains target ( or contains a bag that contains target )
fun findTarget(target: String, bagtype: String, unique_bags: Set<Bag>): Boolean {
    val bag = get(bagtype, unique_bags)
    for (content in bag!!.contents)
        if (findTarget(target, content.first, unique_bags) || content.first == target) return true
    return false
}


// recursively goes through each bag in target and multiplies each of it's contents by the amount held
fun countInTarget(bagtype: String, unique_bags: Set<Bag>): Int {
    var bagCount = 1
    val bag = get(bagtype, unique_bags)
    for (content in bag!!.contents)
        bagCount += content.second * countInTarget(content.first, unique_bags)
    return bagCount
}



fun part1(unique_bags: Set<Bag>): Int {
    var largeEnough = 0
    for (bag in unique_bags)
        if (findTarget("shiny gold", bag.id, unique_bags)) largeEnough++
    return largeEnough
}

fun part2(unique_bags: Set<Bag>): Int {
    return countInTarget("shiny gold", unique_bags) -1
}