func isValidSudoku(board [][]byte) bool {
    if len(board)!=9 || len(board[0])!=9{
        return false
    }

    length:=9

    //square checking number id same or not
     for x:=0;x<length;x++{
        for y:=0;y<length;y++{
            if board[x][y]!=
            '.'{
                start_x:=x-x%3
                start_y:=y-y%3
            
               if !IsValidSquare(start_x,start_y,board,x,y){
                fmt.Println("1","x","y",x,y)
                return false
               }
           
          if !IsVaildVerticallyAndHorizontally(x,y,board){
            fmt.Println("5","x","y",x,y)
            return false
          }
            }
        }
     }

     return true

}

func IsValidSquare(x int,y int, board [][]byte, board_x int, board_y int) bool{

    for i:=x;i<=x+2;i++{
        for j:=y;j<=y+2;j++{
            if i!=board_x && j!=board_y && board[board_x][board_y] == board[i][j]{
                fmt.Println("inside","i","j",i,j,board_x,board_y,x,y)
                return false
            }
        }

    }
    return true
}

func IsVaildVerticallyAndHorizontally(x int, y int, board [][]byte) bool{
    for i:=0;i<9;i++{
        if i!=y && board[x][i] == board[x][y]{
            return false
        }

        if i!=x && board[i][y] == board[x][y]{
            return false
        }
    }
    return true
}