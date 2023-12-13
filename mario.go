package main
import "fmt"

func main() {

    
	var height float64 = get_int("enter the height: ") 
	if height <= 0 || height >18{
		fmt.Printf("error  height can't be more than 18 or less than 1 \n")
	} else {
        print_pyramid(int(height),int(height))
    }

	
}

func get_int( print_statement string ) float64{

	var input_number float64 
    fmt.Printf("%s",print_statement)
	fmt.Scan(&input_number) 
	fmt.Printf("\n")
	return input_number

}

func print_pyramid(height int, line int )  {

	if line == 0{
		return
	} else {
        print_pyramid(height,line - 1)
	}
   
    print_blanSpaces(height-line)
	print_hash(line)
	fmt.Printf("  ")
	print_hash(line)
	fmt.Printf("\n")

}

func print_hash(line int){
	for i := 0; i < line; i++ {
		fmt.Printf("#")
	}
}

func print_blanSpaces(no_of_spaces int){
	for i := 0; i < no_of_spaces; i++ {
		fmt.Printf(" ")
	}
}