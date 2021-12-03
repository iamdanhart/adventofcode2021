import kotlin.system.measureNanoTime
import java.io.File
import java.io.BufferedReader

fun main() {
    val part1Time = measureNanoTime {
        part1()
    }
    val part2Time = measureNanoTime {
        part2()
    }
    println("P1 in Kotlin $part1Time ns, P2 in Kotlin $part2Time ns")
}

fun part1() {
    val p1Lines = readFileAsLinesUsingBufferedReader(false)
    var prev: Int = 0
    var count: Int = 0
    for (line in p1Lines) {
        val digit = line.toIntOrNull()
        if (digit == null) {
            println("Error: failed to parse digit")
            return
        }
        if (prev != 0 && digit > prev) {
            count++;
        }
        prev = digit;
    }
    println("Part 1 solution: $count")
}

fun part2() {
    val p2Lines = readFileAsLinesUsingBufferedReader(false)
    var count: Int = 0
    var w1 = p2Lines.get(0).toInt()
    var w2 = p2Lines.get(1).toInt()
    var w3 = p2Lines.get(2).toInt()
    for (line in p2Lines.subList(3, p2Lines.size)) {
        val digit = line.toInt()
        val oldSum = w1 + w2 + w3
        val newSum = w2 + w3 + digit
        if (w2 != 0 && w3 != 0 && oldSum < newSum) {
            count++
        }
        w1 = w2
        w2 = w3
        w3 = digit
    }
    println("Part 2 solution: $count")
}

fun readFileAsLinesUsingBufferedReader(fileName: String): List<String>
    = File(fileName).bufferedReader().readLines()

fun readFileAsLinesUsingBufferedReader(useTestInput: Boolean): List<String> {
    val fileName: String = if (useTestInput) "inputs/testinput" else "inputs/input"
    return File(fileName).bufferedReader().readLines()
}