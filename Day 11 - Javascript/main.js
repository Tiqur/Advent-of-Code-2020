const fs = require('fs');
let rows = fs.readFileSync("input.txt", 'utf8').split("\r\n").map(r => r.split(""));
let cache = [];


function isOccupied(y, x) {
    if (outOfBounds(y, x)) return false;
    return rows[y][x] == "#";
}

function outOfBounds(y, x) {
    if (y < 0 || y >= rows.length || x < 0 || x >= rows[0].length) return true;
}


function updateDirection(c, i) {
    let y = c[0]
    let x = c[1]
    // Direction
    switch (i) {
        case 0:  // N
            y++;
            break;
        case 1: // E
            x++;
            break;
        case 2: // S
            y--;
            break;
        case 3: // W
            x--;
            break;
        case 4: // NE
            y++;
            x++;
            break;
        case 5: // SE
            x++;
            y--;
            break;
        case 6: // SE
            y--;
            x--;
            break;
        case 7: // NW
            x--;
            y++;
            break;
    }
    return [y, x]
}


function getVisibleSeats(sy, sx) {
    let visible_seats = 0;
    for (k = 0; k < 8; k++) {
        let c = [sy, sx];
        c = updateDirection(c, k)

        while (!outOfBounds(c[0], c[1])) {

            // Current Seat
            if (rows[c[0]][c[1]] == "L") break;
            if (rows[c[0]][c[1]] == "#") {
                visible_seats++
                break;
            };
            // update slope
            c = updateDirection(c, k)
        }
    }
        return visible_seats;
}



function getAdjancent(y, x) {
    let adj = 0;
    let dir = [[1, 0], [-1, 0], [0, 1], [0, -1], [1, 1], [1, -1], [-1, 1], [-1, -1]] 
    for (k = 0; k < 8; k++) 
        if (isOccupied(y+dir[k][0], x+dir[k][1])) adj++;    
    return adj;
}

function deepCopy(data) {
    return JSON.parse(JSON.stringify(data))
}


function countSeats() {
    let count = 0;
    for (i = 0; i < rows.length; i++) 
        for (j = 0; j < rows[i].length; j++) 
            if (rows[i][j] == "#") 
                count++;
    return count;
}


function part1() {
    let previous;
    while (true) {

        // occupy seat if it and surrounding area is unoccupied
        cache = deepCopy(rows);
        for (i = 0; i < rows.length; i++) 
            for (j = 0; j < rows[i].length; j++) 
                if (rows[i][j] == "L" && getAdjancent(i, j) == 0) 
                    cache[i][j] = "#";
                
            
        // deoccupy seats if > 3 adjacent occupied seats
        rows = deepCopy(cache);
        for (i = 0; i < rows.length; i++) 
            for (j = 0; j < rows[i].length; j++) 
                if (isOccupied(i, j) && getAdjancent(i, j) > 3) 
                    cache[i][j] = "L";
                
            
        rows = deepCopy(cache);
        if (rows == previous) break;
        previous = rows.join()    
    }

    return countSeats()
}


function part2() {
    let previous;
    while (true) {
        // occupy seat if it and surrounding area is unoccupied
        cache = deepCopy(rows);
        for (i = 0; i < rows.length; i++) 
            for (j = 0; j < rows[i].length; j++) 
                if (rows[i][j] == "L" && getVisibleSeats(i, j) == 0) 
                    cache[i][j] = "#";
            
            
        // deoccupy seats if > 3 adjacent occupied seats
        rows = deepCopy(cache);
        for (i = 0; i < rows.length; i++) 
            for (j = 0; j < rows[i].length; j++) 
                if (isOccupied(i, j) && getVisibleSeats(i, j) > 4) 
                    cache[i][j] = "L";
                

        rows = deepCopy(cache);
        if (rows == previous) break;
        previous = rows.join()
    }
    return countSeats()
}

console.log(`Part 1: ${part1()}`);
// Pass by reference is resulting in bugs that are not very fun, ill fix this later 
rows = fs.readFileSync("input.txt", 'utf8').split("\r\n").map(r => r.split(""));
console.log(`Part 2: ${part2()}`);
