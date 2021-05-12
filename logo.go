package inc

import "fmt"

const version = "1.0.0"

func Logo(name string) string {
	s := `
  ___  ________   ________     
 |\  \|\   ___  \|\   ____\       
 \ \  \ \  \\ \  \ \  \___|    
  \ \  \ \  \\ \  \ \  \       			
   \ \  \ \  \\ \  \ \  \____      %s
    \ \__\ \__\\ \__\ \_______\    
     \|__|\|__| \|__|\|_______|    Internal Network Channel
`

	return fmt.Sprintf(s, name)
}
