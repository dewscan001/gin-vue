const app = Vue.createApp({
    data() {
        return {
                board: [
                    ['', '', ''],
                    ['', '', ''],
                    ['', '', ''],
                ],
                turn : 'X',
                your_turn : '',
                board_over : false,
                win_index_list : []
                }
            },
    methods: {
        // Manual
        manual_click(row, col){
            if (this.board[row][col] == '' && this.turn == this.your_turn){
                this.board[row][col] = this.your_turn
                this.check_game()
            }
        },

        // Check Game
        check_game() {
            let game_over = false
            let have_value = 0
            if (this.board[0][0] == this.board[0][1] && this.board[0][1] == this.board[0][2] && this.board[0][2] == this.board[0][0] && this.board[0][0] == this.turn) {
                game_over = true
                this.win_index_list_push({ x: 0, y: 0 }, { x: 0, y: 1 }, { x: 0, y: 2 })
            }
            if (this.board[1][0] == this.board[1][1] && this.board[1][1] == this.board[1][2] && this.board[1][2] == this.board[1][0] && this.board[1][0] == this.turn) {
                game_over = true
                this.win_index_list_push({ x: 1, y: 0 }, { x: 1, y: 1 }, { x: 1, y: 2 })
            }
            if (this.board[2][0] == this.board[2][1] && this.board[2][1] == this.board[2][2] && this.board[2][2] == this.board[2][0] && this.board[2][0] == this.turn) {
                game_over = true
                this.win_index_list_push({ x: 2, y: 0 }, { x: 2, y: 1 }, { x: 2, y: 2 })
            }
            if (this.board[0][0] == this.board[1][0] && this.board[1][0] == this.board[2][0] && this.board[2][0] == this.board[0][0] && this.board[0][0] == this.turn) {
                game_over = true
                this.win_index_list_push({ x: 0, y: 0 }, { x: 1, y: 0 }, { x: 2, y: 0 })
            }
            if (this.board[0][1] == this.board[1][1] && this.board[1][1] == this.board[2][1] && this.board[2][1] == this.board[0][1] && this.board[0][1] == this.turn) {
                game_over = true
                this.win_index_list_push({ x: 0, y: 1 }, { x: 1, y: 1 }, { x: 2, y: 1 })
            }
            if (this.board[0][2] == this.board[1][2] && this.board[1][2] == this.board[2][2] && this.board[2][2] == this.board[0][2] && this.board[0][2] == this.turn) {
                game_over = true
                this.win_index_list_push({ x: 0, y: 2 }, { x: 1, y: 2 }, { x: 2, y: 2 })
            }
            if (this.board[0][0] == this.board[1][1] && this.board[1][1] == this.board[2][2] && this.board[2][2] == this.board[0][0] && this.board[0][0] == this.turn) {
                game_over = true
                this.win_index_list_push({ x: 0, y: 0 }, { x: 1, y: 1 }, { x: 2, y: 2 })
            }
            if (this.board[0][2] == this.board[1][1] && this.board[1][1] == this.board[2][0] && this.board[2][0] == this.board[0][2] && this.board[0][2] == this.turn) {
                game_over = true
                this.win_index_list_push({ x: 0, y: 2 }, { x: 1, y: 1 }, { x: 2, y: 0 })
            }

            if (!game_over) {
                for (let i = 0; i < 3; i++) {
                    for (let j = 0; j < 3; j++) {
                        if (this.board[i][j]) {
                            have_value++
                        }
                        else {
                            have_value--
                        }
                    }
                }
                if (have_value == 9) {
                    alert('DRAW!!')
                    have_value = 0
                    this.board_over = true
                }
                else{
                    this.turn = this.turn == 'X' ? 'O' : 'X'
                }
            }
            else {
                alert(this.turn + ' WIN!!')
                game_over = false
                have_value = 0
                this.board_over = true
            }
        },
        win_index_list_push(board1, board2, board3) {
            this.win_index_list.push(board1)
            this.win_index_list.push(board2)
            this.win_index_list.push(board3)
        },
        check_win_index_list(x, y) {
            if (this.win_index_list.length > 0) {
                for (let index = 0; index < 3; index++) {
                    if (x == this.win_index_list[index]["x"] && y == this.win_index_list[index]["y"]) {
                        return true
                    }
                }
            }
            return false
        },

        // Reset
        reset_board(){
            this.board = [
                ['', '', ''],
                ['', '', ''],
                ['', '', '']
            ]
            this.board_over = false
            this.your_turn = ''
            this.turn = 'X'
            this.win_index_list = []
        },
        // API
        turnAuto() {
            if (this.turn != this.your_turn && this.your_turn != '') {

                let formData = new FormData();
                formData.append('board', JSON.stringify(this.board));
                formData.append('turn', this.turn)

                fetch('/TurnAuto', {
                    method: 'POST',
                    body: formData
                })
                .then(response => response.json())
                .then(data => {
                    console.log(data)
                    if (this.board[data["x"]][data["y"]] == '') {
                        this.board[data["x"]][data["y"]] = this.turn
                    } else {
                        alert("ERROR")
                        this.reset_board()
                    }
                    this.check_game()
                })
                .catch(err => {
                    console.log(err)
                })
            }
        }
    },
    watch:{
        turn(){
            this.turnAuto()
        },
        your_turn() {
            this.turnAuto()
        }
    }
})
app.mount('#app');