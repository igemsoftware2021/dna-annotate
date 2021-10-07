package main


import (
	"fmt"
	"io/ioutil"
    "log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
  Execute()
}

var (
	input string
	output string
)


var rootCmd = &cobra.Command{
	Use:   "annotator",
	Short: "Annotator is github action to annotate problematic sequences from given Genbank file.",
	Long: `Annotator is github action to annotate problematic sequences from given Genbank file.`,
	Run: func(cmd *cobra.Command, args []string) {
		Script(input)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Directory where all the input genbank files will be read")
	rootCmd.PersistentFlags().StringVarP(&output, "ouput", "o", "", "Directory where all the output genbank files wil be written")
	
	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("ouput")	
}

func Script(input string) {
	files, err := ioutil.ReadDir("./" + input)
    if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
            fmt.Println(f.Name())
    }
}

