package inputoutput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zaffar-maqbool/Golang/calci/info"
)

var reader = bufio.NewReader(os.Stdin)

func GetUserMetrics() (weight float64, height float64) {

	weight = GetUserInput(info.WeightPrompt)
	height = GetUserInput(info.HeightPrompt)

	return
}
func GetUserInput(WeightPrompt string) (value float64) {
	fmt.Print(WeightPrompt)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", " ", -1)

	value, _ = strconv.ParseFloat(userInput, 64)

	return
}

/* fmt.Print(info.WeightPrompt)
 weightInput, _ := reader.ReadString('\n')
 weightInput = strings.Replace(weightInput, "\n", "", -1)
 weight, _ := strconv.ParseFloat(weightInput, 64)
 fmt.Print(info.HeightPrompt)
 heightInput, _ := reader.ReadString('\n')
 heightInput = strings.Replace(heightInput, "\n", "", -1)
height, _ := strconv.ParseFloat(heightInput, 64)
*/

/* func GetUserMetrics() (weight float64, height float64){
		fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')
	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)
	// converting next line
	weight, _ = strconv.ParseFloat(weightInput, 64)
	height, _ = strconv.ParseFloat(heightInput, 64)
	return
}*/
