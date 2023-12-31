package tpl

func MainTemplate() []byte {
	return []byte(`package main

import (
	"{{ .ModName }}/execute"
)

func main() {
	execute.Execute()
}
`)
}

func ExecuteTemplate() []byte {
	return []byte(`package execute

import (
	"fmt"
	"{{ .ModName }}/comment"

	"github.com/spf13/cobra"
	"github.com/txy2023/potato/execute"
	"github.com/txy2023/potato/register"
)

var rootCmd = &cobra.Command{
	Use: "go run main.go OR ./main(the compiled name)",
	Long: ` + "`" +
		`when flag testcase is specified, the testcase specified will be executed, so is flag testsuite
if flag testcase or flag testsuite is not specified, all testcases will be executed` + "`" + `,
	Run: func(cmd *cobra.Command, args []string) {
		execute.Run(TestCasesSpecified, TestSuitesSpecified)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

var (
	TestCasesSpecified  *string
	TestSuitesSpecified *string
)

func init() {
	TestCasesSpecified = rootCmd.Flags().StringP("testcase", "c", "", 
		fmt.Sprintf("specify the testcases to execute(separated by commas), such as(total: %d):\n%s", 
		    comment.TestCaseCount, comment.TestCaseComment))
	TestSuitesSpecified = rootCmd.Flags().StringP("testsuite", "s", "", 
		fmt.Sprintf("specify the testsuites to execute(separated by commas), such as(total: %d):\n%s", 
		    comment.TestSuiteCount, comment.TestSuiteComment))
}
`)
}

func AddSuiteTemplate() []byte {
	return []byte(`package {{ .PackageName }}

// Here you can define Setup of testsuite optionally
func (d *{{ .StructName }}) Setup() (err error) {
	// implementation
	return 
}	

// Here you can define Teardown of testsuite optionally
func (d *{{ .StructName }}) Teardown() (err error) {
	// implementation
	return
}	

// Here you will define your specific testcase
func (d *{{ .StructName }}) TestCase1() (err error) {
	return 
}

`)
}

func SuiteInitTemplate() []byte {
	return []byte(`package {{ .PackageName }}

import (
	"github.com/txy2023/potato/execute"
)

type {{ .StructName }} struct{}

func (d *{{ .StructName }}) Execute() {
	execute.Execute(d)
}

`)
}

func SuiteRegisteTemplate() []byte {
	return []byte(`package execute

import (
	"{{ .ImportedPath }}"

	"github.com/txy2023/potato/register"
)

func init() {
	register.Registe(new({{ .PackageName }}.{{ .StructName }}))
}
`)
}

func TestSuiteCommentTemplate() []byte {
	return []byte(`package comment

var TestSuiteCount = {{ .TestsuiteCount }}

var TestSuiteComment = ` + "`" + `{{ .Testsuite }}` + "`")
}

func TestCaseCommentTemplate() []byte {
	return []byte(`package comment

var TestCaseCount = {{ .TestcaseCount }}	

var TestCaseComment = ` + "`" + `{{ .Testcase }}` + "`")
}
