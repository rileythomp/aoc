colors = [
    '#F0F8FF',
    '#E6E6FA',
    '#B0E0E6',
    '#ADD8E6',	
    '#87CEFA',
    '#87CEEB',
    '#00BFFF',
    '#1E90FF',
    '#0000FF',
    '#000080'
]
var adjacents = [
	[-1, -1],
	[0, -1],
	[1, -1],
	[-1, 0],
	[1, 0],
	[-1, 1],
	[0, 1],
	[1, 1]
]
function main() {
    lights = document.getElementById('lights')
    board = [
        [5,4,8,3,1,4,3,2,2,3],
        [2,7,4,5,8,5,4,7,1,1],
        [5,2,6,4,5,5,6,1,7,3],
        [6,1,4,1,3,3,6,1,4,6],
        [6,3,5,7,3,8,5,4,7,8],
        [4,1,6,7,5,2,4,6,4,5],
        [2,1,7,6,8,4,1,7,2,1],
        [6,8,8,2,8,8,1,1,3,4],
        [4,8,4,6,8,4,8,5,5,4],
        [5,2,8,3,7,5,1,5,2,6]
    ]

    for (let i = 0; i < board.length; i++) {
        row = document.createElement('tr')
        for (let j = 0; j < board[i].length; j++) {
            td = document.createElement('td')
            // td.innerHTML = board[i][j]
            td.style.backgroundColor = colors[board[i][j]]
            row.appendChild(td)
        }
        lights.appendChild(row)
    }
    step = 0
	flashes = 0
    lights = document.getElementById('lights')
    setInterval(async function(){
        for (y = 0; y < board.length; y++) {
			for (x = 0; x < board[y].length; x++) {
				board[y][x]++
                lights.children[y].children[x].style.backgroundColor = colors[board[y][x]]
			}
		}
		for (y = 0; y < board.length; y++) {
			for (x = 0; x < board[y].length; x++) {
				stack = [[x, y]]
				for (;stack.length != 0;) {
					coord = stack.pop()
					if (board[coord[1]][coord[0]] > 9) {
                        flashes++
						board[coord[1]][coord[0]] = -1
						for (let i = 0; i < adjacents.lengths; i++) {
                            adj = adjacents[i]
							if (coord[1]+adj[1] >= 0 && coord[1]+adj[1] < board.length &&
								coord[0]+adj[0] >= 0 && coord[0]+adj[0] < board.length && board[coord[1]+adj[1]][coord[0]+adj[0]] >= 0){
								board[coord[1]+adj[1]][coord[0]+adj[0]]++
								stack.push([coord[0] + adj[0], coord[1] + adj[1]])
							}
						}
					}
				}
			}
		}
		flashed = 0
		for (y = 0; y < board.length; y++) {
			for (x = 0; x < board[y].length; x++) {
				if (board[y][x] < 0) {
					flashed++
					board[y][x] = 0
				}
			}
		}
        // lights = document.getElementById('lights')
        // lights.innerHTML = ''
        // for (let i = 0; i < board.length; i++) {
        //     row = document.createElement('tr')
        //     for (let j = 0; j < board[i].length; j++) {
        //         td = document.createElement('td')
        //         // td.innerHTML = board[i][j]
        //         td.style.backgroundColor = colors[board[i][j]]
        //         row.appendChild(td)
        //     }
        //     lights.appendChild(row)
        // }
		step++
		if (flashed == 100) {
            // await new Promise(r => setTimeout(r, 2000));

            alert("big flash")
			// return step
		}
    }, 500);
}
main()

