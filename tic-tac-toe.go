package main

import (
	"github.com/gin-gonic/gin" 
	"encoding/json"
	"math/rand"
)

func tic_tac_toe_router(tic_tac_toe *gin.RouterGroup){
	tic_tac_toe.GET("/", func(c *gin.Context) {
		c.HTML(200, "tic-tac-toe.tmpl", nil)
	})
	
	tic_tac_toe.POST("/TurnAuto", func(c *gin.Context){
		turn := c.PostForm("turn")
		board := make([][]string, 3)      
		for i:=0; i<3; i++ {
			board[i] = make([]string, 3)
		}
		json.Unmarshal([]byte(c.PostForm("board")), &board)

		diff_type := ""
		max_score := 0
		max_score_point := score_type{}

		if turn == "X"{
			diff_type = "O"  
		} else {
			diff_type = "X"  
		}

		new_point_list := make([]point_type, 0)

		for index1, value1 := range board {
			for index2, value2 := range value1 {
				if value2 == ""{
					new_point_list = append(new_point_list, point_type{x: index1, y:index2})
				}
			}
		}

		for _, point := range new_point_list {
			x := point.x
			y := point.y
			score := 0

			if (board[0][0] == "" || board[0][1] == "" || board[0][2] == "") && x == 0{
				score += calc(board[0][0], board[0][1], board[0][2], turn, diff_type)
			}
			
			if (board[1][0] == "" || board[1][1] == "" || board[1][2] == "") && x == 1{
				score += calc(board[1][0], board[1][1], board[1][2], turn, diff_type)
			}
				
			if (board[2][0] == "" || board[2][1] == "" || board[2][2] == "") && x == 2{
				score += calc(board[2][0], board[2][1], board[2][2], turn, diff_type)
			}

			if (board[0][0] == "" || board[1][0] == "" || board[2][0] == "") && y == 0{
				score += calc(board[0][0], board[1][0], board[2][0], turn, diff_type)
			}

			if (board[0][1] == "" || board[1][1] == "" || board[2][1] == "") && y == 1{
				score += calc(board[0][1], board[1][1], board[2][1], turn, diff_type)
			}

			if (board[0][2] == "" || board[1][2] == "" || board[2][2] == "") && y == 2{
				score += calc(board[0][2], board[1][2], board[2][2], turn, diff_type)
			}
			
			if (board[0][0] == "" || board[1][1] == "" || board[2][2] == "") && x == y{
				score += calc(board[0][0], board[1][1], board[2][2], turn, diff_type)
			}

			if (board[0][2] == "" || board[1][1] == "" || board[2][0] == "") && x+y == 2{
				score += calc(board[0][2], board[1][1], board[2][0], turn, diff_type)
			}

			board[x][y] = ""
			if score >= max_score{
				max_score_point = score_type{ score: score, point_board: point }
				max_score = score
			}
		}

		if max_score > 0{
			c.JSON(200, gin.H{"x":max_score_point.point_board.x, "y":max_score_point.point_board.y, "score":max_score_point.score})
		} else {
			index := rand.Intn(len(new_point_list))
			c.JSON(200, gin.H{"x":new_point_list[index].x, "y":new_point_list[index].y})
		}
	})
}

func calc(new_board1 string, new_board2 string, new_board3 string, current_type string, diff_type string) int {
	temp_score := 0

    if new_board1 == current_type{
        temp_score += 1
	} 
    if new_board2 == current_type{
        temp_score += 1 
	} 
    if new_board3 == current_type{
        temp_score += 1 
	}

    if new_board1 == diff_type{
        temp_score -= 1 
	} 
    if new_board2 == diff_type{
        temp_score -= 1 
	}
    if new_board3 == diff_type{
        temp_score -= 1
	}

	if temp_score == 2 {
        return 3
	} else if temp_score == -2{
        return 2
	} else if temp_score == 1{
        return 1
	} else {
        return 0
	}
}   