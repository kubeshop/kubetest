/*
 * Kubtest API
 *
 * Kubtest provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: kubtest@kubshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package kubtest

type ExecutionsSummary []ExecutionSummary

func (executions ExecutionsSummary) Table() (header []string, output [][]string) {
	header = []string{"Script", "Type", "Name", "ID", "Status"}

	for _, e := range executions {
		output = append(output, []string{
			e.ScriptName,
			e.ScriptType,
			e.Name,
			e.Id,
			e.Status,
		})
	}

	return
}
