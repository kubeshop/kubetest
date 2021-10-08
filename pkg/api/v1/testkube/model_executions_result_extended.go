package testkube

func (result ExecutionsResult) Table() (header []string, output [][]string) {
	header = []string{"Script", "Type", "Name", "ID", "Status"}

	for _, e := range result.Results {
		var status string
		if e.Status != nil {
			status = string(*e.Status)
		}
		output = append(output, []string{
			e.ScriptName,
			e.ScriptType,
			e.Name,
			e.Id,
			status,
		})
	}

	return
}
